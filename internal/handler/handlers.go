package handler

import (
	"net/http"

	"github.com/alp-tahta/prod-rdy-go-microservice/internal/service"
)

type Handler struct {
	s service.ServiceInterface
}

func NewHandler(s service.ServiceInterface) *Handler {
	return &Handler{s: s}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) HealthChecker(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	h.s.CreateProduct()
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	h.s.GetProduct()
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	h.s.UpdateProduct()
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	h.s.DeleteProduct()
}
