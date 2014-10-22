package hydra_worker_pilot_client_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHydraWorkerPilotClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HydraWorkerPilotClient Suite")
}
