package controllers

import (
	"api/app/controllers/utils"
	"api/app/models"
	"api/config"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
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
	products, err := models.GetAllProducts(mode)
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
	if errors := validate.Struct(product); errors != nil {
		log.Println(errors)
		http.Error(w, fmt.Sprintf("error: %s", errors), http.StatusBadRequest)
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

// 商品の更新
func updateProductHandler(w http.ResponseWriter, r *http.Request) {
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
	if errors := validate.Struct(product); errors != nil {
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
	products := models.GetNewProducts(5)
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

// 商品の削除
func deleteProductImageHandler(w http.ResponseWriter, r *http.Request) {
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
