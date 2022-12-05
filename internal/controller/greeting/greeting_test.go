package greeting

import (
	"github.com/stretchr/testify/assert"
	"go-testing-blueprint/internal/model"
	"testing"
)

func TestGreeting_GetGreeting(t *testing.T) {
	g := Greeting{}
	assert.Equal(t, model.Greeting{Text: "Hallo Samba", Language: "de"}, g.GetGreeting("Samba"), "Greet")
	assert.Equal(t, model.Greeting{Text: "Hallo ", Language: "de"}, g.GetGreeting(""), "empty Greet")
	assert.NotEqual(t, "Hallo Samba", g.GetGreeting("Samba"), "expecting String")
}

func TestGreeting_PostGreeting(t *testing.T) {
	g := Greeting{}
	assert.Equal(t,
		model.Greeting{Text: "You posted a greeting: Hello Samba", Language: "en"},
		g.PostGreeting(model.Greeting{Text: "Hello Samba", Language: "en"}),
		"Greet")
	assert.Equal(t,
		model.Greeting{Text: "You posted a greeting: ", Language: ""},
		g.PostGreeting(model.Greeting{}),
		"empty Greet")
	assert.NotEqual(t,
		"You posted a greeting: Hello Samba",
		g.PostGreeting(model.Greeting{Text: "Hello Samba", Language: "en"}),
		"expecting String")

}
