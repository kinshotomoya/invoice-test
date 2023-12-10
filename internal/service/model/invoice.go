package model

import (
	"invoice-test/internal/repository/model"
	"time"
)

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

func NewInvoice(invoice model.Invoice) Invoice {
	return Invoice{
		InvoiceId:      invoice.InvoiceId,
		CompanyId:      invoice.CompanyId,
		SuppliersId:    invoice.SuppliersId,
		IssueDate:      invoice.IssueDate,
		PaymentAmount:  invoice.PaymentAmount,
		Fee:            invoice.Fee,
		FeeRate:        invoice.FeeRate,
		Tax:            invoice.Tax,
		TaxRate:        invoice.TaxRate,
		TotalAmount:    invoice.TotalAmount,
		PaymentDueDate: invoice.PaymentDueDate,
		Status:         invoice.Status,
	}
}
