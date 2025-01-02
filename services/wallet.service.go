package services

import (
	"context"
	"time"

	"worklinker-api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WalletService struct {
	collection *mongo.Collection
}

func NewWalletService(db *mongo.Database) *WalletService {
	return &WalletService{
		collection: db.Collection("wallets"),
	}
}

func (s *WalletService) CreateWallet(userID string, currency string) (*models.Wallet, error) {
	wallet := &models.Wallet{
		ID:        primitive.NewObjectID().Hex(),
		UserID:    userID,
		Balance:   0.0,
		Currency:  currency,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	_, err := s.collection.InsertOne(context.TODO(), wallet)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}
