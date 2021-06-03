package adders_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAdders(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Adders Suite")
}
