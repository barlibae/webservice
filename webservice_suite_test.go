package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// go install github.com/onsi/ginkgo/v2/ginkgo and ginkgo bootstrap to automatically generate this file
func TestWebservice(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Webservice Suite")
}
