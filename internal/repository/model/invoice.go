package model

import "time"

type Invoice struct {
	InvoiceId      uint64
	CompanyId      uint64
	SuppliersId    uint64
	IssueDate      time.Time
	PaymentAmount  float64
	Fee            float64
	FeeRate        float64
	Tax            float64
	TaxRate        float64
	TotalAmount    float64
	PaymentDueDate time.Time
	Status         string
}
