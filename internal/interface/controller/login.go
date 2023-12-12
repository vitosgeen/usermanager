package controller

import (
	"net/http"
	"time"

	"usermanager/internal/apperrors"
	"usermanager/internal/domain/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func (uc *userController) Login(ctx echo.Context) error {
	loginRequest := &model.LoginRequest{}
	if err := ctx.Bind(loginRequest); err != nil {
		appErr := err.(*apperrors.AppError)
		return ctx.JSON(appErr.HTTPCode, appErr.Error())
	}

	user, err := uc.userUsecase.GetUserByNickname(ctx.Request().Context(), loginRequest.Nickname)
	if err != nil {
		appErr := err.(*apperrors.AppError)
		return ctx.JSON(appErr.HTTPCode, appErr.Error())
	}
	if user == nil {
		appError := apperrors.UserControllerLoginGetUserByNicknameEmpty.AppendMessage(echo.ErrUnauthorized)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	err = user.ComparePasswords(loginRequest.Password)
	if err != nil {
		appErr := err.(*apperrors.AppError)
		return ctx.JSON(appErr.HTTPCode, appErr.Error())
	}

	claims := &model.JwtCustomClaims{
		UserID:   user.UserID,
		Nickname: user.Nickname,
		Role:     user.Role,
	}
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(uc.cfg.Jwt.Ttl)))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(uc.cfg.Jwt.Secret))
	if err != nil {
		appErr := err.(*apperrors.AppError)
		return ctx.JSON(appErr.HTTPCode, appErr.Error())
	}

	return ctx.JSON(http.StatusOK, model.LoginResponse{Token: tokenSigned})
}
