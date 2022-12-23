package fruits_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFruits(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fruits Suite")
}
