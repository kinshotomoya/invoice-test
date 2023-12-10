package model

import (
	"bytes"
	"encoding/json"
	serviceModel "invoice-test/internal/service/model"
	"time"
)

type Invoice struct {
	InvoiceId      uint64    `json:"invoice_id"`
	CompanyId      uint64    `json:"company_id"`
	SuppliersId    uint64    `json:"suppliers_id"`
	IssueDate      time.Time `json:"issue_date"`
	PaymentAmount  float64   `json:"payment_amount"`
	Fee            float64   `json:"fee"`
	FeeRate        float64   `json:"fee_rate"`
	Tax            float64   `json:"tax"`
	TaxRate        float64   `json:"tax_rate"`
	TotalAmount    float64   `json:"total_amount"`
	PaymentDueDate time.Time `json:"payment_due_date"`
	Status         string    `json:"status"`
}

func ConvertToResponse(invoices []serviceModel.Invoice) ([]byte, error) {

	var buf bytes.Buffer
	resInvoices := make([]Invoice, len(invoices))
	for i := range resInvoices {
		resInvoices[i] = Invoice{
			InvoiceId:      resInvoices[i].InvoiceId,
			CompanyId:      resInvoices[i].CompanyId,
			SuppliersId:    resInvoices[i].SuppliersId,
			IssueDate:      resInvoices[i].IssueDate,
			PaymentAmount:  resInvoices[i].PaymentAmount,
			Fee:            resInvoices[i].Fee,
			FeeRate:        resInvoices[i].FeeRate,
			Tax:            resInvoices[i].Tax,
			TaxRate:        resInvoices[i].TaxRate,
			TotalAmount:    resInvoices[i].TotalAmount,
			PaymentDueDate: resInvoices[i].PaymentDueDate,
			Status:         resInvoices[i].Status,
		}
	}
	err := json.NewEncoder(&buf).Encode(resInvoices)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}
