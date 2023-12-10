package service

import "invoice-test/internal/repository"

type ListInvoiceService struct {
	MysqlRepository *repository.MysqlRepository
}

func (l ListInvoiceService) ListInvoices(sessionId string, condition *ListInvoiceCondition) {
	// user取得
	_, err := l.MysqlRepository.FindUser(sessionId)
	if err != nil {
		return
	}

	// TODO: 請求書データの取得
}
