package model

import (
	"fmt"

	"usermanager/internal/apperrors"
)

const (
	RoleUser                          = "user"
	RoleModerator                     = "moderator"
	RoleAdmin                         = "admin"
	PermissionUpdate                  = "update"
	PermissionDelete                  = "delete"
	hasNoPermissionsToUpdateUserError = "auth user can't change this user"
	hasNoPermissionsToDeleteUserError = "auth user can't delete this user"
	hasNoPermissionsError             = "auth user can't delete this user"
)

func (u *User) GetDefaultRole() string {
	return RoleUser
}

func (u *User) GetRoles() []string {
	return []string{RoleUser, RoleModerator, RoleAdmin}
}

func (u *User) IsAdmin() bool {
	return (u.Role == RoleAdmin)
}

func (u *User) Can(permission string) error {
	switch permission {
	case PermissionUpdate:
		return u.HasPermissionsToUpdateUser()
	case PermissionDelete:
		return u.HasPermissionsToDeleteUser()
	default:
		return apperrors.RoleCanNoPermission.AppendMessage(fmt.Errorf(hasNoPermissionsError))
	}
}

func (u *User) HasPermissionsToUpdateUser() error {
	if u.IsAdmin() {
		return nil
	}
	return apperrors.HasPermissionsToUpdateUser.AppendMessage(fmt.Errorf(hasNoPermissionsToUpdateUserError))
}

func (u *User) HasPermissionsToDeleteUser() error {
	if u.IsAdmin() {
		return nil
	}

	return apperrors.HasPermissionsDeleteUser.AppendMessage(fmt.Errorf(hasNoPermissionsToDeleteUserError))
}
