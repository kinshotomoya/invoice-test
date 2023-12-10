package handler

import (
	"fmt"
	"invoice-test/internal/presentaion"
	"invoice-test/internal/presentaion/model"
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
		response, err := model.ConvertToResponse(invoices)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(200)
		w.Write(response)

	} else {
		// postの場合
		// TODO: 請求書データの作成
		body := r.Body
		defer body.Close()
		fmt.Println(body)
	}
}
