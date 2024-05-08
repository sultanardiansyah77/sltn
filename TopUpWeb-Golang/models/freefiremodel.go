package models

import (
	"topup-game/config"
	"topup-game/entities"
)

// GetProductFF mengembalikan semua produk dari database.
func GetProductFF() ([]entities.Product, error) {
	rows, err := config.DB.Query("SELECT product_id, product_name, value, harga, stock FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.ProductID, &product.ProductName, &product.Value, &product.Harga, &product.Stock); err != nil {
			return nil, err
		}

		// Check if the ProductName is "Free Fire" before appending to products
		if product.ProductName == "Free Fire" {
			products = append(products, product)
		}
	}

	return products, nil
}
