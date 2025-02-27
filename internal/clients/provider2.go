package clients

import (
	"encoding/json"
	"net/http"
	"planning/internal/models"
)

type provider2 struct {
	URI string
}

func NewProvider2(uri string) IProvider {
	return &provider2{
		URI: uri,
	}
}

func (p *provider2) FetchTasks() ([]models.Task, error) {
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
