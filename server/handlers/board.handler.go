package handlers

import (
	"encoding/json"
	"go-ai/pkg/services/board"
	"go-ai/pkg/services/board/models"
	"go-ai/server/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
)

type BoardHandler struct {
	boardsvcs board.BoardSvcs
}

func NewBoardHandler(i *do.Injector, r *chi.Mux) {
	h := &BoardHandler{
		boardsvcs: do.MustInvoke[board.BoardSvcs](i),
	}

	r.Route("/boards", func(r chi.Router) {
		r.Post("/", helpers.Make(h.AddBoard))
		r.Post("/{boardId}/records", helpers.Make(h.AddRecord))
	})

}

func (h *BoardHandler) AddBoard(w http.ResponseWriter, r *http.Request) error {

	var data models.BoardDto
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return err
	}

	board, err := h.boardsvcs.AddBoard(r.Context(), &data)
	if err != nil {
		return err
	}

	return helpers.WriteJson(w, http.StatusOK, board)

}

func (h *BoardHandler) AddRecord(w http.ResponseWriter, r *http.Request) error {

	id := chi.URLParam(r, "boardId")

	var data map[string]json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return err
	}

	result, err := h.boardsvcs.AddRecord(r.Context(), id, data)
	if err != nil {
		return err
	}

	return helpers.WriteJson(w, http.StatusOK, result)

}
