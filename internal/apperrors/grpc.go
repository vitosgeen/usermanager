package apperrors

var (
	UserGrpcControllerGetUserUuidParse = AppError{
		Message:  "The get user operation has been failed. Parse uuid has been failed",
		Code:     "USER_GRPC_CONTROLLER_GET_USER_UUID_PARSE",
		HTTPCode: 500,
	}

	UserUsecaseNotFound = AppError{
		Message:  "The user not found",
		Code:     "USER_USECASE_NOT_FOUND",
		HTTPCode: 404,
	}

	UserGrpcControllerCreateUserError = AppError{
		Message:  "The create user operation has been failed",
		Code:     "USER_GRPC_CONTROLLER_CREATE_USER_ERROR",
		HTTPCode: 500,
	}

	UserGrpcControllerUpdateUserUuidParse = AppError{
		Message:  "The update user operation has been failed. Parse uuid has been failed",
		Code:     "USER_GRPC_CONTROLLER_UPDATE_USER_UUID_PARSE",
		HTTPCode: 500,
	}

	UserGrpcControllerUpdateUserUpdateUser = AppError{
		Message:  "The update user operation has been failed. Update user has been failed",
		Code:     "USER_GRPC_CONTROLLER_UPDATE_USER_UPDATE_USER",
		HTTPCode: 500,
	}

	UserGrpcControllerDeleteUserUuidParse = AppError{
		Message:  "The delete user operation has been failed. Parse uuid has been failed",
		Code:     "USER_GRPC_CONTROLLER_DELETE_USER_UUID_PARSE",
		HTTPCode: 500,
	}

	UserGrpcControllerDeleteUser = AppError{
		Message:  "The delete user operation has been failed. Delete user has been failed",
		Code:     "USER_GRPC_CONTROLLER_DELETE_USER",
		HTTPCode: 500,
	}

	UserGrpcControllerGetUsers = AppError{
		Message:  "The get users operation has been failed",
		Code:     "USER_GRPC_CONTROLLER_GET_USERS",
		HTTPCode: 500,
	}

	UserGrpcControllerGetUsersByPaginationQuery = AppError{
		Message:  "The get users by pagination query operation has been failed",
		Code:     "USER_GRPC_CONTROLLER_GET_USERS_BY_PAGINATION_QUERY",
		HTTPCode: 500,
	}

	UserGrpcControllerGetUser = AppError{
		Message:  "The get user operation has been failed",
		Code:     "USER_GRPC_CONTROLLER_GET_USER",
		HTTPCode: 500,
	}

	UserGrpcControllerGetUserByIDUuidParse = AppError{
		Message:  "The get user by id operation has been failed. Parse uuid has been failed",
		Code:     "USER_GRPC_CONTROLLER_GET_USER_BY_ID_UUID_PARSE",
		HTTPCode: 500,
	}

	UserGrpcControllerGetUserByID = AppError{
		Message:  "The get user by id operation has been failed",
		Code:     "USER_GRPC_CONTROLLER_GET_USER_BY_ID",
		HTTPCode: 500,
	}

	UserGrpcControllerGetUserByNickname = AppError{
		Message:  "The get user by nickname operation has been failed",
		Code:     "USER_GRPC_CONTROLLER_GET_USER_BY_NICKNAME",
		HTTPCode: 500,
	}

	UserGrpcControllerCheckUserByNickname = AppError{
		Message:  "The check user by nickname operation has been failed",
		Code:     "USER_GRPC_CONTROLLER_CHECK_USER_BY_NICKNAME",
		HTTPCode: 500,
	}

	UserGrpcControllerVoteUser = AppError{
		Message:  "The vote user operation has been failed",
		Code:     "USER_GRPC_CONTROLLER_VOTE_USER",
		HTTPCode: 500,
	}

	UserGrpcControllerFindExistVoting = AppError{
		Message:  "The find exist voting operation has been failed",
		Code:     "USER_GRPC_CONTROLLER_FIND_EXIST_VOTING",
		HTTPCode: 500,
	}

	UserGrpcControllerFindVotesForUserUuidParse = AppError{
		Message:  "The find votes for user operation has been failed. Parse uuid has been failed",
		Code:     "USER_GRPC_CONTROLLER_FIND_VOTES_FOR_USER_UUID_PARSE",
		HTTPCode: 500,
	}

	UserGrpcControllerFindVotesForUserFindVotesForUser = AppError{
		Message:  "The find votes for user operation has been failed. Find votes for user has been failed",
		Code:     "USER_GRPC_CONTROLLER_FIND_VOTES_FOR_USER_FIND_VOTES_FOR_USER",
		HTTPCode: 500,
	}

	UserGrpcControllerLoadVotesToUsers = AppError{
		Message:  "The load votes to users operation has been failed",
		Code:     "USER_GRPC_CONTROLLER_LOAD_VOTES_TO_USERS",
		HTTPCode: 500,
	}

	UserGrpcControllerGetLastVoteForUserUuidParse = AppError{
		Message:  "The get last vote for user operation has been failed. Parse uuid has been failed",
		Code:     "USER_GRPC_CONTROLLER_GET_LAST_VOTE_FOR_USER_UUID_PARSE",
		HTTPCode: 500,
	}

	UserGrpcControllerGetLastVoteForUser = AppError{
		Message:  "The get last vote for user operation has been failed",
		Code:     "USER_GRPC_CONTROLLER_GET_LAST_VOTE_FOR_USER",
		HTTPCode: 500,
	}

	UserGrpcControllerVote = AppError{
		Message:  "The vote operation has been failed",
		Code:     "USER_GRPC_CONTROLLER_VOTE",
		HTTPCode: 500,
	}
)
