package configparser

import (
	"io/ioutil"

	"github.com/hashicorp/hcl"
	"github.com/pkg/errors"
)

//Config is a representation of a config.hcl file
type Config struct {
	Service map[string]struct {
		RouteName string
		Command   map[string]struct {
			RoutesName string
			ArgNames   []string
		}
	}
}

//ParseConfigFile parses a config.hcl and returns a representation of it
func ParseConfigFile(configFile string) (*Config, error) {
	config := &Config{}
	bts, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed reading config")
	}
	err = hcl.Decode(config, string(bts))
	return config, errors.Wrap(err, "failed decoding config")
}
