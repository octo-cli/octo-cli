package main

import (
	"github.com/WillAbides/go-github-cli/services/issuessvc"
	"github.com/WillAbides/go-github-cli/services/organizationssvc"
	"github.com/alecthomas/kong"
)

type CLI struct {
	Issues        issuessvc.IssuesCmd               `cmd:""`
	Organizations organizationssvc.OrganizationsCmd `cmd:""`
}

func main() {
	cli := &CLI{}
	k := kong.Parse(cli)
	k.Run(k)
}
