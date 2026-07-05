package repository

import (
	"errors"
	"user-api/models"
)

var users []models.User
var NextId int

func CreateUser(user models.User) models.User {
	NextId++
	user.ID = NextId
	users = append(users, user)
	return user
}

func GetUsers() []models.User {
	return users
}
func GetUsersByID(id int) (models.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, errors.New("Not Found")
}

func UpdateUser(id int, newuser models.User) (models.User, error) {

	for i := range users {
		if users[i].ID == id {
			users[i].Name = newuser.Name
			users[i].Email = newuser.Email
			return users[i], nil
		}
	}
	return models.User{}, errors.New("Not Found")
}

func DeleteUser(id int) (bool, error) {
	for i := range users {
		if users[i].ID == id {
			users = append(users[:i], users[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("Not Found")
}
