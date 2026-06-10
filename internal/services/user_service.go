package services

import (
	"github-analyzer/internal/models"
	"github-analyzer/internal/repositories"
)

func CreateUser(user models.User) (interface{}, error) {

	return repositories.CreateUser(user)

}

func GetAllUsers() ([]models.User, error) {

	return repositories.GetAllUsers()

}

func GetUserByID(id string) (*models.User, error) {

	return repositories.GetUserByID(id)

}

func UpdateUser(id string, user models.User) error {

	return repositories.UpdateUser(id, user)

}

func DeleteUser(id string) error {

	return repositories.DeleteUser(id)

}
