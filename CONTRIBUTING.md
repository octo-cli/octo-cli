# Contributing to octo-cli

## Requirements

- Go 1.11+
- dep
- golint
- golangci-lint (installed by script/bootstrap)
- bash shell for `./scripts/*`

## Scripts

Octo-cli is using [Scripts to Rule Them All](https://githubengineering.com/scripts-to-rule-them-all/). Or at least a subset of the idea.

- `script/bootstrap` - installs dependencies. This is currently lacking and only installs golangci-lint.  You are on your own for the rest.
- `script/build` - builds the main binary and writes it to `bin/octo`
- `script/cibuild` - This is what ci runs. You can also run it locally so you aren't surprised by a ci failure.  It currently takes about 13 seconds to run on my laptop.
- `script/generate` - Generates the code in `generated/`.  We'll get to what that is further down the page.
- `script/golint` - runs golint on everything that isn't vendor
- `script/latestversion` - returns the last git tagged version
- `script/lint` - runs all the default linters from golangci-lint
- `script/newversion` - tells you what version to use for a new release flag
- `script/test` - runs tests
- `script/update-routes` - updates `./routes.json` with the latest from https://octokit.github.io/routes/index.json
- `script/update-testdata` - copies `./routes.json` to `./generator/testdata/routes.json` and runs generator to regenerate `./generator/testdata/generated`

## Generated Code


