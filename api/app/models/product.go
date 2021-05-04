package models

import (
	"log"
	"time"
)

type Product struct {
	Id                int                `json:"-"`
	Uuid              string             `json:"uuid,omitempty"`
	Name              string             `json:"name,omitempty"`
	Description       string             `json:"description,omitempty"`
	AccessoryCategory *AccessoryCategory `json:"accessory_category,omitempty"`
	MaterialCategory  *MaterialCategory  `json:"material_category,omitempty"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
}

type Products []Product

type ProductImage struct {
	Id        int       `json:"-"`
	Uuid      string    `json:"uuid,omitempty"`
	Name      string    `json:"name,omitempty"`
	MimeType  string    `json:"-"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetProducts() (products Products, err error) {
	cmd := `SELECT id, uuid, name, description, created_at, updated_at
			FROM product`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		product := Product{}
		err = rows.Scan(
			&product.Id,
			&product.Uuid,
			&product.Name,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		products = append(products, product)
	}
	rows.Close()
	return products, err
}
