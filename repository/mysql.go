package repository

import (
	"database/sql"
	"log/slog"
	"user-api/models"
)

type MySQLRepository struct {
	db  *sql.DB
	log *slog.Logger
}

func NewMySQLRepository(db *sql.DB, log *slog.Logger) *MySQLRepository {
	return &MySQLRepository{db: db, log: log}
}
func (m *MySQLRepository) GetUsers() ([]models.User, error) {
	rows, err := m.db.Query("SELECT id,name,email FROM users")
	if err != nil {
		m.log.Error("Failed to fetch users", "operations", "GetUsers", "table", "users", "error", err)
		return []models.User{}, err
	}
	defer rows.Close()
	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			m.log.Error("failed to scanner", "operations", "GetUsers", "table", "users", "error", err)
			return []models.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}
func (m *MySQLRepository) CreateUser(user models.User) (models.User, error) {
	result, err := m.db.Exec("INSERT INTO users(name,email)VALUES(?,?)", user.Name, user.Email)
	if err != nil {
		m.log.Error("Failed to Insert user details", "operations", "CreateUser", "table", "users", "error", err)
		return models.User{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		m.log.Error("error ", "operations", "CreateUser", "table", "users", "error", err)
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
		m.log.Error("Failed to Get Details", "operations", "GetUserByID", "table", "users", "error", err)
		return models.User{}, err
	}
	return user, nil
}
func (m *MySQLRepository) UpdateUser(id int, user models.User) (models.User, error) {
	_, err := m.db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, id)
	if err != nil {
		m.log.Error("Failed to Update details", "operations", "UpdateUser", "table", "users", "error", err)
		return models.User{}, err
	}
	user.ID = id
	return user, nil
}
func (m *MySQLRepository) DeleteUser(id int) error {
	_, err := m.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		m.log.Error("Failed To delete the user", "operations", "UpdateUser", "table", "users", "error", err)
		return err
	}
	return nil
}
