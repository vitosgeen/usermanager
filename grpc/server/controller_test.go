package server

import (
	"context"
	"fmt"
	"testing"

	grpcUsermanager "usermanager/grpc"
	"usermanager/internal/domain/model"
	"usermanager/internal/usecase/usecase"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserManagerGrpcController_GetUser(t *testing.T) {
	user := model.User{
		UserID:    uuid.New(),
		Nickname:  "testuser",
		FirstName: "Test",
		LastName:  "User",
		Email:     "test@test.com",
		Password:  "password",
		IsPublic:  true,
		Role:      model.UsualUser,
	}
	mockUserUsecase := &UserUsecaseMock{}
	mockUserUsecase.On("GetUser", mock.Anything, user.UserID).Return(&user, nil)
	mockUserUsecaseErr := &UserUsecaseMock{}
	err := fmt.Errorf("error")
	mockUserUsecaseErr.On("GetUser", mock.Anything, user.UserID).Return((*model.User)(nil), err)
	type fields struct {
		usecase usecase.IUserUsecase
	}
	type args struct {
		ctx         context.Context
		userRequest *grpcUsermanager.GetUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpcUsermanager.GetUserResponse
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				usecase: mockUserUsecase,
			},
			args: args{
				ctx: context.Background(),
				userRequest: &grpcUsermanager.GetUserRequest{
					UserId: user.UserID.String(),
				},
			},
			want: &grpcUsermanager.GetUserResponse{
				User: &grpcUsermanager.User{
					UserId:    user.UserID.String(),
					Nickname:  user.Nickname,
					FirstName: user.FirstName,
					LastName:  user.LastName,
					Email:     user.Email,
					Password:  user.Password,
					IsPublic:  user.IsPublic,
					UserRole:  user.Role,
				},
			},
			wantErr: false,
		},
		{
			name: "Error",
			fields: fields{
				usecase: mockUserUsecaseErr,
			},
			args: args{
				ctx: context.Background(),
				userRequest: &grpcUsermanager.GetUserRequest{
					UserId: user.UserID.String(),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := NewUserManagerGrpcController(tt.fields.usecase)
			got, err := ctrl.GetUser(tt.args.ctx, tt.args.userRequest)

			assert.Equal(t, got, tt.want)
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}
}
