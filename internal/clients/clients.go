package clients

import "planning/internal/models"

type IClients interface {
	GetProvider1() IProvider
	GetProvider2() IProvider
}

// IProvider interface
// All the providers should implement this interface
// If provider3, provider4, etc. will be added in the future
// they should implement this interface
type IProvider interface {
	FetchTasks() ([]models.Task, error)
}

type clients struct {
	provider1 IProvider
	provider2 IProvider
}

func New(provider1, provider2 IProvider) IClients {
	return &clients{
		provider1: provider1,
		provider2: provider2,
	}
}

func (c *clients) GetProvider1() IProvider {
	return c.provider1
}

func (c *clients) GetProvider2() IProvider {
	return c.provider2
}
