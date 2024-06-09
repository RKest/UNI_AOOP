package internal

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var secret = []byte("password123")

type Auth struct {
	privateKey *rsa.PrivateKey
}

func (a *Auth) Login(res http.ResponseWriter, req *http.Request) {
	err := req.ParseMultipartForm(32 << 20) // 32 MB of max memory
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	user := req.FormValue("login")
	pass := req.FormValue("password")
	if user == "" || pass != string(secret) {
		http.Error(res, "login or password is wrong",
			http.StatusBadRequest)
		return
	}
	token, err := a.generateJWT(user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Set-Cookie", fmt.Sprint("Token=", token))
	if _, err = res.Write([]byte("Authorized")); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func (a *Auth) generateJWT(username string) (string, error) {
	var err error
	token := jwt.New(jwt.SigningMethodRS256)
	a.privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", err
	}
	if _, err = rsa.EncryptOAEP(sha256.New(), rand.Reader,
		&a.privateKey.PublicKey, secret, nil); err != nil {
		return "", err
	}
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["authorized"] = true
	claims["user"] = username
	if tokenStr, err := token.SignedString(a.privateKey); err != nil {
		return "", err
	} else {
		return tokenStr, nil
	}
}

func (a *Auth) HandleFuncAuth(mux *http.ServeMux, pattern string, next http.HandlerFunc) {
	mux.HandleFunc(pattern, a.verifyJWT(next))
}

func (a *Auth) verifyJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		tokenStr, err := req.Cookie("Token")
		if err != nil {
			http.Error(res, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(tokenStr.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok ||
				a.privateKey == nil {
				return nil, errors.New("invalid token")
			}
			return &a.privateKey.PublicKey, nil
		})
		if err != nil {
			http.Error(res, err.Error(), http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			http.Error(res, "token is invalid", http.StatusUnauthorized)
			return
		}
		next(res, req)
	}
}
