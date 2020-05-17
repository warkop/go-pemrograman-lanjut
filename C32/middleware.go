package main

import (
	"net/http"
)

// CustomMux menampung koneksi mux
type CustomMux struct {
	http.ServeMux
	middlewares []func(next http.Handler) http.Handler
}

// RegisterMiddleware untuk mendaftarkan middleware
func (c *CustomMux) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux

	for _, next := range c.middlewares {
		current = next(current)
	}

	current.ServeHTTP(w, r)
}
