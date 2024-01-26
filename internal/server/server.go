package server

import (
	"context"
	"github.com/evoaway/link-shortener/internal/config"
	"github.com/evoaway/link-shortener/internal/handlers"
	"github.com/evoaway/link-shortener/internal/storage"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func Run(ctx context.Context) error {
	conf := config.LoadConfig()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.MongoURI))
	if err != nil {
		return err
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Print(err)
		}
	}()
	linkStorage := storage.NewStorage(client.Database(conf.DBName))
	linkHandler := handlers.New(linkStorage)
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Post("/", linkHandler.CreateShortLink)
		r.Get("/{link}", linkHandler.GetLink)
	})
	return http.ListenAndServe(":"+conf.Port, r)
}
