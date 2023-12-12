package usecase

import (
	"context"
	"fmt"
	"time"

	"usermanager/internal/apperrors"
	"usermanager/internal/domain/model"
	"usermanager/internal/interface/repository"
	"usermanager/internal/utils"

	"github.com/google/uuid"
)

const (
	votePositive = 1
	voteNegative = -1
	voteWithdraw = 0
	voteInterval = 3600
)

type IUserUsecase interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, userID *uuid.UUID) error
	GetUsers(ctx context.Context, paginationQuery *utils.PaginationQuery) (*model.Users, error)
	GetUsersByPaginationQuery(ctx context.Context, paginationQuery *utils.PaginationQuery) (*model.Users, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error)
	GetUserByNickname(ctx context.Context, nickname string) (*model.User, error)
	CheckUserByNickname(ctx context.Context, user *model.User) (bool, error)
	VoteUser(ctx context.Context, vote *model.Vote, userVote *model.UserVote) (*model.Vote, *model.UserVote, error)
	VoteUserWithdraw(ctx context.Context, vote *model.Vote, userVote *model.UserVote) (*model.Vote, *model.UserVote, error)
	FindExistVoting(ctx context.Context, userID *uuid.UUID, voterUserID *uuid.UUID) (*model.Vote, *model.UserVote, error)
	FindVotesForUser(ctx context.Context, userID *uuid.UUID) ([]*model.Vote, error)
	LoadVotesToUsers(ctx context.Context, users *model.Users) ([]*model.User, error)
	GetLastVoteForUser(ctx context.Context, userID *uuid.UUID) (*model.Vote, error)
	Vote(ctx context.Context, vote *model.Vote, userVote *model.UserVote) (*model.Vote, *model.UserVote, error)
}

type UserUsecase struct {
	UserRepo      repository.UserRepository
	UserRedisRepo repository.UserRedisRepository
	VoteRepo      repository.VoteRepository
	VoteRedisRepo repository.VoteRedisRepository
}

func NewUserUsecase(userRepo repository.UserRepository, voteRepo repository.VoteRepository, userRedisRepo repository.UserRedisRepository, voteRedisRepo repository.VoteRedisRepository) IUserUsecase {
	return &UserUsecase{
		UserRepo:      userRepo,
		VoteRepo:      voteRepo,
		UserRedisRepo: userRedisRepo,
		VoteRedisRepo: voteRedisRepo,
	}
}

func (us *UserUsecase) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	user.Created.At = time.Now()
	user.UserID = uuid.New()
	user.Role = user.GetDefaultRole()
	err := user.HashPassword()
	if err != nil {
		return nil, apperrors.UserUsecaseCreateUserHashPassword.AppendMessage(err)
	}

	existingUser, err := us.UserRepo.FindUserByNickname(ctx, user.Nickname)
	if err != nil {
		appError := err.(*apperrors.AppError)
		if appError.Code != apperrors.UserRepoFindUserByNicknameGetDataNotFound.Code {
			return nil, apperrors.UserUsecaseCreateUserFindUserByNickname.AppendMessage(err)
		}
	}
	if existingUser != nil {
		return existingUser, apperrors.UserUsecaseCreateUserUserExists.AppendMessage(err)
	}

	savedUser, err := us.UserRepo.SaveUser(ctx, user)
	if err != nil {
		return nil, apperrors.UserUsecaseCreateUserSaveUser.AppendMessage(err)
	}

	return savedUser, nil
}

func (us *UserUsecase) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	updatedUser, err := us.UserRepo.UpdateUser(ctx, user)
	if err != nil {
		return nil, apperrors.UserUsecaseUpdateUserUpdateUser.AppendMessage(err)
	}

	return updatedUser, nil
}

func (us *UserUsecase) GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	user, err := us.UserRedisRepo.FindUserByUUID(ctx, userID)
	if err != nil {
		appError := err.(*apperrors.AppError)
		if appError.Code != apperrors.UserRedisRepoFindUserByUUIDGetDataNotFound.Code {
			return nil, apperrors.UserUsecaseGetUserFindUserByUUID.AppendMessage(err)
		}
	}
	if user != nil {
		return user, nil
	}

	user, err = us.UserRepo.FindUserByUUID(ctx, userID)
	if err != nil {
		fmt.Println("TestUserUsecase_GetUser_Error ERROR", err)
		appError := err.(*apperrors.AppError)
		if appError.Code != apperrors.UserRepoFindUserByUUIDGetDataNotFound.Code {
			return nil, apperrors.UserUsecaseGetUserLoad.AppendMessage(err)
		}
	}

	err = us.UserRedisRepo.SetFindUserByUUID(ctx, userID, user)
	if err != nil {
		return nil, apperrors.UserUsecaseGetUserSetFindUserByUUID.AppendMessage(err)
	}

	return user, nil
}

func (us *UserUsecase) GetUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	user, err := us.GetUser(ctx, userID)
	if err != nil {
		return nil, apperrors.UserUsecaseGetUserByIDGetUser.AppendMessage(err)
	}

	user.Votes, err = us.FindVotesForUser(ctx, &user.UserID)
	if err != nil {
		return nil, apperrors.UserUsecaseGetUserByIDFindVotesForUser.AppendMessage(err)
	}

	return user, nil
}

