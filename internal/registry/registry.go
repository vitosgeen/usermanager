package registry

import (
	"usermanager/internal/config"
	"usermanager/internal/infrastructure/datastore"
	"usermanager/internal/interface/controller"
)

type registry struct {
	db    *datastore.DB
	redis *datastore.Redis
	cfg   *config.Config
}

type Registry interface {
	NewAppController() controller.UserManagerController
}

func NewRegistry(db *datastore.DB, redis *datastore.Redis, cfg *config.Config) Registry {
	return &registry{
		db:    db,
		redis: redis,
		cfg:   cfg,
	}
}

func (r *registry) NewAppController() controller.UserManagerController {
	return controller.UserManagerController{
		UserController: r.NewUserController(),
	}
}
