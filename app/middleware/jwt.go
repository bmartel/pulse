package middleware

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/bmartel/pulse/config"
	"github.com/dgrijalva/jwt-go"
)

// Jwt returns a configured json web token middleware
var Jwt = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppKey), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
