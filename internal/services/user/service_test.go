package user

import (
	"context"
	"fmt"
	"github.com/arrrrtur/class-scheduler.git/internal/models"
	"github.com/arrrrtur/class-scheduler.git/internal/services/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewService(t *testing.T) {
	t.Parallel()

	type args struct {
		options ServiceOptions
	}
	tests := []struct {
		name    string
		args    args
		want    *Service
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				options: ServiceOptions{
					Repository: func(test *testing.T) Repository {
						repository := mocks.NewRepository(test)
						assert.NotNil(test, repository)

						return repository
					}(t),
				},
			},
			want: &Service{
				options: ServiceOptions{
					Repository: mocks.NewRepository(t),
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "error/repository is nil",
			args: args{
				options: ServiceOptions{
					Repository: nil,
				},
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := New(tt.args.options)

			tt.wantErr(t, err)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestService_CreateUser(t *testing.T) {
	t.Parallel()

	type fields struct {
		options ServiceOptions
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
				options: ServiceOptions{
					Repository: func(test *testing.T) Repository {
						repository := mocks.NewRepository(test)
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
				options: ServiceOptions{
					Repository: func(test *testing.T) Repository {
						repository := mocks.NewRepository(test)
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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			service := Service{
				options: tt.fields.options,
			}
			err := service.CreateUser(tt.args.ctx, tt.args.username, tt.args.group, tt.args.id)
			tt.wantErr(t, err)
		})
	}
}
