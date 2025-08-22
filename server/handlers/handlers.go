package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
)

func Register(i *do.Injector, r *chi.Mux) {
	NewBoardHandler(i, r)
	NewWorkspaceHandler(i, r)
}
