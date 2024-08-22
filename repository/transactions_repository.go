package repository

import (
	"context"
 	"lmizania/models"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepo struct {
	MongoCollection *mongo.Collection
	UserRepo        *UserRepo
}

func (r *TransactionRepo) AddTransaction(transaction *models.Transaction) (interface{}, error) {
	// Generate IDs and timestamps
	transaction.ID = uuid.New().String()
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	// Find user
	_, err := r.UserRepo.FindUserByID(transaction.UserID)
	if err != nil {
		return nil, err
	}

	// Update user's financial data based on transaction type
	if transaction.Type == "income" {
		err = r.UserRepo.IncreaseIncome(transaction.UserID, transaction.Amount)
	} else if transaction.Type == "expense" {
		err = r.UserRepo.IncreaseExpense(transaction.UserID, transaction.Amount)
	}

	if err != nil {
		return nil, err
	}

	// Insert the transaction into the database
	result, err := r.MongoCollection.InsertOne(context.Background(), transaction)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *TransactionRepo) DeleteTransaction(transactionID string) error {
	// Find the transaction to determine its type and amount
	var transaction models.Transaction
	err := r.MongoCollection.FindOne(context.Background(), bson.M{"_id": transactionID}).Decode(&transaction)
	if err != nil {
		return err
	}

	// Find the user
	_, err = r.UserRepo.FindUserByID(transaction.UserID)
	if err != nil {
		return err
	}

	// Adjust user's financial data based on transaction type
	if transaction.Type == "income" {
		err = r.UserRepo.DecreaseIncome(transaction.UserID, transaction.Amount)
	} else if transaction.Type == "expense" {
		err = r.UserRepo.DecreaseExpense(transaction.UserID, transaction.Amount)
	}

	if err != nil {
		return err
	}

	// Delete the transaction from the database
	_, err = r.MongoCollection.DeleteOne(context.Background(), bson.M{"_id": transactionID})
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepo) UpdateTransaction(transactionID string, updatedData *models.Transaction) (interface{}, error) {
	// Find the existing transaction
	var existingTransaction models.Transaction
	err := r.MongoCollection.FindOne(context.Background(), bson.M{"_id": transactionID}).Decode(&existingTransaction)
	if err != nil {
 
		return nil, err
	}

	// Update the timestamps
	updatedData.UpdatedAt = time.Now()

	// Find the user
	_, err = r.UserRepo.FindUserByID(updatedData.UserID)
 	if err != nil {
 
		return nil, err
	}

	// Adjust user's financial data if the amount or type has changed
	if existingTransaction.Type == "income" {
		err = r.UserRepo.DecreaseIncome(updatedData.UserID, existingTransaction.Amount)
		if err != nil {
			return nil, err
		}
	} else if existingTransaction.Type == "expense" {
		err = r.UserRepo.DecreaseExpense(updatedData.UserID, existingTransaction.Amount)
		if err != nil {
			return nil, err
		}
	}

	if updatedData.Type == "income" {
		err = r.UserRepo.IncreaseIncome(updatedData.UserID, updatedData.Amount)
	} else if updatedData.Type == "expense" {
		err = r.UserRepo.IncreaseExpense(updatedData.UserID, updatedData.Amount)
	}

	if err != nil {
		return nil, err
	}

	// Update the transaction in the database
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
 		transactions = append(transactions, transaction)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
