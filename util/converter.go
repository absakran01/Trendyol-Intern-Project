package util

import (
	"encoding/json"
	"log"
	"strconv"

	"example.com/model"
)



// func toProduct(value []byte) model.Product {
// 	var envelope struct {
// 		Payload struct {
// 			After json.RawMessage `json:"after"`
// 		} `json:"payload"`
// 	}
// 	if err := json.Unmarshal(value, &envelope); err != nil {
// 		log.Println("Error unmarshalling envelope:", err)
// 		return model.Product{}
// 	}
// 	var product model.Product
// 	if len(envelope.Payload.After) == 0 || string(envelope.Payload.After) == "null" {
// 		log.Println("No 'after' data in payload")
// 		return model.Product{}
// 	}
// 	if err := json.Unmarshal(envelope.Payload.After, &product); err != nil {
// 		log.Println("Error unmarshalling product:", err)
// 		return model.Product{}
// 	}
// 	return product
// }

func ToTrendyolProduct(value []byte) model.TrendyolProduct {
	var envelope struct {
		Payload struct {
			After json.RawMessage `json:"after"`
		} `json:"payload"`
	}

	var tproduct model.TrendyolProduct

	if err := json.Unmarshal(value, &envelope); err != nil {
		log.Println("Error unmarshalling envelope:", err)
		return model.TrendyolProduct{}
	}
	var product model.Product
	if len(envelope.Payload.After) == 0 || string(envelope.Payload.After) == "null" {
		log.Println("No 'after' data in payload")
		return model.TrendyolProduct{}
	}
	if err := json.Unmarshal(envelope.Payload.After, &product); err != nil {
		log.Println("Error unmarshalling product:", err)
		return model.TrendyolProduct{}
	}

	tproduct = model.TrendyolProduct{
		ID:          product.Asin,
		Name:        product.Title,
		Description: product.Description,
		Price:       stringToFloat(product.FinalPrice),
		Currency:    product.Currency,
		Stock:       product.Availability,
		Category:    product.Categories,
		Brand:       product.Brand,
		Images:      stringToStringSlice(product.Images),
		Variants:    stringToStringSlice(product.Variations),
		Ratings:     product.Rating,
	}
	return tproduct
}




// data conversion functions

func stringToFloat(price string) float64 {
	if price == "" {
		return 0.0
	}
	value, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("Error converting price to float:", err)
		return 0.0
	}
	return value
}

func stringToStringSlice(txt string) []string {
	if txt == "" {
		return []string{}
	}
	var result []string
	err := json.Unmarshal([]byte(txt), &result)
	if err != nil {
		log.Println("Error converting string to string slice:", err)
		return []string{}
	}
	return result
}