package main

import (
	"github.com/alecthomas/kong"
	"github.com/go-github-cli/go-github-cli/services/issuessvc"
	"github.com/go-github-cli/go-github-cli/services/organizationssvc"
	"github.com/go-github-cli/go-github-cli/services/repositoriessvc"
)

type CLI struct {
	Issues        issuessvc.IssuesCmd               `cmd:""`
	Organizations organizationssvc.OrganizationsCmd `cmd:""`
	Repositories  repositoriessvc.RepositoriesCmd   `cmd:""`
}

func main() {
	cli := &CLI{}
	k := kong.Parse(cli)
	k.FatalIfErrorf(k.Run(k))
}
