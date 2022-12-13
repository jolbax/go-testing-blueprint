package greeting

import "go-testing-blueprint/internal/model"

type Greeting struct {
}

func (g Greeting) GetGreeting(name string) model.Greeting {
	return model.Greeting{Language: "de", Text: "Hallo " + name}
}

func (g Greeting) PostGreeting(greeting model.Greeting) model.Greeting {
	return model.Greeting{Language: greeting.Language, Text: "You posted a greeting: " + greeting.Text}
}
