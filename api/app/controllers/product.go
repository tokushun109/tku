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
)

var typeToExtention = map[string]string{
	"image/jpeg": ".jpg",
	"image/png":  ".png",
	"image/gif":  ".gif",
}

// 商品一覧を取得
func getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		log.Fatalln(err)
	}
}

// 商品詳細を取得
func getProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	product := models.GetProduct(uuid)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(product); err != nil {
		log.Fatalln(err)
	}
}

// 商品の新規作成
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var product models.Product
	if err := json.Unmarshal(reqBody, &product); err != nil {
		log.Fatal(err)
	}
	// modelの呼び出し
	models.InsertProduct(&product)
	responseBody, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

// 商品画像の新規作成
func createProductImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	// requestのuuidから商品のIDを取得しておく
	productId := models.GetProduct(uuid).ID
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
		}
		defer file.Close()

		uuid, err := models.GenerateUuid()
		if err != nil {
			log.Fatal(err)
		}
		savedirectory := fmt.Sprintf("img/%s/%s", uuid[0:1], uuid[1:2])
		// 保存用のディレクトリがない場合は作成する
		if err := os.MkdirAll(savedirectory, 0777); err != nil {
			log.Println(err)
		}
		// fileのMIMETypeを取得
		mimeType := handler.Header["Content-Type"][0]
		savePath := savedirectory + "/" + uuid + typeToExtention[mimeType]
		f, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println(err)
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
		models.InsertProductImage(&productImage)
		i++
	}
}
