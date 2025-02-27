package services

import (
	"log"
	"planning/internal/clients"
	"planning/internal/models"
	"planning/internal/repository"
	"planning/internal/resource/request"
	"sync"
)

type ITaskService interface {
	DistributeTasks(developers []request.Developer) (map[string][]models.Task, int, error)
}

type taskService struct {
	repo          repository.IRepository
	taskScheduler ITaskSchedulerService
	providers     []clients.IProvider
}

func NewTaskService(repo repository.IRepository, taskScheduler ITaskSchedulerService, providers []clients.IProvider) ITaskService {
	return &taskService{
		repo:          repo,
		taskScheduler: taskScheduler,
		providers:     providers,
	}
}

func (t *taskService) DistributeTasks(developers []request.Developer) (map[string][]models.Task, int, error) {
	tasks, err := t.FetchTasks()
	if err != nil {
		log.Println("[TaskService] Failed to fetch tasks", err)
		return nil, 0, err
	}

	if err := t.StoreTasks(tasks); err != nil {
		log.Println("[TaskService] Failed to store tasks", err)
		return nil, 0, err
	}

	developerTasks, totolWeeks := t.taskScheduler.DistributeTasks(developers, tasks)
	return developerTasks, totolWeeks, nil
}

func (t *taskService) FetchTasks() ([]models.Task, error) {
	var tasks []models.Task
	var mutex sync.Mutex
	var wg sync.WaitGroup
	errChan := make(chan error, len(t.providers))

	for _, provider := range t.providers {
		wg.Add(1)
		go func(p clients.IProvider) {
			defer wg.Done()

			providerTasks, err := p.FetchTasks()
			if err != nil {
				log.Println("[TaskService] Failed to fetch tasks from provider", err)
				errChan <- err
				return
			}
			mutex.Lock()
			tasks = append(tasks, providerTasks...)
			mutex.Unlock()
		}(provider)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	close(errChan)

	// Check if any errors occurred
	if err := <-errChan; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *taskService) StoreTasks(tasks []models.Task) error {
	for _, task := range tasks {
		if err := t.repo.GetTaskRepository().StoreTask(&task); err != nil {
			log.Println("[TaskService] Failed to store task:", task, err)
			return err
		}
	}
	return nil
}

func (t *taskService) GetTasks() ([]models.Task, error) {
	tasks, err := t.repo.GetTaskRepository().GetTasks()
	if err != nil {
		log.Println("[TaskService] Failed to get tasks", err)
		return nil, err
	}
	return tasks, nil
}
