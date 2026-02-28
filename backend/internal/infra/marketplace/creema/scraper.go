package creema

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"

	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	"github.com/tokushun109/tku/backend/internal/usecase"
	usecaseProduct "github.com/tokushun109/tku/backend/internal/usecase/product"
)

const (
	defaultTimeout = 10 * time.Second

	creemaProductPageHost = "www.creema.jp"

	maxProductPageBodySize  = 2 << 20  // 2MB
	maxProductImageBodySize = 20 << 20 // 20MB

	// Creemaの商品ページの現在のDOM構造に依存するため、サイト変更時に見直しが必要。
	priceSelector       = "#js-item-detail > aside > div.p-item-detail-info.p-item-detail-info--side > div > div:nth-child(1) > div.p-item-detail-info__item--price--row > span.p-item-detail-info__item--price--row--price"
	descriptionSelector = "#introduction > div > div"
	imageSelector       = "#js-item-detail-centerpiece > img.js-attach-centerpiece"
	tagSelector         = "a.js-item-attributes-tag"
)

var allowedImageHosts = map[string]struct{}{
	"c.p02.c4a.im": {},
}

type Scraper struct {
	client *http.Client
}

var _ usecaseProduct.DuplicateSource = (*Scraper)(nil)

func NewScraper() *Scraper {
	return &Scraper{
		client: &http.Client{Timeout: defaultTimeout},
	}
}

func (s *Scraper) Duplicate(ctx context.Context, rawURL string) (*usecaseProduct.DuplicateProductData, error) {
	productURL, err := validateProductPageURL(rawURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, productURL.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.doRequest(req, validateProductPageRequestURL)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, mapProductPageStatusError(resp.StatusCode)
	}

	body, err := readResponseBodyWithLimit(resp.Body, maxProductPageBodySize)
	if err != nil {
		return nil, err
	}

	document, err := newDocument(body)
	if err != nil {
		return nil, err
	}

	priceText := strings.TrimSpace(document.Find(priceSelector).Text())
	price, err := parsePrice(priceText)
	if err != nil {
		return nil, err
	}

	baseURL, err := url.Parse(productURL.String())
	if err != nil {
		return nil, err
	}

	product := &usecaseProduct.DuplicateProductData{
		Name:        strings.TrimSpace(document.Find("title").Text()),
		Description: strings.TrimSpace(document.Find(descriptionSelector).Text()),
		Price:       price,
		Tags:        extractTags(document),
		Images:      make([]usecaseProduct.DuplicateProductImage, 0),
	}

	var imageErr error
	document.Find(imageSelector).EachWithBreak(func(i int, selection *goquery.Selection) bool {
		src, ok := selection.Attr("src")
		if !ok {
			return true
		}

		imageURL, err := resolveImageURL(baseURL, src)
		if err != nil {
			imageErr = err
			return false
		}

		imageData, err := s.fetchImage(ctx, imageURL)
		if err != nil {
			imageErr = err
			return false
		}

		product.Images = append(product.Images, usecaseProduct.DuplicateProductImage{
			Name: buildImageName(imageURL),
			Data: imageData,
		})
		return true
	})
	if imageErr != nil {
		return nil, imageErr
	}

	return product, nil
}

func validateProductPageURL(rawURL string) (*url.URL, error) {
	trimmed := strings.TrimSpace(rawURL)
	parsedURL, err := primitive.NewURL(trimmed)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(parsedURL.Value())
	if err != nil {
		return nil, primitive.ErrInvalidURL
	}

	if err := validateProductPageRequestURL(u); err != nil {
		return nil, usecase.ErrInvalidInput
	}

	return u, nil
}

func validateProductPageRequestURL(u *url.URL) error {
	if u == nil {
		return primitive.ErrInvalidURL
	}
	if u.Scheme != "https" || strings.ToLower(u.Hostname()) != creemaProductPageHost {
		return usecase.ErrInvalidInput
	}
	return nil
}

func mapProductPageStatusError(statusCode int) error {
	switch statusCode {
	case http.StatusBadRequest, http.StatusUnprocessableEntity:
		return usecase.ErrInvalidInput
	case http.StatusNotFound, http.StatusGone:
		return usecase.ErrNotFound
	default:
		return fmt.Errorf("unexpected status code: %d", statusCode)
	}
}

func newDocument(body []byte) (*goquery.Document, error) {
	reader, err := newDecodedReader(body)
	if err != nil {
		return nil, err
	}
	return goquery.NewDocumentFromReader(reader)
}

func newDecodedReader(body []byte) (io.Reader, error) {
	detector := chardet.NewTextDetector()
	detected, err := detector.DetectBest(body)
	if err != nil {
		return bytes.NewReader(body), nil
	}

	reader, err := charset.NewReaderLabel(detected.Charset, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func parsePrice(price string) (int, error) {
	replacer := strings.NewReplacer("￥", "", "¥", "", ",", "", " ", "", "\u00a0", "")
	normalized := replacer.Replace(strings.TrimSpace(price))
	if normalized == "" {
		return 0, fmt.Errorf("price is empty")
	}

	value, err := strconv.Atoi(normalized)
	if err != nil {
		return 0, fmt.Errorf("parse price: %w", err)
	}
	return value, nil
}

func extractTags(document *goquery.Document) []string {
	tags := make([]string, 0)
	document.Find(tagSelector).Each(func(_ int, selection *goquery.Selection) {
		tagName := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(selection.Text()), "#"))
		if tagName == "" {
			return
		}
		tags = append(tags, tagName)
	})
	return tags
}

func resolveImageURL(baseURL *url.URL, rawURL string) (string, error) {
	parsedURL, err := url.Parse(strings.TrimSpace(rawURL))
	if err != nil {
		return "", err
	}
	resolvedURL := baseURL.ResolveReference(parsedURL)
	if err := validateImageRequestURL(resolvedURL); err != nil {
		return "", err
	}
	return resolvedURL.String(), nil
}

func (s *Scraper) fetchImage(ctx context.Context, imageURL string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, imageURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.doRequest(req, validateImageRequestURL)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected image status code: %d", resp.StatusCode)
	}

	return readResponseBodyWithLimit(resp.Body, maxProductImageBodySize)
}

func validateImageRequestURL(u *url.URL) error {
	if u == nil {
		return primitive.ErrInvalidURL
	}
	if u.Scheme != "https" {
		return usecase.ErrInvalidInput
	}
	if _, ok := allowedImageHosts[strings.ToLower(u.Hostname())]; !ok {
		return usecase.ErrInvalidInput
	}
	return nil
}

func (s *Scraper) doRequest(req *http.Request, validateURL func(*url.URL) error) (*http.Response, error) {
	if err := validateURL(req.URL); err != nil {
		return nil, err
	}

	client := *s.client
	client.CheckRedirect = func(req *http.Request, _ []*http.Request) error {
		return validateURL(req.URL)
	}

	resp, err := client.Do(req)
	if err != nil {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
		return nil, err
	}

	return resp, nil
}

func readResponseBodyWithLimit(body io.Reader, maxSize int64) ([]byte, error) {
	data, err := io.ReadAll(io.LimitReader(body, maxSize+1))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 || int64(len(data)) > maxSize {
		return nil, usecase.ErrInvalidInput
	}
	return data, nil
}

func buildImageName(imageURL string) string {
	parsedURL, err := url.Parse(imageURL)
	if err == nil {
		name := path.Base(parsedURL.Path)
		if name != "" && name != "." && name != "/" {
			return name
		}
	}
	return ""
}
