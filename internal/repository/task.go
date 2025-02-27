package repository

import (
	"database/sql"
	"log"
	"planning/internal/models"
	"planning/pkg/sqlite"
)

type ITasksRepository interface {
	StoreTask(task *models.Task) error
	GetTasks() ([]models.Task, error)
}

type tasksRepository struct {
	db *sql.DB
}

func NewTaskRepository(sqlite sqlite.ISqliteInstance) ITasksRepository {
	return &tasksRepository{
		db: sqlite.Database(),
	}
}

func (t tasksRepository) StoreTask(task *models.Task) error {
	query := `
	INSERT INTO tasks (name, duration, difficulty) VALUES (?, ?, ?);
	`
	_, err := t.db.Exec(query, task.Name, task.Duration, task.Difficulty)

	if err != nil {
		log.Println("[TaskRepository] Failed to store task", err)
	}

	return nil
}

func (t tasksRepository) GetTasks() ([]models.Task, error) {
	query := `SELECT name, duration, difficulty FROM tasks;`
	rows, err := t.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.Name, &task.Duration, &task.Difficulty); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
