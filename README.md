# go-github-cli

`go-github-cli` is a cli client for GitHub's REST API.  It is generated
by inspecting [go-github](https://github.com/google/go-github) and
https://octokit.github.io/routes/.

See ./generator for more on how it is generated.

## Work In Progress

 go-github-cli is a work in progress.  Use it with
caution.

## Usage

There are just a few endpoints covered currently. You can set the
environment variable GITHUB_TOKEN instead of using a `--token` flag to
avoid putting credentials on the command line.

```shell
Usage: go-github-cli <command>

Flags:
  --help    Show context-sensitive help.

Commands:
  issues add-labels-to-issue --token=STRING --owner=STRING --repo=STRING --number=INT --labels=LABELS,...
    Add labels to an issue

  issues create --token=STRING --owner=STRING --repo=STRING --title=STRING
    Create an issue

  issues edit --token=STRING --owner=STRING --repo=STRING --number=INT
    Edit an issue

  issues list --token=STRING
    List all issues assigned to the authenticated user across all visible repositories including owned repositories, member repositories, and organization
    repositories

  issues lock --token=STRING --owner=STRING --repo=STRING --number=INT
    Lock an issue

  organizations get --token=STRING --org=STRING
    Get an organization

  repositories get --token=STRING --owner=STRING --repo=STRING
    Get

Run "go-github-cli <command> --help" for more information on a command.
```



