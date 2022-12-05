package greeting

import "go-testing-blueprint/internal/model"

type Greeting struct {
}

func (g Greeting) GetGreeting(name string) model.Greeting {
	return model.Greeting{"de", "Hallo " + name}
}

func (g Greeting) PostGreeting(greeting model.Greeting) model.Greeting {
	return model.Greeting{greeting.Language, "You posted a greeting: " + greeting.Text}
}
