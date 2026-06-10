package repositories

import (
	"context"
	"os"

	"github-analyzer/internal/database"
	"github-analyzer/internal/models"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateUser(user models.User) (interface{}, error) {

	collection := database.DB.Collection(
		os.Getenv("COLLECTION_NAME"),
	)

	result, err := collection.InsertOne(
		context.Background(),
		user,
	)

	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func GetAllUsers() ([]models.User, error) {

	collection := database.DB.Collection(
		os.Getenv("COLLECTION_NAME"),
	)

	cursor, err := collection.Find(
		context.Background(),
		map[string]interface{}{},
	)

	if err != nil {
		return nil, err
	}

	var users []models.User

	err = cursor.All(
		context.Background(),
		&users,
	)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(id string) (*models.User, error) {

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := database.DB.Collection(
		os.Getenv("COLLECTION_NAME"),
	)

	var user models.User

	err = collection.FindOne(
		context.Background(),
		bson.M{"_id": objectID},
	).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(id string, user models.User) error {

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	collection := database.DB.Collection(
		os.Getenv("COLLECTION_NAME"),
	)

	update := bson.M{
		"$set": bson.M{
			"name":        user.Name,
			"lastname":    user.LastName,
			"githubEmail": user.GithubEmail,
			"birthDate":   user.BirthDate,
		},
	}

	_, err = collection.UpdateOne(
		context.Background(),
		bson.M{"_id": objectID},
		update,
	)

	return err
}

func DeleteUser(id string) error {

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	collection := database.DB.Collection(
		os.Getenv("COLLECTION_NAME"),
	)

	_, err = collection.DeleteOne(
		context.Background(),
		bson.M{"_id": objectID},
	)

	return err
}
