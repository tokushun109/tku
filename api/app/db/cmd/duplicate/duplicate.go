package main

import (
	"api/app/controllers"
	"api/app/controllers/utils"
	"api/app/models"
	"api/config"
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

func duplicateProduct(url string) {
	if !strings.Contains(url, "creema") {
		err := errors.New("invalid url")
		panic(err)
	}

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	buffer, _ := ioutil.ReadAll(res.Body)

	detector := chardet.NewTextDetector()
	detectResult, _ := detector.DetectBest(buffer)

	bufferReader := bytes.NewReader(buffer)
	reader, _ := charset.NewReaderLabel(detectResult.Charset, bufferReader)

	document, _ := goquery.NewDocumentFromReader(reader)

	// 作品名
	title := document.Find("title").Text()
	// 作品説明
	productDetail := document.Find("#introduction > div > div").Text()

	// 価格
	price := strings.Trim(strings.TrimSpace(document.Find("#js-item-detail > aside > div.p-item-detail-info.p-item-detail-info--side > div > div:nth-child(1) > div.p-item-detail-info__item.p-item-detail-info__item--price").Text()), "￥")
	intPrice, _ := strconv.Atoi(strings.Replace(price, ",", "", -1))

	// タグ
	tagsDom := document.Find("#js-item-detail > aside > div:nth-child(5) > ul > li > a")
	tagNames := []string{}
	tagsDom.Each(func(i int, s *goquery.Selection) {
		tagName := s.Text()
		if strings.Contains(tagName, "#") {
			trimTagName := strings.TrimSpace(strings.Trim(strings.TrimSpace(tagName), "#"))
			tagNames = append(tagNames, trimTagName)
			tag := models.Tag{
				Name: trimTagName,
			}
			if err = models.InsertUnDuplicateTag(&tag); err != nil {
				panic(err)
			}
		}
	})
	tags := models.GetTagsByNames(tagNames)
	siteDetails := []*models.SiteDetail{}
	siteDetails = append(siteDetails, &models.SiteDetail{
		DetailUrl: url,
		SalesSite: models.GetSalesSiteByName("creema"),
	})
	isActive := false
	IsRecommend := false
	product := models.Product{
		Name:        strings.TrimSpace(title),
		Description: strings.TrimSpace(productDetail),
		Price:       intPrice,
		IsActive:    &isActive,
		IsRecommend: &IsRecommend,
		CategoryId:  nil,
		TargetId:    nil,
		Tags:        tags,
		SiteDetails: siteDetails,
	}

	// 商品の登録
	err = models.InsertProduct(&product)
	if err != nil {
		panic(err)
	}

	// 画像情報を取得
	img := document.Find("#js-item-detail-centerpiece > img")
	img.Each(func(i int, s *goquery.Selection) {

		func(i int) {
			uuid, err := models.GenerateUuid()
			if err != nil {
				panic(err)
			}

			saveDirectory := fmt.Sprintf("img/%s/%s", uuid[0:1], uuid[1:2])
			// 保存用のディレクトリがない場合は作成する
			if err := os.MkdirAll(saveDirectory, 0777); err != nil {
				panic(err)
			}

			src, _ := s.Attr("src")
			response, err := http.Get(src)
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()
			fileName := fmt.Sprintf("%d.png", i)

			mimeType := "image/png"
			savePath := saveDirectory + "/" + uuid + controllers.TypeToExtension[mimeType]
			// localの場合はプロジェクト内のディレクトリに保存
			f, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			io.Copy(f, response.Body)

			if config.Config.Env != "local" {
				file, err := os.Open(savePath)
				if err != nil {
					panic(err)
				}
				defer file.Close()
				// 本番の場合はS3にアップロード
				if err := utils.UploadS3(&savePath, file); err != nil {
					panic(err)
				}
				os.Remove(savePath)
			}
			// ファイルの情報をsqlに保存する
			var productImage models.ProductImage
			// productImageのfieldを更新する
			productImage.Name = fileName
			productImage.MimeType = mimeType
			productImage.Path = savePath
			productImage.ProductId = product.ID
			productImage.Order = 100 - i
			// sqlにデータを作成する
			err = models.InsertProductImage(&productImage)
			if err != nil {
				panic(err)
			}
		}(i)
	})
	fmt.Println("正常に終了しました")
}

func StrStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput = scanner.Text()

	stringInput = strings.TrimSpace(stringInput)
	return
}

func main() {
	fmt.Println("販売サイトのurlを入力してください")
	url := StrStdin()
	duplicateProduct(url)
}
