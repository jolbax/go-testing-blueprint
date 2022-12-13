package fruits

import (
	fruits_client "go-testing-blueprint/pkg/fruits-client"
)

type FruitsService interface {
	GetFruits() (*fruits_client.FruitsResponse, error)
	PostFruits() (*fruits_client.FruitsResponse, error)
}

func NewFruitsService(fruitsClient FruitClientIf) FruitsService {
	return &Fruits{c: fruitsClient}
}

// Fruits FruitClientIf abstraction
type FruitClientIf interface {
	Get() (*fruits_client.FruitsResponse, error)
	Post() (*fruits_client.FruitsResponse, error)
}
