package Account

import (
	"fmt"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func AuthMiddlewareIntern(next http.Handler) http.Handler {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("intern"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	fmt.Println(&next)
	return jwtMiddleware.Handler(next)
}

func AuthMiddlewareLecturer(next http.Handler) http.Handler {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("lecturer"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	fmt.Println(&next)
	return jwtMiddleware.Handler(next)
}
