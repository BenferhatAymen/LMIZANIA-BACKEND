package repository

import (
	"context"
 
	"lmizania/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	MongoCollection *mongo.Collection
}

// FindUserByID finds a user by their ID
func (r *UserRepo) FindUserByID(userID string) (*models.User, error) {
	var user models.User
	err := r.MongoCollection.FindOne(context.Background(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
 		return nil, err
	}
	return &user, nil
}

// IncreaseIncome increases the user's income and updates the wallet
func (r *UserRepo) IncreaseIncome(userID string, amount float64) error {
	user, err := r.FindUserByID(userID)
	if err != nil {
		return err
	}

	user.IncreaseIncome(amount)

	_, err = r.MongoCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": userID},
		bson.M{"$set": bson.M{"income": user.Income, "wallet": user.Wallet}},
	)
	return err
}

// DecreaseIncome decreases the user's income and updates the wallet
func (r *UserRepo) DecreaseIncome(userID string, amount float64) error {
	user, err := r.FindUserByID(userID)
	if err != nil {
		return err
	}

	err = user.DecreaseIncome(amount)
	if err != nil {
		return err
	}

	_, err = r.MongoCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": userID},
		bson.M{"$set": bson.M{"income": user.Income, "wallet": user.Wallet}},
	)
	return err
}

// IncreaseExpense increases the user's expenses and updates the wallet
func (r *UserRepo) IncreaseExpense(userID string, amount float64) error {
	user, err := r.FindUserByID(userID)
	if err != nil {
		return err
	}

	user.IncreaseExpense(amount)

	_, err = r.MongoCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": userID},
		bson.M{"$set": bson.M{"expenses": user.Expenses, "wallet": user.Wallet}},
	)
	return err
}

// DecreaseExpense decreases the user's expenses and updates the wallet
func (r *UserRepo) DecreaseExpense(userID string, amount float64) error {
	user, err := r.FindUserByID(userID)
	if err != nil {
		return err
	}

	err = user.DecreaseExpense(amount)
	if err != nil {
		return err
	}

	_, err = r.MongoCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": userID},
		bson.M{"$set": bson.M{"expenses": user.Expenses, "wallet": user.Wallet}},
	)
	return err
}

// DepositSavings increases the user's savings
func (r *UserRepo) DepositSavings(userID string, amount float64) error {
	user, err := r.FindUserByID(userID)
	if err != nil {
		return err
	}

	user.DepositSavings(amount)

	_, err = r.MongoCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": userID},
		bson.M{"$set": bson.M{"savings": user.Savings}},
	)
	return err
}

// SetWallet sets the user's wallet value
func (r *UserRepo) SetWallet(userID string, amount float64) error {
	user, err := r.FindUserByID(userID)
	if err != nil {
		return err
	}

	user.SetWallet(amount)

	_, err = r.MongoCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": userID},
		bson.M{"$set": bson.M{"wallet": user.Wallet}},
	)
	return err
}

// GetWallet retrieves the user's wallet value
func (r *UserRepo) GetWallet(userID string) (float64, error) {
	user, err := r.FindUserByID(userID)
	if err != nil {
		return 0, err
	}

	return user.GetWallet(), nil
}

// SetTarget sets the user's savings target
func (r *UserRepo) SetTarget(userID string, target float64) error {
	user, err := r.FindUserByID(userID)
	if err != nil {
		return err
	}

	user.SetTarget(target)

	_, err = r.MongoCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": userID},
		bson.M{"$set": bson.M{"savings": user.Savings}},
	)
	return err
}

// GetTarget retrieves the user's savings target
func (r *UserRepo) GetTarget(userID string) (float64, error) {
	user, err := r.FindUserByID(userID)
	if err != nil {
		return 0, err
	}

	return user.GetTarget(), nil
}
