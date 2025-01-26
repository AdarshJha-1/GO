package router

import (
	"net/http"
	"new-api/internal/handler"
)

func New() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("POST /news", handler.PostNews())
	r.HandleFunc("GET /news", handler.GetAllNews())
	r.HandleFunc("GET /news/{news_id}", handler.GetNewsById())
	r.HandleFunc("PUT /news/{news_id}", handler.UpdateNewsById())
	r.HandleFunc("DELETE /news/{news_id}", handler.DeleteNewsById())
	return r
}
