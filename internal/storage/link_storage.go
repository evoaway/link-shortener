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

func (s *Storage) Create(ctx context.Context, url models.Link) (*models.Link, error) {
	count, err := s.Collection().CountDocuments(ctx, bson.M{"_id": url.Short})
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	if count > 0 {
		return nil, fmt.Errorf("%s", "ID already exists")
	}
	_, err = s.Collection().InsertOne(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return &url, nil
}
func (s *Storage) GetOne(ctx context.Context, url string) (*models.Link, error) {
	var originalURl models.Link
	if err := s.Collection().FindOne(ctx, bson.D{{"_id", url}}).Decode(&originalURl); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("%s", "Record does not exist")
		}
		return nil, fmt.Errorf("%w", err)
	}
	return &originalURl, nil
}
