package service

import (
	"user-api/models"
	"user-api/repository"
)

func CreateUser(user models.User) models.User {
	return repository.CreateUser(user)

}
func GetUsers() []models.User {
	return repository.GetUsers()
}
func GetUsersByID(id int) (models.User, error) {
	return repository.GetUsersByID(id)
}

func UpdateUser(id int, user models.User) (models.User, error) {
	return repository.UpdateUser(id, user)
}

func DeleteUser(id int) (bool, error) {
	return repository.DeleteUser(id)
}
