package lib

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/rs/cors"
)

// CorsMiddleware returns a configured middleware to handle cross origin resource sharing
func CorsMiddleware() *cors.Cors {
	return cors.New(cors.Options{
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
		AllowedMethods: []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
	})
}

// JwtMiddleware returns a configured json web token middleware
func JwtMiddleware() negroni.HandlerFunc {
	return negroni.HandlerFunc(jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(AppKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	}).HandlerWithNext)
}
