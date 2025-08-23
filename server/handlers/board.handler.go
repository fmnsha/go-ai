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
		r.Post("/{boardId}/items", helpers.Make(h.AddRecord))
		r.Put("/{boardId}/items/{itemId}", helpers.Make(h.UpdateItem))
		r.Get("/{boardId}/items", helpers.Make(h.GetAllRecords))
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

func (h *BoardHandler) UpdateItem(w http.ResponseWriter, r *http.Request) error {

	boardId := chi.URLParam(r, "boardId")
	itemId := chi.URLParam(r, "itemId")

	var data map[string]json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return err
	}

	result, err := h.boardsvcs.UpdateRecord(r.Context(), boardId, itemId, data)
	if err != nil {
		return err
	}

	return helpers.WriteJson(w, http.StatusOK, result)

}

func (h *BoardHandler) GetAllRecords(w http.ResponseWriter, r *http.Request) error {

	boardId := chi.URLParam(r, "boardId")

	records, err := h.boardsvcs.GetAllRecords(r.Context(), boardId)
	if err != nil {
		return err
	}

	return helpers.WriteJson(w, http.StatusOK, records)
}
