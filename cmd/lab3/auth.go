package main

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var secret = []byte("password123")

func Auth(res http.ResponseWriter, req *http.Request) {
	err := req.ParseMultipartForm(32 << 20) // 32 MB of max memory
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	user := req.FormValue("login")
	pass := req.FormValue("password")
	if user == "" || pass != string(secret) {
		http.Error(res, "login or password is wrong", http.StatusBadRequest)
		return
	}
	token, err := generateJWT(user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Set-Cookie", fmt.Sprint("Token=", token))
	if _, err = res.Write([]byte("Authorized")); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func generateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = username
	if tokenStr, err := token.SignedString(secret); err != nil {
		return "", err
	} else {
		return tokenStr, nil
	}
}

func HandleFuncAuth(mux *http.ServeMux, pattern string, next http.HandlerFunc) {
	mux.HandleFunc(pattern, verifyJWT(next))
}

func verifyJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Header["Token"] == nil {
			http.Error(res, "missing token", http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(req.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, errors.New("invalid token")
			}
			return "", nil
		})
		if err != nil || !token.Valid {
			http.Error(res, "invalid token", http.StatusUnauthorized)
			return
		}
		next(res, req)
	}
}