func (us *UserUsecase) DeleteUser(ctx context.Context, userID *uuid.UUID) error {
	err := us.UserRepo.DeleteUserByUserID(ctx, userID)
	if err != nil {
		return apperrors.UserUsecaseDeleteUser.AppendMessage(err)
	}

	return nil
}

func (us *UserUsecase) CheckUserByNickname(ctx context.Context, user *model.User) (bool, error) {
	checkedUser, err := us.UserRepo.FindUserByNickname(ctx, user.Nickname)
	if err != nil {
		appError := err.(*apperrors.AppError)
		if appError.Code != apperrors.UserRepoFindUserByNicknameGetDataNotFound.Code {
			return false, apperrors.UserUsecaseCheckProfileByNick.AppendMessage(err)
		}
	}
	if checkedUser != nil && checkedUser.UserID != user.UserID {
		return false, apperrors.UserUsecaseCheckProfileByNickBusy.AppendMessage(err)
	}

	return true, nil
}

func (us *UserUsecase) GetUserByNickname(ctx context.Context, nickname string) (*model.User, error) {
	user, err := us.UserRedisRepo.FindUserByNickname(ctx, nickname)
	if err != nil {
		return nil, apperrors.UserUsecaseGetUserByNicknameUserRedisRepoFindUserByNickname.AppendMessage(err)
	}
	if user != nil {
		return user, nil
	}

	user, err = us.UserRepo.FindUserByNickname(ctx, nickname)
	if err != nil {
		return nil, apperrors.UserUsecaseGetUserByNickname.AppendMessage(err)
	}

	err = us.UserRedisRepo.SetFindUserByNickname(ctx, nickname, user)
	if err != nil {
		return nil, apperrors.UserUsecaseGetUserByNicknameUserRedisRepoSetFindUserByNickname.AppendMessage(err)
	}

	return user, nil
}

func (us *UserUsecase) GetUsers(ctx context.Context, paginationQuery *utils.PaginationQuery) (*model.Users, error) {
	users, err := us.UserRedisRepo.GetUsers(ctx, paginationQuery)
	if err != nil {
		appError := err.(*apperrors.AppError)
		if appError.Code != apperrors.UserRedisRepoGetUsersGetDataNotFound.Code {
			return nil, apperrors.UserUsecaseGetUsersUserRedisRepoGetUsers.AppendMessage(err)
		}
	}
	if users != nil {
		return users, nil
	}

	users, err = us.UserRepo.GetUsers(ctx, paginationQuery)
	if err != nil {
		return nil, apperrors.UserUsecaseGetUsers.AppendMessage(err)
	}

	err = us.UserRedisRepo.SetGetUsers(ctx, paginationQuery, users)
	if err != nil {
		return nil, apperrors.UserUsecaseGetUsersUserRedisRepoSetGetUsers.AppendMessage(err)
	}

	return users, nil
}

func (us *UserUsecase) GetUsersByPaginationQuery(ctx context.Context, paginationQuery *utils.PaginationQuery) (*model.Users, error) {
	users, err := us.GetUsers(ctx, paginationQuery)
	if err != nil {
		return nil, apperrors.UserUsecaseGetUsers.AppendMessage(err)
	}

	users.Users, err = us.LoadVotesToUsers(ctx, users)
	if err != nil {
		return nil, apperrors.UserUsecaseGetUsers.AppendMessage(err)
	}

	return users, nil
}

func (us *UserUsecase) VoteUser(ctx context.Context, vote *model.Vote, userVote *model.UserVote) (*model.Vote, *model.UserVote, error) {
	voteExist, userVoteExist, err := us.FindExistVoting(ctx, &userVote.UserID, &vote.CreatedUserID)
	if err != nil {
		return nil, nil, apperrors.UserUsecaseVoteUserFindExistVoting.AppendMessage(err)
	}
	if voteExist != nil && userVoteExist != nil {
		if vote.Vote == voteExist.Vote {
			return nil, nil, apperrors.UserUsecaseVoteUserVotingExist.AppendMessage(err)
		}
		voteExist.Vote = vote.Vote
		updatedVote, err := us.VoteRepo.UpdateVote(ctx, voteExist)
		if err != nil {
			return nil, nil, apperrors.UserUsecaseVoteUserUpdateVote.AppendMessage(err)
		}

		return updatedVote, userVoteExist, nil
	} else {
		savedVote, err := us.VoteRepo.SaveVote(ctx, vote)
		if err != nil {
			return nil, nil, apperrors.UserUsecaseVoteUserSaveVote.AppendMessage(err)
		}
		userVote.VoteID = savedVote.VoteID
		savedUserVote, err := us.VoteRepo.SaveUserVote(ctx, userVote)
		if err != nil {
			return nil, nil, apperrors.UserUsecaseVoteUserSaveUserVote.AppendMessage(err)
		}

		return savedVote, savedUserVote, nil
	}
}

