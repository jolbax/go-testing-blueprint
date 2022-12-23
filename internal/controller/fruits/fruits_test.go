package fruits_test

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"go-testing-blueprint/internal/controller/fruits"
	"go-testing-blueprint/mocks"
	fruits_client "go-testing-blueprint/pkg/fruits-client"
)

var _ = Describe("Fruits", func() {

	okExpected := fruits_client.FruitsResponse{Fruits: &fruits_client.Fruits{Fruits: []fruits_client.Fruit{"mango", "pineapple", "grapefruit"}}, Code: 200}
	nokExpected := fruits_client.FruitsResponse{Fruits: nil, Code: 400}

	mockClient := new(mocks.FruitClientIf)
	service := fruits.NewFruitsService(mockClient)

	Describe("Getting fruits", func() {
		Context("with expected values", func() {
			It("should return the same object with 200 code and no errors", func() {
				mockClient.On("Get", mock.Anything).Return(&okExpected, nil).Once()
				actual, err := service.GetFruits()
				Expect(*actual).To(Equal(okExpected))
				Expect(actual.Code).To(Equal(200))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with unexpected values", func() {
			It("should return an nil object with an errors", func() {
				mockClient.On("Get", mock.Anything).Return(&nokExpected, errors.New("an error occurred")).Once()
				actual, err := service.GetFruits()
				Expect(actual).To(Equal(&nokExpected))
				Expect(actual.Code).To(Equal(400))
				Expect(actual.Code).NotTo(Equal(200))
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe("Posting fruits", func() {
		Context("with expected values", func() {
			It("should return the same object with 200 code and no errors", func() {
				mockClient.On("Post", mock.Anything).Return(&okExpected, nil).Once()
				actual, err := service.PostFruits()
				Expect(*actual).To(Equal(okExpected))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with expected values", func() {
			It("should return an empty object with an errors", func() {
				mockClient.On("Post", mock.Anything).Return(&nokExpected, errors.New("an error occurred")).Once()
				actual, err := service.PostFruits()
				Expect(actual).To(Equal(&nokExpected))
				Expect(err).Should(HaveOccurred())
			})
		})

	})

	When(" tests finish the mockClient should be called as many times as On() was called", func() {
		It("shouldn't return an error", func() {
			Expect(mockClient.AssertExpectations(GinkgoT())).To(BeTrue())
		})
	})
})
