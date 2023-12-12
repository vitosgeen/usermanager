package repository

import (
	"context"
	"database/sql"
	"time"

	"usermanager/internal/apperrors"
	"usermanager/internal/domain/model"
	"usermanager/internal/infrastructure/datastore"

	"github.com/google/uuid"
)

type VoteRepository interface {
	SaveVote(ctx context.Context, vote *model.Vote) (*model.Vote, error)
	UpdateVote(ctx context.Context, vote *model.Vote) (*model.Vote, error)
	FindVoteByID(ctx context.Context, voteID int64) (*model.Vote, error)
	FindUserVoteByID(ctx context.Context, id int64) (*model.UserVote, error)
	FindVoteByUserID(ctx context.Context, voteUserID *uuid.UUID) ([]*model.Vote, error)
	FindVotesByUserIDs(ctx context.Context, userIDs []*uuid.UUID) ([]*model.Vote, error)
	FindUserVoteByUserID(ctx context.Context, userID *uuid.UUID) ([]*model.UserVote, error)
	SaveUserVote(ctx context.Context, userVote *model.UserVote) (*model.UserVote, error)
	DeleteUserVote(ctx context.Context, userVote *model.UserVote) error
	DeleteVote(ctx context.Context, userVote *model.Vote) error
}
type voteRepo struct {
	db *datastore.DB
}

func NewVoteRepository(db *datastore.DB) VoteRepository {
	return &voteRepo{db: db}
}

func (v *voteRepo) SaveVote(ctx context.Context, vote *model.Vote) (*model.Vote, error) {
	timeNow := time.Now()
	vote.CreatedAt = &timeNow
	err := v.db.SQL.QueryRowxContext(ctx, addVote, &vote.Vote, &vote.CreatedUserID, &vote.CreatedAt).StructScan(vote)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.VoteRepoSaveVoteQueryRowxContextDataNotFound.AppendMessage(err)
		}

		return nil, apperrors.VoteRepoSaveVoteQueryRowxContext.AppendMessage(err)
	}
	return vote, nil
}

func (v *voteRepo) UpdateVote(ctx context.Context, vote *model.Vote) (*model.Vote, error) {
	err := v.db.SQL.QueryRowxContext(ctx, updateVote, &vote.Vote, &vote.VoteID).StructScan(vote)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.VoteRepoUpdateVoteQueryRowxContextDataNotFound.AppendMessage(err)
		}
		return nil, apperrors.VoteRepoUpdateVoteQueryRowxContext.AppendMessage(err)
	}
	return vote, nil
}

func (v *voteRepo) FindVoteByID(ctx context.Context, voteID int64) (*model.Vote, error) {
	if voteID == 0 {
		return nil, apperrors.VoteRepoFindVoteByIDVoteIDEmpty.AppendMessage(nil)
	}
	vote := &model.Vote{}
	err := v.db.SQL.GetContext(ctx, vote, getVoteByID, voteID)
	if err != nil {
		if sql.ErrNoRows == err {
			return nil, apperrors.VoteRepoFindVoteByIDQueryxContextDataNotFound.AppendMessage(err)
		}
		return nil, apperrors.VoteRepoFindVoteByIDGetContext.AppendMessage(err)
	}
	return vote, nil
}

func (v *voteRepo) FindUserVoteByID(ctx context.Context, id int64) (*model.UserVote, error) {
	if id == 0 {
		return nil, apperrors.VoteRepoFindUserVoteByIDVoteIDEmpty.AppendMessage(nil)
	}
	userVote := &model.UserVote{}
	err := v.db.SQL.GetContext(ctx, userVote, getUserVoteByID, id)
	if err != nil {
		if sql.ErrNoRows == err {
			return nil, apperrors.VoteRepoFindUserVoteByIDQueryxContextDataNotFound.AppendMessage(err)
		}
		return nil, apperrors.VoteRepoFindUserVoteByIDGetContext.AppendMessage(err)
	}
	return userVote, nil
}

