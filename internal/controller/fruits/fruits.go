package fruits

import (
	client "go-testing-blueprint/pkg/fruits-client"
)

type Fruits struct {
	c FruitClientIf
}

func (f Fruits) GetFruits() (*client.FruitsResponse, error) {
	r, err := f.c.Get()
	if err != nil {
		return r, err
	}
	return r, nil
}

func (f Fruits) PostFruits() (*client.FruitsResponse, error) {
	r, err := f.c.Post()
	if err != nil {
		return r, err
	}
	return r, nil
}
