package models

import (
	"log"
	"topup-game/config"
	"topup-game/entities"
)

func GetProductMlbb() ([]entities.InvoiceWithProduct, error) {
	rows, err := config.DB.Query("SELECT product_id, product_name, value, harga, stock FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entities.InvoiceWithProduct
	for rows.Next() {
		var product entities.InvoiceWithProduct
		if err := rows.Scan(&product.ProductID, &product.ProductName, &product.Value, &product.Harga, &product.Stock); err != nil {
			return nil, err
		}

		// Add product to the products slice
		products = append(products, product)
	}

	return products, nil
}

func CreateInvoice(invoice *entities.InvoiceWithProduct) bool {
	result, err := config.DB.Exec(`
		INSERT INTO invoices(
			data_id, product_id, payment, phone_number, status, time)
			VALUES (?,?,?,?,?,?)`,
		invoice.DataID,
		invoice.ProductID,
		invoice.Payment,
		invoice.PhoneNumber,
		invoice.Status,
		invoice.Time,
	)
	if err != nil {
		log.Println("Error: Unable to create InvoiceWithProduct:", err)
		return false
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Println("Error: Unable to get Last Insert ID:", err)
		return false
	}

	log.Println("Debug: Created InvoiceWithProduct with ID:", lastInsertID)
	return lastInsertID > 0
}
