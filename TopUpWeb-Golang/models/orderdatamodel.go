package models

import (
	"database/sql"
	"fmt"
	"log"
	"topup-game/config"
	"topup-game/entities"
)

func scanInvoiceWithProduct(row *sql.Row) (*entities.InvoiceWithProduct, error) {
	var invoiceWithProduct entities.InvoiceWithProduct
	err := row.Scan(
		&invoiceWithProduct.InvoiceID,
		&invoiceWithProduct.DataID,
		&invoiceWithProduct.ProductID,
		&invoiceWithProduct.Payment,
		&invoiceWithProduct.Time,
		&invoiceWithProduct.PhoneNumber,
		&invoiceWithProduct.Status,
		&invoiceWithProduct.ProductName,
		&invoiceWithProduct.Value,
		&invoiceWithProduct.Harga,
		&invoiceWithProduct.Stock,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("Error: Unable to scan row:", err)
		return nil, err
	}

	return &invoiceWithProduct, nil
}

func GetInvoiceWithProductByID(invoiceID int) (*entities.InvoiceWithProduct, error) {
	fmt.Println("Debug: InvoiceID received:", invoiceID)

	query := `
		SELECT i.invoice_id, i.data_id, i.product_id, i.payment, i.time, i.phone_number, i.status,
			   f.product_name, f.value AS product_value, f.harga AS product_harga, f.stock AS product_stock
		FROM invoices i
		INNER JOIN product f ON i.product_id = f.product_id
		WHERE i.invoice_id = ?`

	row := config.DB.QueryRow(query, invoiceID)
	invoiceWithProduct, err := scanInvoiceWithProduct(row)

	if err != nil {
		log.Println("Error: Unable to get InvoiceWithProduct by ID:", err)
	} else {
		log.Println("Debug: InvoiceWithProduct retrieved from the database:", invoiceWithProduct)
	}

	return invoiceWithProduct, err
}

func GetLatestInvoiceWithProduct() (*entities.InvoiceWithProduct, error) {
	query := `
        SELECT i.invoice_id, i.data_id, i.product_id, i.payment, i.time, i.phone_number, i.status,
               f.product_name, f.value AS product_value, f.harga AS product_harga, f.stock AS product_stock
        FROM invoices i
        INNER JOIN product f ON i.product_id = f.product_id
        ORDER BY i.time DESC
        LIMIT 1`

	row := config.DB.QueryRow(query)
	invoiceWithProduct, err := scanInvoiceWithProduct(row)

	if err != nil {
		log.Println("Error: Unable to get Latest InvoiceWithProduct:", err)
	} else {
		log.Println("Debug: Latest InvoiceWithProduct retrieved from the database:", invoiceWithProduct)
	}

	return invoiceWithProduct, err
}

func UpdateInvoiceStatus(invoiceID int, status string) bool {
	result, err := config.DB.Exec(`
		UPDATE invoices SET
			status = ?
		WHERE invoice_id = ?
	`,
		status,
		invoiceID,
	)
	if err != nil {
		log.Println("Error: Unable to update invoice status:", err)
		return false
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error: Unable to get RowsAffected:", err)
		return false
	}

	return rowsAffected > 0
}
