package model

import "github.com/google/uuid"

type LoginResponse struct {
	Token string `form:"token" json:"token" binding:"required"`
}

type CreateUserResponse struct {
	UserID    uuid.UUID `json:"user_id" db:"user_id" validate:"omitempty"`
	Nickname  string    `json:"nickname" db:"nickname" validate:"required"`
	FirstName string    `json:"first_name" db:"first_name" validate:"required"`
	LastName  string    `json:"last_name" db:"last_name" validate:"required"`
	Email     string    `json:"email,omitempty" db:"email" redis:"email" validate:"email"`
	IsPublic  bool      `json:"is_public,omitempty" db:"is_public" validate:"omitempty"`
	Role      string    `json:"user_role" db:"user_role" validate:"required"`
}

type UpdateUserResponse struct {
	UserID    uuid.UUID `json:"user_id" db:"user_id" validate:"omitempty"`
	Nickname  string    `json:"nickname" db:"nickname" validate:"required"`
	FirstName string    `json:"first_name" db:"first_name" validate:"required"`
	LastName  string    `json:"last_name" db:"last_name" validate:"required"`
	Email     string    `json:"email,omitempty" db:"email" redis:"email" validate:"email"`
	IsPublic  bool      `json:"is_public,omitempty" db:"is_public" validate:"omitempty"`
	Role      string    `json:"user_role" db:"user_role" validate:"required"`
}

type GetUsersResponse struct {
	Page    int                `json:"page"`
	HasMore bool               `json:"has_more"`
	Users   []*GetUserResponse `json:"users"`
}

type GetUserResponse struct {
	UserID    uuid.UUID `json:"user_id" db:"user_id" validate:"omitempty"`
	Nickname  string    `json:"nickname" db:"nickname" validate:"required"`
	FirstName string    `json:"first_name" db:"first_name" validate:"required"`
	LastName  string    `json:"last_name" db:"last_name" validate:"required"`
	IsPublic  bool      `json:"is_public,omitempty" db:"is_public" validate:"omitempty"`
	Role      string    `json:"user_role" db:"user_role" validate:"required"`
	Rate      int       `json:"user_rate" db:"user_rate" validate:"required"`
}

type VoteUserResponse struct {
	VoteID      int64     `json:"vote_id" validate:"omitempty"`
	VoterUserID uuid.UUID `json:"voter_user_id" validate:"omitempty"`
	VoteUserID  uuid.UUID `json:"vote_user_id" validate:"omitempty"`
	Vote        int       `json:"vote" validate:"omitempty"`
}
