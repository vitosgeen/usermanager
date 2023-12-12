package usecase

import (
	"context"
	"fmt"
	"testing"
	"time"

	"usermanager/internal/apperrors"
	"usermanager/internal/domain/model"
	"usermanager/internal/interface/repository"
	"usermanager/internal/utils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gotest.tools/v3/assert"
)

func TestUserUsecase_CreateUser(t *testing.T) {
	userRepoMock := &UserRepositoryMock{}
	user := &model.User{Nickname: "nickname", FirstName: "fname", LastName: "lname"}
	userRepoMock.On("CreateUser", user).Return(user, nil)
	userRepoMock.On("FindUserByNickname", context.TODO(), user.Nickname).Return((*model.User)(nil), nil)
	userRepoMock.On("SaveUser", context.TODO(), user).Return(user, nil)
	type fields struct {
		UserRepo      repository.UserRepository
		VoteRepo      repository.VoteRepository
		UserRedisRepo repository.UserRedisRepository
		VoteRedisRepo repository.VoteRedisRepository
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{"create profile", fields{UserRepo: userRepoMock}, args{ctx: context.TODO(), user: user}, user, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userusecase := NewUserUsecase(tt.fields.UserRepo, tt.fields.VoteRepo, tt.fields.UserRedisRepo, tt.fields.VoteRedisRepo)
			got, err := userusecase.CreateUser(tt.args.ctx, tt.args.user)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, !(err == nil), tt.wantErr)
		})
	}
}

func TestUserUsecase_CreateUser_Error_ExistUser(t *testing.T) {
	userRepoMock := &UserRepositoryMock{}
	user := &model.User{Nickname: "nickname", FirstName: "fname", LastName: "lname"}
	userRepoMock.On("CreateUser", user).Return(user, nil)
	userRepoMock.On("FindUserByNickname", mock.Anything, user.Nickname).Return(user, nil)
	userRepoMock.On("SaveUser", mock.Anything, user).Return((*model.User)(nil), fmt.Errorf("something went wrong"))
	type fields struct {
		UserRepo      repository.UserRepository
		VoteRepo      repository.VoteRepository
		UserRedisRepo repository.UserRedisRepository
		VoteRedisRepo repository.VoteRedisRepository
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{"create profile exist error", fields{UserRepo: userRepoMock}, args{ctx: context.TODO(), user: user}, user, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userusecase := NewUserUsecase(tt.fields.UserRepo, tt.fields.VoteRepo, tt.fields.UserRedisRepo, tt.fields.VoteRedisRepo)
			got, err := userusecase.CreateUser(tt.args.ctx, tt.args.user)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, (err == nil), tt.wantErr)
		})
	}
}

func TestUserUsecase_CreateUser_Error_SaveUser(t *testing.T) {
	userRepoMock := &UserRepositoryMock{}
	user := &model.User{Nickname: "nickname", FirstName: "fname", LastName: "lname"}
	userRepoMock.On("CreateUser", user).Return(user, nil)
	userRepoMock.On("FindUserByNickname", mock.Anything, user.Nickname).Return((*model.User)(nil), nil)
	userRepoMock.On("SaveUser", mock.Anything, user).Return((*model.User)(nil), fmt.Errorf("something went wrong"))
	type fields struct {
		UserRepo      repository.UserRepository
		VoteRepo      repository.VoteRepository
		UserRedisRepo repository.UserRedisRepository
		VoteRedisRepo repository.VoteRedisRepository
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{"create profile add error", fields{UserRepo: userRepoMock}, args{ctx: context.TODO(), user: user}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userusecase := NewUserUsecase(tt.fields.UserRepo, tt.fields.VoteRepo, tt.fields.UserRedisRepo, tt.fields.VoteRedisRepo)
			got, err := userusecase.CreateUser(tt.args.ctx, tt.args.user)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, (err == nil), tt.wantErr)
		})
	}
}

