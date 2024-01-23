package storage

import (
	"context"
	"errors"
	"github.com/evoaway/link-shortener/internal/handlers"
	"github.com/evoaway/link-shortener/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Storage struct {
	db *mongo.Database
}

func NewStorage(client *mongo.Database) *Storage {
	return &Storage{db: client}
}

func (s *Storage) Collection() *mongo.Collection {
	return s.db.Collection("links")
}

func (s *Storage) Create(ctx context.Context, link models.Link) error {
	if _, err := s.Collection().InsertOne(ctx, link); err != nil {
		log.Printf("%v", err)
		if !mongo.IsDuplicateKeyError(err) {
			return err
		}
	}
	return nil
}
func (s *Storage) GetOne(ctx context.Context, shortLink string) (*models.Link, error) {
	var originalLink models.Link
	if err := s.Collection().FindOne(ctx, bson.D{{"_id", shortLink}}).Decode(&originalLink); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, handlers.ErrNotFound
		}
		return nil, err
	}
	return &originalLink, nil
}
