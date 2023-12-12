package model

import (
	"time"

	"github.com/google/uuid"
)

type Vote struct {
	VoteID        int64      `json:"vote_id" db:"vote_id" validate:"omitempty"`
	Vote          int        `json:"vote" db:"vote" validate:"omitempty"`
	CreatedUserID uuid.UUID  `json:"created_user_id" db:"created_user_id" validate:"omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty" db:"created_at" validate:"omitempty"`
}

type Votes struct {
	Page    int     `json:"page"`
	HasMore bool    `json:"has_more"`
	Votes   []*Vote `json:"votes"`
}

type UserVote struct {
	ID     int64     `json:"id" db:"id" validate:"omitempty"`
	UserID uuid.UUID `json:"user_id" db:"user_id" validate:"omitempty"`
	VoteID int64     `json:"vote_id" db:"vote_id" validate:"omitempty"`
}

func (uv *UserVote) MapUserAndVoteToUserVote(user *User, vote *Vote) {
	uv.UserID = user.UserID
	uv.VoteID = vote.VoteID
}

func (v *Vote) MapVoteUserRequestToVoteModel(req *VoteUserRequest, user *User) {
	v.Vote = req.Vote
	v.CreatedUserID = user.UserID
}

func (uv *UserVote) MapVoteUserVoteToVoteUserResponse(vote *Vote) *VoteUserResponse {
	voteUserResponse := &VoteUserResponse{}
	voteUserResponse.VoteID = vote.VoteID
	voteUserResponse.VoterUserID = vote.CreatedUserID
	voteUserResponse.VoteUserID = uv.UserID
	voteUserResponse.Vote = vote.Vote

	return voteUserResponse
}

func MapVoteUserRequestToVoteModel(req *VoteUserRequest, user *User) *Vote {
	vote := &Vote{}
	vote.Vote = req.Vote
	vote.CreatedUserID = user.UserID
	return vote
}

func MapVoteUserRequestToUserVoteModel(req *VoteUserRequest, vote *Vote) *UserVote {
	userVote := &UserVote{}
	userVote.UserID = req.UserID
	userVote.VoteID = vote.VoteID
	return userVote
}
