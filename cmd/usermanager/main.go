package main

import (
	"usermanager/internal/apperrors"
	"usermanager/internal/config"
	"usermanager/internal/infrastructure/datastore"
	"usermanager/internal/infrastructure/logger"
	"usermanager/internal/infrastructure/router"
	"usermanager/internal/registry"

	"github.com/labstack/echo/v4"
)

const dotEnv = "./configs/.env"

func main() {
	logger := logger.NewLogger()
	cfg, err := config.NewConfig(dotEnv)
	if err != nil {
		logger.Fatal(err)
	}

	db, err := datastore.NewDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	redisClient, err := datastore.NewRedisClient(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	reg := registry.NewRegistry(db, redisClient, cfg)

	e := echo.New()
	e = router.NewRouter(e, reg.NewAppController())

	logger.Println("app starting")

	err = e.Start(cfg.Port)
	if err != nil {
		logger.Fatal(apperrors.ServerStartError.AppendMessage(err))
	}
}
