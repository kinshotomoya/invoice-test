package presentaion

import (
	"context"
	"errors"
	"net/http"
)

func GetAuth(r *http.Request, ctx context.Context) (context.Context, error) {
	userName, password, ok := r.BasicAuth()
	if !ok {
		return nil, errors.New("authorization required")
	}
	ctx = context.WithValue(ctx, "email", userName)
	ctx = context.WithValue(ctx, "password", password)
	return ctx, nil
}
