package controller

import (
	"net/http"

	"usermanager/internal/apperrors"
	"usermanager/internal/config"
	"usermanager/internal/domain/model"
	"usermanager/internal/usecase/usecase"
	"usermanager/internal/utils"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type userController struct {
	userUsecase usecase.IUserUsecase
	cfg         *config.Config
}

type IUserController interface {
	GetUser(ctx echo.Context) error
	GetUsers(ctx echo.Context) error
	CreateUser(ctx echo.Context) error
	UpdateUser(ctx echo.Context) error
	DeleteUser(ctx echo.Context) error
	Login(ctx echo.Context) error
	VoteUser(ctx echo.Context) error
	SetUpJWTConfig() echo.MiddlewareFunc
	BasicAuth() echo.MiddlewareFunc
	JWTAuth(next echo.HandlerFunc) echo.HandlerFunc
	FetchJWTUser(ctx echo.Context) *model.User
	CanUpdateUser() echo.MiddlewareFunc
	CanDeleteUser() echo.MiddlewareFunc
}

func NewUserController(userUsecase usecase.IUserUsecase, cfg *config.Config) IUserController {
	return &userController{userUsecase, cfg}
}

func (uc *userController) CreateUser(ctx echo.Context) error {
	createUser := &model.CreateUserRequest{}
	if err := ctx.Bind(createUser); err != nil {
		appError := apperrors.UserControllerCreateUserBind.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	user := &model.User{}
	authUser := uc.FetchJWTUser(ctx)
	user.MapCreateUserRequestToUserModel(createUser)
	user.Created.By = authUser.UserID.String()

	createdUser, err := uc.userUsecase.CreateUser(ctx.Request().Context(), user)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	return ctx.JSON(http.StatusCreated, createdUser.MapUserModelToCreateUserResponse())
}

func (uc *userController) UpdateUser(ctx echo.Context) error {
	updateUser := &model.UpdateUserRequest{}
	userUUID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		appError := apperrors.UserControllerUpdateUserUuidParse.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	user, err := uc.userUsecase.GetUser(ctx.Request().Context(), userUUID)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}
	if user == nil {
		appError := apperrors.UserControllerUpdateUserUserNotExist
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}
	if err = ctx.Bind(updateUser); err != nil {
		appError := apperrors.UserControllerUpdateUserBind.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}
	user.MapUpdateUserRequestToUserModel(updateUser)

	authUser := uc.FetchJWTUser(ctx)
	if user.IsAdmin() && !authUser.IsAdmin() {
		appError := apperrors.UserControllerTryToSetAdmin.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	err = user.ComparePasswords(user.Password)
	if err != nil {
		err = user.HashPassword()
	}
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	err = ctx.Validate(user)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	_, err = uc.userUsecase.CheckUserByNickname(ctx.Request().Context(), user)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	updatedUser, err := uc.userUsecase.UpdateUser(ctx.Request().Context(), user)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	return ctx.JSON(http.StatusOK, updatedUser.MapUserModelToUpdateUserResponse())
}

func (uc *userController) DeleteUser(ctx echo.Context) error {
	userID := ctx.Param("id")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		appError := apperrors.UserControllerDeleteUserUuidParse.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	err = uc.userUsecase.DeleteUser(ctx.Request().Context(), &userUUID)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	return ctx.JSON(http.StatusOK, userUUID)
}

func (uc *userController) GetUser(ctx echo.Context) error {
	userUUID := ctx.Param("id")

	uid, err := uuid.Parse(userUUID)
	if err != nil {
		appError := apperrors.UserControllerGetUserUuidParse.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	user, err := uc.userUsecase.GetUserByID(ctx.Request().Context(), uid)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}
	if user == nil {
		appError := apperrors.UserControllerGetUserUserNotExist
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}
	return ctx.JSON(http.StatusOK, user.MapUserModelToGetUserResponse())
}

func (uc *userController) GetUsers(ctx echo.Context) error {
	paginationQuery, err := utils.GetPaginationFromCtx(ctx.QueryParam("page"), ctx.QueryParam("size"), ctx.QueryParam("orderBy"))
	if err != nil {
		appError := apperrors.UserControllerGetUsersGetPaginationFromCtx.AppendMessage(err)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	users, err := uc.userUsecase.GetUsersByPaginationQuery(ctx.Request().Context(), paginationQuery)
	if err != nil {
		appError := err.(*apperrors.AppError)
		return ctx.JSON(appError.HTTPCode, appError.Error())
	}

	return ctx.JSON(http.StatusOK, users.MapUserModelToGetUserResponse())
}

func (uc *userController) FetchAuthUser(ctx echo.Context, UserAuthCtx string) *model.User {
	return ctx.Get(UserAuthCtx).(*model.User)
}

func (uc *userController) FetchJWTUser(ctx echo.Context) *model.User {
	userContext := ctx.Get("user").(*jwt.Token)
	claims := userContext.Claims.(*model.JwtCustomClaims)

	if !userContext.Valid {
		apperrors.UserControllerFetchJWTUser.AppendMessage(echo.ErrUnauthorized)
	}

	return &model.User{
		UserID:   claims.UserID,
		Nickname: claims.Nickname,
		Role:     claims.Role,
	}
}