func (us *UserUsecase) VoteUserWithdraw(ctx context.Context, vote *model.Vote, userVote *model.UserVote) (*model.Vote, *model.UserVote, error) {
	voteExist, userVoteExist, err := us.FindExistVoting(ctx, &userVote.UserID, &vote.CreatedUserID)
	if err != nil {
		return nil, nil, apperrors.UserUsecaseVoteUserWithdrawFindExistVoting.AppendMessage(err)
	}
	if voteExist != nil && userVoteExist != nil {
		err := us.VoteRepo.DeleteVote(ctx, voteExist)
		if err != nil {
			return nil, nil, apperrors.UserUsecaseVoteUserWithdrawDeleteVote.AppendMessage(err)
		}

		err = us.VoteRepo.DeleteUserVote(ctx, userVoteExist)
		if err != nil {
			return nil, nil, apperrors.UserUsecaseVoteUserWithdrawDeleteUserVote.AppendMessage(err)
		}

	} else {
		return nil, nil, apperrors.UserUsecaseVoteUserWithdrawVoteNotExist.AppendMessage(err)
	}
	return voteExist, userVoteExist, nil
}

func (us *UserUsecase) FindExistVoting(ctx context.Context, userID *uuid.UUID, voterUserID *uuid.UUID) (*model.Vote, *model.UserVote, error) {
	votes, err := us.VoteRepo.FindVoteByUserID(ctx, voterUserID)
	if err != nil {
		return nil, nil, apperrors.UserUsecaseVoteIsExistFindVoteByUserID.AppendMessage(err)
	}
	if votes == nil {
		return nil, nil, nil
	}

	userVotes, err := us.VoteRepo.FindUserVoteByUserID(ctx, userID)
	if err != nil {
		return nil, nil, apperrors.UserUsecaseVoteIsExistFindUserVoteByUserID.AppendMessage(err)
	}
	if userVotes == nil {
		return nil, nil, nil
	}
	for _, vote := range votes {
		for _, userVote := range userVotes {
			if vote.VoteID == userVote.VoteID {
				return vote, userVote, nil
			}
		}
	}

	return nil, nil, nil
}

func (us *UserUsecase) FindVotesForUser(ctx context.Context, userID *uuid.UUID) ([]*model.Vote, error) {
	votes, err := us.VoteRepo.FindVoteByUserID(ctx, userID)
	if err != nil {
		appError := err.(*apperrors.AppError)
		if appError.Code != apperrors.VoteRepoFindVoteByUserIDQueryxContextDataNotFound.Code {
			return nil, apperrors.UserUsecaseFindVotesForUser.AppendMessage(err)
		}
	}

	return votes, nil
}

func (us *UserUsecase) LoadVotesToUsers(ctx context.Context, users *model.Users) ([]*model.User, error) {
	userIDs := model.MapUsersToIDs(users)
	votes, err := us.VoteRepo.FindVotesByUserIDs(ctx, userIDs)
	if err != nil {
		return nil, apperrors.UserUsecaseLoadVotesToUsers.AppendMessage(err)
	}

	userIDToVotes := model.MapToVotesByUserID(votes)
	return appendVotesToUsers(users, userIDToVotes)
}

func appendVotesToUsers(users *model.Users, userIDToVotes map[string][]*model.Vote) ([]*model.User, error) {
	usersWithVotes := []*model.User{}
	for _, user := range users.Users {
		usersWithVotes = append(usersWithVotes, user)
		votes, ok := userIDToVotes[user.UserID.String()]
		if !ok {
			continue
		}
		user.Votes = votes
	}

	return usersWithVotes, nil
}

func (us *UserUsecase) GetLastVoteForUser(ctx context.Context, userID *uuid.UUID) (*model.Vote, error) {
	votes, err := us.VoteRepo.FindVoteByUserID(ctx, userID)
	if err != nil {
		return nil, apperrors.UserUsecaseFindVotesForUser.AppendMessage(err)
	}
	if votes != nil {
		return votes[0], nil
	}

	return nil, nil
}

func (us *UserUsecase) Vote(ctx context.Context, vote *model.Vote, userVote *model.UserVote) (*model.Vote, *model.UserVote, error) {
	switch vote.Vote {
	case votePositive:
		vote, userVote, err := us.VoteUser(ctx, vote, userVote)
		if err != nil {
			return nil, nil, apperrors.UserUsecaseVotePositiveVoteUser.AppendMessage(err)
		}
		return vote, userVote, nil
	case voteNegative:
		vote, userVote, err := us.VoteUser(ctx, vote, userVote)
		if err != nil {
			return nil, nil, apperrors.UserUsecaseVoteNegativeVoteUser.AppendMessage(err)
		}
		return vote, userVote, nil
	case voteWithdraw:
		vote, userVote, err := us.VoteUserWithdraw(ctx, vote, userVote)
		if err != nil {
			return nil, nil, apperrors.UserUsecaseVoteWithdrawVoteUser.AppendMessage(err)
		}
		return vote, userVote, nil
	default:
		appError := apperrors.UserControllerVoteUserValueOfVoteIsNotRight
		return nil, nil, apperrors.UserControllerVoteUserValueOfVoteIsNotRight.AppendMessage(appError.Error())
	}
}
