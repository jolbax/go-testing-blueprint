package greeting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-testing-blueprint/internal/controller/greeting"
	"go-testing-blueprint/internal/model"
	"testing"
)

func TestGreeting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Greeting Suite")
}

var g = greeting.Greeting{}
var _ = Describe("./Internal/Controller/Greeting", func() {
	Context("GetGreeting", func() {
		It("should get a greeting", func() {
			Expect(g.GetGreeting("Samba")).
				To(Equal(model.Greeting{Text: "Hallo Samba", Language: "de"}))
		})
		It("should get an empty greeting", func() {
			Expect(g.GetGreeting("")).
				To(Equal(model.Greeting{Text: "Hallo ", Language: "de"}))
		})
	})

	Context("PostGretting", func() {
		It("should post a greeting", func() {
			Expect(g.PostGreeting(model.Greeting{Text: "Hello Samba", Language: "en"})).
				To(Equal(model.Greeting{Text: "You posted a greeting: Hello Samba", Language: "en"}))
		})
		It("should post an empty greeting", func() {
			Expect(g.PostGreeting(model.Greeting{Text: "", Language: "en"})).
				To(Equal(model.Greeting{Text: "You posted a greeting: ", Language: "en"}))
		})
	})

})
