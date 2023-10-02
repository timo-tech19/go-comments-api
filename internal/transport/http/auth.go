package http

import (
	"errors"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

// Verifies if user is logged in.
// original handler function is not called is user is unauthorised.
func JWTAuth(
	original func(w http.ResponseWriter, r *http.Request),
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		if validateToken(authHeaderParts[1]) {
			original(w, r)
		} else {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}
	}
}

// verifies a jwt token
func validateToken(accessToken string) bool {
	var mySigningKey = []byte("jwtaccesstokenkey")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Could not validate author token")
		}

		return mySigningKey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}