func TestUserUsecase_UpdateUser(t *testing.T) {
	userRepoMock := &UserRepositoryMock{}
	user := &model.User{Nickname: "nickname", FirstName: "fname", LastName: "lname"}
	userRepoMock.On("UpdateUser", mock.Anything, user).Return(user, nil)
	type fields struct {
		UserRepo      repository.UserRepository
		VoteRepo      repository.VoteRepository
		UserRedisRepo repository.UserRedisRepository
		VoteRedisRepo repository.VoteRedisRepository
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{"update profile", fields{UserRepo: userRepoMock}, args{ctx: context.TODO(), user: user}, user, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userusecase := NewUserUsecase(tt.fields.UserRepo, tt.fields.VoteRepo, tt.fields.UserRedisRepo, tt.fields.VoteRedisRepo)
			got, err := userusecase.UpdateUser(tt.args.ctx, tt.args.user)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, !(err == nil), tt.wantErr)
		})
	}
}

func TestUserUsecase_UpdateUser_Error(t *testing.T) {
	userRepoMock := &UserRepositoryMock{}
	user := &model.User{Nickname: "nickname", FirstName: "fname", LastName: "lname"}
	userRepoMock.On("UpdateUser", mock.Anything, user).Return((*model.User)(nil), nil)
	type fields struct {
		UserRepo      repository.UserRepository
		VoteRepo      repository.VoteRepository
		UserRedisRepo repository.UserRedisRepository
		VoteRedisRepo repository.VoteRedisRepository
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{"update profile", fields{UserRepo: userRepoMock}, args{ctx: context.TODO(), user: user}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userusecase := NewUserUsecase(tt.fields.UserRepo, tt.fields.VoteRepo, tt.fields.UserRedisRepo, tt.fields.VoteRedisRepo)
			got, err := userusecase.UpdateUser(tt.args.ctx, tt.args.user)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, !(err == nil), tt.wantErr)
		})
	}
}

func TestUserUsecase_GetUsers(t *testing.T) {
	userRepoMock := &UserRepositoryMock{}
	userRedisRepoMock := &UserRedisRepositoryMock{}
	voteRepoMock := &VoteRepositoryMock{}
	voteRedisRepoMock := &VoteRedisRepositoryMock{}

	users := &model.Users{Users: []*model.User{}}
	paginationQuery := &utils.PaginationQuery{Size: 10, Page: 0}
	userRepoMock.On("GetUsers", mock.Anything, paginationQuery).Return(users, nil)
	userRedisRepoMock.On("GetUsers", mock.Anything, paginationQuery).Return(users, nil)
	userRedisRepoMock.On("SetGetUsers", mock.Anything, paginationQuery, users).Return(nil)
	type fields struct {
		UserRepo      repository.UserRepository
		VoteRepo      repository.VoteRepository
		UserRedisRepo repository.UserRedisRepository
		VoteRedisRepo repository.VoteRedisRepository
	}
	type args struct {
		ctx             context.Context
		paginationQuery *utils.PaginationQuery
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Users
		wantErr bool
	}{
		{"list profiles", fields{UserRepo: userRepoMock, VoteRepo: voteRepoMock, UserRedisRepo: userRedisRepoMock, VoteRedisRepo: voteRedisRepoMock}, args{ctx: context.TODO(), paginationQuery: paginationQuery}, users, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userusecase := NewUserUsecase(tt.fields.UserRepo, tt.fields.VoteRepo, tt.fields.UserRedisRepo, tt.fields.VoteRedisRepo)
			got, err := userusecase.GetUsers(tt.args.ctx, tt.args.paginationQuery)
			assert.Equal(t, !(err == nil), tt.wantErr)
			assert.Equal(t, users, got)
		})
	}
}

