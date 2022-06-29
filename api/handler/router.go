package handler

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (g *GenericHandler) Routes() *chi.Mux {

	r := chi.NewRouter()

	r.Get("/", g.GetAll)
	r.Post("/", g.Create)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
			g.Get("id", w, r)
		})
		r.Delete("/*", func(w http.ResponseWriter, r *http.Request) {
			g.Delete("id", w, r)
		})

	})
	return r
}
