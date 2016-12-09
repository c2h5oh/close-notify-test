package main

import (
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/pressly/chi/render"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.CloseNotify)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		v := map[string]interface{}{"x": 1}

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, v)
	})
	http.ListenAndServe(":3333", r)
}
