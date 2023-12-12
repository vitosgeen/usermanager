package apperrors

import "net/http"

var (
	UserRepoFindUserByUUIDGetContext = AppError{
		Message:  "FindUserByUUID operation has been failed",
		Code:     "USER_REPO_FIND_USER_BY_UUID_GET_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRepoFindUserByUUIDGetDataNotFound = AppError{
		Message:  "FindUserByUUID operation has been failed. Data not found",
		Code:     "USER_REPO_FIND_USER_BY_UUID_GET_DATA_NOT_FOUND",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRepoFindUserByNicknameGetContext = AppError{
		Message:  "FindUserByNickname user by nickname operation has been failed",
		Code:     "USER_REPO_FIND_USER_BY_NICKNAME_GET_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRepoFindUserByNicknameGetDataNotFound = AppError{
		Message:  "FindUserByNickname user by nickname operation has been failed. Data not found",
		Code:     "USER_REPO_FIND_USER_BY_NICKNAME_GET_DATA_NOT_FOUND",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRepoSaveUserQueryRowxContext = AppError{
		Message:  "Add operation has been failed",
		Code:     "USER_REPO_ADD_USER_QUERY_ROWX_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRepoUpdateUserQueryRowxContext = AppError{
		Message:  "The update operation has been failed",
		Code:     "USER_REPO_UPDATE_USER_QUERY_ROWX_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRepoSoftDeleteUserByUserIDQueryRowxContext = AppError{
		Message:  "The delete soft operation has been failed",
		Code:     "USER_REPO_SOFT_DELETE_USER_BY_USER_ID_QUERY_ROWX_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRepoDeleteUserByUserIDExecContext = AppError{
		Message:  "The delete operation from db has been failed",
		Code:     "USER_REPO_DELETE_USER_BY_USER_ID_EXEC_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRepoDeleteUserByUserIDRowsAffected = AppError{
		Message:  "The delete operation from db affected rows has been failed",
		Code:     "USER_REPO_DELETE_USER_BY_USER_ID_ROWS_AFFECTED",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRepoDeleteUserByUserIDEmptyRowsAffected = AppError{
		Message:  "The delete operation from db doesn't have affected rows",
		Code:     "USER_REPO_DELETE_USER_BY_USER_ID_EMPTY_ROWS_AFFECTED",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRepoGetUsersGetContext = AppError{
		Message:  "The get users total count operation has been failed",
		Code:     "USER_REPO_GET_USERS_GET_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRepoGetUsersQueryxContext = AppError{
		Message:  "The get users query operation has been failed",
		Code:     "USER_REPO_GET_USERS_QUERYX_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRepoGetUsersStructScan = AppError{
		Message:  "The get users struct scan operation has been failed",
		Code:     "USER_REPO_GET_USERS_STRUCT_SCAN",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRepoGetUsersRows = AppError{
		Message:  "The get users rows operation has been failed",
		Code:     "USER_REPO_GET_USERS_ROWS",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoSaveVoteQueryRowxContext = AppError{
		Message:  "The save vote operation has been failed",
		Code:     "VOTE_REPO_SAVE_VOTE_QUERY_ROWX_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoSaveVoteQueryRowxContextDataNotFound = AppError{
		Message:  "The save vote operation has been failed. Data not found",
		Code:     "VOTE_REPO_SAVE_VOTE_QUERY_ROWX_CONTEXT_DATA_NOT_FOUND",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoUpdateVoteQueryRowxContext = AppError{
		Message:  "The update vote operation has been failed",
		Code:     "VOTE_REPO_UPDATE_VOTE_QUERY_ROWX_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoUpdateVoteQueryRowxContextDataNotFound = AppError{
		Message:  "The update vote operation has been failed. Data not found",
		Code:     "VOTE_REPO_UPDATE_VOTE_QUERY_ROWX_CONTEXT_DATA_NOT_FOUND",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoFindVoteByIDGetContext = AppError{
		Message:  "The find vote by id operation has been failed",
		Code:     "VOTE_REPO_FIND_VOTE_BY_ID_GET_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoFindVoteByIDVoteIDEmpty = AppError{
		Message:  "The find vote by id operation has been failed. Vote id is empty",
		Code:     "VOTE_REPO_FIND_VOTE_BY_ID_VOTE_ID_EMPTY",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoFindVoteByIDQueryxContextDataNotFound = AppError{
		Message:  "The find vote by id operation has been failed. Data not found",
		Code:     "VOTE_REPO_FIND_VOTE_BY_ID_QUERYX_CONTEXT_DATA_NOT_FOUND",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoFindUserVoteByIDGetContext = AppError{
		Message:  "The find user's vote by id operation has been failed",
		Code:     "VOTE_REPO_FIND_USER_VOTE_BY_ID_GET_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoFindUserVoteByIDVoteIDEmpty = AppError{
		Message:  "The find user's vote by id operation has been failed. Vote id is empty",
		Code:     "VOTE_REPO_FIND_USER_VOTE_BY_ID_VOTE_ID_EMPTY",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoFindUserVoteByIDQueryxContextDataNotFound = AppError{
		Message:  "The find user's vote by id operation has been failed. Data not found",
		Code:     "VOTE_REPO_FIND_USER_VOTE_BY_ID_QUERYX_CONTEXT_DATA_NOT_FOUND",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoFindVoteByUserIDGetContext = AppError{
		Message:  "The find vote by user id operation has been failed",
		Code:     "VOTE_REPO_FIND_VOTE_BY_USER_ID_GET_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoFindVoteByUserIDQueryxContextDataNotFound = AppError{
		Message:  "The find vote by user id operation has been failed. Data not found",
		Code:     "VOTE_REPO_FIND_VOTE_BY_USER_ID_QUERYX_CONTEXT_DATA_NOT_FOUND",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoFindVoteByUserIDStructScan = AppError{
		Message:  "The find vote by user id struct scan operation has been failed",
		Code:     "VOTE_REPO_FIND_VOTE_BY_USER_ID_STRUCT_SCAN",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoFindVotesByUserIDsQueryxContext = AppError{
		Message:  "The find vote by user ids operation has been failed. QueryxContext has been failed",
		Code:     "VOTE_REPO_FIND_VOTE_BY_USER_IDS_QUERYX_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoFindVotesByUserIDsQueryxContextDataNotFound = AppError{
		Message:  "The find vote by user ids operation has been failed. Data not found",
		Code:     "VOTE_REPO_FIND_VOTE_BY_USER_IDS_QUERYX_CONTEXT_DATA_NOT_FOUND",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoFindVotesByUserIDsStructScan = AppError{
		Message:  "The find vote by user ids operation has been failed. StructScan has been failed",
		Code:     "VOTE_REPO_FIND_VOTE_BY_USER_IDS_STRUCT_SCAN",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoFindVoteByUsersIDGetContext = AppError{
		Message:  "The find vote by users id operation has been failed",
		Code:     "VOTE_REPO_FIND_VOTE_BY_USERS_ID_GET_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoFindVoteByUsersIDStructScan = AppError{
		Message:  "The find vote by users id struct scan operation has been failed",
		Code:     "VOTE_REPO_FIND_VOTE_BY_USERS_ID_STRUCT_SCAN",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoFindUserVoteByUserIDGetContext = AppError{
		Message:  "The find user's vote by user id operation has been failed",
		Code:     "VOTE_REPO_FIND_USER_VOTE_BY_USER_ID_GET_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoFindUserVoteByUserIDQueryxContextDataNotFound = AppError{
		Message:  "The find user's vote by user id operation has been failed. Data not found",
		Code:     "VOTE_REPO_FIND_USER_VOTE_BY_USER_ID_QUERYX_CONTEXT_DATA_NOT_FOUND",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoSaveUserVoteQueryRowxContext = AppError{
		Message:  "The save user vote operation has been failed",
		Code:     "VOTE_REPO_SAVE_USER_VOTE_QUERY_ROWX_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoSaveUserVoteQueryRowxContextDataNotFound = AppError{
		Message:  "The save user vote operation has been failed. Data not found",
		Code:     "VOTE_REPO_SAVE_USER_VOTE_QUERY_ROWX_CONTEXT_DATA_NOT_FOUND",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoDeleteVoteExecContext = AppError{
		Message:  "The delete vote operation has been failed",
		Code:     "VOTE_REPO_DELETE_VOTE_EXEC_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoDeleteVoteRowsAffected = AppError{
		Message:  "The delete vote affected rows operation has been failed",
		Code:     "VOTE_REPO_DELETE_VOTE_ROWS_AFFECTED",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoDeleteVoteEmptyRowsAffected = AppError{
		Message:  "The delete vote empty rows affected operation has been failed",
		Code:     "VOTE_REPO_DELETE_VOTE_EMPTY_ROWS_AFFECTED",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoDeleteUserVoteExecContext = AppError{
		Message:  "The delete user vote operation has been failed",
		Code:     "VOTE_REPO_DELETE_USER_VOTE_EXEC_CONTEXT",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoDeleteUserVoteRowsAffected = AppError{
		Message:  "The delete user vote affected rows operation has been failed",
		Code:     "VOTE_REPO_DELETE_USER_VOTE_ROWS_AFFECTED",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRepoDeleteUserVoteEmptyRowsAffected = AppError{
		Message:  "The delete user vote empty rows affected operation has been failed",
		Code:     "VOTE_REPO_DELETE_USER_VOTE_EMPTY_ROWS_AFFECTED",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRedisRepoFindUserByUUIDGet = AppError{
		Message:  "The find user by uuid operation has been failed. Redis get has been failed",
		Code:     "USER_REDIS_REPO_FIND_USER_BY_UUID_GET",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRedisRepoFindUserByUUIDGetDataNotFound = AppError{
		Message:  "The find user by uuid operation has been failed. Data not found",
		Code:     "USER_REDIS_REPO_FIND_USER_BY_UUID_GET_DATA_NOT_FOUND",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRedisRepoFindUserByUUIDUnmarshal = AppError{
		Message:  "The find user by uuid operation has been failed. Unmarshal has been failed",
		Code:     "USER_REDIS_REPO_FIND_USER_BY_UUID_UNMARSHAL",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRedisRepoFindUserByNicknameGet = AppError{
		Message:  "The find user by nickname operation has been failed. Redis get has been failed",
		Code:     "USER_REDIS_REPO_FIND_USER_BY_NICKNAME_GET",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRedisRepoFindUserByNicknameGetDataNotFound = AppError{
		Message:  "The find user by nickname operation has been failed. Data not found",
		Code:     "USER_REDIS_REPO_FIND_USER_BY_NICKNAME_GET_DATA_NOT_FOUND",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRedisRepoFindUserByNicknameUnmarshal = AppError{
		Message:  "The find user by nickname operation has been failed. Unmarshal has been failed",
		Code:     "USER_REDIS_REPO_FIND_USER_BY_NICKNAME_UNMARSHAL",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRedisRepoGetUsersPrepareKey = AppError{
		Message:  "The get users operation has been failed. Prepare key has been failed",
		Code:     "USER_REDIS_REPO_GET_USERS_PREPARE_KEY",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRedisRepoGetUsersGet = AppError{
		Message:  "The get users operation has been failed. Redis get has been failed",
		Code:     "USER_REDIS_REPO_GET_USERS_GET",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRedisRepoGetUsersGetDataNotFound = AppError{
		Message:  "The get users operation has been failed. Data not found",
		Code:     "USER_REDIS_REPO_GET_USERS_GET_DATA_NOT_FOUND",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRedisRepoGetUsersUnmarshal = AppError{
		Message:  "The get users operation has been failed. Unmarshal has been failed",
		Code:     "USER_REDIS_REPO_GET_USERS_UNMARSHAL",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoFindVoteByUserIDGet = AppError{
		Message:  "The find vote by user id operation has been failed. Redis get has been failed",
		Code:     "VOTE_REDIS_REPO_FIND_VOTE_BY_USER_ID_GET",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoFindVoteByUserIDGetDataNotFound = AppError{
		Message:  "The find vote by user id operation has been failed. Data not found",
		Code:     "VOTE_REDIS_REPO_FIND_VOTE_BY_USER_ID_GET_DATA_NOT_FOUND",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoFindVoteByUserIDUnmarshal = AppError{
		Message:  "The find vote by user id operation has been failed. Unmarshal has been failed",
		Code:     "VOTE_REDIS_REPO_FIND_VOTE_BY_USER_ID_UNMARSHAL",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoFindVotesByUserIDsPrepareKey = AppError{
		Message:  "The find votes by user ids operation has been failed. Prepare key has been failed",
		Code:     "VOTE_REDIS_REPO_FIND_VOTES_BY_USER_IDS_PREPARE_KEY",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoFindVotesByUserIDsGet = AppError{
		Message:  "The find votes by user ids operation has been failed. Redis get has been failed",
		Code:     "VOTE_REDIS_REPO_FIND_VOTES_BY_USER_IDS_GET",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoFindVotesByUserIDsGetDataNotFound = AppError{
		Message:  "The find votes by user ids operation has been failed. Data not found",
		Code:     "VOTE_REDIS_REPO_FIND_VOTES_BY_USER_IDS_GET_DATA_NOT_FOUND",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoFindVotesByUserIDsUnmarshal = AppError{
		Message:  "The find votes by user ids operation has been failed. Unmarshal has been failed",
		Code:     "VOTE_REDIS_REPO_FIND_VOTES_BY_USER_IDS_UNMARSHAL",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoFindUserVoteByUserIDGet = AppError{
		Message:  "The find user's vote by user id operation has been failed. Redis get has been failed",
		Code:     "VOTE_REDIS_REPO_FIND_USER_VOTE_BY_USER_ID_GET",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoFindUserVoteByUserIDGetDataNotFound = AppError{
		Message:  "The find user's vote by user id operation has been failed. Data not found",
		Code:     "VOTE_REDIS_REPO_FIND_USER_VOTE_BY_USER_ID_GET_DATA_NOT_FOUND",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoFindUserVoteByUserIDUnmarshal = AppError{
		Message:  "The find user's vote by user id operation has been failed. Unmarshal has been failed",
		Code:     "VOTE_REDIS_REPO_FIND_USER_VOTE_BY_USER_ID_UNMARSHAL",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoFindUserVoteByIDGet = AppError{
		Message:  "The find user's vote by id operation has been failed. Redis get has been failed",
		Code:     "VOTE_REDIS_REPO_FIND_USER_VOTE_BY_ID_GET",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoFindUserVoteByIDGetDataNotFound = AppError{
		Message:  "The find user's vote by id operation has been failed. Data not found",
		Code:     "VOTE_REDIS_REPO_FIND_USER_VOTE_BY_ID_GET_DATA_NOT_FOUND",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoFindUserVoteByIDUnmarshal = AppError{
		Message:  "The find user's vote by id operation has been failed. Unmarshal has been failed",
		Code:     "VOTE_REDIS_REPO_FIND_USER_VOTE_BY_ID_UNMARSHAL",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRedisRepoSetFindUserByUUIDMarshal = AppError{
		Message:  "The set find user by uuid operation has been failed. Marshal has been failed",
		Code:     "USER_REDIS_REPO_SET_FIND_USER_BY_UUID_MARSHAL",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRedisRepoSetFindUserByUUIDSet = AppError{
		Message:  "The set find user by uuid operation has been failed. Redis set has been failed",
		Code:     "USER_REDIS_REPO_SET_FIND_USER_BY_UUID_SET",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRedisRepoSetFindUserByNicknameMarshal = AppError{
		Message:  "The set find user by nickname operation has been failed. Marshal has been failed",
		Code:     "USER_REDIS_REPO_SET_FIND_USER_BY_NICKNAME_MARSHAL",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRedisRepoSetFindUserByNicknameSet = AppError{
		Message:  "The set find user by nickname operation has been failed. Redis set has been failed",
		Code:     "USER_REDIS_REPO_SET_FIND_USER_BY_NICKNAME_SET",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRedisRepoSetGetUsersPrepareKey = AppError{
		Message:  "The set get users operation has been failed. Prepare key has been failed",
		Code:     "USER_REDIS_REPO_SET_GET_USERS_PREPARE_KEY",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRedisRepoSetGetUsersMarshal = AppError{
		Message:  "The set get users operation has been failed. Marshal has been failed",
		Code:     "USER_REDIS_REPO_SET_GET_USERS_MARSHAL",
		HTTPCode: http.StatusInternalServerError,
	}

	UserRedisRepoSetGetUsersSet = AppError{
		Message:  "The set get users operation has been failed. Redis set has been failed",
		Code:     "USER_REDIS_REPO_SET_GET_USERS_SET",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoSetFindVoteByUserIDMarshal = AppError{
		Message:  "The set find vote by user id operation has been failed. Marshal has been failed",
		Code:     "VOTE_REDIS_REPO_SET_FIND_VOTE_BY_USER_ID_MARSHAL",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoSetFindVoteByUserIDSet = AppError{
		Message:  "The set find vote by user id operation has been failed. Redis set has been failed",
		Code:     "VOTE_REDIS_REPO_SET_FIND_VOTE_BY_USER_ID_SET",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoSetFindVotesByUserIDsPrepareKey = AppError{
		Message:  "The set find votes by user ids operation has been failed. Prepare key has been failed",
		Code:     "VOTE_REDIS_REPO_SET_FIND_VOTES_BY_USER_IDS_PREPARE_KEY",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoSetFindVotesByUserIDsMarshal = AppError{
		Message:  "The set find votes by user ids operation has been failed. Marshal has been failed",
		Code:     "VOTE_REDIS_REPO_SET_FIND_VOTES_BY_USER_IDS_MARSHAL",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoSetFindVotesByUserIDsSet = AppError{
		Message:  "The set find votes by user ids operation has been failed. Redis set has been failed",
		Code:     "VOTE_REDIS_REPO_SET_FIND_VOTES_BY_USER_IDS_SET",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoSetFindUserVoteByUserIDMarshal = AppError{
		Message:  "The set find user's vote by user id operation has been failed. Marshal has been failed",
		Code:     "VOTE_REDIS_REPO_SET_FIND_USER_VOTE_BY_USER_ID_MARSHAL",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoSetFindUserVoteByUserIDSet = AppError{
		Message:  "The set find user's vote by user id operation has been failed. Redis set has been failed",
		Code:     "VOTE_REDIS_REPO_SET_FIND_USER_VOTE_BY_USER_ID_SET",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoSetFindUserVoteByIDMarshal = AppError{
		Message:  "The set find user's vote by id operation has been failed. Marshal has been failed",
		Code:     "VOTE_REDIS_REPO_SET_FIND_USER_VOTE_BY_ID_MARSHAL",
		HTTPCode: http.StatusInternalServerError,
	}

	VoteRedisRepoSetFindUserVoteByIDSet = AppError{
		Message:  "The set find user's vote by id operation has been failed. Redis set has been failed",
		Code:     "VOTE_REDIS_REPO_SET_FIND_USER_VOTE_BY_ID_SET",
		HTTPCode: http.StatusInternalServerError,
	}
)
