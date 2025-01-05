package stop

import (
	"fmt"
	"net/http"

	"github.com/whatdislol/mobility-app/types"
	"github.com/whatdislol/mobility-app/utils"
)

type Handler struct {
	store types.StopStore
}

func NewHandler(store types.StopStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.Handle("POST /api/stops", http.HandlerFunc(h.handleCreateStop))
}

func (h *Handler) handleCreateStop(w http.ResponseWriter, r *http.Request) {
	var stop types.Stop
	err := utils.ParseJSON(r, &stop)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.store.CreateStop(stop); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	utils.WriteJSON(w, http.StatusCreated, fmt.Sprintf("successfully added %+v", stop))
}