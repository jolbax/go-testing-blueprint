package router_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	. "go-testing-blueprint/internal/router"
	"go-testing-blueprint/mocks"
	"testing"
)

func TestNewRouter(t *testing.T) {
	mockRegiterable := &mocks.Registerable{}

	mockRegiterable.
		On("RegisterRoutes", mock.Anything).
		Return().
		Once()

	router := NewRouter(mockRegiterable)

	assert.IsType(t, &gin.Engine{}, router, "Is type gin.Engine")
	mockRegiterable.AssertExpectations(t)

}