func TestUserUsecase_GetUsers_Error(t *testing.T) {
	userRepoMock := &UserRepositoryMock{}
	userRedisRepoMock := &UserRedisRepositoryMock{}
	voteRepoMock := &VoteRepositoryMock{}
	voteRedisRepoMock := &VoteRedisRepositoryMock{}
	paginationQuery := &utils.PaginationQuery{Size: 10, Page: 0}
	userRepoMock.On("GetUsers", mock.Anything, paginationQuery).Return(&model.Users{}, fmt.Errorf("something went wrong"))
	userRedisRepoMock.On("GetUsers", mock.Anything, paginationQuery).Return((*model.Users)(nil), nil)
	userRedisRepoMock.On("SetGetUsers", mock.Anything, paginationQuery, &model.Users{}).Return(nil)
	type fields struct {
		UserRepo      repository.UserRepository
		VoteRepo      repository.VoteRepository
		UserRedisRepo repository.UserRedisRepository
		VoteRedisRepo repository.VoteRedisRepository
	}
	type args struct {
		ctx             context.Context
		paginationQuery *utils.PaginationQuery
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Users
		wantErr bool
	}{
		{"list profiles error", fields{UserRepo: userRepoMock, VoteRepo: voteRepoMock, UserRedisRepo: userRedisRepoMock, VoteRedisRepo: voteRedisRepoMock}, args{ctx: context.TODO(), paginationQuery: paginationQuery}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userusecase := NewUserUsecase(tt.fields.UserRepo, tt.fields.VoteRepo, tt.fields.UserRedisRepo, tt.fields.VoteRedisRepo)
			got, err := userusecase.GetUsers(tt.args.ctx, tt.args.paginationQuery)
			assert.Equal(t, !(err == nil), tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUserUsecase_GetUser(t *testing.T) {
	userRepoMock := &UserRepositoryMock{}
	userRedisRepoMock := &UserRedisRepositoryMock{}
	voteRepoMock := &VoteRepositoryMock{}
	voteRedisRepoMock := &VoteRedisRepositoryMock{}
	user := &model.User{Nickname: "nickname", FirstName: "fname", LastName: "lname"}
	userRepoMock.On("FindUserByUUID", mock.Anything, user.UserID).Return(user, nil)
	userRedisRepoMock.On("FindUserByUUID", mock.Anything, user.UserID).Return((*model.User)(nil), nil)
	userRedisRepoMock.On("SetFindUserByUUID", mock.Anything, user.UserID, user).Return(nil)
	type fields struct {
		UserRepo      repository.UserRepository
		VoteRepo      repository.VoteRepository
		UserRedisRepo repository.UserRedisRepository
		VoteRedisRepo repository.VoteRedisRepository
	}
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{"get user", fields{UserRepo: userRepoMock, VoteRepo: voteRepoMock, UserRedisRepo: userRedisRepoMock, VoteRedisRepo: voteRedisRepoMock}, args{ctx: context.TODO(), userID: user.UserID}, user, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userusecase := NewUserUsecase(tt.fields.UserRepo, tt.fields.VoteRepo, tt.fields.UserRedisRepo, tt.fields.VoteRedisRepo)
			got, err := userusecase.GetUser(tt.args.ctx, tt.args.userID)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, !(err == nil), tt.wantErr)
		})
	}
}

func TestUserUsecase_GetUser_Error(t *testing.T) {
	userRepoMock := &UserRepositoryMock{}
	userRedisRepoMock := &UserRedisRepositoryMock{}
	voteRepoMock := &VoteRepositoryMock{}
	voteRedisRepoMock := &VoteRedisRepositoryMock{}
	user := &model.User{Nickname: "nickname", FirstName: "fname", LastName: "lname"}
	userRepoMock.On("FindUserByUUID", mock.Anything, user.UserID).Return((*model.User)(nil), apperrors.UserRepoFindUserByUUIDGetContext.AppendMessage(fmt.Errorf("something went wrong")))
	userRedisRepoMock.On("FindUserByUUID", mock.Anything, user.UserID).Return((*model.User)(nil), nil)
	userRedisRepoMock.On("SetFindUserByUUID", mock.Anything, user.UserID, user).Return(nil)
	type fields struct {
		UserRepo      repository.UserRepository
		VoteRepo      repository.VoteRepository
		UserRedisRepo repository.UserRedisRepository
		VoteRedisRepo repository.VoteRedisRepository
	}
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{"get user error", fields{UserRepo: userRepoMock, VoteRepo: voteRepoMock, UserRedisRepo: userRedisRepoMock, VoteRedisRepo: voteRedisRepoMock}, args{ctx: context.TODO(), userID: user.UserID}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userusecase := NewUserUsecase(tt.fields.UserRepo, tt.fields.VoteRepo, tt.fields.UserRedisRepo, tt.fields.VoteRedisRepo)
			got, err := userusecase.GetUser(tt.args.ctx, tt.args.userID)
			fmt.Println("TestUserUsecase_GetUser_Error ERROR", err)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, !(err == nil), tt.wantErr)
		})
	}
}

