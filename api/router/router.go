package router

import (
	"api-development/handlers"
	"net/http"
)

func New() *http.ServeMux {
	r := http.NewServeMux()

	r.Handle("GET /", handlers.ServeHome())

	r.Handle("GET /news", handlers.GetAllNews())
	r.Handle("POST /news", handlers.PostNews())
	r.Handle("GET /news/{id}", handlers.GetNewsById())
	r.Handle("DELETE /news/{id}", handlers.DeleteNewsById())

	return r
}
