package repository

import (
	"context"
	"database/sql"
	"time"

	"usermanager/internal/apperrors"
	"usermanager/internal/domain/model"
	"usermanager/internal/infrastructure/datastore"
	"usermanager/internal/utils"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindUserByUUID(ctx context.Context, userID uuid.UUID) (*model.User, error)
	FindUserByNickname(ctx context.Context, nickname string) (*model.User, error)
	GetUsers(ctx context.Context, paginationQuery *utils.PaginationQuery) (*model.Users, error)
	SaveUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	SoftDeleteUserByUserID(ctx context.Context, userID uuid.UUID) (*model.User, error)
	DeleteUserByUserID(ctx context.Context, userID *uuid.UUID) error
}

type userRepo struct {
	db *datastore.DB
}

func NewUserRepository(db *datastore.DB) UserRepository {
	return &userRepo{db: db}
}

func (u *userRepo) FindUserByUUID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	user := &model.User{}
	err := u.db.SQL.GetContext(ctx, user, getUserByID, userID)
	if err != nil {
		if sql.ErrNoRows == err {
			return nil, apperrors.UserRepoFindUserByUUIDGetDataNotFound.AppendMessage(err)
		}
		return nil, apperrors.UserRepoFindUserByUUIDGetContext.AppendMessage(err)
	}

	return user, nil
}

func (u *userRepo) FindUserByNickname(ctx context.Context, nickname string) (*model.User, error) {
	existingUser := &model.User{}
	if err := u.db.SQL.GetContext(ctx, existingUser, getUserByNickname, nickname); err != nil {
		if sql.ErrNoRows == err {
			return nil, apperrors.UserRepoFindUserByNicknameGetDataNotFound.AppendMessage(err)
		}
		return nil, apperrors.UserRepoFindUserByNicknameGetContext.AppendMessage(err)
	}

	return existingUser, nil
}

func (u *userRepo) SaveUser(ctx context.Context, user *model.User) (*model.User, error) {
	err := u.db.SQL.QueryRowxContext(
		ctx,
		addUser,
		&user.UserID,
		&user.Nickname,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.IsPublic,
		&user.Role,
		&user.Created.At,
		&user.UpdatedAt,
		&user.DeletedAt,
		&user.LoginDate,
		&user.Created.By,
	).StructScan(user)
	if err != nil && sql.ErrNoRows != err {
		return nil, apperrors.UserRepoSaveUserQueryRowxContext.AppendMessage(err)
	}
	return user, nil
}

func (u *userRepo) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	err := u.db.SQL.QueryRowxContext(
		ctx,
		updateUser,
		&user.Nickname,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.IsPublic,
		&user.UpdatedAt,
		&user.LoginDate,
		&user.UserID,
	).StructScan(user)
	if err != nil && sql.ErrNoRows != err {
		return nil, apperrors.UserRepoUpdateUserQueryRowxContext.AppendMessage(err)
	}
	return user, nil
}

func (u *userRepo) SoftDeleteUserByUserID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	existingUser := &model.User{}
	deletedAt := time.Now()
	err := u.db.SQL.QueryRowxContext(
		ctx,
		updateDeletedAt,
		deletedAt,
		userID,
	).StructScan(&existingUser)
	if err != nil {
		return nil, apperrors.UserRepoSoftDeleteUserByUserIDQueryRowxContext.AppendMessage(err)
	}
	return existingUser, nil
}

func (u *userRepo) DeleteUserByUserID(ctx context.Context, userID *uuid.UUID) error {
	result, err := u.db.SQL.ExecContext(ctx, deleteUserFromDb, userID)
	if err != nil {
		return apperrors.UserRepoDeleteUserByUserIDExecContext.AppendMessage(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return apperrors.UserRepoDeleteUserByUserIDRowsAffected.AppendMessage(err)
	}
	if rowsAffected == 0 {
		return apperrors.UserRepoDeleteUserByUserIDEmptyRowsAffected.AppendMessage(err)
	}
	return nil
}

func (u *userRepo) GetUsers(ctx context.Context, paginationQuery *utils.PaginationQuery) (*model.Users, error) {
	usersList := &model.Users{
		Page:    paginationQuery.GetPage(),
		HasMore: false,
		Users:   make([]*model.User, 0),
	}

	users := make([]*model.User, 0, paginationQuery.GetSize())

	usersLimit := paginationQuery.GetLimit()
	queryLimit := paginationQuery.GetLimit() + 1
	rows, err := u.db.SQL.QueryxContext(ctx, getUsers, paginationQuery.GetOffset(), queryLimit)
	if err != nil {
		return nil, apperrors.UserRepoGetUsersQueryxContext.AppendMessage(err)
	}
	defer rows.Close()

	userCounter := 0
	for rows.Next() {
		user := &model.User{}

		if err = rows.StructScan(user); err != nil {
			return nil, apperrors.UserRepoGetUsersStructScan.AppendMessage(err)
		}

		users = append(users, user)
		userCounter++
		if userCounter == usersLimit {
			usersList.HasMore = true
			break
		}
	}

	if err = rows.Err(); err != nil {
		return nil, apperrors.UserRepoGetUsersRows.AppendMessage(err)
	}

	usersList.Users = users
	return usersList, nil
}
