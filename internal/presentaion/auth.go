package presentaion

import "net/http"

func GetSessionId(r *http.Request) (*string, error) {
	cookie, err := r.Cookie("SESSION_ID")
	if err != nil || cookie.Valid() != nil {
		return nil, err
	}

	return &cookie.Value, nil
}
