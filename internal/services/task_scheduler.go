package services

import (
	"math"
	"planning/internal/models"
	"planning/internal/resource/request"
	"sort"
)

type ITaskSchedulerService interface {
	DistributeTasks(developers []request.Developer, tasks []models.Task) (map[string][]models.Task, int)
}

type taskSchedulerService struct{}

func NewTaskSchedulerService() ITaskSchedulerService {
	return &taskSchedulerService{}
}

func (s *taskSchedulerService) DistributeTasks(developers []request.Developer, tasks []models.Task) (map[string][]models.Task, int) {
	// Sort tasks by difficulty in descending order
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Difficulty > tasks[j].Difficulty
	})

	// Initialize developer tasks
	developerTasks := make(map[string][]models.Task)
	for _, dev := range developers {
		developerTasks[dev.Name] = []models.Task{}
	}

	totalWeeks := 1
	remainingTasks := make([]models.Task, len(tasks))
	// Make a copy of tasks to avoid modifying the original slice
	copy(remainingTasks, tasks)

	for len(remainingTasks) > 0 {
		taskAssigned := false
		for i := 0; i < len(remainingTasks); i++ {
			task := remainingTasks[i]
			for j := range developers {
				requiredHours := int(math.Ceil(float64(task.Duration) * float64(task.Difficulty) / float64(developers[j].Capacity)))
				if requiredHours <= developers[j].CurrentHours {
					developerTasks[developers[j].Name] = append(developerTasks[developers[j].Name], task)
					developers[j].CurrentHours -= requiredHours
					remainingTasks = append(remainingTasks[:i], remainingTasks[i+1:]...)
					taskAssigned = true
					break
				}
			}
			if taskAssigned {
				break
			}
		}

		if !taskAssigned {
			totalWeeks++
			for j := range developers {
				developers[j].CurrentHours = developers[j].WeeklyHours
			}
		}
	}

	return developerTasks, totalWeeks
}
