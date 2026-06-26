package e2e

import (
	"testing"

	_ "github.com/devsy-org/devsy-provider-kubernetes/e2e/pullsecrets" // Register tests.
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func TestRunE2ETests(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Devsy provider kubernetes e2e suite")
}