func (v *voteRepo) FindVoteByUserID(ctx context.Context, userID *uuid.UUID) ([]*model.Vote, error) {
	rows, err := v.db.SQL.QueryxContext(ctx, getVotesByUserID, userID)
	if err != nil {
		if sql.ErrNoRows == err {
			return nil, apperrors.VoteRepoFindVoteByUserIDQueryxContextDataNotFound.AppendMessage(err)
		}
		return nil, apperrors.VoteRepoFindVoteByUserIDGetContext.AppendMessage(err)
	}

	votes := make([]*model.Vote, 0)
	for rows.Next() {
		vote := &model.Vote{}
		if err = rows.StructScan(vote); err != nil {
			return nil, apperrors.VoteRepoFindVoteByUserIDStructScan.AppendMessage(err)
		}
		votes = append(votes, vote)
	}
	return votes, nil
}

func (v *voteRepo) FindVotesByUserIDs(ctx context.Context, userIDs []*uuid.UUID) ([]*model.Vote, error) {
	rows, err := v.db.SQL.QueryxContext(ctx, getVotesByUserIDs, userIDs)
	if err != nil {
		if sql.ErrNoRows == err {
			return nil, apperrors.VoteRepoFindVotesByUserIDsQueryxContextDataNotFound.AppendMessage(err)
		}
		return nil, apperrors.VoteRepoFindVotesByUserIDsQueryxContext.AppendMessage(err)
	}

	votes := make([]*model.Vote, 0)
	for rows.Next() {
		vote := &model.Vote{}
		if err = rows.StructScan(vote); err != nil {
			return nil, apperrors.VoteRepoFindVotesByUserIDsStructScan.AppendMessage(err)
		}
		votes = append(votes, vote)
	}
	return votes, nil
}

func (v *voteRepo) FindUserVoteByUserID(ctx context.Context, userID *uuid.UUID) ([]*model.UserVote, error) {
	rows, err := v.db.SQL.QueryxContext(ctx, getUserVotesByUserID, userID)
	if err != nil {
		if sql.ErrNoRows == err {
			return nil, apperrors.VoteRepoFindUserVoteByUserIDQueryxContextDataNotFound.AppendMessage(err)
		}
		return nil, apperrors.VoteRepoFindUserVoteByUserIDGetContext.AppendMessage(err)
	}

	userVotes := make([]*model.UserVote, 0)
	for rows.Next() {
		userVote := &model.UserVote{}
		err = rows.StructScan(userVote)
		if err != nil {
			return nil, apperrors.VoteRepoFindVoteByUserIDStructScan.AppendMessage(err)
		}
		userVotes = append(userVotes, userVote)
	}
	return userVotes, nil
}

func (v *voteRepo) SaveUserVote(ctx context.Context, userVote *model.UserVote) (*model.UserVote, error) {
	err := v.db.SQL.QueryRowxContext(ctx, addUserVote, &userVote.UserID, &userVote.VoteID).StructScan(userVote)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.VoteRepoSaveUserVoteQueryRowxContextDataNotFound.AppendMessage(err)
		}
		return nil, apperrors.VoteRepoSaveUserVoteQueryRowxContext.AppendMessage(err)
	}
	return userVote, nil
}

func (v *voteRepo) DeleteUserVote(ctx context.Context, userVote *model.UserVote) error {
	result, err := v.db.SQL.ExecContext(ctx, deleteUserVote, userVote.ID)
	if err != nil {
		return apperrors.VoteRepoDeleteVoteExecContext.AppendMessage(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return apperrors.VoteRepoDeleteVoteRowsAffected.AppendMessage(err)
	}
	if rowsAffected == 0 {
		return apperrors.VoteRepoDeleteVoteEmptyRowsAffected.AppendMessage(err)
	}
	return nil
}

func (v *voteRepo) DeleteVote(ctx context.Context, vote *model.Vote) error {
	result, err := v.db.SQL.ExecContext(ctx, deleteVote, vote.VoteID)
	if err != nil {
		return apperrors.VoteRepoDeleteVoteExecContext.AppendMessage(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return apperrors.VoteRepoDeleteVoteRowsAffected.AppendMessage(err)
	}
	if rowsAffected == 0 {
		return apperrors.VoteRepoDeleteVoteEmptyRowsAffected.AppendMessage(err)
	}
	return nil
}
