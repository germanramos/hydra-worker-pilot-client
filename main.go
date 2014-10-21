package main

import (
	"errors"
	"log"
	"os"

	worker "github.com/innotech/hydra-worker-sort-by-number/vendors/github.com/innotech/hydra-worker-lib"
)

const (
	clientFilterValueKey     string = "clientFilterValue"
	instanceFilterFieldKey   string = "instanceFilterField"
	matchersKey              string = "matchers"
	instanceFilterPatternKey string = "instanceFilterPattern"
	clientFilterPatternsKey  string = "clientFilterPatterns"
)

type Matcher struct {
	instanceFilterPattern string
	clientFilterPatterns  []string
}

func main() {
	// New Worker connected to Hydra Load Balancer
	sortByNumberWorker := worker.NewWorker(os.Args)
	fn := func(instances []interface{}, requestParams map[string][]string, args map[string]interface{}) (finalInstances []interface{}) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Error: ", r)
			}
		}()
		finalInstances = instances
		var tmpInstances []interface{}

		clientFilterField, err := obtainClientFilterField(requestParams, workerArgs)
		if err != nil {
			log.Println(err.Error())
			return instances
		}
		instanceFilterField, err := obtainInstanceFilterField(workerArgs)
		if err != nil {
			log.Println(err.Error())
			return instances
		}
		matchers, err := obtainMatchers(workerArgs)
		if err != nil {
			log.Println(err.Error())
			return instances
		}
		var instanceFilterValue string
		for _, instance := range instances.([]map[string]interface{}) {
			instanceFilterValue = obtainInstanceFilterValue(instance, instanceFilterField)
			for instanceMatchValue, clientMatchPattern := range matchers {

			}
			// for i := 0; i < count; i++ {

			// }
		}
		// return instances
	}
	sortByNumberWorker.Run(fn)
}

func obtainClientFilterField(requestParams map[string][]string, workerArgs map[string]interface{}) (string, error) {
	if val, ok := workerArgs[clientFilterValueKey]; ok && val != "" {
		return val.(string), nil
	}
	return "", errors.New("Invalid clientFilterValue")
}

func obtainInstanceFilterField(workerArgs map[string]interface{}) (string, error) {
	if val, ok := workerArgs[instanceFilterFieldKey]; ok && val != "" {
		return val.(string), nil
	}
	return "", errors.New("Invalid instanceFilterField")
}

// TODO: Change
func obtainMatchers(workerArgs map[string]interface{}) ([]Matcher, error) {
	// TODO: Maybe move the recover function in all obtain (not string casting)
	// if val, ok := workerArgs[matchersKey]; ok && val != "" {

	// 	return val.(map[string][]string), nil
	// }
	// return "", errors.New("Invalid instanceFilterField")
}

func obtainInstanceFilterPattern(matcher map[string]interface{}) string {
	if val, ok := matcher[instanceFilterPatternKey]; ok && val != "" {
		return val.(string), nil
	}
	return "", errors.New("Invalid instanceFilterPattern")
}

func obtainClientFilterPatterns(matcher map[string]interface{}) []string {
	if val, ok := matcher[clientFilterPatternsKey]; ok && val != "" {
		return val.([]string), nil
	}
	return "", errors.New("Invalid clientFilterPatterns")
}

func obtainInstanceFilterValue(instance map[string]interface{}, filterField string) (string, error) {
	if val, ok := instance[filterField]; ok && val != "" {
		return val.(string), nil
	}
	return "", errors.New("Invalid instanceFilterValue")
}
