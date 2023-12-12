package apperrors

import "net/http"

var (
	UserControllerCreateUserBind = AppError{
		Message:  "The create user operation has been failed",
		Code:     "USER_REPO_LOAD_GET_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	UserControllerCreateUserJSON = AppError{
		Message:  "The create user operation has been failed",
		Code:     "USER_CONTROLLER_CREATE_USER_JSON",
		HTTPCode: http.StatusInternalServerError,
	}

	UserControllerDeleteUserHasPermissions = AppError{
		Message:  "The delete user operation has been failed, User doesn't have permission",
		Code:     "USER_CONTROLLER_DELETE_USER_HAS_PERMISSIONS",
		HTTPCode: http.StatusInternalServerError,
	}

	UserControllerDeleteUserUuidParse = AppError{
		Message:  "The delete user operation has been failed",
		Code:     "USER_CONTROLLER_DELETE_USER_UUID_PARSE",
		HTTPCode: http.StatusInternalServerError,
	}

	UserControllerDeleteUserDeleteUser = AppError{
		Message:  "The delete user operation has been failed",
		Code:     "USER_CONTROLLER_DELETE_USER_DELETE_USER",
		HTTPCode: http.StatusInternalServerError,
	}

	UserControllerGetUserUuidParse = AppError{
		Message:  "The get user operation has been failed",
		Code:     "USER_CONTROLLER_GET_USER_UUID_PARSE",
		HTTPCode: http.StatusInternalServerError,
	}

	UserControllerGetUserUserNotExist = AppError{
		Message:  "The get user operation has been failed, user is not exist",
		Code:     "USER_CONTROLLER_GET_USER_USER_NOT_EXIST",
		HTTPCode: http.StatusNotFound,
	}

	UserControllerGetUserGetUser = AppError{
		Message:  "The get user operation has been failed",
		Code:     "USER_CONTROLLER_GET_USER_GET_USER",
		HTTPCode: http.StatusInternalServerError,
	}

	UserControllerUpdateUserUuidParse = AppError{
		Message:  "The update user operation has been failed, the uuid parse has error",
		Code:     "USER_CONTROLLER_UPDATE_USER_USER_NOT_EXIST",
		HTTPCode: http.StatusInternalServerError,
	}

	UserControllerUpdateUserUserNotExist = AppError{
		Message:  "The update user operation has been failed, user is not exist",
		Code:     "USER_CONTROLLER_UPDATE_USER_USER_NOT_EXIST",
		HTTPCode: http.StatusInternalServerError,
	}

	UserControllerUpdateUserHasPermission = AppError{
		Message:  "The update user operation has been failed, user dosn't have a permission",
		Code:     "USER_CONTROLLER_UPDATE_USER_HAS_PERMISSION",
		HTTPCode: http.StatusBadRequest,
	}

	UserControllerUpdateUserBind = AppError{
		Message:  "The update user operation has been failed, user binds error",
		Code:     "USER_CONTROLLER_UPDATE_USER_BIND",
		HTTPCode: http.StatusInternalServerError,
	}

	UserControllerTryToSetAdmin = AppError{
		Message:  "The update user operation has been failed, user had tried to set himself as an admin",
		Code:     "USER_CONTROLLER_TRY_TO_SET_ADMIN",
		HTTPCode: http.StatusBadRequest,
	}

	UserControllerGetUserJSON = AppError{
		Message:  "The get user operation has been failed",
		Code:     "USER_CONTROLLER_GET_USER_JSON",
		HTTPCode: http.StatusInternalServerError,
	}

	UserControllerGetUsersGetPaginationFromCtx = AppError{
		Message:  "The get users operation has been failed",
		Code:     "USER_CONTROLLER_GET_USERS_GET_PAGINATION_FROM_CTX",
		HTTPCode: http.StatusInternalServerError,
	}

	UserControllerGetUsersUsers = AppError{
		Message:  "The get users operation has been failed",
		Code:     "USER_CONTROLLER_GET_USERS_LIST_USERS",
		HTTPCode: http.StatusInternalServerError,
	}

	UserControllerLoginCtxBind = AppError{
		Message:  "The login operation has been failed,  bind user error",
		Code:     "USER_CONTROLLER_LOGIN_CTX_BIND",
		HTTPCode: http.StatusBadRequest,
	}

	UserControllerLoginGetUserByNickname = AppError{
		Message:  "The login operation has been failed,  get user by nickname error",
		Code:     "USER_CONTROLLER_LOGIN_GET_USER_BY_NICKNAME",
		HTTPCode: http.StatusBadRequest,
	}

	UserControllerLoginGetUserByNicknameEmpty = AppError{
		Message:  "The login operation has been failed,  get user by nickname, user not exist",
		Code:     "USER_CONTROLLER_LOGIN_GET_USER_BY_NICKNAME_EMPTY",
		HTTPCode: http.StatusUnauthorized,
	}

	UserControllerLoginComparePasswords = AppError{
		Message:  "The login operation has been failed, password is wrong",
		Code:     "USER_CONTROLLER_LOGIN_COMPARE_PASSWORDS",
		HTTPCode: http.StatusUnauthorized,
	}

	UserControllerLoginTokenSigned = AppError{
		Message:  "The login operation token signed has been failed",
		Code:     "USER_CONTROLLER_LOGIN_TOKEN_SIGNED",
		HTTPCode: http.StatusUnauthorized,
	}

	UserControllerFetchJWTUser = AppError{
		Message:  "The FetchJWTUser operation has been failed",
		Code:     "USER_CONTROLLER_FETCH_JWT_USER",
		HTTPCode: http.StatusUnauthorized,
	}

	UserControllerVoteUserValidate = AppError{
		Message:  "The vote user operation has been failed, validate error",
		Code:     "USER_CONTROLLER_VOTE_USER_VALIDATE",
		HTTPCode: http.StatusBadRequest,
	}

	UserControllerVoteUserBind = AppError{
		Message:  "The vote user operation has been failed, vote user binds error",
		Code:     "USER_CONTROLLER_VOTE_USER_BIND",
		HTTPCode: http.StatusInternalServerError,
	}

	UserControllerVoteUserGetUser = AppError{
		Message:  "The vote user operation has been failed, get user error",
		Code:     "USER_CONTROLLER_VOTE_USER_GET_USER",
		HTTPCode: http.StatusInternalServerError,
	}

	UserControllerVoteUserUserNotExist = AppError{
		Message:  "The vote user operation has been failed, user is not exist",
		Code:     "USER_CONTROLLER_VOTE_USER_USER_NOT_EXIST",
		HTTPCode: http.StatusNotFound,
	}

	ValidatorCustomValidatorValidate = AppError{
		Message:  "The custom validate operation has been failed",
		Code:     "VALIDATOR_CUSTOM_VALIDATOR_VALIDATE",
		HTTPCode: http.StatusBadRequest,
	}

	MiddlewareJWTAuthValid = AppError{
		Message:  "The jwt auth user hasn't been validate",
		Code:     "MIDDLEWARE_JWT_AUTH_VALID",
		HTTPCode: http.StatusUnauthorized,
	}

	MiddlewareJWTAuthVerifyJwtUser = AppError{
		Message:  "The jwt auth verify user hasn't been validate",
		Code:     "MIDDLEWARE_JWT_AUTH_VERIFY_JWT_USER",
		HTTPCode: http.StatusUnauthorized,
	}

	MiddlewareVerifyAuthUserGetUserByNickname = AppError{
		Message:  "The verify user operation has been failed",
		Code:     "MIDDLEWARE_VERIFY_AUTH_USER_GET_USER_BY_NICKNAME",
		HTTPCode: http.StatusUnauthorized,
	}

	MiddlewareVerifyAuthUserComparePasswords = AppError{
		Message:  "The verify user operation has been failed, password isn't correct",
		Code:     "MIDDLEWARE_VERIFY_AUTH_USER_COMPARE_PASSWORDS",
		HTTPCode: http.StatusUnauthorized,
	}

	MiddlewareVerifyJwtUserGetUserByNickname = AppError{
		Message:  "The jwt verify user operation has been failed",
		Code:     "MIDDLEWARE_VERIFY_JWT_USER_GET_USER_BY_NICKNAME",
		HTTPCode: http.StatusUnauthorized,
	}

	MiddlewareVerifyJwtUserRoleGetUserByNickname = AppError{
		Message:  "The jwt verify user role operation hasn't been validate",
		Code:     "MIDDLEWARE_VERIFY_JWT_USER_ROLE_GET_USER_BY_NICKNAME",
		HTTPCode: http.StatusUnauthorized,
	}

	HasPermissionUuidParse = AppError{
		Message:  "The hasPermission operation has been failed, the uuid parse has error",
		Code:     "HAS_PERMISSION_UUID_PARSE",
		HTTPCode: http.StatusUnauthorized,
	}

	UserControllerVoteUserValueOfVoteIsNotRight = AppError{
		Message:  "Value of vote is not right",
		Code:     "USER_CONTROLLER_VOTE_USER_VALUE_OF_VOTE_IS_NOT_RIGHT",
		HTTPCode: http.StatusBadRequest,
	}

	UserControllerVoteUserVoteForYourself = AppError{
		Message:  "You can't vote for yourself",
		Code:     "USER_CONTROLLER_VOTE_USER_VOTE_FOR_YOURSELF",
		HTTPCode: http.StatusBadRequest,
	}

	UserControllerVoteUserVoteInterval = AppError{
		Message:  "You can't vote for user more than once an hour",
		Code:     "USER_CONTROLLER_VOTE_USER_VOTE_INTERVAL",
		HTTPCode: http.StatusBadRequest,
	}
)
