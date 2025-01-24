package user

import (
	"context"
	"fmt"
	"github.com/arrrrtur/class-scheduler.git/internal/models"
	"time"
)

//go:generate mockery --name=Repository --case=underscore --outpkg=mocks --output=./mocks --with-expecter
type Repository interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByID(ctx context.Context, id string) (models.User, error)
}

type ServiceOptions struct {
	Repository Repository
}

type Service struct {
	options ServiceOptions
}

func New(options ServiceOptions) (*Service, error) {
	if options.Repository == nil {
		err := fmt.Errorf("repository is nil")
		return nil, err
	}

	return &Service{
		options: options,
	}, nil
}

func (s Service) CreateUser(ctx context.Context, username string, group string, id string) error {
	user := models.User{
		ID:        id,
		Username:  username,
		Group:     group,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.options.Repository.CreateUser(ctx, user)
	if err != nil {
		err = fmt.Errorf("failed to create user: %w", err)
		return err
	}

	return nil
}
