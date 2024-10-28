package services

import (
	"context"
	"fmt"
	"github.com/arrrrtur/class-scheduler.git/internal/domain/models"
	"time"
)

//go:generate mockery --name=UserRepository --case=underscore --outpkg=mocks --output=./mocks --with-expecter
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

func NewUserService(options UserServiceOptions) (*UserService, error) {
	if options.Repository == nil {
		err := fmt.Errorf("failed to create user service: repository is nil")
		return nil, err
	}

	return &UserService{
		options: options,
	}, nil
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
