package service

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/repository"
	"context"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) RegisterUser(ctx context.Context, name, email string) (*models.User, error) {
	// 可以在这里添加业务逻辑，比如检查email是否已存在等
	user := &models.User{
		Name:  name,
		Email: email,
	}
	err := s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUser(ctx context.Context, id int64) (*models.User, error) {
	return s.userRepo.FindByID(ctx, id)
}
