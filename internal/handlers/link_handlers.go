package handlers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/evoaway/link-shortener/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

type Storage interface {
	Create(ctx context.Context, url models.Link) (*models.Link, error)
	GetOne(ctx context.Context, url string) (*models.Link, error)
}

type Handler struct {
	storage Storage
}

func New(s Storage) *Handler {
	return &Handler{
		storage: s,
	}
}

type Request struct {
	Link string `json:"link" validate:"required,url"`
}

func (h *Handler) CreateShortLink(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		log.Fatal(err)
	}
	var short = HashEncode(req.Link)
	newLink := models.Link{Short: short, Original: req.Link}
	_, err = h.storage.Create(r.Context(), newLink)
	if err != nil {
		log.Fatal(err)
	}
	render.JSON(w, r, newLink)
}

func (h *Handler) CetLink(w http.ResponseWriter, r *http.Request) {
	link := chi.URLParam(r, "link")
	originalLink, err := h.storage.GetOne(r.Context(), link)
	if err != nil {
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}
	render.JSON(w, r, originalLink)
}
func HashEncode(url string) string {
	h := sha256.New()
	h.Write([]byte(url))
	hashValue := h.Sum(nil)
	return hex.EncodeToString(hashValue)
}
