package repository

import (
	"context"
	"errors"
	"lmizania/models"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GoalRepo struct {
	MongoCollection *mongo.Collection
	UserRepo        *UserRepo
}

// AddGoal adds a new goal to the database
func (r *GoalRepo) AddGoal(goal *models.Goal) (interface{}, error) {
	// Generate IDs and timestamps
	goal.ID = uuid.New().String()
	goal.CreatedAt = time.Now()
	goal.UpdatedAt = time.Now()
	goal.CurrentAmount = 0

	// Find the user
	_, err := r.UserRepo.FindUserByID(goal.UserID)
	if err != nil {
		return nil, err
	}

	// Insert the goal into the database
	result, err := r.MongoCollection.InsertOne(context.Background(), goal)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteGoal deletes a goal from the database by its ID
func (r *GoalRepo) DeleteGoal(goalID string) error {
	// Find the goal to ensure it exists
	var goal models.Goal
	err := r.MongoCollection.FindOne(context.Background(), bson.M{"_id": goalID}).Decode(&goal)
	if err != nil {
		return err
	}

	// Delete the goal from the database
	_, err = r.MongoCollection.DeleteOne(context.Background(), bson.M{"_id": goalID})
	if err != nil {
		return err
	}

	return nil
}

// UpdateGoal updates an existing goal's details in the database
func (r *GoalRepo) UpdateGoal(goalID string, updatedData *models.Goal) (interface{}, error) {
	// Find the existing goal
	var existingGoal models.Goal
	err := r.MongoCollection.FindOne(context.Background(), bson.M{"_id": goalID}).Decode(&existingGoal)
	if err != nil {
		return nil, err
	}

	// Update the timestamps
	updatedData.UpdatedAt = time.Now()

	// Ensure the total amount is not less than the current amount
	if updatedData.TotalAmount < existingGoal.CurrentAmount {
		return nil, errors.New("total amount cannot be less than the current amount")
	}

	// Update the goal in the database
	update := bson.M{
		"$set": bson.M{
			"title":        updatedData.Title,
			"total_amount": updatedData.TotalAmount,
			"updated_at":   updatedData.UpdatedAt,
		},
	}

	result, err := r.MongoCollection.UpdateOne(context.Background(), bson.M{"_id": goalID}, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetAllGoals retrieves all goals for a specific user by their userID
func (r *GoalRepo) GetAllGoals(userID string) ([]models.Goal, error) {
	filter := bson.M{"user_id": userID}
	cursor, err := r.MongoCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var goals []models.Goal
	for cursor.Next(context.Background()) {
		var goal models.Goal
		err := cursor.Decode(&goal)
		if err != nil {
			return nil, err
		}
		goals = append(goals, goal)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return goals, nil
}

// DepositGoal increases the current amount of a goal
func (r *GoalRepo) DepositGoal(goalID string, amount float64) error {
	// Find the goal
	var goal models.Goal
	err := r.MongoCollection.FindOne(context.Background(), bson.M{"_id": goalID}).Decode(&goal)
	if err != nil {
		return err
	}

	// Ensure the amount is not negative
	if amount < 0 {
		return errors.New("deposit amount cannot be negative")
	}

	// Ensure the current amount + deposit does not exceed the total amount
	if goal.CurrentAmount+amount > goal.TotalAmount {
		return errors.New("deposit exceeds the total goal amount")
	}

	// Update the current amount and timestamps
	goal.CurrentAmount += amount
	goal.UpdatedAt = time.Now()

	// Update the goal in the database
	update := bson.M{
		"$set": bson.M{
			"current_amount": goal.CurrentAmount,
			"updated_at":     goal.UpdatedAt,
		},
	}

	_, err = r.MongoCollection.UpdateOne(context.Background(), bson.M{"_id": goalID}, update)
	if err != nil {
		return err
	}

	return nil
}


