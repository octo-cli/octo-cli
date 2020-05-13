# Contributing to octo-cli

## Scripts

octo-cli uses [Scripts to Rule Them All](https://githubengineering.com/scripts-to-rule-them-all/). There is also a
 Makefile, but make should primarily be used by scripts.

## Issues before PRs

Please create issues to discuss enhancements before spending any significant
time writing code. This project has a very specific purpose, and pull
requests that don't move that forward will be rejected. We don't want to
say no to a PR that you have spent hours on because it doesn't fit the project's
vision. Please don't make us do that.

## Questions and Help

If you have questions or get stuck, you can create an issue here asking
for clarification.

## Project structure

Let's start with what this project is. Octo-cli is a command-line interface
for GitHub's REST API that is generated from [Octokit routes](https://octokit.github.io/routes/).
That last bit about being generated is key to understanding the project
structure.

### `./api.github.com.json`
`api.github.com.json` is what we download from [Octokit routes](https://octokit.github.io/routes/).
It is an openapi v3 schema that describes all the REST endpoints covered in 
[GitHub's documentation](https://developer.github.com/v3/).
Updating routes.json will eventually be automated. Until then, only trusted
octo-cli developers can update routes.json. PRs with changes to routes.json
will be politely rejected.

### `./internal/generator/`
`internal/generator/` is the code that parses routes.json and creates commands
for each defined endpoint. Most of the action here is in the Generate
function and `const tmplt`. Generate is currently a bit of an oversized
mess. Please bear with us until it is broken up into more easily grokked
pieces.

Generator is not well tested. Tests consist generating new commands and
checking whether they match the expected output. Unit tests are needed here.

### `./internal/generator/testdata`
As you probably guessed this is test data for generator's few tests. This
should only be modified by running `./script/update-testdata`.

### `./internal/generated`

This is where you will find all of our generated code. Generated code
should never be edited manually, and if it doesn't match what is output
by `script/generate` then ci will fail.

These files are used by [kong](https://github.com/alecthomas/kong) to
create octo-cli's sub-commands.

Each struct represents one subcommand.
The top level subcommands are in the `CLI` struct in `cli.go`. These have
their own subcommands that each represent an endpoint. For instance the
`CLI` contains `Issues` with the type `IssuesCmd`. `IssuesCmd` then contains
`Create` with the type `IssuesCreateCmd`.  `IssuesCreateCmd` represents
the command-line `octo issues create`. `IssuesCreateCmd` contains fields
representing the flags for that command.

Each executable command struct has a `Run()` method. This is what is
executed when the command is run. Here flag values are used to modify the
API request then the request is performed by `DoRequest` on the last line.

### `./internal/basecmd.go`

This contains the `BaseCmd` struct and it's receivers. `BaseCmd` is an
anonymous member of every executable command struct.  This can be thought
of as the low-level api client that the generated code uses.

### `./tests/`

`./tests/` contains vcr style tests for a select few commands.
