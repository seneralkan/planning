package planning

import (
	"planning/internal/clients"
	"planning/internal/handler"
	"planning/internal/repository"
	"planning/internal/router"
	"planning/internal/services"
	"planning/pkg/sqlite"
)

func CreateRouter(
	db sqlite.ISqliteInstance,
	providers []clients.IProvider,
) router.IRouter {
	taskRepository := repository.NewTaskRepository(db)
	repository := repository.NewRepository(taskRepository)

	taskScheduler := services.NewTaskSchedulerService()
	taskService := services.NewTaskService(repository, taskScheduler, providers)

	taskHandler := handler.NewTaskHandler(taskService)
	iRouter := router.NewRouter(taskHandler)
	return iRouter
}

func CreateClientProviders(
	provider1URI string,
	provider2URI string,
) []clients.IProvider {
	provider1 := clients.NewProvider1(provider1URI)
	provider2 := clients.NewProvider2(provider2URI)
	return []clients.IProvider{provider1, provider2}
}
