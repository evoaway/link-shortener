package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/evoaway/link-shortener/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (s *Storage) Create(ctx context.Context, link models.Link) (*models.Link, error) {
	_, err := s.Collection().InsertOne(ctx, link)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return &link, nil
}
func (s *Storage) GetOne(ctx context.Context, shortLink string) (*models.Link, error) {
	var originalLink models.Link
	if err := s.Collection().FindOne(ctx, bson.D{{"_id", shortLink}}).Decode(&originalLink); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("%s", "Record does not exist")
		}
		return nil, fmt.Errorf("%w", err)
	}
	return &originalLink, nil
}
