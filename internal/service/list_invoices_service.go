package service

import (
	"context"
	"invoice-test/internal/repository"
)

type ListInvoiceService struct {
	MysqlRepository *repository.MysqlRepository
}

func (l ListInvoiceService) ListInvoices(ctx context.Context, condition *ListInvoiceCondition) {
	// user取得
	_, err := l.MysqlRepository.FindUser(ctx.Value("email").(string), ctx.Value("password").(string))
	if err != nil {
		return
	}

	// TODO: 請求書データの取得
}
