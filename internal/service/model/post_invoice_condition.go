package model

import (
	"invoice-test/internal/repository/model"
)

const FeeRate = 0.04
const TaxRate = 0.1

type PostInvoiceCondition struct {
	SuppliersId    *uint64
	PaymentAmount  *float64
	PaymentDueDate *string
}

func (p *PostInvoiceCondition) NewPostInvoiceCondition(customTime CustomTimeInterface) *model.PostInvoiceCondition {
	now := customTime.NowDateOnly()
	fee := *p.PaymentAmount * FeeRate
	tax := fee * TaxRate
	totalAmount := *p.PaymentAmount + fee + tax

	return &model.PostInvoiceCondition{
		SuppliersId:    *p.SuppliersId,
		IssueDate:      now,
		PaymentAmount:  *p.PaymentAmount,
		Fee:            fee,
		FeeRate:        FeeRate,
		Tax:            tax,
		TaxRate:        TaxRate,
		TotalAmount:    totalAmount,
		PaymentDueDate: *p.PaymentDueDate,
		Status:         "PENDING",
	}

}
