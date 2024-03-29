package controllers

import (
	"api/app/controllers/utils"
	"api/app/models"
	"api/config"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocarina/gocsv"
	"github.com/gorilla/mux"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
	"gopkg.in/go-playground/validator.v9"
)

var TypeToExtension = map[string]string{
	"image/jpeg": ".jpg",
	"image/png":  ".png",
	"image/gif":  ".gif",
	"image/webp": ".webp",
}

type OrderParams struct {
	IsChanged bool
	Order     map[int]int
}

// 商品に紐づく商品画像に画像取得用のapiをつける
func setProductImageApiPath(product *models.Product) error {
	if config.Config.Env == "local" {
		// localの場合はプロジェクト内のディレクトリから取得
		base := config.Config.ApiBaseUrl
		for _, productImage := range product.ProductImages {
			productImage.ApiPath = base + "/product_image/" + productImage.Uuid + "/blob"
		}
	} else {
		// 本番の場合はS3から取得
		var err error
		for _, productImage := range product.ProductImages {
			productImage.ApiPath, err = utils.GetS3Content(&productImage.Path)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// 商品一覧を取得
func getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	mode := r.URL.Query().Get("mode")
	category := r.URL.Query().Get("category")
	target := r.URL.Query().Get("target")

	products, err := models.GetAllProducts(mode, category, target)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	for _, product := range products {
		if err := setProductImageApiPath(&product); err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
}

// 商品詳細を取得
func getProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["product_uuid"]
	product, err := models.GetProduct(uuid)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusNotFound)
		return
	}
	if err := setProductImageApiPath(&product); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(product); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
}

// 商品の新規作成
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	_, err := sessionCheck(r)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	var product models.Product
	if err := json.Unmarshal(reqBody, &product); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	// validationの確認
	validate := validator.New()
	if err := validate.Struct(product); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
		return
	}

	// modelの呼び出し
	err = models.InsertProduct(&product)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	responseBody, err := json.Marshal(product)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func duplicateProductHandler(w http.ResponseWriter, r *http.Request) {
	_, err := sessionCheck(r)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	var params struct {
		Url string
	}

	if err := json.Unmarshal(reqBody, &params); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	url := params.Url
	DuplicateProduct(url)

	// responseBodyで処理の成功を返す
	responseBody := getSuccessResponse()
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)

}

// 商品の更新
func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	_, err := sessionCheck(r)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	vars := mux.Vars(r)
	uuid := vars["product_uuid"]

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	var product models.Product
	if err := json.Unmarshal(reqBody, &product); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	// validationの確認
	validate := validator.New()
	if err := validate.Struct(product); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
		return
	}

	if err = models.UpdateProduct(&product, uuid); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	// responseBodyで処理の成功を返す
	responseBody := getSuccessResponse()
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

// 商品の削除
func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	_, err := sessionCheck(r)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	vars := mux.Vars(r)
	uuid := vars["product_uuid"]

	product, err := models.GetProduct(uuid)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	if err := product.DeleteProduct(); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	// responseBodyで処理の成功を返す
	responseBody := getSuccessResponse()
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func getCarouselImageHandler(w http.ResponseWriter, r *http.Request) {
	products := models.GetRecommendProducts()
	limit := 5
	if len(products) < limit {
		products = models.GetNewProducts(limit)
	}
	for _, product := range products {
		if err := setProductImageApiPath(&product); err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
			return
		}
	}
	type NewProductImage struct {
		Product         models.Product `json:"product"`
		NewImageApiPath string         `json:"apiPath"`
	}
	var newProductImages []NewProductImage
	for _, product := range products {
		if len(product.ProductImages) > 0 {
			newProductImages = append(newProductImages, NewProductImage{Product: product, NewImageApiPath: product.ProductImages[0].ApiPath})
		}
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newProductImages); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
}

// 商品画像のパスからバイナリデータを返す
func getProductImageBlobHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productImageUuid := vars["product_image_uuid"]
	productImage, err := models.GetProductImage(productImageUuid)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	file, err := os.Open(productImage.Path)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	defer file.Close()

	binary, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Type", productImage.MimeType)
	w.Write(binary)
}

// 商品画像の新規作成
func createProductImageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := sessionCheck(r)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	vars := mux.Vars(r)
	uuid := vars["product_uuid"]
	// requestのuuidから商品のIDを取得しておく
	product, err := models.GetProduct(uuid)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	productId := product.ID

	orderJson := r.FormValue("order")
	orderMap := OrderParams{}
	if err := json.Unmarshal([]byte(orderJson), &orderMap); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	i := 0
	for {
		// ファイルをapiディレクトリ内に保存する
		s := strconv.Itoa(i)
		file, handler, err := r.FormFile("file" + s)
		if file == nil {
			break
		}
		if err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
			return
		}
		defer file.Close()

		uuid, err := models.GenerateUuid()
		if err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
			return
		}
		saveDirectory := fmt.Sprintf("img/%s/%s", uuid[0:1], uuid[1:2])
		// 保存用のディレクトリがない場合は作成する
		if err := os.MkdirAll(saveDirectory, 0777); err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
			return
		}
		// fileのMIMETypeを取得
		mimeType := handler.Header["Content-Type"][0]
		savePath := saveDirectory + "/" + uuid + TypeToExtension[mimeType]
		if config.Config.Env == "local" {
			// localの場合はプロジェクト内のディレクトリに保存
			f, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Println(err)
				http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
				return
			}
			defer f.Close()
			io.Copy(f, file)
		} else {
			// 本番の場合はS3にアップロード
			if err := utils.UploadS3(&savePath, file); err != nil {
				log.Println(err)
				http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
				return
			}
		}

		// ファイルの情報をsqlに保存する
		var productImage models.ProductImage
		// productImageのfieldを更新する
		productImage.Name = handler.Filename
		productImage.MimeType = mimeType
		productImage.Path = savePath
		productImage.ProductId = productId
		productImage.Order = 0
		if orderMap.IsChanged {
			productImage.Order = orderMap.Order[i]
		}

		// sqlにデータを作成する
		err = models.InsertProductImage(&productImage)
		if err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
			return
		}
		i++
	}

	// responseBodyで処理の成功を返す
	responseBody := getSuccessResponse()
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

