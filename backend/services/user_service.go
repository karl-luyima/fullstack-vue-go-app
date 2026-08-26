package services

import (
	"context"
	"errors"

	"github.com/karl-luyima/fullstack-vue-go-app/backend/repositories"
)

type UserService interface {
	GetByID(ctx context.Context, id uint) (*UserResponse, error)
}

type userService struct {
	users repositories.UserRepository
}

func NewUserService(users repositories.UserRepository) UserService {
	return &userService{users: users}
}

func (s *userService) GetByID(ctx context.Context, id uint) (*UserResponse, error) {
	user, err := s.users.FindByID(ctx, id)
	if errors.Is(err, repositories.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	return &UserResponse{ID: user.ID, Name: user.Name, Email: user.Email}, nil
}