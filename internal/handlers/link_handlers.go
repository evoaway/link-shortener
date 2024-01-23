package handlers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/evoaway/link-shortener/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

var (
	ErrNotFound = errors.New("not found")
)

type Storage interface {
	Create(ctx context.Context, url models.Link) error
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
		log.Printf("Error when decoding request due to %v", err)
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	var short = HashEncode(req.Link)
	newLink := models.Link{Short: short, Original: req.Link}
	err = h.storage.Create(r.Context(), newLink)
	if err != nil {
		log.Printf("Error when creating short link due to %v", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, newLink)
}

func (h *Handler) GetLink(w http.ResponseWriter, r *http.Request) {
	shortLink := chi.URLParam(r, "link")
	originalLink, err := h.storage.GetOne(r.Context(), shortLink)
	if err != nil {
		log.Printf("Error when getting link due to %v", err)
		if errors.Is(err, ErrNotFound) {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
		}
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
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
