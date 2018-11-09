package ginserver_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinserver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginserver Suite")
}
