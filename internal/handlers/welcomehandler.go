package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HelloRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", getHello)

	return r
}

func getHello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
}
