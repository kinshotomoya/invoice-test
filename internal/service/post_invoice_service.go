package service

import (
	"context"
	"invoice-test/internal/repository"
	serviceModel "invoice-test/internal/service/model"
)

type PostInvoiceService struct {
	MysqlRepository *repository.MysqlRepository
	CustomTime      serviceModel.CustomTimeInterface
}

func (s PostInvoiceService) PostInvoice(ctx context.Context, postInvoiceCondition *serviceModel.PostInvoiceCondition) (*serviceModel.Invoice, error) {
	user, err := s.MysqlRepository.FindUser(ctx.Value("email").(string), ctx.Value("password").(string))
	if err != nil {
		return nil, err
	}

	condition := postInvoiceCondition.NewPostInvoiceCondition(s.CustomTime)

	invoices, err := s.MysqlRepository.PostInvoice(user, condition)

	if err != nil {
		return nil, err
	}
	postedInvoice := serviceModel.NewInvoice(invoices)
	return &postedInvoice, nil
}
