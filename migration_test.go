package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "migration/cmd"
)

var _ = Describe("Adder", func() {
	Describe("Add", func() {
		It("adds two numbers", func() {
			sum := Add(2, 3)
			Expect(sum).To(Equal(5))
		})
	})
})
