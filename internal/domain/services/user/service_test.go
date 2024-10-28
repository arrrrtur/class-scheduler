package user

import (
	"context"
	"fmt"
	"github.com/arrrrtur/class-scheduler.git/internal/domain/models"
	"github.com/arrrrtur/class-scheduler.git/internal/domain/services/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewUserService(t *testing.T) {
	type args struct {
		options UserServiceOptions
	}
	tests := []struct {
		name    string
		args    args
		want    *UserService
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				options: UserServiceOptions{
					Repository: func(test *testing.T) UserRepository {
						repository := mocks.NewUserRepository(test)
						assert.NotNil(test, repository)

						return repository
					}(t),
				},
			},
			want: &UserService{
				options: UserServiceOptions{
					Repository: mocks.NewUserRepository(t),
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "error/repository is nil",
			args: args{
				options: UserServiceOptions{
					Repository: nil,
				},
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserService(tt.args.options)

			tt.wantErr(t, err)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUserService_CreateUser(t *testing.T) {
	type fields struct {
		options UserServiceOptions
	}
	type args struct {
		ctx      context.Context
		username string
		group    string
		id       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fields: fields{
				options: UserServiceOptions{
					Repository: func(test *testing.T) UserRepository {
						repository := mocks.NewUserRepository(test)
						assert.NotNil(test, repository)
						repository.EXPECT().
							CreateUser(
								mock.Anything,
								mock.MatchedBy(
									func(user models.User) bool {
										return assert.Equal(test, "username", user.Username) &&
											assert.Equal(test, "group", user.Group) &&
											assert.Equal(test, "id", user.ID)
									},
								),
							).
							Times(1).
							Return(nil)
						return repository
					}(t),
				},
			},
			args: args{
				ctx:      context.Background(),
				username: "username",
				group:    "group",
				id:       "id",
			},
			wantErr: assert.NoError,
		},
		{
			name: "error/user already exists",
			fields: fields{
				options: UserServiceOptions{
					Repository: func(test *testing.T) UserRepository {
						repository := mocks.NewUserRepository(test)
						assert.NotNil(test, repository)
						repository.EXPECT().
							CreateUser(
								mock.Anything,
								mock.MatchedBy(
									func(user models.User) bool {
										return assert.Equal(test, "username", user.Username) &&
											assert.Equal(test, "group", user.Group) &&
											assert.Equal(test, "id", user.ID)
									},
								),
							).
							Times(1).
							Return(fmt.Errorf("user already exists"))
						return repository
					}(t),
				},
			},
			args: args{
				ctx:      context.Background(),
				username: "username",
				group:    "group",
				id:       "id",
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := UserService{
				options: tt.fields.options,
			}
			err := service.CreateUser(tt.args.ctx, tt.args.username, tt.args.group, tt.args.id)
			tt.wantErr(t, err)
		})
	}
}