// 商品画像の削除
func deleteProductImageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := sessionCheck(r)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	vars := mux.Vars(r)
	productUuid := vars["product_uuid"]
	productImageUuid := vars["product_image_uuid"]
	if _, err := models.GetProduct(productUuid); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	productImage, err := models.GetProductImage(productImageUuid)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	if err := productImage.DeleteProductImage(); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	// responseBodyで処理の成功を返す
	responseBody := getSuccessResponse()
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

// 商品一覧をCSVで出力する
func getProductsCsvHandler(w http.ResponseWriter, r *http.Request) {
	_, err := sessionCheck(r)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	products, err := models.GetAllProducts("all", "all", "all")
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	csv := []*models.ProductCsv{}

	for _, product := range products {
		c := product.ProductToProductCsv()
		csv = append(csv, &c)
	}
	w.Header().Set("Content-Type", "text/csv; charset=utf8")
	gocsv.Marshal(csv, w)
}

// アップロードしたCSVを元に商品レコードを更新する
func uploadProductsCsvHandler(w http.ResponseWriter, r *http.Request) {
	_, err := sessionCheck(r)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	csv, handler, err := r.FormFile("csv")
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	defer csv.Close()

	mimeType := handler.Header["Content-Type"][0]
	if mimeType != "text/csv" {
		err = errors.New("mime-type is invalid")
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	buffer, _ := ioutil.ReadAll(csv)

	productCsv := []*models.ProductCsv{}
	if err := gocsv.UnmarshalBytes(buffer, &productCsv); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
	checkCategoryNames := map[string]struct{}{}
	checkTargetNames := map[string]struct{}{}
	validate := validator.New()

	// TODO エラーハンドリングを見直し
	for _, pc := range productCsv {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			if pc.CategoryName == "" {
				return
			}

			if _, ok := checkCategoryNames[pc.CategoryName]; ok {
				return
			}

			// 新規カテゴリーの作成
			category := models.Category{
				Name: pc.CategoryName,
			}
			if errors := validate.Struct(category); errors != nil {
				log.Println(err)
				http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
				return
			}
			if err = models.InsertUnDuplicateCategory(&category); err != nil {
				log.Println(err)
				http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
				return
			}
			checkCategoryNames[pc.CategoryName] = struct{}{}
		}()

		go func() {
			defer wg.Done()
			if pc.TargetName == "" {
				return
			}

			if _, ok := checkTargetNames[pc.TargetName]; ok {
				return
			}

			// 新規ターゲットの作成
			target := models.Target{
				Name: pc.TargetName,
			}
			if err := validate.Struct(target); err != nil {
				log.Println(err)
				http.Error(w, fmt.Sprintf("error: %s", err), http.StatusBadRequest)
				return
			}
			if err = models.InsertUnDuplicateTarget(&target); err != nil {
				log.Println(err)
				http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
				return
			}
			checkTargetNames[pc.TargetName] = struct{}{}

		}()
		wg.Wait()

		if err := pc.UpdateProductByCsv(); err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
			return
		}
	}

	responseBody := getSuccessResponse()
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

// カテゴリー別商品一覧を取得
func getAllCategoryProductsHandler(w http.ResponseWriter, r *http.Request) {
	mode := r.URL.Query().Get("mode")
	categoryUuid := r.URL.Query().Get("category")
	targetUuid := r.URL.Query().Get("target")

	if categoryUuid == "" || targetUuid == "" {
		err := errors.New("invalid params")
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}

	categoryProductsList := models.CategoryProductsList{}
	if categoryUuid == "all" {
		categories := models.GetUsedCategories()
		for _, category := range categories {
			products, err := models.GetAllProducts(mode, category.Uuid, targetUuid)
			if err != nil {
				log.Println(err)
				http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
				return
			}
			for _, product := range products {
				if err := setProductImageApiPath(&product); err != nil {
					log.Println(err)
					http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
					return
				}
			}
			categoryProduct := models.CategoryProducts{
				Category: category,
				Products: products,
			}
			categoryProductsList = append(categoryProductsList, categoryProduct)
		}
	} else {
		products, err := models.GetAllProducts(mode, categoryUuid, targetUuid)
		if err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
			return
		}
		for _, product := range products {
			if err := setProductImageApiPath(&product); err != nil {
				log.Println(err)
				http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
				return
			}
		}
		category := models.GetCategory(categoryUuid)
		categoryProduct := models.CategoryProducts{
			Category: category,
			Products: products,
		}
		categoryProductsList = append(categoryProductsList, categoryProduct)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(categoryProductsList); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
	}
}

func DuplicateProduct(url string) {
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
	tagsDom := document.Find("#js-item-detail > aside > div:nth-child(5) > ul:nth-child(3) > li > a")
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
			savePath := saveDirectory + "/" + uuid + TypeToExtension[mimeType]
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
