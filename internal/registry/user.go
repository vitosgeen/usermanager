package registry

import (
	"usermanager/internal/interface/controller"
	"usermanager/internal/interface/repository"
	"usermanager/internal/usecase/usecase"
)

func (r *registry) NewUserController() controller.IUserController {
	userUsecase := usecase.NewUserUsecase(
		repository.NewUserRepository(r.db),
		repository.NewVoteRepository(r.db),
		repository.NewUserRedisRepository(r.redis),
		repository.NewVoteRedisRepository(r.redis),
	)

	return controller.NewUserController(userUsecase, r.cfg)
}
