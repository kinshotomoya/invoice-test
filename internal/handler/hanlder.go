package handler

import "invoice-test/internal/service"

type Handler struct {
	ListInvoiceService *service.ListInvoiceService
	PostInvoiceService *service.PostInvoiceService
}
