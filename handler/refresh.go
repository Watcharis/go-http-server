package handler

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (h *handleService) Refresh(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
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
		tkns, err := jwt.ParseWithClaims(rawToken, claims, func(token *jwt.Token) (interface{}, error) {
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

		if !tkns.Valid {
			HttpResponse(w, r, false, nil, ErrorExtractTokenSignatureFail, http.StatusUnauthorized)
			return
		}

		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
			HttpResponse(w, r, false, nil, ErrorOldTokenIsNotExipred, http.StatusBadRequest)
			return
		}

		expirationTime := time.Now().Add(5 * time.Minute)
		claims.ExpiresAt = expirationTime.Unix()
		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		refreshToken, err := newToken.SignedString(JwtKey)
		if err != nil {
			HttpResponse(w, r, false, nil, nil, http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"accessToken":  rawToken,
			"refreshToken": refreshToken,
			"expiredAt":    claims.ExpiresAt,
		}
		HttpResponse(w, r, true, data, nil, http.StatusOK)
		return
	}
}
