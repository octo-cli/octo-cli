package configparser_test

import (
	"testing"

	. "github.com/go-github-cli/go-github-cli/generator/internal/configparser"
	"github.com/stretchr/testify/assert"
)

func TestParseConfigFile(t *testing.T) {
	config, err := ParseConfigFile("testdata/config.hcl")
	assert.Nil(t, err)
	assert.Equal(t, []string{"Owner", "Repo", "Number"}, config.Service["Issues"].Command["Edit"].ArgNames)
}