func TestUserUsecase_GetUserByNickname(t *testing.T) {
	userRepoMock := &UserRepositoryMock{}
	userRedisRepoMock := &UserRedisRepositoryMock{}
	voteRepoMock := &VoteRepositoryMock{}
	voteRedisRepoMock := &VoteRedisRepositoryMock{}
	user := &model.User{Nickname: "nickname", FirstName: "fname", LastName: "lname"}
	userRepoMock.On("FindUserByNickname", mock.Anything, user.Nickname).Return(user, nil)
	userRedisRepoMock.On("FindUserByNickname", mock.Anything, user.Nickname).Return((*model.User)(nil), nil)
	userRedisRepoMock.On("SetFindUserByNickname", mock.Anything, user.Nickname, user).Return(nil)
	type fields struct {
		UserRepo      repository.UserRepository
		VoteRepo      repository.VoteRepository
		UserRedisRepo repository.UserRedisRepository
		VoteRedisRepo repository.VoteRedisRepository
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{"load profile", fields{UserRepo: userRepoMock, VoteRepo: voteRepoMock, UserRedisRepo: userRedisRepoMock, VoteRedisRepo: voteRedisRepoMock}, args{ctx: context.TODO(), user: user}, user, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userusecase := NewUserUsecase(tt.fields.UserRepo, tt.fields.VoteRepo, tt.fields.UserRedisRepo, tt.fields.VoteRedisRepo)
			got, err := userusecase.GetUserByNickname(tt.args.ctx, tt.args.user.Nickname)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, !(err == nil), tt.wantErr)
		})
	}
}

func TestUserUsecase_GetUserByNickname_Error(t *testing.T) {
	userRepoMock := &UserRepositoryMock{}
	userRedisRepoMock := &UserRedisRepositoryMock{}
	voteRepoMock := &VoteRepositoryMock{}
	voteRedisRepoMock := &VoteRedisRepositoryMock{}
	user := &model.User{Nickname: "nickname", FirstName: "fname", LastName: "lname"}
	userRepoMock.On("FindUserByNickname", mock.Anything, user.Nickname).Return((*model.User)(nil), fmt.Errorf("something went wrong"))
	userRedisRepoMock.On("FindUserByNickname", mock.Anything, user.Nickname).Return((*model.User)(nil), nil)
	userRedisRepoMock.On("SetFindUserByNickname", mock.Anything, user.Nickname, user).Return(nil)
	type fields struct {
		UserRepo      repository.UserRepository
		VoteRepo      repository.VoteRepository
		UserRedisRepo repository.UserRedisRepository
		VoteRedisRepo repository.VoteRedisRepository
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{"load profile error", fields{UserRepo: userRepoMock, VoteRepo: voteRepoMock, UserRedisRepo: userRedisRepoMock, VoteRedisRepo: voteRedisRepoMock}, args{ctx: context.TODO(), user: user}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userusecase := NewUserUsecase(tt.fields.UserRepo, tt.fields.VoteRepo, tt.fields.UserRedisRepo, tt.fields.VoteRedisRepo)
			got, err := userusecase.GetUserByNickname(tt.args.ctx, tt.args.user.Nickname)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, !(err == nil), tt.wantErr)
		})
	}
}

func TestUserUsecase_CheckUserByNickname(t *testing.T) {
	userRepoMock := &UserRepositoryMock{}
	user := &model.User{UserID: uuid.New(), Nickname: "nickname", FirstName: "fname", LastName: "lname"}
	userRepoMock.On("FindUserByNickname", mock.Anything, user.Nickname).Return(user, nil)
	type fields struct {
		UserRepo      repository.UserRepository
		VoteRepo      repository.VoteRepository
		UserRedisRepo repository.UserRedisRepository
		VoteRedisRepo repository.VoteRedisRepository
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{"check profile by nick", fields{UserRepo: userRepoMock}, args{ctx: context.TODO(), user: user}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userusecase := NewUserUsecase(tt.fields.UserRepo, tt.fields.VoteRepo, tt.fields.UserRedisRepo, tt.fields.VoteRedisRepo)
			got, err := userusecase.CheckUserByNickname(tt.args.ctx, tt.args.user)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, !(err == nil), tt.wantErr)
		})
	}
}

