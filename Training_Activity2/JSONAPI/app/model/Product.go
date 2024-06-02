package model

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Seller struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		CountryCode string `json:"countryCode"`
	} `json:"seller"`
	Price int `json:"price"`
}

var Products []Product

func init() {
	jsonInput := `
    [
        {
            "id": 1,
            "name": "Book1",
            "seller": {
                "id": 1,
                "name": "ABC company",
                "countryCode": "US"
            },
            "price": 100
        },
        {
            "id": 2,
            "name": "Book2",
            "seller": {
                "id": 2,
                "name": "ABD company",
                "countryCode": "Japan"
            },
            "price": 200
        }
    ]
    `

	err := json.Unmarshal([]byte(jsonInput), &Products)
	if err != nil {
		fmt.Println("Error initializing products:", err)
		return
	}
	fmt.Println("Products initialized:", Products)
}
