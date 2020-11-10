package server

import (
	"fmt"
	"net/http"

	"datnguyen.stably.test/api"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//StartServer -
func StartServer() {
	apiHandler := api.New()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/stably", apiHandler)
	fmt.Printf("starting the server: %s", "3100")
	http.ListenAndServe(":3100", r)
}
