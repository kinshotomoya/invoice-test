package handler

import (
	"invoice-test/internal/presentaion"
	"invoice-test/internal/presentaion/model"
	serviceModel "invoice-test/internal/service/model"
	"net/http"
)

func (h *Handler) InvoiceHandler(w http.ResponseWriter, r *http.Request) {
	ctx, err := presentaion.GetAuth(r, r.Context())
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	if r.Method == http.MethodGet {
		// getの場合
		// url parameterの取得
		condition, err := presentaion.NewListInvoiceCondition(r)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		invoices, err := h.ListInvoiceService.ListInvoices(ctx, condition)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		response, err := model.ConvertToListResponse(invoices)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(200)
		w.Write(response)

	} else {
		// postの場合
		condition, err := model.NewPostInvoiceCondition(r)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		serviceCondition := &serviceModel.PostInvoiceCondition{
			SuppliersId:    condition.SuppliersId,
			PaymentAmount:  condition.PaymentAmount,
			PaymentDueDate: condition.PaymentDueDate,
		}

		invoice, err := h.PostInvoiceService.PostInvoice(ctx, serviceCondition)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		response, err := model.ConvertToResponse(invoice)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(200)
		w.Write(response)

	}
}
