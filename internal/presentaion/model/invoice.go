package model

import (
	"bytes"
	"encoding/json"
	serviceModel "invoice-test/internal/service/model"
)

type Invoice struct {
	InvoiceId      uint64  `json:"invoice_id"`
	CompanyId      uint64  `json:"company_id"`
	SuppliersId    uint64  `json:"suppliers_id"`
	IssueDate      string  `json:"issue_date"`
	PaymentAmount  float64 `json:"payment_amount"`
	Fee            float64 `json:"fee"`
	FeeRate        float64 `json:"fee_rate"`
	Tax            float64 `json:"tax"`
	TaxRate        float64 `json:"tax_rate"`
	TotalAmount    float64 `json:"total_amount"`
	PaymentDueDate string  `json:"payment_due_date"`
	Status         string  `json:"status"`
}

func ConvertToResponse(invoices []serviceModel.Invoice) ([]byte, error) {

	var buf bytes.Buffer
	resInvoices := make([]Invoice, len(invoices))
	for i := range resInvoices {
		resInvoices[i] = Invoice{
			InvoiceId:      invoices[i].InvoiceId,
			CompanyId:      invoices[i].CompanyId,
			SuppliersId:    invoices[i].SuppliersId,
			IssueDate:      invoices[i].IssueDate,
			PaymentAmount:  invoices[i].PaymentAmount,
			Fee:            invoices[i].Fee,
			FeeRate:        invoices[i].FeeRate,
			Tax:            invoices[i].Tax,
			TaxRate:        invoices[i].TaxRate,
			TotalAmount:    invoices[i].TotalAmount,
			PaymentDueDate: invoices[i].PaymentDueDate,
			Status:         invoices[i].Status,
		}
	}
	err := json.NewEncoder(&buf).Encode(resInvoices)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}
