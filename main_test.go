package hydra_worker_pilot_client_test

import (
	. "github.com/innotech/hydra-worker-pilot-client"

	. "github.com/innotech/hydra-worker-pilot-client/vendors/github.com/onsi/ginkgo"
	. "github.com/innotech/hydra-worker-pilot-client/vendors/github.com/onsi/gomega"
)

var _ = Describe("PilotClient", func() {
	var (
		inputInstances []interface{}
		workerArgs     map[string]interface{}
		requestParams  map[string][]string
	)

	BeforeEach(func() {
		inputInstances = []interface{}{
			map[string]interface{}{
				"Info": map[string]interface{}{
					"version": "1.0.0",
				},
			},
			map[string]interface{}{
				"Info": map[string]interface{}{
					"version": "1.1.0",
				},
			},
		}
		workerArgs = map[string]interface{}{
			InstanceFilterFieldKey: "version",
			MatchersKey: []map[string]interface{}{
				map[string]interface{}{
					InstanceFilterPatternKey: "1.*",
					ClientFilterPatternsKey:  []string{"xe4([0-9]+)", "xe5([0-9]+)"},
				},
				map[string]interface{}{
					InstanceFilterPatternKey: `1\.0\.0`,
					ClientFilterPatternsKey:  []string{".*"},
				},
			},
			ClientFilterFieldKey: "client_uuid",
		}
		requestParams = map[string][]string{
			"client_uuid": []string{"xe47030"},
		}
	})

	Context("when doesn't exist client filter field", func() {
		It("should return input instances", func() {
			delete(workerArgs, "clientFilterField")
			outputInstances := Pilot(inputInstances, requestParams, workerArgs)
			Expect(outputInstances).To(Equal(inputInstances))
		})
	})
	Context("when invalid client filter field", func() {
		It("should return input instances", func() {
			workerArgs["clientFilterField"] = ""
			outputInstances := Pilot(inputInstances, requestParams, workerArgs)
			Expect(outputInstances).To(Equal(inputInstances))
		})
	})
	Context("when doesn't exist instance filter field", func() {
		It("should return input instances", func() {
			delete(workerArgs, "instanceFilterField")
			outputInstances := Pilot(inputInstances, requestParams, workerArgs)
			Expect(outputInstances).To(Equal(inputInstances))
		})
	})
	Context("when invalid instance filter field", func() {
		It("should return input instances", func() {
			workerArgs["instanceFilterField"] = ""
			outputInstances := Pilot(inputInstances, requestParams, workerArgs)
			Expect(outputInstances).To(Equal(inputInstances))
		})
	})
	Context("when doesn't exist matchers", func() {
		It("should return input instances", func() {
			delete(workerArgs, "matchers")
			outputInstances := Pilot(inputInstances, requestParams, workerArgs)
			Expect(outputInstances).To(Equal(inputInstances))
		})
	})
	Context("when invalid matchers", func() {
		It("should return input instances", func() {
			workerArgs["matchers"] = ""
			outputInstances := Pilot(inputInstances, requestParams, workerArgs)
			Expect(outputInstances).To(Equal(inputInstances))
		})
	})
	Context("when doesn't exist client filter value", func() {
		It("should return input instances", func() {
			delete(requestParams, "client_uuid")
			outputInstances := Pilot(inputInstances, requestParams, workerArgs)
			Expect(outputInstances).To(Equal(inputInstances))
		})
	})
	Context("when invalid client filter value", func() {
		It("should return input instances", func() {
			requestParams["client_uuid"] = []string{}
			outputInstances := Pilot(inputInstances, requestParams, workerArgs)
			Expect(outputInstances).To(Equal(inputInstances))
		})
	})
	Context("when client parameter doesn't match with any configured worker matcher", func() {
		It("should return input instances", func() {
			requestParams["client_uuid"] = []string{"xe10234"}
			workerArgs[MatchersKey].([]map[string]interface{})[1][ClientFilterPatternsKey] = []string{"xe6([0-9]+)"}
			outputInstances := Pilot(inputInstances, requestParams, workerArgs)
			Expect(outputInstances).To(Equal(inputInstances))
		})
	})
	Context("when client parameter matchs with some configured worker matcher", func() {
		Context("when doesn't exist instance filter value", func() {
			It("should return input instances", func() {
				for i := 0; i < len(inputInstances); i++ {
					delete(inputInstances[i].(map[string]interface{})["Info"].(map[string]interface{}), "version")
				}
				outputInstances := Pilot(inputInstances, requestParams, workerArgs)
				Expect(outputInstances).To(Equal(inputInstances))
			})
		})
		Context("when invalid instance filter value", func() {
			It("should return input instances", func() {
				for i := 0; i < len(inputInstances); i++ {
					inputInstances[i].(map[string]interface{})["Info"].(map[string]interface{})["version"] = ""
				}
				outputInstances := Pilot(inputInstances, requestParams, workerArgs)
				Expect(outputInstances).To(Equal(inputInstances))
			})
		})
		Context("when doesn't exist compatible instances", func() {
			It("should return input instances", func() {
				workerArgs[MatchersKey].([]map[string]interface{})[1][InstanceFilterPatternKey] = "2.*"
				outputInstances := Pilot(inputInstances, requestParams, workerArgs)
				Expect(outputInstances).To(Equal(inputInstances))
			})
		})
		Context("when exists compatible instances", func() {
			It("should return input instances", func() {
				requestParams["client_uuid"] = []string{"xe10234"}
				outputInstances := Pilot(inputInstances, requestParams, workerArgs)
				Expect(outputInstances).To(HaveLen(1))
				Expect(outputInstances).To(Equal(inputInstances[0:1]))
			})
		})
	})
})
