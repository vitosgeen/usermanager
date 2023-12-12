package controller

import (
	"usermanager/internal/apperrors"
	"usermanager/internal/domain/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	UserAuthCtx      = "userAuth"
	updatePermission = "update"
	deletePermission = "delete"
)

func (uc *userController) SetUpJWTConfig() echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.JwtCustomClaims)
		},
		SigningKey: []byte(uc.cfg.Jwt.Secret),
	}

	return echojwt.WithConfig(config)
}

func (uc *userController) CanUpdateUser() echo.MiddlewareFunc {
	return uc.hasPermission(updatePermission)
}

func (uc *userController) CanDeleteUser() echo.MiddlewareFunc {
	return uc.hasPermission(deletePermission)
}

func (uc *userController) hasPermission(permission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			userUUID, err := uuid.Parse(ctx.Param("id"))
			if err != nil {
				appError := apperrors.HasPermissionUuidParse.AppendMessage(err)
				return ctx.JSON(appError.HTTPCode, appError.Error())
			}

			authUser := uc.FetchJWTUser(ctx)
			if userUUID == authUser.UserID {
				return next(ctx)
			}

			err = authUser.Can(permission)
			if err != nil {
				appError := err.(*apperrors.AppError)
				return ctx.JSON(appError.HTTPCode, appError.Error())
			}

			return next(ctx)
		}
	}
}

func (uc *userController) JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user := ctx.Get("user").(*jwt.Token)
		claims := user.Claims.(*model.JwtCustomClaims)

		if !user.Valid {
			apperrors.MiddlewareJWTAuthValid.AppendMessage(echo.ErrUnauthorized)
		}

		_, err := uc.VerifyJwtUser(ctx, claims.Nickname, claims.Role)
		if err != nil {
			apperrors.MiddlewareJWTAuthVerifyJwtUser.AppendMessage(echo.ErrUnauthorized)
		}

		return next(ctx)
	}
}

func (uc *userController) BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(uc.VerifyAuthUser())
}

func (uc *userController) VerifyJwtUser(ctx echo.Context, nickname, role string) (bool, error) {
	user, err := uc.userUsecase.GetUserByNickname(ctx.Request().Context(), nickname)
	if err != nil {
		return false, apperrors.MiddlewareVerifyJwtUserGetUserByNickname.AppendMessage(err)
	}
	if user == nil {
		return false, apperrors.MiddlewareVerifyJwtUserGetUserByNickname.AppendMessage(err)
	}
	if user.Role != role {
		return false, apperrors.MiddlewareVerifyAuthUserGetUserByNickname.AppendMessage(err)
	}
	ctx.Set(UserAuthCtx, user)

	return true, nil
}

func (uc *userController) VerifyAuthUser() func(username, password string, ctx echo.Context) (bool, error) {
	return func(username, password string, ctx echo.Context) (bool, error) {
		user, err := uc.userUsecase.GetUserByNickname(ctx.Request().Context(), username)
		if err != nil {
			return false, apperrors.MiddlewareVerifyAuthUserGetUserByNickname.AppendMessage(err)
		}
		if user == nil {
			return false, apperrors.MiddlewareVerifyAuthUserGetUserByNickname.AppendMessage(err)
		}

		err = user.ComparePasswords(password)
		if err != nil {
			return false, apperrors.MiddlewareVerifyAuthUserComparePasswords.AppendMessage(err)
		}

		ctx.Set(UserAuthCtx, user)

		return true, nil
	}
}
