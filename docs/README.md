# octo-cli

[![Stability: Experimental](https://masterminds.github.io/stability/experimental.svg)](https://masterminds.github.io/stability/experimental.html)

Octo-cli is a shell client for GitHub's REST API.

Think of it as [octokit](http://octokit.github.io/) for shell scripts, and you won't be far off.

If you are looking for a GitHub client to use interactively, [gh](https://cli.github.com/) should be your first choice. 
`gh` has a more refined user interface and demands at least 80% less typing in day to day use. Octo-cli wants to
 be a more convenient alternative to using raw curl commands in your shell scripts -- although if you want to go the
 curl route, check out octo-cli's `--curl` option.

Octo-cli is generated from the OpenAPI spec published at https://github.com/github/rest-api-description

## Walk Through

A quick walk through is often the best way to get a feel for things, so here we are walking through octo-cli. This
 should cover most of what you need to get started.

Start by setting the `GITHUB_TOKEN` environment variable to a personal access token for your account.

```shell-session
$ export GITHUB_TOKEN=mypersonalaccesstoken
```

Now make sure it works.

```shell-session
$ octo users get-authenticated
{
...
  "login": "WillAbides",
...
}
```

Yep. That's me.

Shell completions makes everything easier, so let's get that set up.

```shell-session
$ `octo --install-completions`
$ octo <tab><tab>
actions               emojis                licenses              projects              search
activity              gists                 markdown              pulls                 teams
apps                  git                   meta                  rate-limit            users
checks                gitignore             migrations            reactions
code-scanning         interactions          oauth-authorizations  repos
codes-of-conduct      issues                orgs                  scim
$ octo git <tab><tab>
create-blob         create-tag          get-blob            get-tag             update-ref
create-commit       create-tree         get-commit          get-tree
create-ref          delete-ref          get-ref             list-matching-refs
```

Octo has hundreds of sub commands. Shell completion is a good way to remember what's available. You may want to add 
`$(octo --install-completions)` to .bash_profile or your shell's equivalent config file.

Now let's do something more fun like create a release on `octo-cli-testorg/scratch`.  First we need to figure out
 what we are releasing. Let's start by finding the latest release.

```shell-session
$ octo repos get-latest-release --repo octo-cli-testorg/scratch
{
  "assets": [],
  "assets_url": "https://api.github.com/repos/octo-cli-testorg/scratch/releases/26402951/assets",
  "author": {
    "avatar_url": "https://avatars3.githubusercontent.com/u/233500?v=4",
    "events_url": "https://api.github.com/users/WillAbides/events{/privacy}",
...
  "tag_name": "v0.0.1",
  "tarball_url": "https://api.github.com/repos/octo-cli-testorg/scratch/tarball/v0.0.1",
  "target_commitish": "master",
  "upload_url": "https://uploads.github.com/repos/octo-cli-testorg/scratch/releases/26402951/assets{?name,label}",
  "url": "https://api.github.com/repos/octo-cli-testorg/scratch/releases/26402951",
  "zipball_url": "https://api.github.com/repos/octo-cli-testorg/scratch/zipball/v0.0.1"
}
```

That's right...v0.0.1. I remember it fondly. It's the release I made a few minutes ago to prep for this walk through. 
But I just need to know the tag name, not all that other information. Let's try again using `--format` to provide a
go template for the output, and we'll capture it in a variable for later.

```shell-session
$ last_release_tag="$(octo repos get-latest-release --repo octo-cli-testorg/scratch --format {{.tag_name}})"
$ echo $last_release_tag
v0.0.1
``` 

What's changed since v0.0.1? We can check commit messages from the diff to find out.

```shell-session
$ octo repos compare-commits --repo octo-cli-testorg/scratch --base "$last_release_tag" --head master | jq '.commits[].commit.message'
"Update README.md"
"add greeting capability"
"Merge pull request #2 from octo-cli-testorg/greetme\n\nadd greeting capability"
```

I used the `$last_release_tag` variable from earlier. I also used `jq` to process the output instead of `--format`.

Let's get this new greeting capability released stat. We'll call it `v0.0.2`.

```shell-session
octo repos create-release --repo octo-cli-testorg/scratch \
--tag_name v0.0.2 \
--name "v0.0.2 - The Greeting" \
--body "## Features
Greeting capabilities

## Chores
Updated README.md
"
{
  "assets": [],
  "assets_url": "https://api.github.com/repos/octo-cli-testorg/scratch/releases/26474415/assets",
...
}
```

Nice. Our release is live. We should add some binaries to it. We need to get the release id to upload assets, and I
 forgot to get that from the output before.
 
```shell-session
$ release_id=$(octo repos get-release-by-tag --repo octo-cli-testorg/scratch --tag v0.0.2 | jq '.id')
$ echo $release_id
26474415
```

I've built all the binaries, and they're in `./dist`. Now let's add them all to the release.

```shell-session
for upload in dist/*; do
>   octo repos upload-release-asset --repo octo-cli-testorg/scratch \
>   --release_id "$release_id" \
>   --file "$upload" \
>   --name "$(echo $upload | cut -d "/" -f2)" \
>   --format "{{.name}}"
> done
scratch_0.0.2_Darwin_i386.tar.gz
scratch_0.0.2_Darwin_x86_64.tar.gz
scratch_0.0.2_Linux_i386.tar.gz
scratch_0.0.2_Linux_x86_64.tar.gz
scratch_0.0.2_Windows_i386.tar.gz
scratch_0.0.2_Windows_x86_64.tar.gz
```

We did it! Scratch v0.0.2 is available for download.

## Installation

#### Homebrew

To install with [Homebrew](https://brew.sh/) on macOS or Linux:

```
brew install octo-cli/octo/octo
```

#### The easy and overly trusting way

Do you trust me? Do you even know me? Do you trust me anyway? Do you also
happen to have curl available on your system? I may have the answer for you.

Just run the following command to download octo to your current directory:

```
curl https://raw.githubusercontent.com/octo-cli/octo-cli/master/download.sh | sh
```

#### The slightly harder and slightly less trusting way

Go to the [latest release](https://github.com/octo-cli/octo-cli/releases/latest), download
the archive for your operating system and architecture (if you don't know your
architecture it's probably `x86_64`). Extract the archive using `tar -xzf` or
equivalent.

#### The least trusting way

I get it. You don't trust me. I wouldn't trust me either. If you have golang
on your system you can install from source cloning this repo and running
`script/build`.  This will create `bin/octo`.  You could also run
`go install github.com/octo-cli/octo-cli`. This will install `octo-cli`
for you. Note that the binary will be named `octo-cli` instead of `octo`.
You can rename it if you want.

#### Scoop and Snapcraft

Wouldn't it be cool if you could install with scoop or snapcraft?
I think so, but haven't spent the time to set this up yet. If you have the
know-how and time to spare, we could [use your help](https://github.com/octo-cli/octo-cli/issues/90).

## Help

Like many command line utilities, you can get context sensitive help with `--help`. Unlike them, the help in octo-cli
is generated from api documentation. That means some of the wording doesn't quite fit in a command line context.

## Untested commands

Because octo-cli is automatically generated and covers so many endpoints, 
most of the commands haven't been tested or even run. If you run into a
problem with a command, please create an issue.

## Credentials

Octo-cli looks for a personal access token in the `GITHUB_TOKEN` environment variable.

That is the only form of authentication that is currently available. Unfortunately this means that authenticating as
 a GitHub App is not possible for now.

## GitHub Enterprise Server

Use octo-cli with GitHub enterprise by setting the environment variable
GITHUB_API_BASE_URL. Something like `export GITHUB_API_BASE_URL=https://ghe.example.com/api/v3`.
You can also set this with a flag on each command: `--api-base-url="https://ghe.example.com/api/v3"`.

## Curl

Octo-cli provides a `--curl` flag that causes it to output a curl request. This could be useful if you want to use
 octo-cli to build a request but can't or don't want to include octo-cli in your project.
 
```shell-session
$ octo repos create-release --repo foo/bar --tag_name v0.0.1 --name "my release" --draft --curl

curl -X 'POST' -d '{"draft":true,"name":"my release","tag_name":"v0.0.1"}' -H 'Accept: application/vnd.github.v3+json' 
-H 'Content-Type: application/json' -H 'User-Agent: octo-cli/0.11.0' -H "Authorization: token $GITHUB_TOKEN" 
'https://api.github.com/repos/foo/bar/releases'
```

## Output

__See [Formatting Output](format.md)__

Octo-cli outputs prettified json. You can modify the output with `--format`, `--output-each` and 
`--raw-output` flags. 

## Preview flags

GitHub uses [preview flags](https://developer.github.com/v3/previews/) when
introducing API changes. These require you to send a preview header when
submitting a request. Octo-cli will not set these headers automatically.
Instead it provides flags for you to enable the relevant previews for each
command.

The decision to not automatically set preview headers is intended to prevent
users from unknowingly becoming reliant on APIs that are subject to change.
