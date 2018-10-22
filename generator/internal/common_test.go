package internal_test

import (
	"testing"

	. "github.com/WillAbides/go-github-cli/generator/internal"
	"github.com/stretchr/testify/assert"
)

func TestToArgName(t *testing.T) {
	assert.Equal(t, "FooBar", ToArgName("foo-bar"))
}

func TestUnexport(t *testing.T) {
	assert.Equal(t, "issueListOptionsFlags", Unexport("Issue", "listOptionsFlags"))
}