func TestUserUsecase_CheckUserByNickname_Error(t *testing.T) {
	userRepoMock := &UserRepositoryMock{}
	user := &model.User{UserID: uuid.New(), Nickname: "nickname", FirstName: "fname", LastName: "lname"}
	userRepoMock.On("FindUserByNickname", mock.Anything, user.Nickname).Return((*model.User)(nil), apperrors.UserUsecaseCheckProfileByNick.AppendMessage(fmt.Errorf("something went wrong")))
	type fields struct {
		UserRepo      repository.UserRepository
		VoteRepo      repository.VoteRepository
		UserRedisRepo repository.UserRedisRepository
		VoteRedisRepo repository.VoteRedisRepository
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{"check profile by nick error", fields{UserRepo: userRepoMock}, args{ctx: context.TODO(), user: user}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userusecase := NewUserUsecase(tt.fields.UserRepo, tt.fields.VoteRepo, tt.fields.UserRedisRepo, tt.fields.VoteRedisRepo)
			got, err := userusecase.CheckUserByNickname(tt.args.ctx, tt.args.user)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, !(err == nil), tt.wantErr)
		})
	}
}

func TestUserUsecase_CheckUserByNickname_Error_NicknameExist(t *testing.T) {
	userRepoMock := &UserRepositoryMock{}
	user := &model.User{UserID: uuid.New(), Nickname: "nickname", FirstName: "fname", LastName: "lname"}
	checkedUser := &model.User{UserID: uuid.New(), Nickname: "nickname", FirstName: "fname", LastName: "lname"}
	userRepoMock.On("FindUserByNickname", mock.Anything, user.Nickname).Return(checkedUser, nil)
	type fields struct {
		UserRepo      repository.UserRepository
		VoteRepo      repository.VoteRepository
		UserRedisRepo repository.UserRedisRepository
		VoteRedisRepo repository.VoteRedisRepository
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{"check profile by nick error nickname exit", fields{UserRepo: userRepoMock}, args{ctx: context.TODO(), user: user}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userusecase := NewUserUsecase(tt.fields.UserRepo, tt.fields.VoteRepo, tt.fields.UserRedisRepo, tt.fields.VoteRedisRepo)
			got, err := userusecase.CheckUserByNickname(tt.args.ctx, tt.args.user)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, !(err == nil), tt.wantErr)
		})
	}
}

func TestUserUsecase_FindExistVoting(t *testing.T) {
	userRepoMock := &UserRepositoryMock{}
	voteRepoMock := &VoteRepositoryMock{}
	voteRepoMockNil := &VoteRepositoryMock{}
	voteRepoMockErr := &VoteRepositoryMock{}
	userID := uuid.New()
	userID1 := uuid.New()
	userID2 := uuid.New()
	userID3 := uuid.New()
	voterID := uuid.New()
	voterID1 := uuid.New()
	voterID2 := uuid.New()
	voterID3 := uuid.New()
	timeNow := time.Now()
	votes := make([]*model.Vote, 0)
	votes = append(votes, &model.Vote{VoteID: 1, Vote: 1, CreatedUserID: voterID, CreatedAt: &timeNow})
	votes = append(votes, &model.Vote{VoteID: 2, Vote: -1, CreatedUserID: voterID1, CreatedAt: &timeNow})
	votes = append(votes, &model.Vote{VoteID: 3, Vote: 1, CreatedUserID: voterID2, CreatedAt: &timeNow})
	userVotes := make([]*model.UserVote, 0)
	userVotes = append(userVotes, &model.UserVote{ID: 1, VoteID: 1, UserID: userID})
	userVotes = append(userVotes, &model.UserVote{ID: 2, VoteID: 2, UserID: userID1})
	userVotes = append(userVotes, &model.UserVote{ID: 3, VoteID: 3, UserID: userID2})
	userVotesEmpty := make([]*model.UserVote, 0)

	votesNil := make([]*model.Vote, 0)
	errorTest := fmt.Errorf("something went wrong")
	voteRepoMock.On("FindVoteByUserID", mock.Anything, &voterID).Return(votes, nil)
	voteRepoMock.On("FindUserVoteByUserID", mock.Anything, &userID).Return(userVotes, nil)
	voteRepoMockNil.On("FindVoteByUserID", mock.Anything, &voterID3).Return(votesNil, nil)
	voteRepoMockNil.On("FindUserVoteByUserID", mock.Anything, &userID3).Return(userVotesEmpty, nil)
	voteRepoMockErr.On("FindVoteByUserID", mock.Anything, &voterID).Return(([]*model.Vote)(nil), errorTest)
	voteRepoMockErr.On("FindUserVoteByUserID", mock.Anything, &userID).Return(([]*model.UserVote)(nil), errorTest)
	type fields struct {
		UserRepo      repository.UserRepository
		VoteRepo      repository.VoteRepository
		UserRedisRepo repository.UserRedisRepository
		VoteRedisRepo repository.VoteRedisRepository
	}
	type args struct {
		ctx     context.Context
		userID  *uuid.UUID
		voterID *uuid.UUID
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantVote     *model.Vote
		wantUserVote *model.UserVote
		wantErr      bool
	}{
		{
			"find exist voting user",
			fields{UserRepo: userRepoMock, VoteRepo: voteRepoMock},
			args{ctx: context.TODO(), userID: &userID, voterID: &voterID},
			votes[0],
			userVotes[0],
			false,
		},
		{
			"find exist voting user nil",
			fields{UserRepo: userRepoMock, VoteRepo: voteRepoMockNil},
			args{ctx: context.TODO(), userID: &userID3, voterID: &voterID3},
			(*model.Vote)(nil),
			(*model.UserVote)(nil),
			false,
		},
		{
			"find exist voting user error",
			fields{UserRepo: userRepoMock, VoteRepo: voteRepoMockErr},
			args{ctx: context.TODO(), userID: &userID, voterID: &voterID},
			(*model.Vote)(nil),
			(*model.UserVote)(nil),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userusecase := NewUserUsecase(tt.fields.UserRepo, tt.fields.VoteRepo, tt.fields.UserRedisRepo, tt.fields.VoteRedisRepo)
			gotVote, gotUserVote, err := userusecase.FindExistVoting(tt.args.ctx, tt.args.userID, tt.args.voterID)
			assert.DeepEqual(t, gotVote, tt.wantVote)
			assert.DeepEqual(t, gotUserVote, tt.wantUserVote)
			assert.Equal(t, !(err == nil), tt.wantErr)
		})
	}
}

