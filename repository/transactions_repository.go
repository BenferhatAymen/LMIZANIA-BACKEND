package repository

import (
	"context"
	"fmt"
	"lmizania/models"
	"time"

	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepo struct {
	MongoCollection *mongo.Collection
}

func (r *TransactionRepo) AddTransaction(transaction *models.Transaction) (interface{}, error) {
	transaction.ID = uuid.New().String()
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	result, err := r.MongoCollection.InsertOne(context.Background(), transaction)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *TransactionRepo) DeleteTransaction(transactionID string) error {

	_, err := r.MongoCollection.DeleteOne(context.Background(), bson.M{"_id": transactionID})
	if err != nil {
		return err
	}

	return nil
}
func (r *TransactionRepo) UpdateTransaction(transactionID string, updatedData *models.Transaction) (interface{}, error) {

	updatedData.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"title":       updatedData.Title,
			"type":        updatedData.Type,
			"amount":      updatedData.Amount,
			"date":        updatedData.Date,
			"description": updatedData.Description,
			"updated_at":  updatedData.UpdatedAt,
		},
	}

	result, err := r.MongoCollection.UpdateOne(context.Background(), bson.M{"_id": transactionID}, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (r *TransactionRepo) GetAllTransactions(userID string) ([]models.Transaction, error) {
	fmt.Println("userIDrepo", userID)

	filter := bson.M{"user_id": userID}
	cursor, err := r.MongoCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var transactions []models.Transaction
	for cursor.Next(context.Background()) {
		var transaction models.Transaction
		err := cursor.Decode(&transaction)
		if err != nil {
			return nil, err
		}
		fmt.Println("transaction", transaction)
		transactions = append(transactions, transaction)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
