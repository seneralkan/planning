package repository

type IRepository interface {
	GetTaskRepository() ITasksRepository
}

type repository struct {
	tasksRepository ITasksRepository
}

func NewRepository(tasksRepository ITasksRepository) IRepository {
	return &repository{
		tasksRepository: tasksRepository,
	}
}

func (r *repository) GetTaskRepository() ITasksRepository {
	return r.tasksRepository
}