func TestUserUsecase_FindVotesForUser(t *testing.T) {
	userRepoMock := &UserRepositoryMock{}
	voteRepoMock := &VoteRepositoryMock{}
	voteRepoMockNil := &VoteRepositoryMock{}
	voteRepoMockErr := &VoteRepositoryMock{}
	userID := uuid.New()
	timeNow := time.Now()
	votes := make([]*model.Vote, 0)
	votes = append(votes, &model.Vote{VoteID: 1, Vote: 1, CreatedUserID: userID, CreatedAt: &timeNow})
	votes = append(votes, &model.Vote{VoteID: 2, Vote: -1, CreatedUserID: userID, CreatedAt: &timeNow})
	votes = append(votes, &model.Vote{VoteID: 3, Vote: 1, CreatedUserID: userID, CreatedAt: &timeNow})

	votesNil := make([]*model.Vote, 0)
	errorTest := fmt.Errorf("something went wrong")

	voteRepoMock.On("FindVoteByUserID", mock.Anything, &userID).Return(votes, nil)
	voteRepoMockNil.On("FindVoteByUserID", mock.Anything, &userID).Return(votesNil, nil)
	voteRepoMockErr.On("FindVoteByUserID", mock.Anything, &userID).Return(([]*model.Vote)(nil), apperrors.UserUsecaseFindVotesForUser.AppendMessage(errorTest))
	type fields struct {
		UserRepo      repository.UserRepository
		VoteRepo      repository.VoteRepository
		UserRedisRepo repository.UserRedisRepository
		VoteRedisRepo repository.VoteRedisRepository
	}
	type args struct {
		ctx    context.Context
		userID *uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Vote
		wantErr bool
	}{
		{"find votes for user", fields{UserRepo: userRepoMock, VoteRepo: voteRepoMock}, args{ctx: context.TODO(), userID: &userID}, votes, false},
		{"find votes for user nil", fields{UserRepo: userRepoMock, VoteRepo: voteRepoMockNil}, args{ctx: context.TODO(), userID: &userID}, votesNil, false},
		{"find votes for user error", fields{UserRepo: userRepoMock, VoteRepo: voteRepoMockErr}, args{ctx: context.TODO(), userID: &userID}, ([]*model.Vote)(nil), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userusecase := NewUserUsecase(tt.fields.UserRepo, tt.fields.VoteRepo, tt.fields.UserRedisRepo, tt.fields.VoteRedisRepo)
			got, err := userusecase.FindVotesForUser(tt.args.ctx, tt.args.userID)
			assert.DeepEqual(t, got, tt.want)
			assert.Equal(t, !(err == nil), tt.wantErr)
		})
	}
}
