package specrunner_test

import (
	. "github.com/innotech/hydra-worker-pilot-client/vendors/github.com/onsi/ginkgo"
	. "github.com/innotech/hydra-worker-pilot-client/vendors/github.com/onsi/gomega"

	"testing"
)

func TestSpecRunner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Spec Runner Suite")
}
