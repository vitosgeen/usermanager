package router

import (
	"usermanager/internal/interface/controller"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, c controller.UserManagerController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = &controller.CustomValidator{Validator: validator.New()}

	e.POST("/user/login", func(context echo.Context) error { return c.UserController.Login(context) })
	e.GET("/user/:id", func(context echo.Context) error { return c.UserController.GetUser(context) })
	e.GET("/users", func(context echo.Context) error { return c.UserController.GetUsers(context) })

	userGroup := e.Group("/user")
	userGroup.Use(c.UserController.SetUpJWTConfig())
	userGroup.POST("", func(context echo.Context) error { return c.UserController.CreateUser(context) })
	userGroup.DELETE("/:id", func(context echo.Context) error { return c.UserController.DeleteUser(context) }, c.UserController.CanDeleteUser())
	userGroup.PUT("/:id", func(context echo.Context) error { return c.UserController.UpdateUser(context) }, c.UserController.CanUpdateUser())
	userGroup.POST("/vote", func(context echo.Context) error { return c.UserController.VoteUser(context) })

	return e
}
