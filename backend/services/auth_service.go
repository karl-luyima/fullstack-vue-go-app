package services

import (
	"context"
	"errors"

	"github.com/karl-luyima/fullstack-vue-go-app/backend/models"
	"github.com/karl-luyima/fullstack-vue-go-app/backend/repositories"
	"github.com/karl-luyima/fullstack-vue-go-app/backend/utils"
)

type AuthService interface {
	Register(ctx context.Context, input RegisterInput) (*AuthResponse, error)
	Login(ctx context.Context, input LoginInput) (*AuthResponse, error)
}

type authService struct {
	users repositories.UserRepository
	jwt   *utils.JWTManager
}

func NewAuthService(users repositories.UserRepository, jwt *utils.JWTManager) AuthService {
	return &authService{users: users, jwt: jwt}
}

func (s *authService) Register(ctx context.Context, input RegisterInput) (*AuthResponse, error) {
	exists, err := s.users.ExistsByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("an account with this email already exists")
	}

	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: hash,
	}
	if err := s.users.Create(ctx, user); err != nil {
		return nil, err
	}

	return s.buildResponse(user)
}

func (s *authService) Login(ctx context.Context, input LoginInput) (*AuthResponse, error) {
	user, err := s.users.FindByEmail(ctx, input.Email)
	if errors.Is(err, repositories.ErrRecordNotFound) {
		return nil, errors.New("invalid email or password")
	}
	if err != nil {
		return nil, err
	}

	if !utils.CheckPassword(user.PasswordHash, input.Password) {
		return nil, errors.New("invalid email or password")
	}

	return s.buildResponse(user)
}

func (s *authService) buildResponse(user *models.User) (*AuthResponse, error) {
	token, err := s.jwt.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	return &AuthResponse{
		User:  UserResponse{ID: user.ID, Name: user.Name, Email: user.Email},
		Token: token,
	}, nil
}