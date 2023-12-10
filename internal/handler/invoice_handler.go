package handler

import (
	"fmt"
	"invoice-test/internal/presentaion"
	"net/http"
)

func (h *Handler) InvoiceHandler(w http.ResponseWriter, r *http.Request) {
	// cookieからSESSION_ID取得
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
		h.ListInvoiceService.ListInvoices(ctx, condition)
		fmt.Println(r.URL.RawQuery)
	} else {
		// postの場合
		body := r.Body
		defer body.Close()
		fmt.Println(body)
	}
}
