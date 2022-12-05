package greeting

import "go-testing-blueprint/internal/model"

type GreetingService interface {
	GetGreeting(name string) model.Greeting
	PostGreeting(greeting model.Greeting) model.Greeting
}

func NewGreetingService() GreetingService {
	return &Greeting{}
}
