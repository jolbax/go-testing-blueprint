package fruits

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-testing-blueprint/mocks"
	fruits_client "go-testing-blueprint/pkg/fruits-client"
	"testing"
)

func TestFruits_GetFruits(t *testing.T) {
	mockClient := new(mocks.FruitClientIf)

	okExpected := fruits_client.FruitsResponse{Fruits: &fruits_client.Fruits{Fruits: []fruits_client.Fruit{"mango", "pineapple", "grapefruit"}}, Code: 200}
	nokExpected := fruits_client.FruitsResponse{Fruits: nil, Code: 400}

	service := NewFruitsService(mockClient)

	mockClient.On("Get", mock.Anything).
		Return(&okExpected, nil).
		Once()

	{
		actual, err := service.GetFruits()

		assert.Equal(t, okExpected.Fruits, actual.Fruits)
		assert.Equal(t, 200, actual.Code)
		assert.Equal(t, okExpected.Code, 200)
		assert.NotEqual(t, 400, actual.Code)
		assert.IsType(t, fruits_client.FruitsResponse{}, okExpected)
		assert.Nil(t, err)

	}

	mockClient.On("Get", mock.Anything).
		Return(&nokExpected, errors.New("Something went wrong")).
		Once()

	{
		actual, err := service.GetFruits()

		assert.Equal(t, &nokExpected, actual)
		assert.Equal(t, nokExpected.Fruits, actual.Fruits)
		assert.Equal(t, 400, actual.Code)
		assert.Equal(t, nokExpected.Code, 400)
		assert.NotEqual(t, 200, actual.Code)
		assert.NotNil(t, err)

	}

	mockClient.AssertExpectations(t)
}

func TestFruits_PostFruits(t *testing.T) {
	mockClient := new(mocks.FruitClientIf)

	okExpected := fruits_client.FruitsResponse{Fruits: &fruits_client.Fruits{Fruits: []fruits_client.Fruit{"mango", "pineapple", "grapefruit"}}, Code: 200}
	nokExpected := fruits_client.FruitsResponse{Fruits: nil, Code: 400}

	service := NewFruitsService(mockClient)

	mockClient.On("Post", mock.Anything).
		Return(&okExpected, nil).
		Once()

	{
		actual, err := service.PostFruits()

		assert.Equal(t, okExpected.Fruits, actual.Fruits)
		assert.Equal(t, 200, actual.Code)
		assert.Equal(t, okExpected.Code, 200)
		assert.NotEqual(t, 400, actual.Code)
		assert.IsType(t, fruits_client.FruitsResponse{}, okExpected)
		assert.Nil(t, err)

	}

	mockClient.On("Post", mock.Anything).
		Return(&nokExpected, errors.New("Something went wrong")).
		Once()

	{
		actual, err := service.PostFruits()

		assert.Equal(t, &nokExpected, actual)
		assert.Equal(t, nokExpected.Fruits, actual.Fruits)
		assert.Equal(t, 400, actual.Code)
		assert.Equal(t, nokExpected.Code, 400)
		assert.NotEqual(t, 200, actual.Code)
		assert.NotNil(t, err)

	}

	mockClient.AssertExpectations(t)

}
