package controllers

import (
	"api/app/models"
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

var typeToExtention = map[string]string{
	"image/jpeg": ".jpg",
	"image/png":  ".png",
	"image/gif":  ".gif",
}

// 商品一覧を取得
func getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetAllProducts()
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)
		return
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
	w.Header().Set("Content-Type", "image/jpeg")
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
		savedirectory := fmt.Sprintf("img/%s/%s", uuid[0:1], uuid[1:2])
		// 保存用のディレクトリがない場合は作成する
		if err := os.MkdirAll(savedirectory, 0777); err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)

			return
		}
		// fileのMIMETypeを取得
		mimeType := handler.Header["Content-Type"][0]
		savePath := savedirectory + "/" + uuid + typeToExtention[mimeType]
		f, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)

			return
		}
		defer f.Close()
		io.Copy(f, file)

		// ファイルの情報をsqlに保存する
		var productImage models.ProductImage
		// productImageのfieldを更新する
		productImage.Name = handler.Filename
		productImage.MimeType = mimeType
		productImage.Path = savePath
		productImage.ProductId = productId
		// sqlにデータを作成する
		err = models.InsertProductImage(&productImage)
		if err != nil {
			log.Println(err)
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusForbidden)

			return
		}
		i++
	}
}
