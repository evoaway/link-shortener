package server

import (
	"context"
	"github.com/evoaway/link-shortener/internal/handlers"
	"github.com/evoaway/link-shortener/internal/storage"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func Run(ctx context.Context) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		return err
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Print(err)
		}
	}()
	linkStorage := storage.NewStorage(client.Database("links_db"))
	linkHandler := handlers.New(linkStorage)
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Post("/", linkHandler.CreateShortLink)
		r.Get("/{link}", linkHandler.GetLink)
	})
	return http.ListenAndServe(":3000", r)
}
