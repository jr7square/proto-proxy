package httpserver

import (
	"net/http"
)

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("hello world"))
	})
}
