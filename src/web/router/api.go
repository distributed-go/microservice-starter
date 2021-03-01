package router

import "github.com/go-chi/chi"

// Router interface
type Router interface {
	Router(enableCORS bool) *chi.Mux
}
