package apperrors

import "net/http"

var (
	UserUsecaseCreateUserHashPassword = AppError{
		Message:  "The hash password operation has been failed",
		Code:     "USER_USECASE_CREATE_USER_HASH_PASSWORD",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseCreateUserFindUserByNickname = AppError{
		Message:  "The load by nickname operation has been failed",
		Code:     "USER_USECASE_CREATE_USER_LOAD_BY_NICK_NAME",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseCreateUserSaveUser = AppError{
		Message:  "The add user operation has been failed",
		Code:     "USER_USECASE_CREATE_USER_ADD_USER",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseCreateUserUserExists = AppError{
		Message:  "The user exists",
		Code:     "USER_USECASE_CREATE_USER_USER_EXISTS",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseUpdateUserUpdateUser = AppError{
		Message:  "The user update operaion has been failed",
		Code:     "USER_USECASE_UPDATE_USER_UPDATE_USER",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseGetUserLoad = AppError{
		Message:  "The get user operaion has been failed",
		Code:     "USER_USECASE_GET_USER_LOAD",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseGetUserFindUserByUUID = AppError{
		Message:  "The get user operaion has been failed. Find user by uuid has been failed",
		Code:     "USER_USECASE_GET_USER_FIND_USER_BY_UUID",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseGetUserSetFindUserByUUID = AppError{
		Message:  "The get user operaion has been failed. Set find user by uuid has been failed",
		Code:     "USER_USECASE_GET_USER_SET_FIND_USER_BY_UUID",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseGetUserByNicknameUserRedisRepoFindUserByNickname = AppError{
		Message:  "The get user by nickname operaion has been failed. Find user by nickname has been failed",
		Code:     "USER_USECASE_GET_USER_BY_NICKNAME_USER_REDIS_REPO_FIND_USER_BY_NICKNAME",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseGetUserByNicknameUserRedisRepoSetFindUserByNickname = AppError{
		Message:  "The get user by nickname operaion has been failed. Set find user by nickname has been failed",
		Code:     "USER_USECASE_GET_USER_BY_NICKNAME_USER_REDIS_REPO_SET_FIND_USER_BY_NICKNAME",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseGetUsersUserRedisRepoGetUsers = AppError{
		Message:  "The get users operaion has been failed. Get users has been failed",
		Code:     "USER_USECASE_GET_USERS_USER_REDIS_REPO_GET_USERS",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseGetUsersUserRedisRepoSetGetUsers = AppError{
		Message:  "The get users operaion has been failed. Set get users has been failed",
		Code:     "USER_USECASE_GET_USERS_USER_REDIS_REPO_SET_GET_USERS",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseGetUserByIDGetUser = AppError{
		Message:  "The get user by id operaion has been failed. Get user has been failed",
		Code:     "USER_USECASE_GET_USER_BY_ID_GET_USER",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseGetUserByIDFindVotesForUser = AppError{
		Message:  "The get user by id operaion has been failed. Find votes for user has been failed",
		Code:     "USER_USECASE_GET_USER_BY_ID_FIND_VOTES_FOR_USER",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseDeleteUser = AppError{
		Message:  "The delete user operaion has been failed",
		Code:     "USER_USECASE_DELETE_USER",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseCheckProfileByNick = AppError{
		Message:  "The check user by nickname operaion has been failed",
		Code:     "USER_USECASE_CHECK_PROFILE_BY_NICK",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseCheckProfileByNickBusy = AppError{
		Message:  "user's nickname is busy",
		Code:     "USER_USECASE_CHECK_PROFILE_BY_NICK_BUSY",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseGetUserByNickname = AppError{
		Message:  "The get user by nickname operaion has been failed",
		Code:     "USER_USECASE_GET_USER_BY_NICKNAME",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseGetUsers = AppError{
		Message:  "The get users operaion has been failed",
		Code:     "USER_USECASE_GET_USERS",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseVoteUserFindExistVoting = AppError{
		Message:  "The vote user operaion has been failed. Find exist vote has been failed",
		Code:     "USER_USECASE_VOTE_USER_FIND_EXIST_VOTING",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseVoteUserVotingExist = AppError{
		Message:  "The vote user operaion has been failed. Vote is exist",
		Code:     "USER_USECASE_VOTE_USER_VOTING_EXIST",
		HTTPCode: http.StatusOK,
	}

	UserUsecaseVoteUserSaveVote = AppError{
		Message:  "The vote user save operaion has been failed.",
		Code:     "USER_USECASE_VOTE_USER_SAVE_VOTE",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseVoteUserSaveUserVote = AppError{
		Message:  "The vote user save operaion has been failed.",
		Code:     "USER_USECASE_VOTE_USER_SAVE_USER_VOTE",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseVoteUserUpdateVote = AppError{
		Message:  "The vote user update operaion has been failed.",
		Code:     "USER_USECASE_VOTE_USER_UPDATE_VOTE",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseVoteUserNegativeVoteRepoFindExistVoteGetContext = AppError{
		Message:  "The vote user negative operaion has been failed. Find exist vote has been failed",
		Code:     "USER_USECASE_VOTE_USER_NEGATIVE_FIND_VOTE_BY_VOTER_USER_ID_VOTE_USER_ID",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseVoteUserNegativeVoteExist = AppError{
		Message:  "The vote user negative operaion has been failed. Vote is exist",
		Code:     "USER_USECASE_VOTE_USER_NEGATIVE_VOTE_EXIST",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseVoteUserNegativeSaveVote = AppError{
		Message:  "The vote user negative operaion has been failed. Vote is exist",
		Code:     "USER_USECASE_VOTE_USER_NEGATIVE_SAVE_VOTE",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseVoteUserWithdrawFindExistVoting = AppError{
		Message:  "The vote user withdraw operaion has been failed. Find exist vote has been failed",
		Code:     "USER_USECASE_VOTE_USER_WITHDRAW_FIND_EXIST_VOTING",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseVoteUserWithdrawDeleteVote = AppError{
		Message:  "The vote user withdraw operaion has been failed. Delete vote has been failed",
		Code:     "USER_USECASE_VOTE_USER_WITHDRAW_DELETE_VOTE",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseVoteUserWithdrawDeleteUserVote = AppError{
		Message:  "The vote user withdraw operaion has been failed. Delete user vote has been failed",
		Code:     "USER_USECASE_VOTE_USER_WITHDRAW_DELETE_USER_VOTE",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseVoteUserWithdrawVoteNotExist = AppError{
		Message:  "The vote user withdraw operaion has been failed. Vote is not exist",
		Code:     "USER_USECASE_VOTE_USER_WITHDRAW_VOTE_NOT_EXIST",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseVoteIsExistFindVoteByUserID = AppError{
		Message:  "The find votes by user id operaion has been failed.",
		Code:     "USER_USECASE_VOTE_IS_EXIST_FIND_VOTE_BY_USER_ID",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseVoteIsExistFindUserVoteByUserID = AppError{
		Message:  "The find user votes by user id operaion has been failed.",
		Code:     "USER_USECASE_VOTE_IS_EXIST_FIND_USER_VOTE_BY_USER_ID",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseFindVotesForUser = AppError{
		Message:  "The find votes for user operaion has been failed.",
		Code:     "USER_USECASE_FIND_VOTES_FOR_USER",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseFindVotesForUsers = AppError{
		Message:  "The find votes for users operaion has been failed.",
		Code:     "USER_USECASE_FIND_VOTES_FOR_USERS",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseLoadVotesToUsers = AppError{
		Message:  "The load votes to users operaion has been failed.",
		Code:     "USER_USECASE_LOAD_VOTES_TO_USERS",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseVotePositiveVoteUser = AppError{
		Message:  "The vote user positive operaion has been failed.",
		Code:     "USER_USECASE_VOTE_POSITIVE_VOTE_USER",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseVoteNegativeVoteUser = AppError{
		Message:  "The vote user negative operaion has been failed.",
		Code:     "USER_USECASE_VOTE_NEGATIVE_VOTE_USER",
		HTTPCode: http.StatusInternalServerError,
	}

	UserUsecaseVoteWithdrawVoteUser = AppError{
		Message:  "The vote user withdraw operaion has been failed.",
		Code:     "USER_USECASE_VOTE_WITHDRAW_VOTE_USER",
		HTTPCode: http.StatusInternalServerError,
	}
)
