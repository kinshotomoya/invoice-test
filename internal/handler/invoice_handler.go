package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) InvoiceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// getの場合
		fmt.Println(r.URL.RawQuery)
	} else {
		// postの場合
		body := r.Body
		defer body.Close()
		fmt.Println(body)
	}
}
