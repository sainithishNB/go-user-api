package repository

import (
	"database/sql"
	"fmt"
	"user-api/models"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}
func (m *MySQLRepository) GetUsers() ([]models.User, error) {
	rows, err := m.db.Query("SELECT id,name,email FROM users")
	if err != nil {
		return []models.User{}, err
	}
	defer rows.Close()
	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return []models.User{}, err
		}
		users = append(users, user)
	}
	fmt.Println("mysql")
	return users, nil
}
func (m *MySQLRepository) CreateUser(user models.User) (models.User, error) {
	result, err := m.db.Exec("INSERT INTO users(name,email)VALUES(?,?)", user.Name, user.Email)
	if err != nil {
		return models.User{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return models.User{}, err
	}
	user.ID = int(id)
	return user, nil
}
func (m *MySQLRepository) GetUserByID(id int) (models.User, error) {
	var user models.User
	row := m.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (m *MySQLRepository) UpdateUser(id int, user models.User) (models.User, error) {
	_, err := m.db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, id)
	if err != nil {
		return models.User{}, err
	}
	user.ID = id
	return user, nil
}
func (m *MySQLRepository) DeleteUser(id int) error {
	_, err := m.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
