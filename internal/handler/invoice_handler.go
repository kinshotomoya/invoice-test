package handler

import (
	"fmt"
	"invoice-test/internal/presentaion"
	"net/http"
)

func (h *Handler) InvoiceHandler(w http.ResponseWriter, r *http.Request) {
	// cookieからSESSION_ID取得
	sessionId, err := presentaion.GetSessionId(r)
	if err != nil || sessionId == nil {
		http.Error(w, err.Error(), 403)
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
		h.ListInvoiceService.ListInvoices(*sessionId, condition)
		fmt.Println(r.URL.RawQuery)
	} else {
		// postの場合
		body := r.Body
		defer body.Close()
		fmt.Println(body)
	}
}
