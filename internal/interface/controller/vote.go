package controller

import (
	"net/http"
	"time"

	"usermanager/internal/apperrors"
	"usermanager/internal/domain/model"

	"github.com/labstack/echo/v4"
)

const (
	votePositive = 1
	voteNegative = -1
	voteWithdraw = 0
	voteInterval = 3600
)

func (uc *userController) VoteUser(ctx echo.Context) error {
	voteUserRequest := &model.VoteUserRequest{}
	err := ctx.Bind(voteUserRequest)
	if err != nil {
		appError := apperrors.UserControllerVoteUserBind.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}
	err = ctx.Validate(voteUserRequest)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	user, err := uc.userUsecase.GetUser(ctx.Request().Context(), voteUserRequest.UserID)
	if err != nil {
		appError := apperrors.UserControllerVoteUserGetUser.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}
	if user == nil {
		appError := apperrors.UserControllerVoteUserUserNotExist
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	authUser := uc.FetchJWTUser(ctx)
	if authUser.UserID == voteUserRequest.UserID {
		appError := apperrors.UserControllerVoteUserVoteForYourself
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}
	lastVoteUser, err := uc.userUsecase.GetLastVoteForUser(ctx.Request().Context(), &authUser.UserID)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}
	if lastVoteUser != nil && lastVoteUser.CreatedAt.Add(voteInterval).After(time.Now()) {
		appError := apperrors.UserControllerVoteUserVoteInterval
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	vote := model.MapVoteUserRequestToVoteModel(voteUserRequest, authUser)
	userVote := model.MapVoteUserRequestToUserVoteModel(voteUserRequest, vote)
	vote, userVote, err = uc.userUsecase.Vote(ctx.Request().Context(), vote, userVote)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	return ctx.JSON(http.StatusOK, userVote.MapVoteUserVoteToVoteUserResponse(vote))
}
