package repository

import (
	"awesomeProject/internal/models"
	"context"
	"database/sql"
)

// UserRepository 接口定义了用户数据的所有操作，便于测试和解耦
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id int64) (*models.User, error)
}

// mysqlUserRepository 是 UserRepository 的 MySQL 实现
type mysqlUserRepository struct {
	db *sql.DB
}

// NewUserRepository 创建一个新的 UserRepository 实例
func NewUserRepository(db *sql.DB) UserRepository {
	return &mysqlUserRepository{db: db}
}

func (r *mysqlUserRepository) Create(ctx context.Context, user *models.User) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, err := r.db.ExecContext(ctx, query, user.Name, user.Email)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}

func (r *mysqlUserRepository) FindByID(ctx context.Context, id int64) (*models.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, id)

	var user models.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Or a custom not found error
		}
		return nil, err
	}
	return &user, nil
}
