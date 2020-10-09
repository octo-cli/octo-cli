// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import internal "github.com/octo-cli/octo-cli/internal"

type CodeScanningCmd struct {
	GetAlert           CodeScanningGetAlertCmd           `cmd:""`
	ListAlertsForRepo  CodeScanningListAlertsForRepoCmd  `cmd:""`
	ListRecentAnalyses CodeScanningListRecentAnalysesCmd `cmd:""`
	UpdateAlert        CodeScanningUpdateAlertCmd        `cmd:""`
	UploadSarif        CodeScanningUploadSarifCmd        `cmd:""`
}

type CodeScanningGetAlertCmd struct {
	Repo        string `name:"repo" required:"true"`
	AlertNumber int64  `name:"alert_number" required:"true"`
	internal.BaseCmd
}

func (c *CodeScanningGetAlertCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/code-scanning/alerts/{alert_number}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("alert_number", c.AlertNumber)
	return c.DoRequest("GET")
}

type CodeScanningListAlertsForRepoCmd struct {
	Repo  string `name:"repo" required:"true"`
	Ref   string `name:"ref"`
	State string `name:"state"`
	internal.BaseCmd
}

func (c *CodeScanningListAlertsForRepoCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/code-scanning/alerts")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("state", c.State)
	c.UpdateURLQuery("ref", c.Ref)
	return c.DoRequest("GET")
}

type CodeScanningListRecentAnalysesCmd struct {
	Repo     string `name:"repo" required:"true"`
	Ref      string `name:"ref"`
	ToolName string `name:"tool_name"`
	internal.BaseCmd
}

func (c *CodeScanningListRecentAnalysesCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/code-scanning/analyses")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("ref", c.Ref)
	c.UpdateURLQuery("tool_name", c.ToolName)
	return c.DoRequest("GET")
}

type CodeScanningUpdateAlertCmd struct {
	Repo            string `name:"repo" required:"true"`
	AlertNumber     int64  `name:"alert_number" required:"true"`
	DismissedReason string `name:"dismissed_reason"`
	State           string `name:"state" required:"true"`
	internal.BaseCmd
}

func (c *CodeScanningUpdateAlertCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/code-scanning/alerts/{alert_number}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("alert_number", c.AlertNumber)
	c.UpdateBody("dismissed_reason", c.DismissedReason)
	c.UpdateBody("state", c.State)
	return c.DoRequest("PATCH")
}

type CodeScanningUploadSarifCmd struct {
	Repo        string `name:"repo" required:"true"`
	CheckoutUri string `name:"checkout_uri"`
	StartedAt   string `name:"started_at"`
	CommitSha   string `name:"commit_sha" required:"true"`
	Ref         string `name:"ref" required:"true"`
	Sarif       string `name:"sarif" required:"true"`
	ToolName    string `name:"tool_name" required:"true"`
	internal.BaseCmd
}

func (c *CodeScanningUploadSarifCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/code-scanning/sarifs")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("checkout_uri", c.CheckoutUri)
	c.UpdateBody("commit_sha", c.CommitSha)
	c.UpdateBody("ref", c.Ref)
	c.UpdateBody("sarif", c.Sarif)
	c.UpdateBody("started_at", c.StartedAt)
	c.UpdateBody("tool_name", c.ToolName)
	return c.DoRequest("POST")
}
