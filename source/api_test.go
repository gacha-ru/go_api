package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"os"
	"fmt"
)

var _ = Describe("Book", func() {

	response, err := http.Get("http://localhost:3000")
	if err != nil {
		fmt.Println("Connection Failed!!")
		os.Exit(2)
	}

	Context("Check response", func() {
		It("/", func() {
			Expect(response.Status).To(Equal("200 OK"))
		})
	})

	happy_response, happy_err := http.Get("http://localhost:3000/happy")
	if happy_err != nil {
		os.Exit(2)
	}

	Context("Check response", func() {
		It("/happy", func() {
			Expect(happy_response.Status).To(Equal("200 OK"))
		})
	})

})
