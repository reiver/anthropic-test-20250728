package cfg

import (
	"os"
)

var ApiKey string

func init() {
	const name string = "ANTHROPIC_API_KEY"

	var found bool

	ApiKey, found = os.LookupEnv(name)
	if !found {
		panic("environment-variable "+name+" not set")
	}
	if "" == ApiKey {
		panic("environment-variable "+name+" empty")
	}
}
