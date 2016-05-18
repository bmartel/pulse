package middleware

import "github.com/rs/cors"

// Cors returns a configured middleware to handle cross origin resource sharing
var Cors = cors.New(cors.Options{
	AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	AllowedMethods: []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
})
