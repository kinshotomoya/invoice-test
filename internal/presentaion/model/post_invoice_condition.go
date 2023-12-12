package model

import (
	"encoding/json"
	"errors"
	"net/http"
)

type PostInvoiceCondition struct {
	SuppliersId    *uint64  `json:"suppliers_id"`
	PaymentAmount  *float64 `json:"payment_amount"`
	PaymentDueDate *string  `json:"payment_due_date"`
}

func NewPostInvoiceCondition(r *http.Request) (*PostInvoiceCondition, error) {
	body := r.Body
	defer body.Close()
	var postInvoiceCondition PostInvoiceCondition
	err := json.NewDecoder(body).Decode(&postInvoiceCondition)
	if err != nil {
		return nil, err
	}

	if postInvoiceCondition.SuppliersId == nil || postInvoiceCondition.PaymentAmount == nil || postInvoiceCondition.PaymentDueDate == nil {
		return nil, errors.New("")
	}

	return &postInvoiceCondition, nil

}
