package model

import (
	"time"

	"usermanager/internal/apperrors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	superuserAdmin = "admin"
	superuserRoot  = "root"
	UsualUser      = "user"
	TableName      = "users"
)

type User struct {
	UserID    uuid.UUID `json:"user_id" db:"user_id" validate:"omitempty"`
	Nickname  string    `json:"nickname" db:"nickname" validate:"required"`
	FirstName string    `json:"first_name" db:"first_name" validate:"required"`
	LastName  string    `json:"last_name" db:"last_name" validate:"required"`
	Email     string    `json:"email,omitempty" db:"email" redis:"email" validate:"email"`
	Password  string    `json:"password,omitempty" db:"password" validate:"omitempty,required,gte=6"`
	IsPublic  bool      `json:"is_public,omitempty" db:"is_public" validate:"omitempty"`
	Role      string    `json:"user_role" db:"user_role" validate:"required"`
	Created
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	LoginDate *time.Time `json:"login_date,omitempty" db:"login_date"`
	Votes     []*Vote    `json:"votes,omitempty" db:"votes"`
}

type Created struct {
	By string    `json:"created_by,omitempty" db:"created_by" validate:"omitempty"`
	At time.Time `json:"created_at,omitempty" db:"created_at" validate:"omitempty"`
}

type Users struct {
	Page    int     `json:"page"`
	HasMore bool    `json:"has_more"`
	Users   []*User `json:"users"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return apperrors.UserHashPasswordGenerateFromPassword.AppendMessage(err)
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return apperrors.UserComparePasswordsCompareHashAndPassword.AppendMessage(err)
	}
	return nil
}

func (u *User) MapCreateUserRequestToUserModel(req *CreateUserRequest) {
	u.UserID = req.UserID
	u.Nickname = req.Nickname
	u.FirstName = req.FirstName
	u.LastName = req.LastName
	u.Email = req.Email
	u.Password = req.Password
	u.IsPublic = req.IsPublic
	u.Role = req.Role
}

func (u *User) MapUserModelToCreateUserResponse() *CreateUserResponse {
	createUserResponse := &CreateUserResponse{}
	createUserResponse.UserID = u.UserID
	createUserResponse.Nickname = u.Nickname
	createUserResponse.FirstName = u.FirstName
	createUserResponse.LastName = u.LastName
	createUserResponse.Email = u.Email
	createUserResponse.IsPublic = u.IsPublic
	createUserResponse.Role = u.Role

	return createUserResponse
}

func (u *User) MapUpdateUserRequestToUserModel(req *UpdateUserRequest) {
	u.Nickname = req.Nickname
	u.FirstName = req.FirstName
	u.LastName = req.LastName
	u.Email = req.Email
	u.Password = req.Password
	u.IsPublic = req.IsPublic
	u.Role = req.Role
}

func (u *User) MapUserModelToUpdateUserResponse() *UpdateUserResponse {
	updateUserResponse := &UpdateUserResponse{}
	updateUserResponse.UserID = u.UserID
	updateUserResponse.Nickname = u.Nickname
	updateUserResponse.FirstName = u.FirstName
	updateUserResponse.LastName = u.LastName
	updateUserResponse.Email = u.Email
	updateUserResponse.IsPublic = u.IsPublic
	updateUserResponse.Role = u.Role

	return updateUserResponse
}

func (u *User) MapUserModelToGetUserResponse() *GetUserResponse {
	GetUserResponse := &GetUserResponse{}
	GetUserResponse.UserID = u.UserID
	GetUserResponse.Nickname = u.Nickname
	GetUserResponse.FirstName = u.FirstName
	GetUserResponse.LastName = u.LastName
	GetUserResponse.IsPublic = u.IsPublic
	GetUserResponse.Role = u.Role
	if len(u.Votes) > 0 {
		rate := 0
		for _, vote := range u.Votes {
			rate += vote.Vote
		}
		GetUserResponse.Rate = rate
	} else {
		GetUserResponse.Rate = 0
	}
	return GetUserResponse
}

func (us *Users) MapUserModelToGetUserResponse() *GetUsersResponse {
	getUsersResponse := &GetUsersResponse{}
	getUsersResponse.HasMore = us.HasMore
	getUsersResponse.Page = us.Page
	usersResponse := make([]*GetUserResponse, 0, len(us.Users))
	for _, userValue := range us.Users {
		usersResponse = append(usersResponse, userValue.MapUserModelToGetUserResponse())
	}
	getUsersResponse.Users = usersResponse
	return getUsersResponse
}

func (u *User) TableName() string {
	return TableName
}

func MapUsersToIDs(users *Users) []*uuid.UUID {
	ids := make([]*uuid.UUID, 0, len(users.Users))
	for _, userValue := range users.Users {
		ids = append(ids, &userValue.UserID)
	}
	return ids
}

func MapToVotesByUserID(votes []*Vote) map[string][]*Vote {
	userIDToVotes := make(map[string][]*Vote)
	for _, vote := range votes {
		userIDToVotes[vote.CreatedUserID.String()] = append(userIDToVotes[vote.CreatedUserID.String()], vote)
	}

	return userIDToVotes
}
