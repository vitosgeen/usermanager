package model

import "github.com/google/uuid"

type LoginRequest struct {
	Nickname string `form:"nickname" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type CreateUserRequest struct {
	UserID    uuid.UUID `json:"user_id" db:"user_id" validate:"omitempty"`
	Nickname  string    `json:"nickname" db:"nickname" validate:"required"`
	FirstName string    `json:"first_name" db:"first_name" validate:"required"`
	LastName  string    `json:"last_name" db:"last_name" validate:"required"`
	Email     string    `json:"email,omitempty" db:"email" redis:"email" validate:"email"`
	Password  string    `json:"password,omitempty" db:"password" validate:"omitempty,required,gte=6"`
	IsPublic  bool      `json:"is_public,omitempty" db:"is_public" validate:"omitempty"`
	Role      string    `json:"user_role" db:"user_role" validate:"required"`
}

type UpdateUserRequest struct {
	Nickname  string `json:"nickname" db:"nickname" validate:"required"`
	FirstName string `json:"first_name" db:"first_name" validate:"required"`
	LastName  string `json:"last_name" db:"last_name" validate:"required"`
	Email     string `json:"email,omitempty" db:"email" redis:"email" validate:"email"`
	Password  string `json:"password,omitempty" db:"password" validate:"omitempty,required,gte=6"`
	IsPublic  bool   `json:"is_public,omitempty" db:"is_public" validate:"omitempty"`
	Role      string `json:"user_role" db:"user_role" validate:"required"`
}

type VoteUserRequest struct {
	UserID uuid.UUID `json:"user_id" validate:"required" valid:"-"`
	Vote   int       `json:"vote" validate:"required" valid:"-"`
}
