package api

import (
	chi "github.com/go-chi/chi"
	"net/http"
)

// defaultRoutes returns the minimal routes for the api to work properly.
// This is used when the user does not provide a custom router.
// It is not intended to be used directly, but rather as a helper function for the api package.
func defaultRoutes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
		w.WriteHeader(http.StatusOK)
	})
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
		w.WriteHeader(http.StatusOK)
	})
	return r
}
