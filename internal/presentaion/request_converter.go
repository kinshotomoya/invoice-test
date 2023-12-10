package presentaion

import (
	"errors"
	"invoice-test/internal/service"
	"net/http"
)

func NewListInvoiceCondition(r *http.Request) (*service.ListInvoiceCondition, error) {
	params := r.URL.Query()
	var errs error
	from := params.Get("from")
	if from == "" {
		errs = errors.Join(errs, errors.New("from is required"))
	}
	to := params.Get("to")
	if to == "" {
		errs = errors.Join(errs, errors.New("to is required"))
	}
	if errs != nil {
		return nil, errs
	}
	condition := &service.ListInvoiceCondition{
		From: from,
		To:   to,
	}

	return condition, nil
}
