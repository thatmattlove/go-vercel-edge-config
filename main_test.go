package edgeconfig_test

import (
	"os"

	"github.com/stellaraf/go-utils/environment"
)

type Environment struct {
	EdgeConfigID    string `env:"EDGE_CONFIG_ID"`
	APIToken        string `env:"API_TOKEN"`
	EdgeConfigToken string `env:"EDGE_CONFIG_TOKEN"`
	Digest          string `env:"EDGE_CONFIG_DIGEST"`
	TestKey         string `env:"TEST_KEY"`
	TestValue       string `env:"TEST_VALUE"`
}

func LoadEnv() (env Environment, err error) {
	dotenv := os.Getenv("CI") != "true"
	err = environment.Load(&env, &environment.EnvironmentOptions{DotEnv: dotenv})
	return
}
