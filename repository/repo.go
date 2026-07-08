package repository

import (
	"errors"
	"user-api/models"
)

var users []models.User
var NextId int

type MemoryRepository struct{}

func (m *MemoryRepository) CreateUser(user models.User) (models.User, error) {
	NextId++
	user.ID = NextId
	users = append(users, user)
	return user, nil
}

func (m *MemoryRepository) GetUsers() []models.User {
	return users
}
func (m *MemoryRepository) GetUserByID(id int) (models.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, errors.New("Not Found")
}

func (m *MemoryRepository) UpdateUser(id int, newuser models.User) (models.User, error) {

	for i := range users {
		if users[i].ID == id {
			users[i].Name = newuser.Name
			users[i].Email = newuser.Email
			return users[i], nil
		}
	}
	return models.User{}, errors.New("Not Found")
}

func (m *MemoryRepository) DeleteUser(id int) error {
	for i := range users {
		if users[i].ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return errors.New("Not Found")
}
