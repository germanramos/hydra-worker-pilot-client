package asyncassertion_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/innotech/hydra-worker-pilot-client/vendors/github.com/onsi/gomega"

	"testing"
)

func TestAsyncAssertion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AsyncAssertion Suite")
}
