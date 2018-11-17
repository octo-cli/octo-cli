# octo-cli

Octo-cli is a command line client for GitHub's REST API. It is intended to make
it easier to interact with GitHub in shell scripts. In most cases, it should
be more convenient than curl and more scriptable than [hub](https://hub.github.com/).

If you are looking for a command-line client to use interactively, please
try [hub](https://hub.github.com/) first. Octo-cli is primarily intended
for scripting.

## Installation

#### Homebrew

To install with [homebrew](https://brew.sh/):

```
brew tap octo-cli/octo
brew install octo
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
know-how and time to spare, we could [use your help](https://github.com/octo-cli/octo-cli/issues/45).

## Untested commands

Because octo-cli is still a young project and the subcommands are generated,
most of the commands haven't been tested or even run. If you run into a
problem with a command, there is a good chance it's a problem with octo-cli,
not just your usage. Please create an issue for any problems you run across.

## Contributing

See [CONTRIBUTING.md]

## Usage

#### Credentials

You can set the environment variable GITHUB_TOKEN instead of using a `--token` flag to
avoid putting credentials on the command line.

#### GitHub Enterprise

Use octo-cli with GitHub enterprise by setting the environment variable
GITHUB_API_BASE_URL. Something like `export GITHUB_API_BASE_URL=https://ghe.example.com/v3/api`.
You can also set this with a flag on each command: `--api-base-url="https://ghe.example.com/v3/api"`.

#### Formatting output

By default, octo-cli outputs formatted json results with line-breaks and
indenting. You can get unformatted json by adding the `--raw-output` flag.

You can further format output using [go templates](https://golang.org/pkg/text/template/).
This idea is taken from the docker cli, but not yet implemented as well.
We don't have all the extra functions that docker provides...yet.

If you have [jq](https://stedolan.github.io/jq/) available, you can format
results with it. Especially if you aren't a go developer or aren't already
familiar with go templates, you may be better off spending time learning
jq syntax that you can use in other places instead of go templates which
you may not use for anything else.

#### Preview flags

GitHub uses [preview flags](https://developer.github.com/v3/previews/) when
introducing API changes. These require you to send a preview header when
submitting a request. Octo-cli will not set these headers automatically.
Instead it provides flags for you to enable the relevant previews for each
command.

The decision to not automatically set preview headers is intended to prevent
users from unknowingly becoming reliant on APIs that are subject to change.

#### Errors

Octo-cli exits non-zero when it sees an http status >= 400. It does no
error handling beyond this, and it hasn't been decided how we want to
handle errors.

#### Debugging

There's not currently any way to make octo-cli output any debugging data.

#### Examples

###### Use GITHUB_TOKEN environment variable to set your credentials
```
$ octo users get-authenticated
octo: error: missing flags: --token=STRING

$ export GITHUB_TOKEN=yourpersonalaccesstokenhere
$ octo users get-authenticated
{
  "avatar_url": "https://avatars3.githubusercontent.com/u/56260?v=4",
...
}
```

###### Get an issue

```
$ octo issues get --owner octocat --repo Hello-World --number 7
{
  "assignee": null,
...
}
```

###### Get just the closer and title of an issue using --format

```
$ octo issues get --owner octocat --repo Hello-World --number 7 \
> --format '{{.closed_by.login}} - {{.title}}'
octocat - Hello World in all programming languages
```

###### Use a preview flag to see a preview flag
```
$ octo issues get --owner octocat --repo Hello-World --number 7 \
> --squirrel-girl-preview \
> --format {{.reactions.total_count}}
1
```

#### Help output

This only shows the required flags for each command.  You can find all available
flags with the full command plus --help (e.g. `octo orgs edit --help` )

<!--- START HELP OUTPUT --->
```
Usage: octo <command>

Flags:
  --help       Show context-sensitive help.
  --version

Commands:
  activity list-public-events --token=STRING
    List public events -
    https://developer.github.com/v3/activity/events/#list-public-events

  activity list-repo-events --token=STRING --repo=STRING
    List repository events -
    https://developer.github.com/v3/activity/events/#list-repository-events

  activity list-public-events-for-repo-network --token=STRING --repo=STRING
    List public events for a network of repositories -
    https://developer.github.com/v3/activity/events/#list-public-events-for-a-network-of-repositories

  activity list-public-events-for-org --token=STRING --org=STRING
    List public events for an organization -
    https://developer.github.com/v3/activity/events/#list-public-events-for-an-organization

  activity list-received-events-for-user --token=STRING --username=STRING
    List events that a user has received -
    https://developer.github.com/v3/activity/events/#list-events-that-a-user-has-received

  activity list-received-public-events-for-user --token=STRING --username=STRING
    List public events that a user has received -
    https://developer.github.com/v3/activity/events/#list-public-events-that-a-user-has-received

  activity list-events-for-user --token=STRING --username=STRING
    List events performed by a user -
    https://developer.github.com/v3/activity/events/#list-events-performed-by-a-user

  activity list-public-events-for-user --token=STRING --username=STRING
    List public events performed by a user -
    https://developer.github.com/v3/activity/events/#list-public-events-performed-by-a-user

  activity list-events-for-org --token=STRING --username=STRING --org=STRING
    List events for an organization -
    https://developer.github.com/v3/activity/events/#list-events-for-an-organization

  activity list-feeds --token=STRING
    List feeds - https://developer.github.com/v3/activity/feeds/#list-feeds

  activity list-notifications --token=STRING
    List your notifications -
    https://developer.github.com/v3/activity/notifications/#list-your-notifications

  activity list-notifications-for-repo --token=STRING --repo=STRING
    List your notifications in a repository -
    https://developer.github.com/v3/activity/notifications/#list-your-notifications-in-a-repository

  activity mark-as-read --token=STRING
    Mark as read -
    https://developer.github.com/v3/activity/notifications/#mark-as-read

  activity mark-notifications-as-read-for-repo --token=STRING --repo=STRING
    Mark notifications as read in a repository -
    https://developer.github.com/v3/activity/notifications/#mark-notifications-as-read-in-a-repository

  activity get-thread --token=STRING --thread_id=INT-64
    View a single thread -
    https://developer.github.com/v3/activity/notifications/#view-a-single-thread

  activity mark-thread-as-read --token=STRING --thread_id=INT-64
    Mark a thread as read -
    https://developer.github.com/v3/activity/notifications/#mark-a-thread-as-read

  activity get-thread-subscription --token=STRING --thread_id=INT-64
    Get a thread subscription -
    https://developer.github.com/v3/activity/notifications/#get-a-thread-subscription

  activity set-thread-subscription --token=STRING --thread_id=INT-64
    Set a thread subscription -
    https://developer.github.com/v3/activity/notifications/#set-a-thread-subscription

  activity delete-thread-subscription --token=STRING --thread_id=INT-64
    Delete a thread subscription -
    https://developer.github.com/v3/activity/notifications/#delete-a-thread-subscription

  activity list-stargazers-for-repo --token=STRING --repo=STRING
    List Stargazers -
    https://developer.github.com/v3/activity/starring/#list-stargazers

  activity list-repos-starred-by-user --token=STRING --username=STRING
    List repositories being starred by a user -
    https://developer.github.com/v3/activity/starring/#list-repositories-being-starred

  activity list-repos-starred-by-authenticated-user --token=STRING
    List repositories being starred by the authenticated user -
    https://developer.github.com/v3/activity/starring/#list-repositories-being-starred

  activity check-starring-repo --token=STRING --repo=STRING
    Check if you are starring a repository -
    https://developer.github.com/v3/activity/starring/#check-if-you-are-starring-a-repository

  activity star-repo --token=STRING --repo=STRING
    Star a repository -
    https://developer.github.com/v3/activity/starring/#star-a-repository

  activity unstar-repo --token=STRING --repo=STRING
    Unstar a repository -
    https://developer.github.com/v3/activity/starring/#unstar-a-repository

  activity list-watchers-for-repo --token=STRING --repo=STRING
    List watchers -
    https://developer.github.com/v3/activity/watching/#list-watchers

  activity list-repos-watched-by-user --token=STRING --username=STRING
    List repositories being watched by a user -
    https://developer.github.com/v3/activity/watching/#list-repositories-being-watched

  activity list-watched-repos-for-authenticated-user --token=STRING
    List repositories being watched by the authenticated user -
    https://developer.github.com/v3/activity/watching/#list-repositories-being-watched

  activity get-repo-subscription --token=STRING --repo=STRING
    Get a Repository Subscription -
    https://developer.github.com/v3/activity/watching/#get-a-repository-subscription

  activity set-repo-subscription --token=STRING --repo=STRING
    Set a Repository Subscription -
    https://developer.github.com/v3/activity/watching/#set-a-repository-subscription

  activity delete-repo-subscription --token=STRING --repo=STRING
    Delete a Repository Subscription -
    https://developer.github.com/v3/activity/watching/#delete-a-repository-subscription

  activity check-watching-repo-legacy --token=STRING --repo=STRING
    Check if you are watching a repository (LEGACY) -
    https://developer.github.com/v3/activity/watching/#check-if-you-are-watching-a-repository-legacy

  activity watch-repo-legacy --token=STRING --repo=STRING
    Watch a repository (LEGACY) -
    https://developer.github.com/v3/activity/watching/#watch-a-repository-legacy

  activity stop-watching-repo-legacy --token=STRING --repo=STRING
    Stop watching a repository (LEGACY) -
    https://developer.github.com/v3/activity/watching/#stop-watching-a-repository-legacy

  apps get-by-slug --token=STRING --app_slug=STRING
    Get a single GitHub App -
    https://developer.github.com/v3/apps/#get-a-single-github-app

  apps get-authenticated --token=STRING --machine-man-preview
    Get the authenticated GitHub App -
    https://developer.github.com/v3/apps/#get-the-authenticated-github-app

  apps list-installations --token=STRING --machine-man-preview
    Find installations -
    https://developer.github.com/v3/apps/#find-installations

  apps get-installation --token=STRING --machine-man-preview --installation_id=INT-64
    Get a single installation -
    https://developer.github.com/v3/apps/#get-a-single-installation

  apps list-installations-for-authenticated-user --token=STRING --machine-man-preview
    List installations for user -
    https://developer.github.com/v3/apps/#list-installations-for-user

  apps create-installation-token --token=STRING --machine-man-preview --installation_id=INT-64
    Create a new installation token -
    https://developer.github.com/v3/apps/#create-a-new-installation-token

  apps find-org-installation --token=STRING --machine-man-preview --org=STRING
    Find organization installation -
    https://developer.github.com/v3/apps/#find-organization-installation

  apps find-repo-installation --token=STRING --machine-man-preview --repo=STRING
    Find repository installation -
    https://developer.github.com/v3/apps/#find-repository-installation

  apps find-user-installation --token=STRING --machine-man-preview --username=STRING
    Find user installation -
    https://developer.github.com/v3/apps/#find-user-installation

  apps create-from-manifest --token=STRING --fury-preview --code=STRING
    Create a GitHub App from a manifest -
    https://developer.github.com/v3/apps/#create-a-github-app-from-a-manifest

  apps list-repos --token=STRING
    List repositories -
    https://developer.github.com/v3/apps/installations/#list-repositories

  apps list-installation-repos-for-authenticated-user --token=STRING --machine-man-preview --installation_id=INT-64
    List repositories accessible to the user for an installation -
    https://developer.github.com/v3/apps/installations/#list-repositories-accessible-to-the-user-for-an-installation

  apps add-repo-to-installation --token=STRING --machine-man-preview --installation_id=INT-64 --repository_id=INT-64
    Add repository to installation -
    https://developer.github.com/v3/apps/installations/#add-repository-to-installation

  apps remove-repo-from-installation --token=STRING --machine-man-preview --installation_id=INT-64 --repository_id=INT-64
    Remove repository from installation -
    https://developer.github.com/v3/apps/installations/#remove-repository-from-installation

  apps list-plans --token=STRING
    List all plans for your Marketplace listing -
    https://developer.github.com/v3/apps/marketplace/#list-all-plans-for-your-marketplace-listing

  apps list-plans-stubbed --token=STRING
    List all plans for your Marketplace listing (stubbed) -
    https://developer.github.com/v3/apps/marketplace/#list-all-plans-for-your-marketplace-listing

  apps list-accounts-user-or-org-on-plan --token=STRING --plan_id=INT-64
    List all GitHub accounts (user or organization) on a specific plan -
    https://developer.github.com/v3/apps/marketplace/#list-all-github-accounts-user-or-organization-on-a-specific-plan

  apps list-accounts-user-or-org-on-plan-stubbed --token=STRING --plan_id=INT-64
    List all GitHub accounts (user or organization) on a specific plan (stubbed)
    -
    https://developer.github.com/v3/apps/marketplace/#list-all-github-accounts-user-or-organization-on-a-specific-plan

  apps check-account-is-associated-with-any --token=STRING --account_id=INT-64
    Check if a GitHub account is associated with any Marketplace listing -
    https://developer.github.com/v3/apps/marketplace/#check-if-a-github-account-is-associated-with-any-marketplace-listing

  apps check-account-is-associated-with-any-stubbed --token=STRING --account_id=INT-64
    Check if a GitHub account is associated with any Marketplace listing
    (stubbed) -
    https://developer.github.com/v3/apps/marketplace/#check-if-a-github-account-is-associated-with-any-marketplace-listing

  apps list-marketplace-purchases-for-authenticated-user --token=STRING
    Get a user's Marketplace purchases -
    https://developer.github.com/v3/apps/marketplace/#get-a-users-marketplace-purchases

  apps list-marketplace-purchases-for-authenticated-user-stubbed --token=STRING
    Get a user's Marketplace purchases (stubbed) -
    https://developer.github.com/v3/apps/marketplace/#get-a-users-marketplace-purchases

  checks list-for-ref --token=STRING --antiope-preview --repo=STRING --ref=STRING
    List check runs for a specific ref -
    https://developer.github.com/v3/checks/runs/#list-check-runs-for-a-specific-ref

  checks list-for-suite --token=STRING --antiope-preview --repo=STRING --check_suite_id=INT-64
    List check runs in a check suite -
    https://developer.github.com/v3/checks/runs/#list-check-runs-in-a-check-suite

  checks get --token=STRING --antiope-preview --repo=STRING --check_run_id=INT-64
    Get a single check run -
    https://developer.github.com/v3/checks/runs/#get-a-single-check-run

  checks list-annotations --token=STRING --antiope-preview --repo=STRING --check_run_id=INT-64
    List annotations for a check run -
    https://developer.github.com/v3/checks/runs/#list-annotations-for-a-check-run

  checks get-suite --token=STRING --antiope-preview --repo=STRING --check_suite_id=INT-64
    Get a single check suite -
    https://developer.github.com/v3/checks/suites/#get-a-single-check-suite

  checks list-suites-for-ref --token=STRING --antiope-preview --repo=STRING --ref=STRING
    List check suites for a specific ref -
    https://developer.github.com/v3/checks/suites/#list-check-suites-for-a-specific-ref

  checks create-suite --token=STRING --antiope-preview --repo=STRING --head_sha=STRING
    Create a check suite -
    https://developer.github.com/v3/checks/suites/#create-a-check-suite

  checks rerequest-suite --token=STRING --antiope-preview --repo=STRING --check_suite_id=INT-64
    Rerequest check suite -
    https://developer.github.com/v3/checks/suites/#rerequest-check-suite

  codes-of-conduct list-conduct-codes --token=STRING --scarlet-witch-preview
    List all codes of conduct -
    https://developer.github.com/v3/codes_of_conduct/#list-all-codes-of-conduct

  codes-of-conduct get-conduct-code --token=STRING --scarlet-witch-preview --key=STRING
    Get an individual code of conduct -
    https://developer.github.com/v3/codes_of_conduct/#get-an-individual-code-of-conduct

  codes-of-conduct get-for-repo --token=STRING --scarlet-witch-preview --repo=STRING
    Get the contents of a repository's code of conduct -
    https://developer.github.com/v3/codes_of_conduct/#get-the-contents-of-a-repositorys-code-of-conduct

  emojis get --token=STRING
    Get - https://developer.github.com/v3/emojis/#emojis

  gists list-public-for-user --token=STRING --username=STRING
    List public gists for the specified user -
    https://developer.github.com/v3/gists/#list-a-users-gists

  gists list --token=STRING
    List the authenticated user's gists or if called anonymously, this will
    return all public gists -
    https://developer.github.com/v3/gists/#list-a-users-gists

  gists list-public --token=STRING
    List all public gists -
    https://developer.github.com/v3/gists/#list-all-public-gists

  gists list-starred --token=STRING
    List starred gists -
    https://developer.github.com/v3/gists/#list-starred-gists

  gists get --token=STRING --gist_id=STRING
    Get a single gist - https://developer.github.com/v3/gists/#get-a-single-gist

  gists get-revision --token=STRING --gist_id=STRING --sha=STRING
    Get a specific revision of a gist -
    https://developer.github.com/v3/gists/#get-a-specific-revision-of-a-gist

  gists list-commits --token=STRING --gist_id=STRING
    List gist commits - https://developer.github.com/v3/gists/#list-gist-commits

  gists star --token=STRING --gist_id=STRING
    Star a gist - https://developer.github.com/v3/gists/#star-a-gist

  gists unstar --token=STRING --gist_id=STRING
    Unstar a gist - https://developer.github.com/v3/gists/#unstar-a-gist

  gists check-is-starred --token=STRING --gist_id=STRING
    Check if a gist is starred -
    https://developer.github.com/v3/gists/#check-if-a-gist-is-starred

  gists fork --token=STRING --gist_id=STRING
    Fork a gist - https://developer.github.com/v3/gists/#fork-a-gist

  gists list-forks --token=STRING --gist_id=STRING
    List gist forks - https://developer.github.com/v3/gists/#list-gist-forks

  gists delete --token=STRING --gist_id=STRING
    Delete a gist - https://developer.github.com/v3/gists/#delete-a-gist

  gists list-comments --token=STRING --gist_id=STRING
    List comments on a gist -
    https://developer.github.com/v3/gists/comments/#list-comments-on-a-gist

  gists get-comment --token=STRING --gist_id=STRING --comment_id=INT-64
    Get a single comment -
    https://developer.github.com/v3/gists/comments/#get-a-single-comment

  gists create-comment --token=STRING --gist_id=STRING --body=STRING
    Create a comment -
    https://developer.github.com/v3/gists/comments/#create-a-comment

  gists edit-comment --token=STRING --gist_id=STRING --comment_id=INT-64 --body=STRING
    Edit a comment -
    https://developer.github.com/v3/gists/comments/#edit-a-comment

  gists delete-comment --token=STRING --gist_id=STRING --comment_id=INT-64
    Delete a comment -
    https://developer.github.com/v3/gists/comments/#delete-a-comment

  git get-blob --token=STRING --repo=STRING --file_sha=STRING
    Get a blob - https://developer.github.com/v3/git/blobs/#get-a-blob

  git create-blob --token=STRING --repo=STRING --content=STRING
    Create a blob - https://developer.github.com/v3/git/blobs/#create-a-blob

  git get-commit --token=STRING --repo=STRING --commit_sha=STRING
    Get a commit - https://developer.github.com/v3/git/commits/#get-a-commit

  git get-ref --token=STRING --repo=STRING --ref=STRING
    Get a reference - https://developer.github.com/v3/git/refs/#get-a-reference

  git list-refs --token=STRING --repo=STRING
    Get all references -
    https://developer.github.com/v3/git/refs/#get-all-references

  git create-ref --token=STRING --repo=STRING --ref=STRING --sha=STRING
    Create a reference -
    https://developer.github.com/v3/git/refs/#create-a-reference

  git update-ref --token=STRING --repo=STRING --ref=STRING --sha=STRING
    Update a reference -
    https://developer.github.com/v3/git/refs/#update-a-reference

  git delete-ref --token=STRING --repo=STRING --ref=STRING
    Delete a reference -
    https://developer.github.com/v3/git/refs/#delete-a-reference

  git get-tag --token=STRING --repo=STRING --tag_sha=STRING
    Get a tag - https://developer.github.com/v3/git/tags/#get-a-tag

  git get-tree --token=STRING --repo=STRING --tree_sha=STRING
    Get a tree - https://developer.github.com/v3/git/trees/#get-a-tree

  gitignore list-templates --token=STRING
    Listing available templates -
    https://developer.github.com/v3/gitignore/#listing-available-templates

  gitignore get-template --token=STRING --name=STRING
    Get a single template -
    https://developer.github.com/v3/gitignore/#get-a-single-template

  issues list --token=STRING
    List all issues assigned to the authenticated user across all visible
    repositories including owned repositories, member repositories, and
    organization repositories -
    https://developer.github.com/v3/issues/#list-issues

  issues list-for-authenticated-user --token=STRING
    List all issues across owned and member repositories assigned to the
    authenticated user - https://developer.github.com/v3/issues/#list-issues

  issues list-for-org --token=STRING --org=STRING
    List all issues for a given organization assigned to the authenticated user
    - https://developer.github.com/v3/issues/#list-issues

  issues list-for-repo --token=STRING --repo=STRING
    List issues for a repository -
    https://developer.github.com/v3/issues/#list-issues-for-a-repository

  issues get --token=STRING --repo=STRING --number=INT-64
    Get a single issue -
    https://developer.github.com/v3/issues/#get-a-single-issue

  issues create --token=STRING --repo=STRING --title=STRING
    Create an issue - https://developer.github.com/v3/issues/#create-an-issue

  issues edit --token=STRING --repo=STRING --number=INT-64
    Edit an issue - https://developer.github.com/v3/issues/#edit-an-issue

  issues lock --token=STRING --repo=STRING --number=INT-64
    Lock an issue - https://developer.github.com/v3/issues/#lock-an-issue

  issues unlock --token=STRING --repo=STRING --number=INT-64
    Unlock an issue - https://developer.github.com/v3/issues/#unlock-an-issue

  issues list-assignees --token=STRING --repo=STRING
    List assignees -
    https://developer.github.com/v3/issues/assignees/#list-assignees

  issues check-assignee --token=STRING --repo=STRING --assignee=STRING
    Check assignee -
    https://developer.github.com/v3/issues/assignees/#check-assignee

  issues add-assignees --token=STRING --repo=STRING --number=INT-64
    Add assignees to an issue -
    https://developer.github.com/v3/issues/assignees/#add-assignees-to-an-issue

  issues remove-assignees --token=STRING --repo=STRING --number=INT-64
    Remove assignees from an issue -
    https://developer.github.com/v3/issues/assignees/#remove-assignees-from-an-issue

  issues list-comments --token=STRING --repo=STRING --number=INT-64
    List comments on an issue -
    https://developer.github.com/v3/issues/comments/#list-comments-on-an-issue

  issues list-comments-for-repo --token=STRING --repo=STRING
    List comments in a repository -
    https://developer.github.com/v3/issues/comments/#list-comments-in-a-repository

  issues get-comment --token=STRING --repo=STRING --comment_id=INT-64
    Get a single comment -
    https://developer.github.com/v3/issues/comments/#get-a-single-comment

  issues create-comment --token=STRING --repo=STRING --number=INT-64 --body=STRING
    Create a comment -
    https://developer.github.com/v3/issues/comments/#create-a-comment

  issues edit-comment --token=STRING --repo=STRING --comment_id=INT-64 --body=STRING
    Edit a comment -
    https://developer.github.com/v3/issues/comments/#edit-a-comment

  issues delete-comment --token=STRING --repo=STRING --comment_id=INT-64
    Delete a comment -
    https://developer.github.com/v3/issues/comments/#delete-a-comment

  issues list-events --token=STRING --repo=STRING --number=INT-64
    List events for an issue -
    https://developer.github.com/v3/issues/events/#list-events-for-an-issue

  issues list-events-for-repo --token=STRING --repo=STRING
    List events for a repository -
    https://developer.github.com/v3/issues/events/#list-events-for-a-repository

  issues get-event --token=STRING --repo=STRING --event_id=INT-64
    Get a single event -
    https://developer.github.com/v3/issues/events/#get-a-single-event

  issues list-labels-for-repo --token=STRING --repo=STRING
    List all labels for this repository -
    https://developer.github.com/v3/issues/labels/#list-all-labels-for-this-repository

  issues get-label --token=STRING --repo=STRING --name=STRING
    Get a single label -
    https://developer.github.com/v3/issues/labels/#get-a-single-label

  issues create-label --token=STRING --repo=STRING --name=STRING --color=STRING
    Create a label -
    https://developer.github.com/v3/issues/labels/#create-a-label

  issues update-label --token=STRING --repo=STRING --current_name=STRING
    Update a label -
    https://developer.github.com/v3/issues/labels/#update-a-label

  issues delete-label --token=STRING --repo=STRING --name=STRING
    Delete a label -
    https://developer.github.com/v3/issues/labels/#delete-a-label

  issues list-labels-on-issue --token=STRING --repo=STRING --number=INT-64
    List labels on an issue -
    https://developer.github.com/v3/issues/labels/#list-labels-on-an-issue

  issues add-labels --token=STRING --repo=STRING --number=INT-64 --labels=LABELS,...
    Add labels to an issue -
    https://developer.github.com/v3/issues/labels/#add-labels-to-an-issue

  issues remove-label --token=STRING --repo=STRING --number=INT-64 --name=STRING
    Remove a label from an issue -
    https://developer.github.com/v3/issues/labels/#remove-a-label-from-an-issue

  issues replace-labels --token=STRING --repo=STRING --number=INT-64
    Replace all labels for an issue -
    https://developer.github.com/v3/issues/labels/#replace-all-labels-for-an-issue

  issues remove-labels --token=STRING --repo=STRING --number=INT-64
    Remove all labels from an issue -
    https://developer.github.com/v3/issues/labels/#remove-all-labels-from-an-issue

  issues list-labels-for-milestone --token=STRING --repo=STRING --number=INT-64
    Get labels for every issue in a milestone -
    https://developer.github.com/v3/issues/labels/#get-labels-for-every-issue-in-a-milestone

  issues list-milestones-for-repo --token=STRING --repo=STRING
    List milestones for a repository -
    https://developer.github.com/v3/issues/milestones/#list-milestones-for-a-repository

  issues get-milestone --token=STRING --repo=STRING --number=INT-64
    Get a single milestone -
    https://developer.github.com/v3/issues/milestones/#get-a-single-milestone

  issues create-milestone --token=STRING --repo=STRING --title=STRING
    Create a milestone -
    https://developer.github.com/v3/issues/milestones/#create-a-milestone

  issues update-milestone --token=STRING --repo=STRING --number=INT-64
    Update a milestone -
    https://developer.github.com/v3/issues/milestones/#update-a-milestone

  issues delete-milestone --token=STRING --repo=STRING --number=INT-64
    Delete a milestone -
    https://developer.github.com/v3/issues/milestones/#delete-a-milestone

  issues list-events-for-timeline --token=STRING --mockingbird-preview --repo=STRING --number=INT-64
    List events for an issue -
    https://developer.github.com/v3/issues/timeline/#list-events-for-an-issue

  licenses list --token=STRING
    List all licenses -
    https://developer.github.com/v3/licenses/#list-all-licenses

  licenses get --token=STRING --license=STRING
    Get an individual license -
    https://developer.github.com/v3/licenses/#get-an-individual-license

  licenses get-for-repo --token=STRING --repo=STRING
    Get the contents of a repository's license -
    https://developer.github.com/v3/licenses/#get-the-contents-of-a-repositorys-license

  markdown render --token=STRING --text=STRING
    Render an arbitrary Markdown document -
    https://developer.github.com/v3/markdown/#render-an-arbitrary-markdown-document

  markdown render-raw --token=STRING --data=STRING
    Render a Markdown document in raw mode -
    https://developer.github.com/v3/markdown/#render-a-markdown-document-in-raw-mode

  meta get --token=STRING
    Get - https://developer.github.com/v3/meta/#meta

  migrations start-for-org --token=STRING --org=STRING --repositories=REPOSITORIES,...
    Start an organization migration -
    https://developer.github.com/v3/migrations/orgs/#start-an-organization-migration

  migrations list-for-org --token=STRING --org=STRING
    Get a list of organization migrations -
    https://developer.github.com/v3/migrations/orgs/#get-a-list-of-organization-migrations

  migrations get-status-for-org --token=STRING --org=STRING --migration_id=INT-64
    Get the status of an organization migration -
    https://developer.github.com/v3/migrations/orgs/#get-the-status-of-an-organization-migration

  migrations get-archive-for-org --token=STRING --org=STRING --migration_id=INT-64
    Download an organization migration archive -
    https://developer.github.com/v3/migrations/orgs/#download-an-organization-migration-archive

  migrations delete-archive-for-org --token=STRING --org=STRING --migration_id=INT-64
    Delete an organization migration archive -
    https://developer.github.com/v3/migrations/orgs/#delete-an-organization-migration-archive

  migrations unlock-repo-for-org --token=STRING --org=STRING --migration_id=INT-64 --repo_name=STRING
    Unlock an organization repository -
    https://developer.github.com/v3/migrations/orgs/#unlock-an-organization-repository

  migrations start-import --token=STRING --barred-rock-preview --repo=STRING --vcs_url=STRING
    Start an import -
    https://developer.github.com/v3/migrations/source_imports/#start-an-import

  migrations get-import-progress --token=STRING --barred-rock-preview --repo=STRING
    Get import progress -
    https://developer.github.com/v3/migrations/source_imports/#get-import-progress

  migrations update-import --token=STRING --barred-rock-preview --repo=STRING
    Update existing import -
    https://developer.github.com/v3/migrations/source_imports/#update-existing-import

  migrations get-commit-authors --token=STRING --barred-rock-preview --repo=STRING
    Get commit authors -
    https://developer.github.com/v3/migrations/source_imports/#get-commit-authors

  migrations map-commit-author --token=STRING --barred-rock-preview --repo=STRING --author_id=INT-64
    Map a commit author -
    https://developer.github.com/v3/migrations/source_imports/#map-a-commit-author

  migrations set-lfs-preference --token=STRING --barred-rock-preview --repo=STRING --use_lfs=STRING
    Set Git LFS preference -
    https://developer.github.com/v3/migrations/source_imports/#set-git-lfs-preference

  migrations get-large-files --token=STRING --barred-rock-preview --repo=STRING
    Get large files -
    https://developer.github.com/v3/migrations/source_imports/#get-large-files

  migrations cancel-import --token=STRING --barred-rock-preview --repo=STRING
    Cancel an import -
    https://developer.github.com/v3/migrations/source_imports/#cancel-an-import

  migrations start-for-authenticated-user --token=STRING --repositories=REPOSITORIES,...
    Start a user migration -
    https://developer.github.com/v3/migrations/users/#start-a-user-migration

  migrations list-for-authenticated-user --token=STRING
    Get a list of user migrations -
    https://developer.github.com/v3/migrations/users/#get-a-list-of-user-migrations

  migrations get-status-for-authenticated-user --token=STRING --migration_id=INT-64
    Get the status of a user migration -
    https://developer.github.com/v3/migrations/users/#get-the-status-of-a-user-migration

  migrations get-archive-for-authenticated-user --token=STRING --migration_id=INT-64
    Download a user migration archive -
    https://developer.github.com/v3/migrations/users/#download-a-user-migration-archive

  migrations delete-archive-for-authenticated-user --token=STRING --migration_id=INT-64
    Delete a user migration archive -
    https://developer.github.com/v3/migrations/users/#delete-a-user-migration-archive

  migrations unlock-repo-for-authenticated-user --token=STRING --migration_id=INT-64 --repo_name=STRING
    Unlock a user repository -
    https://developer.github.com/v3/migrations/users/#unlock-a-user-repository

  oauth-authorizations list-grants --token=STRING
    List your grants -
    https://developer.github.com/v3/oauth_authorizations/#list-your-grants

  oauth-authorizations get-grant --token=STRING --grant_id=INT-64
    Get a single grant -
    https://developer.github.com/v3/oauth_authorizations/#get-a-single-grant

  oauth-authorizations delete-grant --token=STRING --grant_id=INT-64
    Delete a grant -
    https://developer.github.com/v3/oauth_authorizations/#delete-a-grant

  oauth-authorizations list-authorizations --token=STRING
    List your authorizations -
    https://developer.github.com/v3/oauth_authorizations/#list-your-authorizations

  oauth-authorizations get-authorization --token=STRING --authorization_id=INT-64
    Get a single authorization -
    https://developer.github.com/v3/oauth_authorizations/#get-a-single-authorization

  oauth-authorizations create-authorization --token=STRING --note=STRING
    Create a new authorization -
    https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization

  oauth-authorizations get-or-create-authorization-for-app --token=STRING --client_id=STRING --client_secret=STRING
    Get-or-create an authorization for a specific app -
    https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app

  oauth-authorizations get-or-create-authorization-for-app-fingerprint --token=STRING --client_id=STRING --fingerprint=STRING --client_secret=STRING
    Get-or-create an authorization for a specific app and fingerprint -
    https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app-and-fingerprint

  oauth-authorizations update-authorization --token=STRING --authorization_id=INT-64
    Update an existing authorization -
    https://developer.github.com/v3/oauth_authorizations/#update-an-existing-authorization

  oauth-authorizations delete-authorization --token=STRING --authorization_id=INT-64
    Delete an authorization -
    https://developer.github.com/v3/oauth_authorizations/#delete-an-authorization

  oauth-authorizations check-authorization --token=STRING --client_id=STRING --access_token=STRING
    Check an authorization -
    https://developer.github.com/v3/oauth_authorizations/#check-an-authorization

  oauth-authorizations reset-authorization --token=STRING --client_id=STRING --access_token=STRING
    Reset an authorization -
    https://developer.github.com/v3/oauth_authorizations/#reset-an-authorization

  oauth-authorizations revoke-authorization-for-application --token=STRING --client_id=STRING --access_token=STRING
    Revoke an authorization for an application -
    https://developer.github.com/v3/oauth_authorizations/#revoke-an-authorization-for-an-application

  oauth-authorizations revoke-grant-for-application --token=STRING --client_id=STRING --access_token=STRING
    Revoke a grant for an application -
    https://developer.github.com/v3/oauth_authorizations/#revoke-a-grant-for-an-application

  orgs list-for-authenticated-user --token=STRING
    List your organizations -
    https://developer.github.com/v3/orgs/#list-your-organizations

  orgs list --token=STRING
    List all organizations -
    https://developer.github.com/v3/orgs/#list-all-organizations

  orgs list-for-user --token=STRING --username=STRING
    List user organizations -
    https://developer.github.com/v3/orgs/#list-user-organizations

  orgs get --token=STRING --org=STRING
    Get an organization -
    https://developer.github.com/v3/orgs/#get-an-organization

  orgs edit --token=STRING --org=STRING
    Edit an organization -
    https://developer.github.com/v3/orgs/#edit-an-organization

  orgs list-blocked-users --token=STRING --org=STRING
    List blocked users -
    https://developer.github.com/v3/orgs/blocking/#list-blocked-users

  orgs check-blocked-user --token=STRING --org=STRING --username=STRING
    Check whether a user is blocked from an organization -
    https://developer.github.com/v3/orgs/blocking/#check-whether-a-user-is-blocked-from-an-organization

  orgs block-user --token=STRING --org=STRING --username=STRING
    Block a user - https://developer.github.com/v3/orgs/blocking/#block-a-user

  orgs unblock-user --token=STRING --org=STRING --username=STRING
    Unblock a user -
    https://developer.github.com/v3/orgs/blocking/#unblock-a-user

  orgs list-members --token=STRING --org=STRING
    Members list - https://developer.github.com/v3/orgs/members/#members-list

  orgs check-membership --token=STRING --org=STRING --username=STRING
    Check membership -
    https://developer.github.com/v3/orgs/members/#check-membership

  orgs remove-member --token=STRING --org=STRING --username=STRING
    Remove a member -
    https://developer.github.com/v3/orgs/members/#remove-a-member

  orgs list-public-members --token=STRING --org=STRING
    Public members list -
    https://developer.github.com/v3/orgs/members/#public-members-list

  orgs check-public-membership --token=STRING --org=STRING --username=STRING
    Check public membership -
    https://developer.github.com/v3/orgs/members/#check-public-membership

  orgs publicize-membership --token=STRING --org=STRING --username=STRING
    Publicize a user's membership -
    https://developer.github.com/v3/orgs/members/#publicize-a-users-membership

  orgs conceal-membership --token=STRING --org=STRING --username=STRING
    Conceal a user's membership -
    https://developer.github.com/v3/orgs/members/#conceal-a-users-membership

  orgs get-membership-for-user --token=STRING --org=STRING --username=STRING
    Get organization membership -
    https://developer.github.com/v3/orgs/members/#get-organization-membership

  orgs add-or-update-membership --token=STRING --org=STRING --username=STRING
    Add or update organization membership -
    https://developer.github.com/v3/orgs/members/#add-or-update-organization-membership

  orgs remove-membership --token=STRING --org=STRING --username=STRING
    Remove organization membership -
    https://developer.github.com/v3/orgs/members/#remove-organization-membership

  orgs list-invitation-teams --token=STRING --org=STRING --invitation_id=INT-64
    List organization invitation teams -
    https://developer.github.com/v3/orgs/members/#list-organization-invitation-teams

  orgs list-pending-invitations --token=STRING --org=STRING
    List pending organization invitations -
    https://developer.github.com/v3/orgs/members/#list-pending-organization-invitations

  orgs create-invitation --token=STRING --org=STRING
    Create organization invitation -
    https://developer.github.com/v3/orgs/members/#create-organization-invitation

  orgs list-memberships --token=STRING
    List your organization memberships -
    https://developer.github.com/v3/orgs/members/#list-your-organization-memberships

  orgs get-membership-for-authenticated-user --token=STRING --org=STRING
    Get your organization membership -
    https://developer.github.com/v3/orgs/members/#get-your-organization-membership

  orgs edit-membership --token=STRING --org=STRING --state=STRING
    Edit your organization membership -
    https://developer.github.com/v3/orgs/members/#edit-your-organization-membership

  orgs list-outside-collaborators --token=STRING --org=STRING
    List outside collaborators -
    https://developer.github.com/v3/orgs/outside_collaborators/#list-outside-collaborators

  orgs remove-outside-collaborator --token=STRING --org=STRING --username=STRING
    Remove outside collaborator -
    https://developer.github.com/v3/orgs/outside_collaborators/#remove-outside-collaborator

  orgs convert-member-to-outside-collaborator --token=STRING --org=STRING --username=STRING
    Convert member to outside collaborator -
    https://developer.github.com/v3/orgs/outside_collaborators/#convert-member-to-outside-collaborator

  orgs list-hooks --token=STRING --org=STRING
    List hooks - https://developer.github.com/v3/orgs/hooks/#list-hooks

  orgs get-hook --token=STRING --org=STRING --hook_id=INT-64
    Get single hook -
    https://developer.github.com/v3/orgs/hooks/#get-single-hook

  orgs ping-hook --token=STRING --org=STRING --hook_id=INT-64
    Ping a hook - https://developer.github.com/v3/orgs/hooks/#ping-a-hook

  orgs delete-hook --token=STRING --org=STRING --hook_id=INT-64
    Delete a hook - https://developer.github.com/v3/orgs/hooks/#delete-a-hook

  projects list-for-repo --token=STRING --inertia-preview --repo=STRING
    List repository projects -
    https://developer.github.com/v3/projects/#list-repository-projects

  projects list-for-org --token=STRING --inertia-preview --org=STRING
    List organization projects -
    https://developer.github.com/v3/projects/#list-organization-projects

  projects get --token=STRING --inertia-preview --project_id=INT-64
    Get a project - https://developer.github.com/v3/projects/#get-a-project

  projects create-for-repo --token=STRING --inertia-preview --repo=STRING --name=STRING
    Create a repository project -
    https://developer.github.com/v3/projects/#create-a-repository-project

  projects create-for-org --token=STRING --inertia-preview --org=STRING --name=STRING
    Create an organization project -
    https://developer.github.com/v3/projects/#create-an-organization-project

  projects update --token=STRING --inertia-preview --project_id=INT-64
    Update a project -
    https://developer.github.com/v3/projects/#update-a-project

  projects delete --token=STRING --inertia-preview --project_id=INT-64
    Delete a project -
    https://developer.github.com/v3/projects/#delete-a-project

  projects list-cards --token=STRING --inertia-preview --column_id=INT-64
    List project cards -
    https://developer.github.com/v3/projects/cards/#list-project-cards

  projects get-card --token=STRING --inertia-preview --card_id=INT-64
    Get a project card -
    https://developer.github.com/v3/projects/cards/#get-a-project-card

  projects create-card --token=STRING --inertia-preview --column_id=INT-64
    Create a project card -
    https://developer.github.com/v3/projects/cards/#create-a-project-card

  projects update-card --token=STRING --inertia-preview --card_id=INT-64
    Update a project card -
    https://developer.github.com/v3/projects/cards/#update-a-project-card

  projects delete-card --token=STRING --inertia-preview --card_id=INT-64
    Delete a project card -
    https://developer.github.com/v3/projects/cards/#delete-a-project-card

  projects move-card --token=STRING --inertia-preview --card_id=INT-64 --position=STRING
    Move a project card -
    https://developer.github.com/v3/projects/cards/#move-a-project-card

  projects list-collaborators --token=STRING --inertia-preview --project_id=INT-64
    List collaborators -
    https://developer.github.com/v3/projects/collaborators/#list-collaborators

  projects review-user-permission-level --token=STRING --inertia-preview --project_id=INT-64 --username=STRING
    Review a user's permission level -
    https://developer.github.com/v3/projects/collaborators/#review-a-users-permission-level

  projects add-collaborator --token=STRING --inertia-preview --project_id=INT-64 --username=STRING
    Add user as a collaborator -
    https://developer.github.com/v3/projects/collaborators/#add-user-as-a-collaborator

  projects remove-collaborator --token=STRING --inertia-preview --project_id=INT-64 --username=STRING
    Remove user as a collaborator -
    https://developer.github.com/v3/projects/collaborators/#remove-user-as-a-collaborator

  projects list-columns --token=STRING --inertia-preview --project_id=INT-64
    List project columns -
    https://developer.github.com/v3/projects/columns/#list-project-columns

  projects get-column --token=STRING --inertia-preview --column_id=INT-64
    Get a project column -
    https://developer.github.com/v3/projects/columns/#get-a-project-column

  projects create-column --token=STRING --inertia-preview --project_id=INT-64 --name=STRING
    Create a project column -
    https://developer.github.com/v3/projects/columns/#create-a-project-column

  projects update-column --token=STRING --inertia-preview --column_id=INT-64 --name=STRING
    Update a project column -
    https://developer.github.com/v3/projects/columns/#update-a-project-column

  projects delete-column --token=STRING --inertia-preview --column_id=INT-64
    Delete a project column -
    https://developer.github.com/v3/projects/columns/#delete-a-project-column

  projects move-column --token=STRING --inertia-preview --column_id=INT-64 --position=STRING
    Move a project column -
    https://developer.github.com/v3/projects/columns/#move-a-project-column

  pulls list --token=STRING --repo=STRING
    List pull requests -
    https://developer.github.com/v3/pulls/#list-pull-requests

  pulls get --token=STRING --repo=STRING --number=INT-64
    Get a single pull request -
    https://developer.github.com/v3/pulls/#get-a-single-pull-request

  pulls create --token=STRING --repo=STRING --title=STRING --head=STRING --base=STRING
    Create a pull request -
    https://developer.github.com/v3/pulls/#create-a-pull-request

  pulls create-from-issue --token=STRING --repo=STRING --issue=INT-64 --head=STRING --base=STRING
    Create a Pull Request from an Issue -
    https://developer.github.com/v3/pulls/#create-a-pull-request

  pulls update --token=STRING --repo=STRING --number=INT-64
    Update a pull request -
    https://developer.github.com/v3/pulls/#update-a-pull-request

  pulls list-commits --token=STRING --repo=STRING --number=INT-64
    List commits on a pull request -
    https://developer.github.com/v3/pulls/#list-commits-on-a-pull-request

  pulls list-files --token=STRING --repo=STRING --number=INT-64
    List pull requests files -
    https://developer.github.com/v3/pulls/#list-pull-requests-files

  pulls check-if-merged --token=STRING --repo=STRING --number=INT-64
    Get if a pull request has been merged -
    https://developer.github.com/v3/pulls/#get-if-a-pull-request-has-been-merged

  pulls merge --token=STRING --repo=STRING --number=INT-64
    Merge a pull request (Merge Button) -
    https://developer.github.com/v3/pulls/#merge-a-pull-request-merge-button

  pulls list-reviews --token=STRING --repo=STRING --number=INT-64
    List reviews on a pull request -
    https://developer.github.com/v3/pulls/reviews/#list-reviews-on-a-pull-request

  pulls get-review --token=STRING --repo=STRING --number=INT-64 --review_id=INT-64
    Get a single review -
    https://developer.github.com/v3/pulls/reviews/#get-a-single-review

  pulls delete-pending-review --token=STRING --repo=STRING --number=INT-64 --review_id=INT-64
    Delete a pending review -
    https://developer.github.com/v3/pulls/reviews/#delete-a-pending-review

  pulls get-comments-for-review --token=STRING --repo=STRING --number=INT-64 --review_id=INT-64
    Get comments for a single review -
    https://developer.github.com/v3/pulls/reviews/#get-comments-for-a-single-review

  pulls submit-review --token=STRING --repo=STRING --number=INT-64 --review_id=INT-64 --event=STRING
    Submit a pull request review -
    https://developer.github.com/v3/pulls/reviews/#submit-a-pull-request-review

  pulls dismiss-review --token=STRING --repo=STRING --number=INT-64 --review_id=INT-64 --message=STRING
    Dismiss a pull request review -
    https://developer.github.com/v3/pulls/reviews/#dismiss-a-pull-request-review

  pulls list-comments --token=STRING --repo=STRING --number=INT-64
    List comments on a pull request -
    https://developer.github.com/v3/pulls/comments/#list-comments-on-a-pull-request

  pulls list-comments-for-repo --token=STRING --repo=STRING
    List comments in a repository -
    https://developer.github.com/v3/pulls/comments/#list-comments-in-a-repository

  pulls get-comment --token=STRING --repo=STRING --comment_id=INT-64
    Get a single comment -
    https://developer.github.com/v3/pulls/comments/#get-a-single-comment

  pulls create-comment --token=STRING --repo=STRING --number=INT-64 --body=STRING --commit_id=STRING --path=STRING --position=INT-64
    Create a comment -
    https://developer.github.com/v3/pulls/comments/#create-a-comment

  pulls create-comment-reply --token=STRING --repo=STRING --number=INT-64 --body=STRING --in_reply_to=INT-64
    Create a comment reply -
    https://developer.github.com/v3/pulls/comments/#create-a-comment

  pulls edit-comment --token=STRING --repo=STRING --comment_id=INT-64 --body=STRING
    Edit a comment -
    https://developer.github.com/v3/pulls/comments/#edit-a-comment

  pulls delete-comment --token=STRING --repo=STRING --comment_id=INT-64
    Delete a comment -
    https://developer.github.com/v3/pulls/comments/#delete-a-comment

  pulls list-review-requests --token=STRING --repo=STRING --number=INT-64
    List review requests -
    https://developer.github.com/v3/pulls/review_requests/#list-review-requests

  pulls create-review-request --token=STRING --repo=STRING --number=INT-64
    Create a review request -
    https://developer.github.com/v3/pulls/review_requests/#create-a-review-request

  pulls delete-review-request --token=STRING --repo=STRING --number=INT-64
    Delete a review request -
    https://developer.github.com/v3/pulls/review_requests/#delete-a-review-request

  rate-limit get --token=STRING
    Get your current rate limit status -
    https://developer.github.com/v3/rate_limit/#get-your-current-rate-limit-status

  reactions list-for-commit-comment --token=STRING --squirrel-girl-preview --repo=STRING --comment_id=INT-64
    List reactions for a commit comment -
    https://developer.github.com/v3/reactions/#list-reactions-for-a-commit-comment

  reactions create-for-commit-comment --token=STRING --squirrel-girl-preview --repo=STRING --comment_id=INT-64 --content=STRING
    Create reaction for a commit comment -
    https://developer.github.com/v3/reactions/#create-reaction-for-a-commit-comment

  reactions list-for-issue --token=STRING --squirrel-girl-preview --repo=STRING --number=INT-64
    List reactions for an issue -
    https://developer.github.com/v3/reactions/#list-reactions-for-an-issue

  reactions create-for-issue --token=STRING --squirrel-girl-preview --repo=STRING --number=INT-64 --content=STRING
    Create reaction for an issue -
    https://developer.github.com/v3/reactions/#create-reaction-for-an-issue

  reactions list-for-issue-comment --token=STRING --squirrel-girl-preview --repo=STRING --comment_id=INT-64
    List reactions for an issue comment -
    https://developer.github.com/v3/reactions/#list-reactions-for-an-issue-comment

  reactions create-for-issue-comment --token=STRING --squirrel-girl-preview --repo=STRING --comment_id=INT-64 --content=STRING
    Create reaction for an issue comment -
    https://developer.github.com/v3/reactions/#create-reaction-for-an-issue-comment

  reactions list-for-pull-request-review-comment --token=STRING --squirrel-girl-preview --repo=STRING --comment_id=INT-64
    List reactions for a pull request review comment -
    https://developer.github.com/v3/reactions/#list-reactions-for-a-pull-request-review-comment

  reactions create-for-pull-request-review-comment --token=STRING --squirrel-girl-preview --repo=STRING --comment_id=INT-64 --content=STRING
    Create reaction for a pull request review comment -
    https://developer.github.com/v3/reactions/#create-reaction-for-a-pull-request-review-comment

  reactions list-for-team-discussion --token=STRING --echo-preview --squirrel-girl-preview --team_id=INT-64 --discussion_number=INT-64
    List reactions for a team discussion -
    https://developer.github.com/v3/reactions/#list-reactions-for-a-team-discussion

  reactions create-for-team-discussion --token=STRING --echo-preview --squirrel-girl-preview --team_id=INT-64 --discussion_number=INT-64 --content=STRING
    Create reaction for a team discussion -
    https://developer.github.com/v3/reactions/#create-reaction-for-a-team-discussion

  reactions list-for-team-discussion-comment --token=STRING --echo-preview --squirrel-girl-preview --team_id=INT-64 --discussion_number=INT-64 --comment_number=INT-64
    List reactions for a team discussion comment -
    https://developer.github.com/v3/reactions/#list-reactions-for-a-team-discussion-comment

  reactions create-for-team-discussion-comment --token=STRING --echo-preview --squirrel-girl-preview --team_id=INT-64 --discussion_number=INT-64 --comment_number=INT-64 --content=STRING
    Create reaction for a team discussion comment -
    https://developer.github.com/v3/reactions/#create-reaction-for-a-team-discussion-comment

  reactions delete --token=STRING --echo-preview --squirrel-girl-preview --reaction_id=INT-64
    Delete a reaction -
    https://developer.github.com/v3/reactions/#delete-a-reaction

  repos list --token=STRING
    List your repositories -
    https://developer.github.com/v3/repos/#list-your-repositories

  repos list-for-user --token=STRING --username=STRING
    List user repositories -
    https://developer.github.com/v3/repos/#list-user-repositories

  repos list-for-org --token=STRING --org=STRING
    List organization repositories -
    https://developer.github.com/v3/repos/#list-organization-repositories

  repos list-public --token=STRING
    List all public repositories -
    https://developer.github.com/v3/repos/#list-all-public-repositories

  repos create-for-authenticated-user --token=STRING --name=STRING
    Create a new repository for the authenticated user -
    https://developer.github.com/v3/repos/#create

  repos create-in-org --token=STRING --org=STRING --name=STRING
    Create a new repository in this organization -
    https://developer.github.com/v3/repos/#create

  repos get --token=STRING --repo=STRING
    Get - https://developer.github.com/v3/repos/#get

  repos edit --token=STRING --repo=STRING --name=STRING
    Edit - https://developer.github.com/v3/repos/#edit

  repos list-topics --token=STRING --repo=STRING
    List all topics for a repository -
    https://developer.github.com/v3/repos/#list-all-topics-for-a-repository

  repos replace-topics --token=STRING --repo=STRING --names=NAMES,...
    Replace all topics for a repository -
    https://developer.github.com/v3/repos/#replace-all-topics-for-a-repository

  repos list-contributors --token=STRING --repo=STRING
    List contributors - https://developer.github.com/v3/repos/#list-contributors

  repos list-languages --token=STRING --repo=STRING
    List languages - https://developer.github.com/v3/repos/#list-languages

  repos list-teams --token=STRING --repo=STRING
    List teams - https://developer.github.com/v3/repos/#list-teams

  repos list-tags --token=STRING --repo=STRING
    List tags - https://developer.github.com/v3/repos/#list-tags

  repos delete --token=STRING --repo=STRING
    Delete a repository -
    https://developer.github.com/v3/repos/#delete-a-repository

  repos transfer --token=STRING --nightshade-preview --repo=STRING
    Transfer a repository -
    https://developer.github.com/v3/repos/#transfer-a-repository

  repos list-branches --token=STRING --repo=STRING
    List branches -
    https://developer.github.com/v3/repos/branches/#list-branches

  repos get-branch --token=STRING --repo=STRING --branch=STRING
    Get branch - https://developer.github.com/v3/repos/branches/#get-branch

  repos get-branch-protection --token=STRING --repo=STRING --branch=STRING
    Get branch protection -
    https://developer.github.com/v3/repos/branches/#get-branch-protection

  repos remove-branch-protection --token=STRING --repo=STRING --branch=STRING
    Remove branch protection -
    https://developer.github.com/v3/repos/branches/#remove-branch-protection

  repos get-protected-branch-required-status-checks --token=STRING --repo=STRING --branch=STRING
    Get required status checks of protected branch -
    https://developer.github.com/v3/repos/branches/#get-required-status-checks-of-protected-branch

  repos update-protected-branch-required-status-checks --token=STRING --repo=STRING --branch=STRING
    Update required status checks of protected branch -
    https://developer.github.com/v3/repos/branches/#update-required-status-checks-of-protected-branch

  repos remove-protected-branch-required-status-checks --token=STRING --repo=STRING --branch=STRING
    Remove required status checks of protected branch -
    https://developer.github.com/v3/repos/branches/#remove-required-status-checks-of-protected-branch

  repos list-protected-branch-required-status-checks-contexts --token=STRING --repo=STRING --branch=STRING
    List required status checks contexts of protected branch -
    https://developer.github.com/v3/repos/branches/#list-required-status-checks-contexts-of-protected-branch

  repos replace-protected-branch-required-status-checks-contexts --token=STRING --repo=STRING --branch=STRING --contexts=CONTEXTS,...
    Replace required status checks contexts of protected branch -
    https://developer.github.com/v3/repos/branches/#replace-required-status-checks-contexts-of-protected-branch

  repos add-protected-branch-required-status-checks-contexts --token=STRING --repo=STRING --branch=STRING --contexts=CONTEXTS,...
    Add required status checks contexts of protected branch -
    https://developer.github.com/v3/repos/branches/#add-required-status-checks-contexts-of-protected-branch

  repos remove-protected-branch-required-status-checks-contexts --token=STRING --repo=STRING --branch=STRING --contexts=CONTEXTS,...
    Remove required status checks contexts of protected branch -
    https://developer.github.com/v3/repos/branches/#remove-required-status-checks-contexts-of-protected-branch

  repos get-protected-branch-pull-request-review-enforcement --token=STRING --repo=STRING --branch=STRING
    Get pull request review enforcement of protected branch -
    https://developer.github.com/v3/repos/branches/#get-pull-request-review-enforcement-of-protected-branch

  repos remove-protected-branch-pull-request-review-enforcement --token=STRING --repo=STRING --branch=STRING
    Remove pull request review enforcement of protected branch -
    https://developer.github.com/v3/repos/branches/#remove-pull-request-review-enforcement-of-protected-branch

  repos get-protected-branch-required-signatures --token=STRING --zzzax-preview --repo=STRING --branch=STRING
    Get required signatures of protected branch -
    https://developer.github.com/v3/repos/branches/#get-required-signatures-of-protected-branch

  repos add-protected-branch-required-signatures --token=STRING --zzzax-preview --repo=STRING --branch=STRING
    Add required signatures of protected branch -
    https://developer.github.com/v3/repos/branches/#add-required-signatures-of-protected-branch

  repos remove-protected-branch-required-signatures --token=STRING --zzzax-preview --repo=STRING --branch=STRING
    Remove required signatures of protected branch -
    https://developer.github.com/v3/repos/branches/#remove-required-signatures-of-protected-branch

  repos get-protected-branch-admin-enforcement --token=STRING --repo=STRING --branch=STRING
    Get admin enforcement of protected branch -
    https://developer.github.com/v3/repos/branches/#get-admin-enforcement-of-protected-branch

  repos add-protected-branch-admin-enforcement --token=STRING --repo=STRING --branch=STRING
    Add admin enforcement of protected branch -
    https://developer.github.com/v3/repos/branches/#add-admin-enforcement-of-protected-branch

  repos remove-protected-branch-admin-enforcement --token=STRING --repo=STRING --branch=STRING
    Remove admin enforcement of protected branch -
    https://developer.github.com/v3/repos/branches/#remove-admin-enforcement-of-protected-branch

  repos get-protected-branch-restrictions --token=STRING --repo=STRING --branch=STRING
    Get restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#get-restrictions-of-protected-branch

  repos remove-protected-branch-restrictions --token=STRING --repo=STRING --branch=STRING
    Remove restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#remove-restrictions-of-protected-branch

  repos list-protected-branch-team-restrictions --token=STRING --repo=STRING --branch=STRING
    List team restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#list-team-restrictions-of-protected-branch

  repos replace-protected-branch-team-restrictions --token=STRING --repo=STRING --branch=STRING --teams=TEAMS,...
    Replace team restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#replace-team-restrictions-of-protected-branch

  repos add-protected-branch-team-restrictions --token=STRING --repo=STRING --branch=STRING --teams=TEAMS,...
    Add team restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#add-team-restrictions-of-protected-branch

  repos remove-protected-branch-team-restrictions --token=STRING --repo=STRING --branch=STRING --teams=TEAMS,...
    Remove team restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#remove-team-restrictions-of-protected-branch

  repos list-protected-branch-user-restrictions --token=STRING --repo=STRING --branch=STRING
    List user restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#list-user-restrictions-of-protected-branch

  repos replace-protected-branch-user-restrictions --token=STRING --repo=STRING --branch=STRING --users=USERS,...
    Replace user restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#replace-user-restrictions-of-protected-branch

  repos add-protected-branch-user-restrictions --token=STRING --repo=STRING --branch=STRING --users=USERS,...
    Add user restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#add-user-restrictions-of-protected-branch

  repos remove-protected-branch-user-restrictions --token=STRING --repo=STRING --branch=STRING --users=USERS,...
    Remove user restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#remove-user-restrictions-of-protected-branch

  repos list-collaborators --token=STRING --repo=STRING
    List collaborators -
    https://developer.github.com/v3/repos/collaborators/#list-collaborators

  repos check-collaborator --token=STRING --repo=STRING --username=STRING
    Check if a user is a collaborator -
    https://developer.github.com/v3/repos/collaborators/#check-if-a-user-is-a-collaborator

  repos get-collaborator-permission-level --token=STRING --repo=STRING --username=STRING
    Review a user's permission level -
    https://developer.github.com/v3/repos/collaborators/#review-a-users-permission-level

  repos add-collaborator --token=STRING --repo=STRING --username=STRING
    Add user as a collaborator -
    https://developer.github.com/v3/repos/collaborators/#add-user-as-a-collaborator

  repos remove-collaborator --token=STRING --repo=STRING --username=STRING
    Remove user as a collaborator -
    https://developer.github.com/v3/repos/collaborators/#remove-user-as-a-collaborator

  repos list-commit-comments --token=STRING --repo=STRING
    List commit comments for a repository -
    https://developer.github.com/v3/repos/comments/#list-commit-comments-for-a-repository

  repos list-comments-for-commit --token=STRING --repo=STRING --ref=STRING
    List comments for a single commit -
    https://developer.github.com/v3/repos/comments/#list-comments-for-a-single-commit

  repos create-commit-comment --token=STRING --repo=STRING --sha=STRING --body=STRING
    Create a commit comment -
    https://developer.github.com/v3/repos/comments/#create-a-commit-comment

  repos get-commit-comment --token=STRING --repo=STRING --comment_id=INT-64
    Get a single commit comment -
    https://developer.github.com/v3/repos/comments/#get-a-single-commit-comment

  repos update-commit-comment --token=STRING --repo=STRING --comment_id=INT-64 --body=STRING
    Update a commit comment -
    https://developer.github.com/v3/repos/comments/#update-a-commit-comment

  repos delete-commit-comment --token=STRING --repo=STRING --comment_id=INT-64
    Delete a commit comment -
    https://developer.github.com/v3/repos/comments/#delete-a-commit-comment

  repos list-commits --token=STRING --repo=STRING
    List commits on a repository -
    https://developer.github.com/v3/repos/commits/#list-commits-on-a-repository

  repos get-commit --token=STRING --repo=STRING --sha=STRING
    Get a single commit -
    https://developer.github.com/v3/repos/commits/#get-a-single-commit

  repos get-commit-ref-sha --token=STRING --repo=STRING --ref=STRING
    Get the SHA-1 of a commit reference -
    https://developer.github.com/v3/repos/commits/#get-the-sha-1-of-a-commit-reference

  repos compare-commits --token=STRING --repo=STRING --base=STRING --head=STRING
    Compare two commits -
    https://developer.github.com/v3/repos/commits/#compare-two-commits

  repos retrieve-community-profile-metrics --token=STRING --repo=STRING
    Retrieve community profile metrics -
    https://developer.github.com/v3/repos/community/#retrieve-community-profile-metrics

  repos get-readme --token=STRING --repo=STRING
    Get the README -
    https://developer.github.com/v3/repos/contents/#get-the-readme

  repos get-contents --token=STRING --repo=STRING --path=STRING
    Get contents - https://developer.github.com/v3/repos/contents/#get-contents

  repos get-archive-link --token=STRING --repo=STRING --archive_format=STRING --ref=STRING
    Get archive link -
    https://developer.github.com/v3/repos/contents/#get-archive-link

  repos list-deploy-keys --token=STRING --repo=STRING
    List deploy keys -
    https://developer.github.com/v3/repos/keys/#list-deploy-keys

  repos get-deploy-key --token=STRING --repo=STRING --key_id=INT-64
    Get a deploy key -
    https://developer.github.com/v3/repos/keys/#get-a-deploy-key

  repos add-deploy-key --token=STRING --repo=STRING --key=STRING
    Add a new deploy key -
    https://developer.github.com/v3/repos/keys/#add-a-new-deploy-key

  repos remove-deploy-key --token=STRING --repo=STRING --key_id=INT-64
    Remove a deploy key -
    https://developer.github.com/v3/repos/keys/#remove-a-deploy-key

  repos list-deployments --token=STRING --repo=STRING
    List deployments -
    https://developer.github.com/v3/repos/deployments/#list-deployments

  repos get-deployment --token=STRING --repo=STRING --deployment_id=INT-64
    Get a single deployment -
    https://developer.github.com/v3/repos/deployments/#get-a-single-deployment

  repos create-deployment --token=STRING --repo=STRING --ref=STRING
    Create a deployment -
    https://developer.github.com/v3/repos/deployments/#create-a-deployment

  repos list-deployment-statuses --token=STRING --repo=STRING --deployment_id=INT-64
    List deployment statuses -
    https://developer.github.com/v3/repos/deployments/#list-deployment-statuses

  repos get-deployment-status --token=STRING --repo=STRING --deployment_id=INT-64 --status_id=INT-64
    Get a single deployment status -
    https://developer.github.com/v3/repos/deployments/#get-a-single-deployment-status

  repos create-deployment-status --token=STRING --repo=STRING --deployment_id=INT-64 --state=STRING
    Create a deployment status -
    https://developer.github.com/v3/repos/deployments/#create-a-deployment-status

  repos list-downloads --token=STRING --repo=STRING
    List downloads for a repository -
    https://developer.github.com/v3/repos/downloads/#list-downloads-for-a-repository

  repos get-download --token=STRING --repo=STRING --download_id=INT-64
    Get a single download -
    https://developer.github.com/v3/repos/downloads/#get-a-single-download

  repos delete-download --token=STRING --repo=STRING --download_id=INT-64
    Delete a download -
    https://developer.github.com/v3/repos/downloads/#delete-a-download

  repos list-forks --token=STRING --repo=STRING
    List forks - https://developer.github.com/v3/repos/forks/#list-forks

  repos create-fork --token=STRING --repo=STRING
    Create a fork - https://developer.github.com/v3/repos/forks/#create-a-fork

  repos list-invitations --token=STRING --repo=STRING
    List invitations for a repository -
    https://developer.github.com/v3/repos/invitations/#list-invitations-for-a-repository

  repos delete-invitation --token=STRING --repo=STRING --invitation_id=INT-64
    Delete a repository invitation -
    https://developer.github.com/v3/repos/invitations/#delete-a-repository-invitation

  repos update-invitation --token=STRING --repo=STRING --invitation_id=INT-64
    Update a repository invitation -
    https://developer.github.com/v3/repos/invitations/#update-a-repository-invitation

  repos list-invitations-for-authenticated-user --token=STRING
    List a user's repository invitations -
    https://developer.github.com/v3/repos/invitations/#list-a-users-repository-invitations

  repos accept-invitation --token=STRING --invitation_id=INT-64
    Accept a repository invitation -
    https://developer.github.com/v3/repos/invitations/#accept-a-repository-invitation

  repos decline-invitation --token=STRING --invitation_id=INT-64
    Decline a repository invitation -
    https://developer.github.com/v3/repos/invitations/#decline-a-repository-invitation

  repos merge --token=STRING --repo=STRING --base=STRING --head=STRING
    Perform a merge -
    https://developer.github.com/v3/repos/merging/#perform-a-merge

  repos get-pages --token=STRING --mister-fantastic-preview --repo=STRING
    Get information about a Pages site -
    https://developer.github.com/v3/repos/pages/#get-information-about-a-pages-site

  repos update-information-about-pages-site --token=STRING --mister-fantastic-preview --repo=STRING
    Update information about a Pages site -
    https://developer.github.com/v3/repos/pages/#update-information-about-a-pages-site

  repos request-page-build --token=STRING --mister-fantastic-preview --repo=STRING
    Request a page build -
    https://developer.github.com/v3/repos/pages/#request-a-page-build

  repos list-pages-builds --token=STRING --repo=STRING
    List Pages builds -
    https://developer.github.com/v3/repos/pages/#list-pages-builds

  repos get-latest-pages-build --token=STRING --repo=STRING
    Get latest Pages build -
    https://developer.github.com/v3/repos/pages/#get-latest-pages-build

  repos get-pages-build --token=STRING --repo=STRING --build_id=INT-64
    Get a specific Pages build -
    https://developer.github.com/v3/repos/pages/#get-a-specific-pages-build

  repos list-releases --token=STRING --repo=STRING
    List releases for a repository -
    https://developer.github.com/v3/repos/releases/#list-releases-for-a-repository

  repos get-release --token=STRING --repo=STRING --release_id=INT-64
    Get a single release -
    https://developer.github.com/v3/repos/releases/#get-a-single-release

  repos get-latest-release --token=STRING --repo=STRING
    Get the latest release -
    https://developer.github.com/v3/repos/releases/#get-the-latest-release

  repos get-release-by-tag --token=STRING --repo=STRING --tag=STRING
    Get a release by tag name -
    https://developer.github.com/v3/repos/releases/#get-a-release-by-tag-name

  repos create-release --token=STRING --repo=STRING --tag_name=STRING
    Create a release -
    https://developer.github.com/v3/repos/releases/#create-a-release

  repos edit-release --token=STRING --repo=STRING --release_id=INT-64
    Edit a release -
    https://developer.github.com/v3/repos/releases/#edit-a-release

  repos delete-release --token=STRING --repo=STRING --release_id=INT-64
    Delete a release -
    https://developer.github.com/v3/repos/releases/#delete-a-release

  repos list-assets-for-release --token=STRING --repo=STRING --release_id=INT-64
    List assets for a release -
    https://developer.github.com/v3/repos/releases/#list-assets-for-a-release

  repos get-release-asset --token=STRING --repo=STRING --asset_id=INT-64
    Get a single release asset -
    https://developer.github.com/v3/repos/releases/#get-a-single-release-asset

  repos edit-release-asset --token=STRING --repo=STRING --asset_id=INT-64
    Edit a release asset -
    https://developer.github.com/v3/repos/releases/#edit-a-release-asset

  repos delete-release-asset --token=STRING --repo=STRING --asset_id=INT-64
    Delete a release asset -
    https://developer.github.com/v3/repos/releases/#delete-a-release-asset

  repos get-contributors-stats --token=STRING --repo=STRING
    Get contributors list with additions, deletions, and commit counts -
    https://developer.github.com/v3/repos/statistics/#get-contributors-list-with-additions-deletions-and-commit-counts

  repos get-commit-activity-stats --token=STRING --repo=STRING
    Get the last year of commit activity data -
    https://developer.github.com/v3/repos/statistics/#get-the-last-year-of-commit-activity-data

  repos get-code-frequency-stats --token=STRING --repo=STRING
    Get the number of additions and deletions per week -
    https://developer.github.com/v3/repos/statistics/#get-the-number-of-additions-and-deletions-per-week

  repos get-participation-stats --token=STRING --repo=STRING
    Get the weekly commit count for the repository owner and everyone else -
    https://developer.github.com/v3/repos/statistics/#get-the-weekly-commit-count-for-the-repository-owner-and-everyone-else

  repos get-punch-card-stats --token=STRING --repo=STRING
    Get the number of commits per hour in each day -
    https://developer.github.com/v3/repos/statistics/#get-the-number-of-commits-per-hour-in-each-day

  repos create-status --token=STRING --repo=STRING --sha=STRING --state=STRING
    Create a status -
    https://developer.github.com/v3/repos/statuses/#create-a-status

  repos list-statuses-for-ref --token=STRING --repo=STRING --ref=STRING
    List statuses for a specific ref -
    https://developer.github.com/v3/repos/statuses/#list-statuses-for-a-specific-ref

  repos get-combined-status-for-ref --token=STRING --repo=STRING --ref=STRING
    Get the combined status for a specific ref -
    https://developer.github.com/v3/repos/statuses/#get-the-combined-status-for-a-specific-ref

  repos get-top-referrers --token=STRING --repo=STRING
    List referrers -
    https://developer.github.com/v3/repos/traffic/#list-referrers

  repos get-top-paths --token=STRING --repo=STRING
    List paths - https://developer.github.com/v3/repos/traffic/#list-paths

  repos get-views --token=STRING --repo=STRING
    Views - https://developer.github.com/v3/repos/traffic/#views

  repos get-clones --token=STRING --repo=STRING
    Clones - https://developer.github.com/v3/repos/traffic/#clones

  repos list-hooks --token=STRING --repo=STRING
    List hooks - https://developer.github.com/v3/repos/hooks/#list-hooks

  repos get-hook --token=STRING --repo=STRING --hook_id=INT-64
    Get single hook -
    https://developer.github.com/v3/repos/hooks/#get-single-hook

  repos test-push-hook --token=STRING --repo=STRING --hook_id=INT-64
    Test a push hook -
    https://developer.github.com/v3/repos/hooks/#test-a-push-hook

  repos ping-hook --token=STRING --repo=STRING --hook_id=INT-64
    Ping a hook - https://developer.github.com/v3/repos/hooks/#ping-a-hook

  repos delete-hook --token=STRING --repo=STRING --hook_id=INT-64
    Delete a hook - https://developer.github.com/v3/repos/hooks/#delete-a-hook

  scim list-provisioned-identities --token=STRING --cloud-9-preview --org=STRING
    Get a list of provisioned identities -
    https://developer.github.com/v3/scim/#get-a-list-of-provisioned-identities

  scim get-provisioning-details-for-user --token=STRING --cloud-9-preview --org=STRING --external_identity_guid=STRING
    Get provisioning details for a single user -
    https://developer.github.com/v3/scim/#get-provisioning-details-for-a-single-user

  scim provision-invite-users --token=STRING --cloud-9-preview --org=STRING
    Provision and invite users -
    https://developer.github.com/v3/scim/#provision-and-invite-users

  scim update-provisioned-org-membership --token=STRING --cloud-9-preview --org=STRING --external_identity_guid=STRING
    Update a provisioned organization membership -
    https://developer.github.com/v3/scim/#update-a-provisioned-organization-membership

  scim update-user-attribute --token=STRING --cloud-9-preview --org=STRING --external_identity_guid=STRING
    Update a user attribute -
    https://developer.github.com/v3/scim/#update-a-user-attribute

  scim remove-user-from-org --token=STRING --cloud-9-preview --org=STRING --external_identity_guid=STRING
    Remove a user from the organization -
    https://developer.github.com/v3/scim/#remove-a-user-from-the-organization

  search repos --token=STRING --q=STRING
    Search repositories -
    https://developer.github.com/v3/search/#search-repositories

  search commits --token=STRING --cloak-preview --q=STRING
    Search commits - https://developer.github.com/v3/search/#search-commits

  search code --token=STRING --q=STRING
    Search code - https://developer.github.com/v3/search/#search-code

  search issues --token=STRING --q=STRING
    Search issues - https://developer.github.com/v3/search/#search-issues

  search users --token=STRING --q=STRING
    Search users - https://developer.github.com/v3/search/#search-users

  search topics --token=STRING --q=STRING
    Search topics - https://developer.github.com/v3/search/#search-topics

  search labels --token=STRING --repository_id=INT-64 --q=STRING
    Search labels - https://developer.github.com/v3/search/#search-labels

  search issues-legacy --token=STRING --owner=STRING --repository=STRING --state=STRING --keyword=STRING
    Search issues - https://developer.github.com/v3/search/legacy/#search-issues

  search repos-legacy --token=STRING --keyword=STRING
    Search repositories -
    https://developer.github.com/v3/search/legacy/#search-repositories

  search users-legacy --token=STRING --keyword=STRING
    Search users - https://developer.github.com/v3/search/legacy/#search-users

  search email-legacy --token=STRING --email=STRING
    Email search - https://developer.github.com/v3/search/legacy/#email-search

  teams list --token=STRING --org=STRING
    List teams - https://developer.github.com/v3/teams/#list-teams

  teams get --token=STRING --team_id=INT-64
    Get team - https://developer.github.com/v3/teams/#get-team

  teams create --token=STRING --org=STRING --name=STRING
    Create team - https://developer.github.com/v3/teams/#create-team

  teams edit --token=STRING --team_id=INT-64 --name=STRING
    Edit team - https://developer.github.com/v3/teams/#edit-team

  teams delete --token=STRING --team_id=INT-64
    Delete team - https://developer.github.com/v3/teams/#delete-team

  teams list-child --token=STRING --hellcat-preview --team_id=INT-64
    List child teams - https://developer.github.com/v3/teams/#list-child-teams

  teams list-repos --token=STRING --team_id=INT-64
    List team repos - https://developer.github.com/v3/teams/#list-team-repos

  teams check-manages-repo --token=STRING --team_id=INT-64 --repo=STRING
    Check if a team manages a repository -
    https://developer.github.com/v3/teams/#check-if-a-team-manages-a-repository

  teams add-or-update-repo --token=STRING --team_id=INT-64 --repo=STRING
    Add or update team repository -
    https://developer.github.com/v3/teams/#add-or-update-team-repository

  teams remove-repo --token=STRING --team_id=INT-64 --repo=STRING
    Remove team repository -
    https://developer.github.com/v3/teams/#remove-team-repository

  teams list-for-authenticated-user --token=STRING
    List user teams - https://developer.github.com/v3/teams/#list-user-teams

  teams list-projects --token=STRING --inertia-preview --team_id=INT-64
    List team projects -
    https://developer.github.com/v3/teams/#list-team-projects

  teams review-project --token=STRING --inertia-preview --team_id=INT-64 --project_id=INT-64
    Review a team project -
    https://developer.github.com/v3/teams/#review-a-team-project

  teams add-or-update-project --token=STRING --inertia-preview --team_id=INT-64 --project_id=INT-64
    Add or update team project -
    https://developer.github.com/v3/teams/#add-or-update-team-project

  teams remove-project --token=STRING --team_id=INT-64 --project_id=INT-64
    Remove team project -
    https://developer.github.com/v3/teams/#remove-team-project

  teams list-discussions --token=STRING --echo-preview --team_id=INT-64
    List discussions -
    https://developer.github.com/v3/teams/discussions/#list-discussions

  teams get-discussion --token=STRING --echo-preview --team_id=INT-64 --discussion_number=INT-64
    Get a single discussion -
    https://developer.github.com/v3/teams/discussions/#get-a-single-discussion

  teams create-discussion --token=STRING --echo-preview --team_id=INT-64 --title=STRING --body=STRING
    Create a discussion -
    https://developer.github.com/v3/teams/discussions/#create-a-discussion

  teams edit-discussion --token=STRING --echo-preview --team_id=INT-64 --discussion_number=INT-64
    Edit a discussion -
    https://developer.github.com/v3/teams/discussions/#edit-a-discussion

  teams delete-discussion --token=STRING --echo-preview --team_id=INT-64 --discussion_number=INT-64
    Delete a discussion -
    https://developer.github.com/v3/teams/discussions/#delete-a-discussion

  teams list-discussion-comments --token=STRING --echo-preview --team_id=INT-64 --discussion_number=INT-64
    List comments -
    https://developer.github.com/v3/teams/discussion_comments/#list-comments

  teams get-discussion-comment --token=STRING --echo-preview --team_id=INT-64 --discussion_number=INT-64 --comment_number=INT-64
    Get a single comment -
    https://developer.github.com/v3/teams/discussion_comments/#get-a-single-comment

  teams create-discussion-comment --token=STRING --echo-preview --team_id=INT-64 --discussion_number=INT-64 --body=STRING
    Create a comment -
    https://developer.github.com/v3/teams/discussion_comments/#create-a-comment

  teams edit-discussion-comment --token=STRING --echo-preview --team_id=INT-64 --discussion_number=INT-64 --comment_number=INT-64 --body=STRING
    Edit a comment -
    https://developer.github.com/v3/teams/discussion_comments/#edit-a-comment

  teams delete-discussion-comment --token=STRING --echo-preview --team_id=INT-64 --discussion_number=INT-64 --comment_number=INT-64
    Delete a comment -
    https://developer.github.com/v3/teams/discussion_comments/#delete-a-comment

  teams list-members --token=STRING --team_id=INT-64
    List team members -
    https://developer.github.com/v3/teams/members/#list-team-members

  teams get-member --token=STRING --team_id=INT-64 --username=STRING
    Get team member -
    https://developer.github.com/v3/teams/members/#get-team-member

  teams add-member --token=STRING --team_id=INT-64 --username=STRING
    Add team member -
    https://developer.github.com/v3/teams/members/#add-team-member

  teams remove-member --token=STRING --team_id=INT-64 --username=STRING
    Remove team member -
    https://developer.github.com/v3/teams/members/#remove-team-member

  teams get-membership --token=STRING --team_id=INT-64 --username=STRING
    Get team membership -
    https://developer.github.com/v3/teams/members/#get-team-membership

  teams add-or-update-membership --token=STRING --team_id=INT-64 --username=STRING
    Add or update team membership -
    https://developer.github.com/v3/teams/members/#add-or-update-team-membership

  teams remove-membership --token=STRING --team_id=INT-64 --username=STRING
    Remove team membership -
    https://developer.github.com/v3/teams/members/#remove-team-membership

  teams list-pending-invitations --token=STRING --team_id=INT-64
    List pending team invitations -
    https://developer.github.com/v3/teams/members/#list-pending-team-invitations

  users get-by-username --token=STRING --username=STRING
    Get a single user - https://developer.github.com/v3/users/#get-a-single-user

  users get-authenticated --token=STRING
    Get the authenticated user -
    https://developer.github.com/v3/users/#get-the-authenticated-user

  users update-authenticated --token=STRING
    Update the authenticated user -
    https://developer.github.com/v3/users/#update-the-authenticated-user

  users get-context-for-user --token=STRING --hagar-preview --username=STRING
    Get contextual information about a user -
    https://developer.github.com/v3/users/#get-contextual-information-about-a-user

  users list --token=STRING
    Get all users - https://developer.github.com/v3/users/#get-all-users

  users list-blocked --token=STRING
    List blocked users -
    https://developer.github.com/v3/users/blocking/#list-blocked-users

  users check-blocked --token=STRING --username=STRING
    Check whether you've blocked a user -
    https://developer.github.com/v3/users/blocking/#check-whether-youve-blocked-a-user

  users block --token=STRING --username=STRING
    Block a user - https://developer.github.com/v3/users/blocking/#block-a-user

  users unblock --token=STRING --username=STRING
    Unblock a user -
    https://developer.github.com/v3/users/blocking/#unblock-a-user

  users list-emails --token=STRING
    List email addresses for a user -
    https://developer.github.com/v3/users/emails/#list-email-addresses-for-a-user

  users list-public-emails --token=STRING
    List public email addresses for a user -
    https://developer.github.com/v3/users/emails/#list-public-email-addresses-for-a-user

  users add-emails --token=STRING --emails=EMAILS,...
    Add email address(es) -
    https://developer.github.com/v3/users/emails/#add-email-addresses

  users delete-emails --token=STRING --emails=EMAILS,...
    Delete email address(es) -
    https://developer.github.com/v3/users/emails/#delete-email-addresses

  users toggle-primary-email-visibility --token=STRING --email=STRING --visibility=STRING
    Toggle primary email visibility -
    https://developer.github.com/v3/users/emails/#toggle-primary-email-visibility

  users list-followers-for-user --token=STRING --username=STRING
    List a user's followers -
    https://developer.github.com/v3/users/followers/#list-followers-of-a-user

  users list-followers-for-authenticated-user --token=STRING
    List the authenticated user's followers -
    https://developer.github.com/v3/users/followers/#list-followers-of-a-user

  users list-following-for-user --token=STRING --username=STRING
    List who a user is following -
    https://developer.github.com/v3/users/followers/#list-users-followed-by-another-user

  users list-following-for-authenticated-user --token=STRING
    List who the authenticated user is following -
    https://developer.github.com/v3/users/followers/#list-users-followed-by-another-user

  users check-following --token=STRING --username=STRING
    Check if you are following a user -
    https://developer.github.com/v3/users/followers/#check-if-you-are-following-a-user

  users check-following-for-user --token=STRING --username=STRING --target_user=STRING
    Check if one user follows another -
    https://developer.github.com/v3/users/followers/#check-if-one-user-follows-another

  users follow --token=STRING --username=STRING
    Follow a user -
    https://developer.github.com/v3/users/followers/#follow-a-user

  users unfollow --token=STRING --username=STRING
    Unfollow a user -
    https://developer.github.com/v3/users/followers/#unfollow-a-user

  users list-public-keys-for-user --token=STRING --username=STRING
    List public keys for a user -
    https://developer.github.com/v3/users/keys/#list-public-keys-for-a-user

  users list-public-keys --token=STRING
    List your public keys -
    https://developer.github.com/v3/users/keys/#list-your-public-keys

  users get-public-key --token=STRING --key_id=INT-64
    Get a single public key -
    https://developer.github.com/v3/users/keys/#get-a-single-public-key

  users create-public-key --token=STRING
    Create a public key -
    https://developer.github.com/v3/users/keys/#create-a-public-key

  users delete-public-key --token=STRING --key_id=INT-64
    Delete a public key -
    https://developer.github.com/v3/users/keys/#delete-a-public-key

  users list-gpg-keys-for-user --token=STRING --username=STRING
    List GPG keys for a user -
    https://developer.github.com/v3/users/gpg_keys/#list-gpg-keys-for-a-user

  users list-gpg-keys --token=STRING
    List your GPG keys -
    https://developer.github.com/v3/users/gpg_keys/#list-your-gpg-keys

  users get-gpg-key --token=STRING --gpg_key_id=INT-64
    Get a single GPG key -
    https://developer.github.com/v3/users/gpg_keys/#get-a-single-gpg-key

  users create-gpg-key --token=STRING
    Create a GPG key -
    https://developer.github.com/v3/users/gpg_keys/#create-a-gpg-key

  users delete-gpg-key --token=STRING --gpg_key_id=INT-64
    Delete a GPG key -
    https://developer.github.com/v3/users/gpg_keys/#delete-a-gpg-key

Run "octo <command> --help" for more information on a command.

```
<!--- END HELP OUTPUT --->
