package clients

import (
	"encoding/json"
	"net/http"
	"planning/internal/models"
)

type provider1 struct {
	URI string
}

func NewProvider1(uri string) IProvider {
	return &provider1{
		URI: uri,
	}
}

func (p *provider1) FetchTasks() ([]models.Task, error) {
	resp, err := http.Get(p.URI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tasks []models.Task
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
