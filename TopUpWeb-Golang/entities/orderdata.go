package entities

import "time"

type InvoiceWithProduct struct {
	InvoiceID   int
	DataID      int
	ProductID   int
	Payment     string
	Time        time.Time
	PhoneNumber int64
	Status      StatusType
	ProductName string
	Value       string
	Harga       int64
	Stock       int
}

type StatusType string

const (
	Sukses     StatusType = "Sukses"
	Proses     StatusType = "Proses"
	BelumBayar StatusType = "Belum di Bayar"
	Gagal      StatusType = "Gagal"
)
