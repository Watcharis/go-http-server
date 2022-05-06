package handler

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var JwtKey = []byte("allexo-project")

var (
	ErrorMethodNotfounds           = errors.New("methods not found.")
	ErrorInvalidJsonBody           = errors.New("request body is invalid.")
	ErrorInvalidUsernameOrPassword = errors.New("authen fail invalid username or password")
	ErrorTokenNotFound             = errors.New("not found token.")
	ErrorInvalidFormatToken        = errors.New("invalid format token.")
	ErrorExtractTokenSignatureFail = errors.New("tkn valid false")
	ErrorOldTokenIsNotExipred      = errors.New("token is not within 30 seconds of expiry")
	ErrorFileNotExists             = errors.New("file is not exists.")
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type ResponseApiNasa struct {
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	Hdurl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	Url            string `json:"url"`
}

type ResponseErrorRateLimit struct {
	Error ErrorRateLimit `json:"error"`
}
type ErrorRateLimit struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
