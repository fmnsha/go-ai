package handlers

import (
	"go-ai/pkg/services/workspace"
	"go-ai/server/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
)

type WorkspaceHandler struct {
	workspacesvcs workspace.WorkspaceSvcs
}

func NewWorkspaceHandler(i *do.Injector, r *chi.Mux) {
	h := &WorkspaceHandler{
		workspacesvcs: do.MustInvoke[workspace.WorkspaceSvcs](i),
	}

	r.Route("/workspaces", func(r chi.Router) {
		r.Post("/", helpers.Make(h.AddWorkspace))
	})

}

func (h *WorkspaceHandler) AddWorkspace(w http.ResponseWriter, r *http.Request) error {
	return nil
}
