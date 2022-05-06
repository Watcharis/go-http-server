package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func (h *handleService) ExtractToken(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			HttpResponse(w, r, false, nil, ErrorMethodNotfounds, http.StatusBadRequest)
			return
		}

		authorization := r.Header.Get("Authorization")
		if validateToken(authorization) {
			HttpResponse(w, r, false, nil, ErrorTokenNotFound, http.StatusUnauthorized)
			return
		}

		token := strings.Split(authorization, " ")
		rawToken := token[len(token)-1]
		if validateToken(rawToken) {
			HttpResponse(w, r, false, nil, ErrorInvalidFormatToken, http.StatusUnauthorized)
			return
		}

		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(rawToken, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				HttpResponse(w, r, false, nil, err, http.StatusUnauthorized)
				return
			}
			HttpResponse(w, r, false, nil, err, http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			HttpResponse(w, r, false, nil, ErrorExtractTokenSignatureFail, http.StatusUnauthorized)
			return
		}

		HttpResponse(w, r, true, claims, nil, http.StatusOK)
		return
	}
}

func validateToken(authorization string) bool {
	return authorization == ""
}
