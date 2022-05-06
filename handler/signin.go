package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// referens :: https://www.sohamkamani.com/golang/jwt-authentication/

type HandleService interface {
	Signin(ctx context.Context) http.HandlerFunc
	ExtractToken(ctx context.Context) http.HandlerFunc
	Refresh(ctx context.Context) http.HandlerFunc
	SendImageFormDirectory(ctx context.Context) http.HandlerFunc
	GetUrlImages(ctx context.Context) http.Handler
}

type handleService struct{}

func NewHttp() HandleService {
	return &handleService{}
}

func (h *handleService) Signin(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			HttpResponse(w, r, false, nil, ErrorMethodNotfounds, http.StatusBadRequest)
			return
		}

		signinBindRequest := SigninRequest{}
		if err := json.NewDecoder(r.Body).Decode(&signinBindRequest); err != nil {
			HttpResponse(w, r, false, nil, ErrorInvalidJsonBody, http.StatusBadRequest)
			return
		}

		expectedPassword := VerifyPassword(signinBindRequest.Username, signinBindRequest.Password)
		if !expectedPassword {
			HttpResponse(w, r, false, nil, ErrorInvalidUsernameOrPassword, http.StatusUnauthorized)
			return
		}

		// Declare the expiration time of the token
		// here, we have kept it as 5 minutes
		expirationTime := time.Now().Add(5 * time.Minute)

		// Create the JWT claims, which includes the username and expiry time
		claims := &Claims{
			Username: signinBindRequest.Username,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(JwtKey)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			HttpResponse(w, r, false, nil, nil, http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"accessToken": tokenString,
			"expiredAt":   claims.ExpiresAt,
		}
		HttpResponse(w, r, true, data, nil, http.StatusOK)
		return
	}
}

func VerifyPassword(users string, pass string) bool {
	auth := map[string]string{
		"test":  "123456",
		"test1": "234567",
	}
	_, ok := auth[users]
	return ok
}

func HttpResponse(w http.ResponseWriter, r *http.Request, success bool, data interface{}, err error, statusCode int) {
	res := map[string]interface{}{}
	if err != nil {
		res["error"] = err.Error()
	}
	res["success"] = success
	res["data"] = data
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
	return
}

func (h *handleService) SendImageFormDirectory(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			HttpResponse(w, r, false, nil, ErrorMethodNotfounds, http.StatusBadRequest)
			return
		}
		fileName := r.URL.Query().Get("filename")
		dir, _ := os.Getwd()
		send := path.Join(dir, "images", fileName)

		if _, err := FileExists(send); err != nil {
			HttpResponse(w, r, false, nil, err, http.StatusBadRequest)
			return
		}
		http.ServeFile(w, r, send)
		return
	}
}

func (h *handleService) GetUrlImages(ctx context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			HttpResponse(w, r, false, nil, ErrorMethodNotfounds, http.StatusBadRequest)
			return
		}

		fileName := r.URL.Query().Get("filename")

		dir, err := os.Getwd()
		if err != nil {
			HttpResponse(w, r, false, nil, err, http.StatusBadRequest)
			return
		}

		send := path.Join(dir, "images", fileName)
		file, err := FileExists(send)
		if err != nil {
			HttpResponse(w, r, false, nil, err, http.StatusBadRequest)
			return
		}

		data := map[string]interface{}{
			"filename": file.Name(),
			"filesize": file.Size(),
			"url":      fmt.Sprintf("http://localhost:8080/images?filename=%s", fileName),
		}
		HttpResponse(w, r, true, data, nil, http.StatusOK)
		return
	})
}

func FileExists(path string) (fs.FileInfo, error) {
	file, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = ErrorFileNotExists
			return nil, err
		}
	}
	return file, nil
}
