package handler

import (
	"encoding/json"
	"net/http"
	"product-service/internal/service"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(
	service *service.ProductService,
) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	product, err := h.service.Create(
		r.Context(),
		req.Name,
		req.Description,
		req.Price,
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(product)
}