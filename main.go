package hydra_worker_pilot_client

import (
	"errors"
	"log"
	"os"
	"regexp"

	worker "github.com/innotech/hydra-worker-pilot-client/vendors/github.com/innotech/hydra-worker-lib"
)

const (
	clientFilterFieldKey     string = "clientFilterField"
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
	pilotClientWorker := worker.NewWorker(os.Args)
	fn := func(instances []interface{}, requestParams map[string][]string, workerArgs map[string]interface{}) (finalInstances []interface{}) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Error: ", r)
			}
		}()
		finalInstances = instances
		var tmpInstances []interface{}

		// TODO: Maybe only need call to obtainClientFilterValue
		clientFilterField, err := obtainClientFilterField(workerArgs)
		if err != nil {
			log.Println(err.Error())
			return instances
		}
		clientFilterValue, err := obtainClientFilterValue(requestParams, clientFilterField)
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
		for _, matcher := range matchers {
			for _, pattern := range matcher.clientFilterPatterns {
				r, err := regexp.Compile(pattern)
				if err != nil {
					log.Println("Invalid regexp pattern: " + pattern)
				}
				matched := r.MatchString(clientFilterValue)
				// TODO:
				if matched == false {
					continue
				}
				tmpInstances = findCompatibleInstances(instances, instanceFilterField, matcher.instanceFilterPattern)
			}
		}
		finalInstances = tmpInstances
		return
	}
	pilotClientWorker.Run(fn)
}

func findCompatibleInstances(instances []interface{}, instanceFilterField string, instanceFilterPattern string) []interface{} {
	finalInstances := make([]interface{}, 0)
	var err error
	var instanceFilterValue string
	var instance map[string]interface{}
	for _, rawInstance := range instances {
		instance = rawInstance.(map[string]interface{})
		instanceFilterValue, err = obtainInstanceFilterValue(instance, instanceFilterField)
		if err != nil {
			continue
		}
		r, err := regexp.Compile(instanceFilterPattern)
		if err != nil {
			log.Println("Invalid regexp pattern: " + instanceFilterPattern)
		}
		matched := r.MatchString(instanceFilterValue)
		if matched == false {
			continue
		}
		finalInstances = append(finalInstances, instance)
	}
	return finalInstances
}

func obtainClientFilterField(workerArgs map[string]interface{}) (string, error) {
	if val, ok := workerArgs[clientFilterFieldKey]; ok && val != "" {
		return val.(string), nil
	}
	return "", errors.New("Invalid clientFilterField")
}

func obtainClientFilterValue(requestParams map[string][]string, param string) (string, error) {
	if val, ok := requestParams[param]; ok && len(val) > 0 && val[0] != "" {
		return val[0], nil
	}
	return "", errors.New("Invalid clientFilterValue")
}

func obtainInstanceFilterField(workerArgs map[string]interface{}) (string, error) {
	if val, ok := workerArgs[instanceFilterFieldKey]; ok && val != "" {
		return val.(string), nil
	}
	return "", errors.New("Invalid instanceFilterField")
}

func obtainMatchers(workerArgs map[string]interface{}) ([]Matcher, error) {
	// TODO: Maybe move the recover function in all obtain (not string casting)
	if val, ok := workerArgs[matchersKey]; ok && val != "" {
		rawMatchers := val.([]map[string]interface{})
		rawMatchersLen := len(rawMatchers)
		matchers := make([]Matcher, rawMatchersLen, rawMatchersLen)
		for i := 0; i < rawMatchersLen; i++ {
			ifp, _ := obtainInstanceFilterPattern(rawMatchers[i])
			cfp, _ := obtainClientFilterPatterns(rawMatchers[i])
			matchers[i] = Matcher{
				instanceFilterPattern: ifp,
				clientFilterPatterns:  cfp,
			}
		}
		return matchers, nil
	}
	return nil, errors.New("Invalid matchers")
}

func obtainInstanceFilterPattern(matcher map[string]interface{}) (string, error) {
	if val, ok := matcher[instanceFilterPatternKey]; ok && val != "" {
		return val.(string), nil
	}
	return "", errors.New("Invalid instanceFilterPattern")
}

func obtainClientFilterPatterns(matcher map[string]interface{}) ([]string, error) {
	if val, ok := matcher[clientFilterPatternsKey]; ok && val != "" {
		return val.([]string), nil
	}
	return nil, errors.New("Invalid clientFilterPatterns")
}

func obtainInstanceFilterValue(instance map[string]interface{}, filterField string) (string, error) {
	if val, ok := instance[filterField]; ok && val != "" {
		return val.(string), nil
	}
	return "", errors.New("Invalid instanceFilterValue")
}
