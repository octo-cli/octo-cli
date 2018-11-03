package tests

import (
	"bytes"
	"encoding/json"
	"github.com/alecthomas/kong"
	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/octo-cli/octo-cli/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func init() {
	tkn, ok := os.LookupEnv("TESTUSER_TOKEN")
	if !ok {
		tkn = "deadbeef"
	}
	_ = os.Setenv("GITHUB_TOKEN", tkn)
}

func startVCR(t *testing.T, recPath string) func() {
	t.Helper()
	var err error
	rec, err := recorder.New(recPath)
	rec.SetMatcher(func(r *http.Request, i cassette.Request) bool {
		var b bytes.Buffer
		if _, err := b.ReadFrom(r.Body); err != nil {
			return false
		}
		r.Body = ioutil.NopCloser(&b)
		return cassette.DefaultMatcher(r, i) &&
			(b.String() == "" || b.String() == i.Body) &&
			r.Header.Get("Accept") == i.Headers.Get("Accept")
	})
	require.Nil(t, err)
	services.TransportWrapper = rec
	return func() {
		t.Helper()
		require.Nil(t, rec.Stop())
	}
}

func testCmdLine(t *testing.T, fixtureName string, cmdStruct interface{}, cmdline ...string) (sout bytes.Buffer, serr bytes.Buffer, err error) {
	t.Helper()
	recCleanup := startVCR(t, filepath.Join("testdata", "fixtures", fixtureName))
	defer recCleanup()
	p, e := kong.New(cmdStruct)

	require.NoError(t, e)
	services.Stdout = &sout

	k, e := p.Parse(cmdline)
	require.NoError(t, e)
	valueIsSetMap := map[string]bool{}
	for _, flag := range k.Flags() {
		valueIsSetMap[flag.Name] = flag.Set
	}
	err = k.Run(valueIsSetMap)
	return
}

func TestCreate(t *testing.T) {
	stdout, stderr, err := testCmdLine(t, "test_issues_create", &services.IssuesCmd{},
		`create`,
		`--owner=octo-cli-testorg`,
		`--repo=test-create-issue`,
		`--title=test this`,
		`--body=test this body`,
		`--labels=label1`,
		`--labels=label2`,
		`--milestone=1`,
		`--assignees=octo-cli-testuser`,
		`--raw-output`,
	)
	assert.NoError(t, err)
	assert.Empty(t, stderr.String())
	var got map[string]interface{}
	err = json.Unmarshal(stdout.Bytes(), &got)
	assert.NoError(t, err)
	assert.Equal(t, "test this", got["title"])
	assert.Equal(t, "test this body", got["body"])
	assert.Len(t, got["labels"], 2)
	assert.EqualValues(t, 1, got["milestone"].(map[string]interface{})["number"])
}