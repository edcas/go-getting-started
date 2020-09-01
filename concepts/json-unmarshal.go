package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	ProductID      int    `json:"productiId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"qualityOnHand"`
	ProductName    string `json:"productName"`
}

func main() {

	productJSON := `{
		"ProductID":      123,
		"Manufacturer":   "Big Box Company",
		"PricePerUnit":   "12.99",
		"Sku":            "4561qHJK",
		"Upc":            "414654444566",
		"QuantityOnHand": 28,
		"ProductName":    "Gizmo"
	}`

	product := Product{}
	err := json.Unmarshal([]byte(productJSON), &product)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("json unmashal product: %s", product.ProductName)
}
