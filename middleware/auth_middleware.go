package middleware

import (
	"net/http"

	"github.com/NurFirdausR/golang-restfull-api/helper"
	"github.com/NurFirdausR/golang-restfull-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}
func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "RAHASIA" == r.Header.Get("X-API-KEY") {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Add("Content-Type", "Aplication/Json")
		w.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(w, webResponse)
	}
}
