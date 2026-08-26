// Package repositories is the only layer allowed to talk to the database
// directly. Everything else in the app asks this layer for data instead
// of writing its own database queries.
package repositories

import (
	"context"
	"errors"

	"github.com/karl-luyima/fullstack-vue-go-app/backend/models"
	"gorm.io/gorm"
)

var ErrRecordNotFound = errors.New("record not found")

// UserRepository is an interface — a list of things something can do,
// without saying exactly how. This lets other code depend on "something
// that can find/create users" rather than on GORM specifically.
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id uint) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

// userRepository is the actual implementation, using GORM.
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a userRepository. Returning the interface type
// (UserRepository) rather than the concrete struct is a common Go pattern —
// callers only see what they're allowed to do, not how it's implemented.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) FindByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}