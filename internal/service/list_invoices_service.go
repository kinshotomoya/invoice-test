package service

import (
	"context"
	"invoice-test/internal/repository"
	"invoice-test/internal/repository/model"
	serviceModel "invoice-test/internal/service/model"
)

type ListInvoiceService struct {
	MysqlRepository *repository.MysqlRepository
}

func (l ListInvoiceService) ListInvoices(ctx context.Context, condition *serviceModel.ListInvoiceCondition) ([]serviceModel.Invoice, error) {
	user, err := l.MysqlRepository.FindUser(ctx.Value("email").(string), ctx.Value("password").(string))
	if err != nil {
		return nil, err
	}

	repositoryCondition := &model.ListInvoiceCondition{
		From: condition.From,
		To:   condition.To,
	}

	invoices, err := l.MysqlRepository.ListInvoices(user, repositoryCondition)
	if err != nil {
		return nil, err
	}

	resInvoices := make([]serviceModel.Invoice, len(invoices))
	for i := range invoices {
		resInvoices[i] = serviceModel.NewInvoice(invoices[i])
	}
	return resInvoices, nil
}
