package services

import (
	"context"
	"fmt"
	"github.com/arrrrtur/class-scheduler.git/internal/domain/models"
	"time"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByID(ctx context.Context, id string) (models.User, error)
}

type UserServiceOptions struct {
	Repository UserRepository
}

type UserService struct {
	options UserServiceOptions
}

func NewUserService(options UserServiceOptions) *UserService {
	return &UserService{
		options: options,
	}
}

func (service UserService) CreateUser(ctx context.Context, username string, group string, id string) error {
	user := models.User{
		ID:        id,
		Username:  username,
		Group:     group,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := service.options.Repository.CreateUser(ctx, user)
	if err != nil {
		err = fmt.Errorf("failed to create user: %w", err)
		return err
	}

	return nil
}
