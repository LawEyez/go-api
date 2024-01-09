package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/laweyez/go-api/internal/middleware"
)

func Handler(r *chi.Mux){
	// Global middleware
	r.Use(chimiddle.StripSlashes)
}