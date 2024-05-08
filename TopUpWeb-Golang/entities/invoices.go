package entities

import "time"

type Invoices struct {
	InvoiceID   int
	DataID      int
	ProductID   int
	Payment     string
	PhoneNumber int64
	Status      StatusType
	Time        time.Time
}
