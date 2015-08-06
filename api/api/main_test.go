package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api Suite")
}

var _ = Describe("main", func() {
	It("add 1 + 1 = 2", func() {
		expect := 2
		actual := Add(1, 1)

		Expect(actual).To(Equal(expect))
	})
})
