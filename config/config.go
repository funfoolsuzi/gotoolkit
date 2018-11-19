package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const configf string = "config.json"

// Configuration describe configuration of the app
type Configuration struct {
	ProjectID string
	Database  struct {
		Limit int
	}
	MinutesPerSession int
}

// GetConfigNameFromArgs will return the configuration name. Default is "dev"
func GetConfigNameFromArgs() string {
	args := os.Args[1:]

	if len(args) <= 0 {
		return "dev"
	}

	envArg := args[0]

	return envArg
}

// GetConfigFromArgs will return the Configuration.
// Because it is a main process. It would panic if anything fail during the process.
func GetConfigFromArgs() (*Configuration, string) {
	envName := GetConfigNameFromArgs()

	p, _ := os.Getwd()

	envs := GetConfigMap(path.Join(p, configf))

	env, envExist := (*envs)[envName]
	if !envExist {
		panic("Can't find envName:" + envName + " among all configs in " + configf)
	}
	return &env, envName
}

// GetConfigMap will return a map of Configuration
func GetConfigMap(path string) *map[string]Configuration {
	rawConfig, errRead := ioutil.ReadFile(path)
	if errRead != nil {
		fmt.Println(errRead)
		panic("Can't find " + configf + ". Quiting.")
	}

	envs := map[string]Configuration{}
	if err := json.Unmarshal(rawConfig, &envs); err != nil {
		fmt.Println(err)
		panic("Can't parse content from " + configf + " to Configuration. Quiting.")
	}
	return &envs
}
