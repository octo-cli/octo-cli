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
  actions cancel-workflow-run --token=STRING --repo=STRING --run_id=INT-64
    Cancel a workflow run -
    https://developer.github.com/v3/actions/workflow-runs/#cancel-a-workflow-run

  actions create-or-update-secret-for-repo --token=STRING --name=STRING --repo=STRING
    Create or update a secret for a repository -
    https://developer.github.com/v3/actions/secrets/#create-or-update-a-secret-for-a-repository

  actions create-registration-token-for-org --token=STRING --org=STRING
    Create a registration token for an organization -
    https://developer.github.com/v3/actions/self-hosted-runners/#create-a-registration-token-for-an-organization

  actions create-registration-token-for-repo --token=STRING --repo=STRING
    Create a registration token for a repository -
    https://developer.github.com/v3/actions/self-hosted-runners/#create-a-registration-token-for-a-repository

  actions create-remove-token-for-org --token=STRING --org=STRING
    Create a remove token for an organization -
    https://developer.github.com/v3/actions/self-hosted-runners/#create-a-remove-token-for-an-organization

  actions create-remove-token-for-repo --token=STRING --repo=STRING
    Create a remove token for a repository -
    https://developer.github.com/v3/actions/self-hosted-runners/#create-a-remove-token-for-a-repository

  actions delete-artifact --token=STRING --artifact_id=INT-64 --repo=STRING
    Delete an artifact -
    https://developer.github.com/v3/actions/artifacts/#delete-an-artifact

  actions delete-secret-from-repo --token=STRING --name=STRING --repo=STRING
    Delete a secret from a repository -
    https://developer.github.com/v3/actions/secrets/#delete-a-secret-from-a-repository

  actions delete-self-hosted-runner-from-org --token=STRING --org=STRING --runner_id=INT-64
    Delete a self-hosted runner from an organization -
    https://developer.github.com/v3/actions/self-hosted-runners/#delete-a-self-hosted-runner-from-an-organization

  actions delete-self-hosted-runner-from-repo --token=STRING --repo=STRING --runner_id=INT-64
    Delete a self-hosted runner from a repository -
    https://developer.github.com/v3/actions/self-hosted-runners/#delete-a-self-hosted-runner-from-a-repository

  actions delete-workflow-run-logs --token=STRING --repo=STRING --run_id=INT-64
    Delete workflow run logs -
    https://developer.github.com/v3/actions/workflow-runs/#delete-workflow-run-logs

  actions download-artifact --token=STRING --archive_format=STRING --artifact_id=INT-64 --repo=STRING
    Download an artifact -
    https://developer.github.com/v3/actions/artifacts/#download-an-artifact

  actions get-artifact --token=STRING --artifact_id=INT-64 --repo=STRING
    Get an artifact -
    https://developer.github.com/v3/actions/artifacts/#get-an-artifact

  actions get-public-key --token=STRING --repo=STRING
    Get your public key -
    https://developer.github.com/v3/actions/secrets/#get-your-public-key

  actions get-secret --token=STRING --name=STRING --repo=STRING
    Get a secret - https://developer.github.com/v3/actions/secrets/#get-a-secret

  actions get-self-hosted-runner-for-org --token=STRING --org=STRING --runner_id=INT-64
    Get a self-hosted runner for an organization -
    https://developer.github.com/v3/actions/self-hosted-runners/#get-a-self-hosted-runner-for-an-organization

  actions get-self-hosted-runner-for-repo --token=STRING --repo=STRING --runner_id=INT-64
    Get a self-hosted runner for a repository -
    https://developer.github.com/v3/actions/self-hosted-runners/#get-a-self-hosted-runner-for-a-repository

  actions get-workflow --token=STRING --repo=STRING --workflow_id=INT-64
    Get a workflow -
    https://developer.github.com/v3/actions/workflows/#get-a-workflow

  actions get-workflow-job --token=STRING --job_id=INT-64 --repo=STRING
    Get a workflow job -
    https://developer.github.com/v3/actions/workflow-jobs/#get-a-workflow-job

  actions get-workflow-run --token=STRING --repo=STRING --run_id=INT-64
    Get a workflow run -
    https://developer.github.com/v3/actions/workflow-runs/#get-a-workflow-run

  actions list-artifacts-for-repo --token=STRING --repo=STRING
    List artifacts for a repository -
    https://developer.github.com/v3/actions/artifacts/#list-artifacts-for-a-repository

  actions list-jobs-for-workflow-run --token=STRING --repo=STRING --run_id=INT-64
    List jobs for a workflow run -
    https://developer.github.com/v3/actions/workflow-jobs/#list-jobs-for-a-workflow-run

  actions list-repo-workflow-runs --token=STRING --repo=STRING
    List repository workflow runs -
    https://developer.github.com/v3/actions/workflow-runs/#list-repository-workflow-runs

  actions list-repo-workflows --token=STRING --repo=STRING
    List repository workflows -
    https://developer.github.com/v3/actions/workflows/#list-repository-workflows

  actions list-runner-applications-for-org --token=STRING --org=STRING
    List runner applications for an organization -
    https://developer.github.com/v3/actions/self-hosted-runners/#list-runner-applications-for-an-organization

  actions list-runner-applications-for-repo --token=STRING --repo=STRING
    List runner applications for a repository -
    https://developer.github.com/v3/actions/self-hosted-runners/#list-runner-applications-for-a-repository

  actions list-secrets-for-repo --token=STRING --repo=STRING
    List secrets for a repository -
    https://developer.github.com/v3/actions/secrets/#list-secrets-for-a-repository

  actions list-self-hosted-runners-for-org --token=STRING --org=STRING
    List self-hosted runners for an organization -
    https://developer.github.com/v3/actions/self-hosted-runners/#list-self-hosted-runners-for-an-organization

  actions list-self-hosted-runners-for-repo --token=STRING --repo=STRING
    List self-hosted runners for a repository -
    https://developer.github.com/v3/actions/self-hosted-runners/#list-self-hosted-runners-for-a-repository

  actions list-workflow-job-logs --token=STRING --job_id=INT-64 --repo=STRING
    List workflow job logs -
    https://developer.github.com/v3/actions/workflow-jobs/#list-workflow-job-logs

  actions list-workflow-run-artifacts --token=STRING --repo=STRING --run_id=INT-64
    List workflow run artifacts -
    https://developer.github.com/v3/actions/artifacts/#list-workflow-run-artifacts

  actions list-workflow-run-logs --token=STRING --repo=STRING --run_id=INT-64
    List workflow run logs -
    https://developer.github.com/v3/actions/workflow-runs/#list-workflow-run-logs

  actions list-workflow-runs --token=STRING --repo=STRING --workflow_id=INT-64
    List workflow runs -
    https://developer.github.com/v3/actions/workflow-runs/#list-workflow-runs

  actions re-run-workflow --token=STRING --repo=STRING --run_id=INT-64
    Re-run a workflow -
    https://developer.github.com/v3/actions/workflow-runs/#re-run-a-workflow

  activity check-repo-is-starred-by-authenticated-user --token=STRING --repo=STRING
    Check if a repository is starred by the authenticated user -
    https://developer.github.com/v3/activity/starring/#check-if-a-repository-is-starred-by-the-authenticated-user

  activity check-watching-repo-legacy --token=STRING --repo=STRING
    Check if you are watching a repository (LEGACY) -
    https://developer.github.com/v3/activity/watching/#check-if-you-are-watching-a-repository-legacy

  activity delete-repo-subscription --token=STRING --repo=STRING
    Delete a repository subscription -
    https://developer.github.com/v3/activity/watching/#delete-a-repository-subscription

  activity delete-thread-subscription --token=STRING --thread_id=INT-64
    Delete a thread subscription -
    https://developer.github.com/v3/activity/notifications/#delete-a-thread-subscription

  activity get-feeds --token=STRING
    Get feeds - https://developer.github.com/v3/activity/feeds/#get-feeds

  activity get-repo-subscription --token=STRING --repo=STRING
    Get a repository subscription -
    https://developer.github.com/v3/activity/watching/#get-a-repository-subscription

  activity get-thread --token=STRING --thread_id=INT-64
    Get a thread -
    https://developer.github.com/v3/activity/notifications/#get-a-thread

  activity get-thread-subscription-for-authenticated-user --token=STRING --thread_id=INT-64
    Get a thread subscription for the authenticated user -
    https://developer.github.com/v3/activity/notifications/#get-a-thread-subscription-for-the-authenticated-user

  activity list-events-for-authenticated-user --token=STRING --username=STRING
    List events for the authenticated user -
    https://developer.github.com/v3/activity/events/#list-events-for-the-authenticated-user

  activity list-notifications-for-authenticated-user --token=STRING
    List notifications for the authenticated user -
    https://developer.github.com/v3/activity/notifications/#list-notifications-for-the-authenticated-user

  activity list-org-events-for-authenticated-user --token=STRING --org=STRING --username=STRING
    List organization events for the authenticated user -
    https://developer.github.com/v3/activity/events/#list-organization-events-for-the-authenticated-user

  activity list-public-events --token=STRING
    List public events -
    https://developer.github.com/v3/activity/events/#list-public-events

  activity list-public-events-for-repo-network --token=STRING --repo=STRING
    List public events for a network of repositories -
    https://developer.github.com/v3/activity/events/#list-public-events-for-a-network-of-repositories

  activity list-public-events-for-user --token=STRING --username=STRING
    List public events for a user -
    https://developer.github.com/v3/activity/events/#list-public-events-for-a-user

  activity list-public-org-events --token=STRING --org=STRING
    List public organization events -
    https://developer.github.com/v3/activity/events/#list-public-organization-events

  activity list-received-events-for-user --token=STRING --username=STRING
    List events received by the authenticated user -
    https://developer.github.com/v3/activity/events/#list-events-received-by-the-authenticated-user

  activity list-received-public-events-for-user --token=STRING --username=STRING
    List public events received by a user -
    https://developer.github.com/v3/activity/events/#list-public-events-received-by-a-user

  activity list-repo-events --token=STRING --repo=STRING
    List repository events -
    https://developer.github.com/v3/activity/events/#list-repository-events

  activity list-repo-notifications-for-authenticated-user --token=STRING --repo=STRING
    List repository notifications for the authenticated user -
    https://developer.github.com/v3/activity/notifications/#list-repository-notifications-for-the-authenticated-user

  activity list-repos-starred-by-authenticated-user --token=STRING
    List repositories starred by the authenticated user -
    https://developer.github.com/v3/activity/starring/#list-repositories-starred-by-the-authenticated-user

  activity list-repos-starred-by-user --token=STRING --username=STRING
    List repositories starred by a user -
    https://developer.github.com/v3/activity/starring/#list-repositories-starred-by-a-user

  activity list-repos-watched-by-user --token=STRING --username=STRING
    List repositories watched by a user -
    https://developer.github.com/v3/activity/watching/#list-repositories-watched-by-a-user

  activity list-stargazers-for-repo --token=STRING --repo=STRING
    List stargazers -
    https://developer.github.com/v3/activity/starring/#list-stargazers

  activity list-watched-repos-for-authenticated-user --token=STRING
    List repositories watched by the authenticated user -
    https://developer.github.com/v3/activity/watching/#list-repositories-watched-by-the-authenticated-user

  activity list-watchers-for-repo --token=STRING --repo=STRING
    List watchers -
    https://developer.github.com/v3/activity/watching/#list-watchers

  activity mark-notifications-as-read --token=STRING
    Mark notifications as read -
    https://developer.github.com/v3/activity/notifications/#mark-notifications-as-read

  activity mark-repo-notifications-as-read --token=STRING --repo=STRING
    Mark repository notifications as read -
    https://developer.github.com/v3/activity/notifications/#mark-repository-notifications-as-read

  activity mark-thread-as-read --token=STRING --thread_id=INT-64
    Mark a thread as read -
    https://developer.github.com/v3/activity/notifications/#mark-a-thread-as-read

  activity set-repo-subscription --token=STRING --repo=STRING
    Set a repository subscription -
    https://developer.github.com/v3/activity/watching/#set-a-repository-subscription

  activity set-thread-subscription --token=STRING --thread_id=INT-64
    Set a thread subscription -
    https://developer.github.com/v3/activity/notifications/#set-a-thread-subscription

  activity star-repo-for-authenticated-user --token=STRING --repo=STRING
    Star a repository for the authenticated user -
    https://developer.github.com/v3/activity/starring/#star-a-repository-for-the-authenticated-user

  activity stop-watching-repo-legacy --token=STRING --repo=STRING
    Stop watching a repository (LEGACY) -
    https://developer.github.com/v3/activity/watching/#stop-watching-a-repository-legacy

  activity unstar-repo-for-authenticated-user --token=STRING --repo=STRING
    Unstar a repository for the authenticated user -
    https://developer.github.com/v3/activity/starring/#unstar-a-repository-for-the-authenticated-user

  activity watch-repo-legacy --token=STRING --repo=STRING
    Watch a repository (LEGACY) -
    https://developer.github.com/v3/activity/watching/#watch-a-repository-legacy

  apps add-repo-to-installation --token=STRING --installation_id=INT-64 --machine-man-preview --repository_id=INT-64
    Add repository to installation -
    https://developer.github.com/v3/apps/installations/#add-repository-to-installation

  apps check-authorization --token=STRING --access_token=STRING --client_id=STRING
    Check an authorization -
    https://developer.github.com/v3/apps/oauth_applications/#check-an-authorization

  apps check-token --token=STRING --client_id=STRING
    Check a token -
    https://developer.github.com/v3/apps/oauth_applications/#check-a-token

  apps create-content-attachment --token=STRING --body=STRING --content_reference_id=INT-64 --corsair-preview --title=STRING
    Create a content attachment -
    https://developer.github.com/v3/apps/installations/#create-a-content-attachment

  apps create-from-manifest --token=STRING --code=STRING
    Create a GitHub App from a manifest -
    https://developer.github.com/v3/apps/#create-a-github-app-from-a-manifest

  apps delete-authorization --token=STRING --client_id=STRING
    Delete an app authorization -
    https://developer.github.com/v3/apps/oauth_applications/#delete-an-app-authorization

  apps delete-installation --token=STRING --installation_id=INT-64 --machine-man-preview
    Delete an installation -
    https://developer.github.com/v3/apps/#delete-an-installation

  apps delete-token --token=STRING --client_id=STRING
    Delete an app token -
    https://developer.github.com/v3/apps/oauth_applications/#delete-an-app-token

  apps get-authenticated --token=STRING --machine-man-preview
    Get the authenticated GitHub App -
    https://developer.github.com/v3/apps/#get-the-authenticated-github-app

  apps get-by-slug --token=STRING --app_slug=STRING --machine-man-preview
    Get a single GitHub App -
    https://developer.github.com/v3/apps/#get-a-single-github-app

  apps get-installation --token=STRING --installation_id=INT-64 --machine-man-preview
    Get an installation -
    https://developer.github.com/v3/apps/#get-an-installation

  apps get-org-installation --token=STRING --machine-man-preview --org=STRING
    Get an organization installation -
    https://developer.github.com/v3/apps/#get-an-organization-installation

  apps get-repo-installation --token=STRING --machine-man-preview --repo=STRING
    Get a repository installation -
    https://developer.github.com/v3/apps/#get-a-repository-installation

  apps get-subscription-plan-for-account --token=STRING --account_id=INT-64
    Get a subscription plan for an account -
    https://developer.github.com/v3/apps/marketplace/#get-a-subscription-plan-for-an-account

  apps get-subscription-plan-for-account-stubbed --token=STRING --account_id=INT-64
    Get a subscription plan for an account (stubbed) -
    https://developer.github.com/v3/apps/marketplace/#get-a-subscription-plan-for-an-account-stubbed

  apps get-user-installation --token=STRING --machine-man-preview --username=STRING
    Get a user installation -
    https://developer.github.com/v3/apps/#get-a-user-installation

  apps list-accounts-for-plan --token=STRING --plan_id=INT-64
    List accounts for a plan -
    https://developer.github.com/v3/apps/marketplace/#list-accounts-for-a-plan

  apps list-accounts-for-plan-stubbed --token=STRING --plan_id=INT-64
    List accounts for a plan (stubbed) -
    https://developer.github.com/v3/apps/marketplace/#list-accounts-for-a-plan-stubbed

  apps list-installation-repos-for-authenticated-user --token=STRING --installation_id=INT-64 --machine-man-preview
    List repositories accessible to the user for an installation -
    https://developer.github.com/v3/apps/installations/#list-repositories-accessible-to-the-user-for-an-installation

  apps list-installations --token=STRING --machine-man-preview
    List installations -
    https://developer.github.com/v3/apps/#list-installations

  apps list-installations-for-authenticated-user --token=STRING --machine-man-preview
    List installations for a user -
    https://developer.github.com/v3/apps/installations/#list-installations-for-a-user

  apps list-plans --token=STRING
    List plans - https://developer.github.com/v3/apps/marketplace/#list-plans

  apps list-plans-stubbed --token=STRING
    List plans (stubbed) -
    https://developer.github.com/v3/apps/marketplace/#list-plans-stubbed

  apps list-repos --token=STRING --machine-man-preview
    List repositories -
    https://developer.github.com/v3/apps/installations/#list-repositories

  apps list-subscriptions-for-authenticated-user --token=STRING
    List subscriptions for the authenticated user -
    https://developer.github.com/v3/apps/marketplace/#list-subscriptions-for-the-authenticated-user

  apps list-subscriptions-for-authenticated-user-stubbed --token=STRING
    List subscriptions for the authenticated user (stubbed) -
    https://developer.github.com/v3/apps/marketplace/#list-subscriptions-for-the-authenticated-user-stubbed

  apps remove-repo-from-installation --token=STRING --installation_id=INT-64 --machine-man-preview --repository_id=INT-64
    Remove repository from installation -
    https://developer.github.com/v3/apps/installations/#remove-repository-from-installation

  apps reset-authorization --token=STRING --access_token=STRING --client_id=STRING
    Reset an authorization -
    https://developer.github.com/v3/apps/oauth_applications/#reset-an-authorization

  apps reset-token --token=STRING --client_id=STRING
    Reset a token -
    https://developer.github.com/v3/apps/oauth_applications/#reset-a-token

  apps revoke-authorization-for-application --token=STRING --access_token=STRING --client_id=STRING
    Revoke an authorization for an application -
    https://developer.github.com/v3/apps/oauth_applications/#revoke-an-authorization-for-an-application

  apps revoke-grant-for-application --token=STRING --access_token=STRING --client_id=STRING
    Revoke a grant for an application -
    https://developer.github.com/v3/apps/oauth_applications/#revoke-a-grant-for-an-application

  apps revoke-installation-token --token=STRING
    Revoke an installation token -
    https://developer.github.com/v3/apps/installations/#revoke-an-installation-token

  checks create-suite --token=STRING --antiope-preview --head_sha=STRING --repo=STRING
    Create a check suite -
    https://developer.github.com/v3/checks/suites/#create-a-check-suite

  checks get --token=STRING --antiope-preview --check_run_id=INT-64 --repo=STRING
    Get a check run -
    https://developer.github.com/v3/checks/runs/#get-a-check-run

  checks get-suite --token=STRING --antiope-preview --check_suite_id=INT-64 --repo=STRING
    Get a check suite -
    https://developer.github.com/v3/checks/suites/#get-a-check-suite

  checks list-annotations --token=STRING --antiope-preview --check_run_id=INT-64 --repo=STRING
    List check run annotations -
    https://developer.github.com/v3/checks/runs/#list-check-run-annotations

  checks list-for-ref --token=STRING --antiope-preview --ref=STRING --repo=STRING
    List check runs for a Git reference -
    https://developer.github.com/v3/checks/runs/#list-check-runs-for-a-git-reference

  checks list-for-suite --token=STRING --antiope-preview --check_suite_id=INT-64 --repo=STRING
    List check runs in a check suite -
    https://developer.github.com/v3/checks/runs/#list-check-runs-in-a-check-suite

  checks list-suites-for-ref --token=STRING --antiope-preview --ref=STRING --repo=STRING
    List check suites for a Git reference -
    https://developer.github.com/v3/checks/suites/#list-check-suites-for-a-git-reference

  checks rerequest-suite --token=STRING --antiope-preview --check_suite_id=INT-64 --repo=STRING
    Rerequest a check suite -
    https://developer.github.com/v3/checks/suites/#rerequest-a-check-suite

  codes-of-conduct get-all-codes-of-conduct --token=STRING --scarlet-witch-preview
    List all codes of conduct -
    https://developer.github.com/v3/codes_of_conduct/#list-all-codes-of-conduct

  codes-of-conduct get-conduct-code --token=STRING --key=STRING --scarlet-witch-preview
    Get an individual code of conduct -
    https://developer.github.com/v3/codes_of_conduct/#get-an-individual-code-of-conduct

  codes-of-conduct get-for-repo --token=STRING --repo=STRING --scarlet-witch-preview
    Get the contents of a repository's code of conduct -
    https://developer.github.com/v3/codes_of_conduct/#get-the-contents-of-a-repositorys-code-of-conduct

  emojis get --token=STRING
    Get - https://developer.github.com/v3/emojis/#emojis

  gists check-is-starred --token=STRING --gist_id=STRING
    Check if a gist is starred -
    https://developer.github.com/v3/gists/#check-if-a-gist-is-starred

  gists create-comment --token=STRING --body=STRING --gist_id=STRING
    Create a comment -
    https://developer.github.com/v3/gists/comments/#create-a-comment

  gists delete --token=STRING --gist_id=STRING
    Delete a gist - https://developer.github.com/v3/gists/#delete-a-gist

  gists delete-comment --token=STRING --comment_id=INT-64 --gist_id=STRING
    Delete a comment -
    https://developer.github.com/v3/gists/comments/#delete-a-comment

  gists fork --token=STRING --gist_id=STRING
    Fork a gist - https://developer.github.com/v3/gists/#fork-a-gist

  gists get --token=STRING --gist_id=STRING
    Get a gist - https://developer.github.com/v3/gists/#get-a-gist

  gists get-comment --token=STRING --comment_id=INT-64 --gist_id=STRING
    Get a single comment -
    https://developer.github.com/v3/gists/comments/#get-a-single-comment

  gists get-revision --token=STRING --gist_id=STRING --sha=STRING
    Get a specific revision of a gist -
    https://developer.github.com/v3/gists/#get-a-specific-revision-of-a-gist

  gists list --token=STRING
    List gists for the authenticated user -
    https://developer.github.com/v3/gists/#list-gists-for-the-authenticated-user

  gists list-comments --token=STRING --gist_id=STRING
    List comments on a gist -
    https://developer.github.com/v3/gists/comments/#list-comments-on-a-gist

  gists list-commits --token=STRING --gist_id=STRING
    List gist commits - https://developer.github.com/v3/gists/#list-gist-commits

  gists list-for-user --token=STRING --username=STRING
    List gists for a user -
    https://developer.github.com/v3/gists/#list-gists-for-a-user

  gists list-forks --token=STRING --gist_id=STRING
    List gist forks - https://developer.github.com/v3/gists/#list-gist-forks

  gists list-public --token=STRING
    List public gists - https://developer.github.com/v3/gists/#list-public-gists

  gists list-starred --token=STRING
    List starred gists -
    https://developer.github.com/v3/gists/#list-starred-gists

  gists star --token=STRING --gist_id=STRING
    Star a gist - https://developer.github.com/v3/gists/#star-a-gist

  gists unstar --token=STRING --gist_id=STRING
    Unstar a gist - https://developer.github.com/v3/gists/#unstar-a-gist

  gists update-comment --token=STRING --body=STRING --comment_id=INT-64 --gist_id=STRING
    Edit a comment -
    https://developer.github.com/v3/gists/comments/#edit-a-comment

  git create-blob --token=STRING --content=STRING --repo=STRING
    Create a blob - https://developer.github.com/v3/git/blobs/#create-a-blob

  git create-ref --token=STRING --ref=STRING --repo=STRING --sha=STRING
    Create a reference -
    https://developer.github.com/v3/git/refs/#create-a-reference

  git delete-ref --token=STRING --ref=STRING --repo=STRING
    Delete a reference -
    https://developer.github.com/v3/git/refs/#delete-a-reference

  git get-blob --token=STRING --file_sha=STRING --repo=STRING
    Get a blob - https://developer.github.com/v3/git/blobs/#get-a-blob

  git get-commit --token=STRING --commit_sha=STRING --repo=STRING
    Get a commit - https://developer.github.com/v3/git/commits/#get-a-commit

  git get-ref --token=STRING --ref=STRING --repo=STRING
    Get a single reference -
    https://developer.github.com/v3/git/refs/#get-a-single-reference

  git get-tag --token=STRING --repo=STRING --tag_sha=STRING
    Get a tag - https://developer.github.com/v3/git/tags/#get-a-tag

  git get-tree --token=STRING --repo=STRING --tree_sha=STRING
    Get a tree - https://developer.github.com/v3/git/trees/#get-a-tree

  git list-matching-refs --token=STRING --ref=STRING --repo=STRING
    List matching references -
    https://developer.github.com/v3/git/refs/#list-matching-references

  git update-ref --token=STRING --ref=STRING --repo=STRING --sha=STRING
    Update a reference -
    https://developer.github.com/v3/git/refs/#update-a-reference

  gitignore get-template --token=STRING --name=STRING
    Get a single template -
    https://developer.github.com/v3/gitignore/#get-a-single-template

  gitignore list-templates --token=STRING
    Listing available templates -
    https://developer.github.com/v3/gitignore/#listing-available-templates

  interactions add-or-update-restrictions-for-org --token=STRING --limit=STRING --org=STRING --sombra-preview
    Add or update interaction restrictions for an organization -
    https://developer.github.com/v3/interactions/orgs/#add-or-update-interaction-restrictions-for-an-organization

  interactions add-or-update-restrictions-for-repo --token=STRING --limit=STRING --repo=STRING --sombra-preview
    Add or update interaction restrictions for a repository -
    https://developer.github.com/v3/interactions/repos/#add-or-update-interaction-restrictions-for-a-repository

  interactions get-restrictions-for-org --token=STRING --org=STRING --sombra-preview
    Get interaction restrictions for an organization -
    https://developer.github.com/v3/interactions/orgs/#get-interaction-restrictions-for-an-organization

  interactions get-restrictions-for-repo --token=STRING --repo=STRING --sombra-preview
    Get interaction restrictions for a repository -
    https://developer.github.com/v3/interactions/repos/#get-interaction-restrictions-for-a-repository

  interactions remove-restrictions-for-org --token=STRING --org=STRING --sombra-preview
    Remove interaction restrictions for an organization -
    https://developer.github.com/v3/interactions/orgs/#remove-interaction-restrictions-for-an-organization

  interactions remove-restrictions-for-repo --token=STRING --repo=STRING --sombra-preview
    Remove interaction restrictions for a repository -
    https://developer.github.com/v3/interactions/repos/#remove-interaction-restrictions-for-a-repository

  issues add-assignees --token=STRING --issue_number=INT-64 --repo=STRING
    Add assignees to an issue -
    https://developer.github.com/v3/issues/assignees/#add-assignees-to-an-issue

  issues add-labels --token=STRING --issue_number=INT-64 --labels=LABELS,... --repo=STRING
    Add labels to an issue -
    https://developer.github.com/v3/issues/labels/#add-labels-to-an-issue

  issues check-assignee --token=STRING --assignee=STRING --repo=STRING
    Check assignee -
    https://developer.github.com/v3/issues/assignees/#check-assignee

  issues create --token=STRING --repo=STRING --title=STRING
    Create an issue - https://developer.github.com/v3/issues/#create-an-issue

  issues create-comment --token=STRING --body=STRING --issue_number=INT-64 --repo=STRING
    Create a comment -
    https://developer.github.com/v3/issues/comments/#create-a-comment

  issues create-label --token=STRING --color=STRING --name=STRING --repo=STRING
    Create a label -
    https://developer.github.com/v3/issues/labels/#create-a-label

  issues create-milestone --token=STRING --repo=STRING --title=STRING
    Create a milestone -
    https://developer.github.com/v3/issues/milestones/#create-a-milestone

  issues delete-comment --token=STRING --comment_id=INT-64 --repo=STRING
    Delete a comment -
    https://developer.github.com/v3/issues/comments/#delete-a-comment

  issues delete-label --token=STRING --name=STRING --repo=STRING
    Delete a label -
    https://developer.github.com/v3/issues/labels/#delete-a-label

  issues delete-milestone --token=STRING --milestone_number=INT-64 --repo=STRING
    Delete a milestone -
    https://developer.github.com/v3/issues/milestones/#delete-a-milestone

  issues get --token=STRING --issue_number=INT-64 --repo=STRING
    Get an issue - https://developer.github.com/v3/issues/#get-an-issue

  issues get-comment --token=STRING --comment_id=INT-64 --repo=STRING
    Get a single comment -
    https://developer.github.com/v3/issues/comments/#get-a-single-comment

  issues get-event --token=STRING --event_id=INT-64 --repo=STRING
    Get a single event -
    https://developer.github.com/v3/issues/events/#get-a-single-event

  issues get-label --token=STRING --name=STRING --repo=STRING
    Get a single label -
    https://developer.github.com/v3/issues/labels/#get-a-single-label

  issues get-milestone --token=STRING --milestone_number=INT-64 --repo=STRING
    Get a single milestone -
    https://developer.github.com/v3/issues/milestones/#get-a-single-milestone

  issues list --token=STRING
    List issues assigned to the authenticated user -
    https://developer.github.com/v3/issues/#list-issues-assigned-to-the-authenticated-user

  issues list-assignees --token=STRING --repo=STRING
    List assignees -
    https://developer.github.com/v3/issues/assignees/#list-assignees

  issues list-comments --token=STRING --issue_number=INT-64 --repo=STRING
    List comments on an issue -
    https://developer.github.com/v3/issues/comments/#list-comments-on-an-issue

  issues list-comments-for-repo --token=STRING --repo=STRING
    List comments in a repository -
    https://developer.github.com/v3/issues/comments/#list-comments-in-a-repository

  issues list-events --token=STRING --issue_number=INT-64 --repo=STRING
    List events for an issue -
    https://developer.github.com/v3/issues/events/#list-events-for-an-issue

  issues list-events-for-repo --token=STRING --repo=STRING
    List events for a repository -
    https://developer.github.com/v3/issues/events/#list-events-for-a-repository

  issues list-events-for-timeline --token=STRING --issue_number=INT-64 --mockingbird-preview --repo=STRING
    List events for an issue -
    https://developer.github.com/v3/issues/timeline/#list-events-for-an-issue

  issues list-for-authenticated-user --token=STRING
    List user account issues assigned to the authenticated user -
    https://developer.github.com/v3/issues/#list-user-account-issues-assigned-to-the-authenticated-user

  issues list-for-org --token=STRING --org=STRING
    List organization issues assigned to the authenticated user -
    https://developer.github.com/v3/issues/#list-organization-issues-assigned-to-the-authenticated-user

  issues list-for-repo --token=STRING --repo=STRING
    List repository issues -
    https://developer.github.com/v3/issues/#list-repository-issues

  issues list-labels-for-milestone --token=STRING --milestone_number=INT-64 --repo=STRING
    Get labels for every issue in a milestone -
    https://developer.github.com/v3/issues/labels/#get-labels-for-every-issue-in-a-milestone

  issues list-labels-for-repo --token=STRING --repo=STRING
    List all labels for this repository -
    https://developer.github.com/v3/issues/labels/#list-all-labels-for-this-repository

  issues list-labels-on-issue --token=STRING --issue_number=INT-64 --repo=STRING
    List labels on an issue -
    https://developer.github.com/v3/issues/labels/#list-labels-on-an-issue

  issues list-milestones-for-repo --token=STRING --repo=STRING
    List milestones for a repository -
    https://developer.github.com/v3/issues/milestones/#list-milestones-for-a-repository

  issues lock --token=STRING --issue_number=INT-64 --repo=STRING
    Lock an issue - https://developer.github.com/v3/issues/#lock-an-issue

  issues remove-all-labels --token=STRING --issue_number=INT-64 --repo=STRING
    Remove all labels from an issue -
    https://developer.github.com/v3/issues/labels/#remove-all-labels-from-an-issue

  issues remove-assignees --token=STRING --issue_number=INT-64 --repo=STRING
    Remove assignees from an issue -
    https://developer.github.com/v3/issues/assignees/#remove-assignees-from-an-issue

  issues remove-label --token=STRING --issue_number=INT-64 --name=STRING --repo=STRING
    Remove a label from an issue -
    https://developer.github.com/v3/issues/labels/#remove-a-label-from-an-issue

  issues replace-all-labels --token=STRING --issue_number=INT-64 --repo=STRING
    Replace all labels for an issue -
    https://developer.github.com/v3/issues/labels/#replace-all-labels-for-an-issue

  issues unlock --token=STRING --issue_number=INT-64 --repo=STRING
    Unlock an issue - https://developer.github.com/v3/issues/#unlock-an-issue

  issues update --token=STRING --issue_number=INT-64 --repo=STRING
    Update an issue - https://developer.github.com/v3/issues/#update-an-issue

  issues update-comment --token=STRING --body=STRING --comment_id=INT-64 --repo=STRING
    Edit a comment -
    https://developer.github.com/v3/issues/comments/#edit-a-comment

  issues update-label --token=STRING --name=STRING --repo=STRING
    Update a label -
    https://developer.github.com/v3/issues/labels/#update-a-label

  issues update-milestone --token=STRING --milestone_number=INT-64 --repo=STRING
    Update a milestone -
    https://developer.github.com/v3/issues/milestones/#update-a-milestone

  licenses get --token=STRING --license=STRING
    Get an individual license -
    https://developer.github.com/v3/licenses/#get-an-individual-license

  licenses get-for-repo --token=STRING --repo=STRING
    Get the contents of a repository's license -
    https://developer.github.com/v3/licenses/#get-the-contents-of-a-repositorys-license

  licenses list-commonly-used --token=STRING
    List commonly used licenses -
    https://developer.github.com/v3/licenses/#list-commonly-used-licenses

  markdown render --token=STRING --text=STRING
    Render an arbitrary Markdown document -
    https://developer.github.com/v3/markdown/#render-an-arbitrary-markdown-document

  markdown render-raw --token=STRING
    Render a Markdown document in raw mode -
    https://developer.github.com/v3/markdown/#render-a-markdown-document-in-raw-mode

  meta get --token=STRING
    Get - https://developer.github.com/v3/meta/#meta

  migrations cancel-import --token=STRING --repo=STRING
    Cancel an import -
    https://developer.github.com/v3/migrations/source_imports/#cancel-an-import

  migrations delete-archive-for-authenticated-user --token=STRING --migration_id=INT-64 --wyandotte-preview
    Delete a user migration archive -
    https://developer.github.com/v3/migrations/users/#delete-a-user-migration-archive

  migrations delete-archive-for-org --token=STRING --migration_id=INT-64 --org=STRING --wyandotte-preview
    Delete an organization migration archive -
    https://developer.github.com/v3/migrations/orgs/#delete-an-organization-migration-archive

  migrations download-archive-for-org --token=STRING --migration_id=INT-64 --org=STRING --wyandotte-preview
    Download an organization migration archive -
    https://developer.github.com/v3/migrations/orgs/#download-an-organization-migration-archive

  migrations get-archive-for-authenticated-user --token=STRING --migration_id=INT-64 --wyandotte-preview
    Download a user migration archive -
    https://developer.github.com/v3/migrations/users/#download-a-user-migration-archive

  migrations get-commit-authors --token=STRING --repo=STRING
    Get commit authors -
    https://developer.github.com/v3/migrations/source_imports/#get-commit-authors

  migrations get-import-progress --token=STRING --repo=STRING
    Get import progress -
    https://developer.github.com/v3/migrations/source_imports/#get-import-progress

  migrations get-large-files --token=STRING --repo=STRING
    Get large files -
    https://developer.github.com/v3/migrations/source_imports/#get-large-files

  migrations get-status-for-authenticated-user --token=STRING --migration_id=INT-64 --wyandotte-preview
    Get the status of a user migration -
    https://developer.github.com/v3/migrations/users/#get-the-status-of-a-user-migration

  migrations get-status-for-org --token=STRING --migration_id=INT-64 --org=STRING --wyandotte-preview
    Get the status of an organization migration -
    https://developer.github.com/v3/migrations/orgs/#get-the-status-of-an-organization-migration

  migrations list-for-authenticated-user --token=STRING --wyandotte-preview
    List user migrations -
    https://developer.github.com/v3/migrations/users/#list-user-migrations

  migrations list-for-org --token=STRING --org=STRING --wyandotte-preview
    List organization migrations -
    https://developer.github.com/v3/migrations/orgs/#list-organization-migrations

  migrations list-repos-for-org --token=STRING --migration_id=INT-64 --org=STRING --wyandotte-preview
    List repositories in an organization migration -
    https://developer.github.com/v3/migrations/orgs/#list-repositories-in-an-organization-migration

  migrations list-repos-for-user --token=STRING --migration_id=INT-64 --wyandotte-preview
    List repositories for a user migration -
    https://developer.github.com/v3/migrations/users/#list-repositories-for-a-user-migration

  migrations map-commit-author --token=STRING --author_id=INT-64 --repo=STRING
    Map a commit author -
    https://developer.github.com/v3/migrations/source_imports/#map-a-commit-author

  migrations set-lfs-preference --token=STRING --repo=STRING --use_lfs=STRING
    Set Git LFS preference -
    https://developer.github.com/v3/migrations/source_imports/#set-git-lfs-preference

  migrations start-for-authenticated-user --token=STRING --repositories=REPOSITORIES,...
    Start a user migration -
    https://developer.github.com/v3/migrations/users/#start-a-user-migration

  migrations start-for-org --token=STRING --org=STRING --repositories=REPOSITORIES,...
    Start an organization migration -
    https://developer.github.com/v3/migrations/orgs/#start-an-organization-migration

  migrations start-import --token=STRING --repo=STRING --vcs_url=STRING
    Start an import -
    https://developer.github.com/v3/migrations/source_imports/#start-an-import

  migrations unlock-repo-for-authenticated-user --token=STRING --migration_id=INT-64 --repo_name=STRING --wyandotte-preview
    Unlock a user repository -
    https://developer.github.com/v3/migrations/users/#unlock-a-user-repository

  migrations unlock-repo-for-org --token=STRING --migration_id=INT-64 --org=STRING --repo_name=STRING --wyandotte-preview
    Unlock an organization repository -
    https://developer.github.com/v3/migrations/orgs/#unlock-an-organization-repository

  migrations update-import --token=STRING --repo=STRING
    Update existing import -
    https://developer.github.com/v3/migrations/source_imports/#update-existing-import

  oauth-authorizations create-authorization --token=STRING --note=STRING
    Create a new authorization -
    https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization

  oauth-authorizations delete-authorization --token=STRING --authorization_id=INT-64
    Delete an authorization -
    https://developer.github.com/v3/oauth_authorizations/#delete-an-authorization

  oauth-authorizations delete-grant --token=STRING --grant_id=INT-64
    Delete a grant -
    https://developer.github.com/v3/oauth_authorizations/#delete-a-grant

  oauth-authorizations get-authorization --token=STRING --authorization_id=INT-64
    Get a single authorization -
    https://developer.github.com/v3/oauth_authorizations/#get-a-single-authorization

  oauth-authorizations get-grant --token=STRING --grant_id=INT-64
    Get a single grant -
    https://developer.github.com/v3/oauth_authorizations/#get-a-single-grant

  oauth-authorizations get-or-create-authorization-for-app --token=STRING --client_id=STRING --client_secret=STRING
    Get-or-create an authorization for a specific app -
    https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app

  oauth-authorizations get-or-create-authorization-for-app-and-fingerprint --token=STRING --client_id=STRING --client_secret=STRING --fingerprint=STRING
    Get-or-create an authorization for a specific app and fingerprint -
    https://developer.github.com/v3/oauth_authorizations/#get-or-create-an-authorization-for-a-specific-app-and-fingerprint

  oauth-authorizations list-authorizations --token=STRING
    List your authorizations -
    https://developer.github.com/v3/oauth_authorizations/#list-your-authorizations

  oauth-authorizations list-grants --token=STRING
    List your grants -
    https://developer.github.com/v3/oauth_authorizations/#list-your-grants

  oauth-authorizations update-authorization --token=STRING --authorization_id=INT-64
    Update an existing authorization -
    https://developer.github.com/v3/oauth_authorizations/#update-an-existing-authorization

  orgs add-or-update-membership --token=STRING --org=STRING --username=STRING
    Add or update organization membership -
    https://developer.github.com/v3/orgs/members/#add-or-update-organization-membership

  orgs block-user --token=STRING --org=STRING --username=STRING
    Block a user - https://developer.github.com/v3/orgs/blocking/#block-a-user

  orgs check-blocked-user --token=STRING --org=STRING --username=STRING
    Check whether a user is blocked from an organization -
    https://developer.github.com/v3/orgs/blocking/#check-whether-a-user-is-blocked-from-an-organization

  orgs check-membership --token=STRING --org=STRING --username=STRING
    Check membership -
    https://developer.github.com/v3/orgs/members/#check-membership

  orgs check-public-membership --token=STRING --org=STRING --username=STRING
    Check public membership -
    https://developer.github.com/v3/orgs/members/#check-public-membership

  orgs conceal-membership --token=STRING --org=STRING --username=STRING
    Conceal a user's membership -
    https://developer.github.com/v3/orgs/members/#conceal-a-users-membership

  orgs convert-member-to-outside-collaborator --token=STRING --org=STRING --username=STRING
    Convert member to outside collaborator -
    https://developer.github.com/v3/orgs/outside_collaborators/#convert-member-to-outside-collaborator

  orgs create-invitation --token=STRING --org=STRING
    Create organization invitation -
    https://developer.github.com/v3/orgs/members/#create-organization-invitation

  orgs delete-hook --token=STRING --hook_id=INT-64 --org=STRING
    Delete a hook - https://developer.github.com/v3/orgs/hooks/#delete-a-hook

  orgs get --token=STRING --org=STRING
    Get an organization -
    https://developer.github.com/v3/orgs/#get-an-organization

  orgs get-hook --token=STRING --hook_id=INT-64 --org=STRING
    Get single hook -
    https://developer.github.com/v3/orgs/hooks/#get-single-hook

  orgs get-membership-for-authenticated-user --token=STRING --org=STRING
    Get your organization membership -
    https://developer.github.com/v3/orgs/members/#get-your-organization-membership

  orgs get-membership-for-user --token=STRING --org=STRING --username=STRING
    Get organization membership -
    https://developer.github.com/v3/orgs/members/#get-organization-membership

  orgs list --token=STRING
    List all organizations -
    https://developer.github.com/v3/orgs/#list-all-organizations

  orgs list-blocked-users --token=STRING --org=STRING
    List blocked users -
    https://developer.github.com/v3/orgs/blocking/#list-blocked-users

  orgs list-credential-authorizations --token=STRING --org=STRING
    List credential authorizations for an organization -
    https://developer.github.com/v3/orgs/#list-credential-authorizations-for-an-organization

  orgs list-for-authenticated-user --token=STRING
    List your organizations -
    https://developer.github.com/v3/orgs/#list-your-organizations

  orgs list-for-user --token=STRING --username=STRING
    List user organizations -
    https://developer.github.com/v3/orgs/#list-user-organizations

  orgs list-hooks --token=STRING --org=STRING
    List hooks - https://developer.github.com/v3/orgs/hooks/#list-hooks

  orgs list-installations --token=STRING --machine-man-preview --org=STRING
    List installations for an organization -
    https://developer.github.com/v3/orgs/#list-installations-for-an-organization

  orgs list-invitation-teams --token=STRING --invitation_id=INT-64 --org=STRING
    List organization invitation teams -
    https://developer.github.com/v3/orgs/members/#list-organization-invitation-teams

  orgs list-members --token=STRING --org=STRING
    Members list - https://developer.github.com/v3/orgs/members/#members-list

  orgs list-memberships --token=STRING
    List your organization memberships -
    https://developer.github.com/v3/orgs/members/#list-your-organization-memberships

  orgs list-outside-collaborators --token=STRING --org=STRING
    List outside collaborators -
    https://developer.github.com/v3/orgs/outside_collaborators/#list-outside-collaborators

  orgs list-pending-invitations --token=STRING --org=STRING
    List pending organization invitations -
    https://developer.github.com/v3/orgs/members/#list-pending-organization-invitations

  orgs list-public-members --token=STRING --org=STRING
    Public members list -
    https://developer.github.com/v3/orgs/members/#public-members-list

  orgs ping-hook --token=STRING --hook_id=INT-64 --org=STRING
    Ping a hook - https://developer.github.com/v3/orgs/hooks/#ping-a-hook

  orgs publicize-membership --token=STRING --org=STRING --username=STRING
    Publicize a user's membership -
    https://developer.github.com/v3/orgs/members/#publicize-a-users-membership

  orgs remove-credential-authorization --token=STRING --credential_id=INT-64 --org=STRING
    Remove a credential authorization for an organization -
    https://developer.github.com/v3/orgs/#remove-a-credential-authorization-for-an-organization

  orgs remove-member --token=STRING --org=STRING --username=STRING
    Remove a member -
    https://developer.github.com/v3/orgs/members/#remove-a-member

  orgs remove-membership --token=STRING --org=STRING --username=STRING
    Remove organization membership -
    https://developer.github.com/v3/orgs/members/#remove-organization-membership

  orgs remove-outside-collaborator --token=STRING --org=STRING --username=STRING
    Remove outside collaborator -
    https://developer.github.com/v3/orgs/outside_collaborators/#remove-outside-collaborator

  orgs unblock-user --token=STRING --org=STRING --username=STRING
    Unblock a user -
    https://developer.github.com/v3/orgs/blocking/#unblock-a-user

  orgs update --token=STRING --org=STRING
    Edit an organization -
    https://developer.github.com/v3/orgs/#edit-an-organization

  orgs update-membership --token=STRING --org=STRING --state=STRING
    Edit your organization membership -
    https://developer.github.com/v3/orgs/members/#edit-your-organization-membership

  projects add-collaborator --token=STRING --inertia-preview --project_id=INT-64 --username=STRING
    Add user as a collaborator -
    https://developer.github.com/v3/projects/collaborators/#add-user-as-a-collaborator

  projects create-card --token=STRING --column_id=INT-64 --inertia-preview
    Create a project card -
    https://developer.github.com/v3/projects/cards/#create-a-project-card

  projects create-column --token=STRING --inertia-preview --name=STRING --project_id=INT-64
    Create a project column -
    https://developer.github.com/v3/projects/columns/#create-a-project-column

  projects create-for-authenticated-user --token=STRING --inertia-preview --name=STRING
    Create a user project -
    https://developer.github.com/v3/projects/#create-a-user-project

  projects create-for-org --token=STRING --inertia-preview --name=STRING --org=STRING
    Create an organization project -
    https://developer.github.com/v3/projects/#create-an-organization-project

  projects create-for-repo --token=STRING --inertia-preview --name=STRING --repo=STRING
    Create a repository project -
    https://developer.github.com/v3/projects/#create-a-repository-project

  projects delete --token=STRING --inertia-preview --project_id=INT-64
    Delete a project -
    https://developer.github.com/v3/projects/#delete-a-project

  projects delete-card --token=STRING --card_id=INT-64 --inertia-preview
    Delete a project card -
    https://developer.github.com/v3/projects/cards/#delete-a-project-card

  projects delete-column --token=STRING --column_id=INT-64 --inertia-preview
    Delete a project column -
    https://developer.github.com/v3/projects/columns/#delete-a-project-column

  projects get --token=STRING --inertia-preview --project_id=INT-64
    Get a project - https://developer.github.com/v3/projects/#get-a-project

  projects get-card --token=STRING --card_id=INT-64 --inertia-preview
    Get a project card -
    https://developer.github.com/v3/projects/cards/#get-a-project-card

  projects get-column --token=STRING --column_id=INT-64 --inertia-preview
    Get a project column -
    https://developer.github.com/v3/projects/columns/#get-a-project-column

  projects list-cards --token=STRING --column_id=INT-64 --inertia-preview
    List project cards -
    https://developer.github.com/v3/projects/cards/#list-project-cards

  projects list-collaborators --token=STRING --inertia-preview --project_id=INT-64
    List collaborators -
    https://developer.github.com/v3/projects/collaborators/#list-collaborators

  projects list-columns --token=STRING --inertia-preview --project_id=INT-64
    List project columns -
    https://developer.github.com/v3/projects/columns/#list-project-columns

  projects list-for-org --token=STRING --inertia-preview --org=STRING
    List organization projects -
    https://developer.github.com/v3/projects/#list-organization-projects

  projects list-for-repo --token=STRING --inertia-preview --repo=STRING
    List repository projects -
    https://developer.github.com/v3/projects/#list-repository-projects

  projects list-for-user --token=STRING --inertia-preview --username=STRING
    List user projects -
    https://developer.github.com/v3/projects/#list-user-projects

  projects move-card --token=STRING --card_id=INT-64 --inertia-preview --position=STRING
    Move a project card -
    https://developer.github.com/v3/projects/cards/#move-a-project-card

  projects move-column --token=STRING --column_id=INT-64 --inertia-preview --position=STRING
    Move a project column -
    https://developer.github.com/v3/projects/columns/#move-a-project-column

  projects remove-collaborator --token=STRING --inertia-preview --project_id=INT-64 --username=STRING
    Remove user as a collaborator -
    https://developer.github.com/v3/projects/collaborators/#remove-user-as-a-collaborator

  projects review-user-permission-level --token=STRING --inertia-preview --project_id=INT-64 --username=STRING
    Review a user's permission level -
    https://developer.github.com/v3/projects/collaborators/#review-a-users-permission-level

  projects update --token=STRING --inertia-preview --project_id=INT-64
    Update a project -
    https://developer.github.com/v3/projects/#update-a-project

  projects update-card --token=STRING --card_id=INT-64 --inertia-preview
    Update a project card -
    https://developer.github.com/v3/projects/cards/#update-a-project-card

  projects update-column --token=STRING --column_id=INT-64 --inertia-preview --name=STRING
    Update a project column -
    https://developer.github.com/v3/projects/columns/#update-a-project-column

  pulls check-if-merged --token=STRING --pull_number=INT-64 --repo=STRING
    Get if a pull request has been merged -
    https://developer.github.com/v3/pulls/#get-if-a-pull-request-has-been-merged

  pulls create --token=STRING --base=STRING --head=STRING --repo=STRING --title=STRING
    Create a pull request -
    https://developer.github.com/v3/pulls/#create-a-pull-request

  pulls create-comment --token=STRING --body=STRING --commit_id=STRING --path=STRING --pull_number=INT-64 --repo=STRING
    Create a comment -
    https://developer.github.com/v3/pulls/comments/#create-a-comment

  pulls create-review-comment-reply --token=STRING --body=STRING --comment_id=INT-64 --pull_number=INT-64 --repo=STRING
    Create a review comment reply -
    https://developer.github.com/v3/pulls/comments/#create-a-review-comment-reply

  pulls create-review-request --token=STRING --pull_number=INT-64 --repo=STRING
    Create a review request -
    https://developer.github.com/v3/pulls/review_requests/#create-a-review-request

  pulls delete-comment --token=STRING --comment_id=INT-64 --repo=STRING
    Delete a comment -
    https://developer.github.com/v3/pulls/comments/#delete-a-comment

  pulls delete-pending-review --token=STRING --pull_number=INT-64 --repo=STRING --review_id=INT-64
    Delete a pending review -
    https://developer.github.com/v3/pulls/reviews/#delete-a-pending-review

  pulls delete-review-request --token=STRING --pull_number=INT-64 --repo=STRING
    Delete a review request -
    https://developer.github.com/v3/pulls/review_requests/#delete-a-review-request

  pulls dismiss-review --token=STRING --message=STRING --pull_number=INT-64 --repo=STRING --review_id=INT-64
    Dismiss a pull request review -
    https://developer.github.com/v3/pulls/reviews/#dismiss-a-pull-request-review

  pulls get --token=STRING --pull_number=INT-64 --repo=STRING
    Get a single pull request -
    https://developer.github.com/v3/pulls/#get-a-single-pull-request

  pulls get-comment --token=STRING --comment_id=INT-64 --repo=STRING
    Get a single comment -
    https://developer.github.com/v3/pulls/comments/#get-a-single-comment

  pulls get-comments-for-review --token=STRING --pull_number=INT-64 --repo=STRING --review_id=INT-64
    Get comments for a single review -
    https://developer.github.com/v3/pulls/reviews/#get-comments-for-a-single-review

  pulls get-review --token=STRING --pull_number=INT-64 --repo=STRING --review_id=INT-64
    Get a single review -
    https://developer.github.com/v3/pulls/reviews/#get-a-single-review

  pulls list --token=STRING --repo=STRING
    List pull requests -
    https://developer.github.com/v3/pulls/#list-pull-requests

  pulls list-comments --token=STRING --pull_number=INT-64 --repo=STRING
    List comments on a pull request -
    https://developer.github.com/v3/pulls/comments/#list-comments-on-a-pull-request

  pulls list-comments-for-repo --token=STRING --repo=STRING
    List comments in a repository -
    https://developer.github.com/v3/pulls/comments/#list-comments-in-a-repository

  pulls list-commits --token=STRING --pull_number=INT-64 --repo=STRING
    List commits on a pull request -
    https://developer.github.com/v3/pulls/#list-commits-on-a-pull-request

  pulls list-files --token=STRING --pull_number=INT-64 --repo=STRING
    List pull requests files -
    https://developer.github.com/v3/pulls/#list-pull-requests-files

  pulls list-review-requests --token=STRING --pull_number=INT-64 --repo=STRING
    List review requests -
    https://developer.github.com/v3/pulls/review_requests/#list-review-requests

  pulls list-reviews --token=STRING --pull_number=INT-64 --repo=STRING
    List reviews on a pull request -
    https://developer.github.com/v3/pulls/reviews/#list-reviews-on-a-pull-request

  pulls merge --token=STRING --pull_number=INT-64 --repo=STRING
    Merge a pull request (Merge Button) -
    https://developer.github.com/v3/pulls/#merge-a-pull-request-merge-button

  pulls submit-review --token=STRING --event=STRING --pull_number=INT-64 --repo=STRING --review_id=INT-64
    Submit a pull request review -
    https://developer.github.com/v3/pulls/reviews/#submit-a-pull-request-review

  pulls update --token=STRING --pull_number=INT-64 --repo=STRING
    Update a pull request -
    https://developer.github.com/v3/pulls/#update-a-pull-request

  pulls update-branch --token=STRING --lydian-preview --pull_number=INT-64 --repo=STRING
    Update a pull request branch -
    https://developer.github.com/v3/pulls/#update-a-pull-request-branch

  pulls update-comment --token=STRING --body=STRING --comment_id=INT-64 --repo=STRING
    Edit a comment -
    https://developer.github.com/v3/pulls/comments/#edit-a-comment

  pulls update-review --token=STRING --body=STRING --pull_number=INT-64 --repo=STRING --review_id=INT-64
    Update a pull request review -
    https://developer.github.com/v3/pulls/reviews/#update-a-pull-request-review

  rate-limit get --token=STRING
    Get your current rate limit status -
    https://developer.github.com/v3/rate_limit/#get-your-current-rate-limit-status

  reactions create-for-commit-comment --token=STRING --comment_id=INT-64 --content=STRING --repo=STRING --squirrel-girl-preview
    Create reaction for a commit comment -
    https://developer.github.com/v3/reactions/#create-reaction-for-a-commit-comment

  reactions create-for-issue --token=STRING --content=STRING --issue_number=INT-64 --repo=STRING --squirrel-girl-preview
    Create reaction for an issue -
    https://developer.github.com/v3/reactions/#create-reaction-for-an-issue

  reactions create-for-issue-comment --token=STRING --comment_id=INT-64 --content=STRING --repo=STRING --squirrel-girl-preview
    Create reaction for an issue comment -
    https://developer.github.com/v3/reactions/#create-reaction-for-an-issue-comment

  reactions create-for-pull-request-review-comment --token=STRING --comment_id=INT-64 --content=STRING --repo=STRING --squirrel-girl-preview
    Create reaction for a pull request review comment -
    https://developer.github.com/v3/reactions/#create-reaction-for-a-pull-request-review-comment

  reactions create-for-team-discussion-comment-in-org --token=STRING --comment_number=INT-64 --content=STRING --discussion_number=INT-64 --org=STRING --squirrel-girl-preview --team_slug=STRING
    Create reaction for a team discussion comment -
    https://developer.github.com/v3/reactions/#create-reaction-for-a-team-discussion-comment

  reactions create-for-team-discussion-comment-legacy --token=STRING --comment_number=INT-64 --content=STRING --discussion_number=INT-64 --squirrel-girl-preview --team_id=INT-64
    Create reaction for a team discussion comment (Legacy) -
    https://developer.github.com/v3/reactions/#create-reaction-for-a-team-discussion-comment-legacy

  reactions create-for-team-discussion-in-org --token=STRING --content=STRING --discussion_number=INT-64 --org=STRING --squirrel-girl-preview --team_slug=STRING
    Create reaction for a team discussion -
    https://developer.github.com/v3/reactions/#create-reaction-for-a-team-discussion

  reactions create-for-team-discussion-legacy --token=STRING --content=STRING --discussion_number=INT-64 --squirrel-girl-preview --team_id=INT-64
    Create reaction for a team discussion (Legacy) -
    https://developer.github.com/v3/reactions/#create-reaction-for-a-team-discussion-legacy

  reactions delete-for-commit-comment --token=STRING --comment_id=INT-64 --reaction_id=INT-64 --repo=STRING --squirrel-girl-preview
    Delete a commit comment reaction -
    https://developer.github.com/v3/reactions/#delete-a-commit-comment-reaction

  reactions delete-for-issue --token=STRING --issue_number=INT-64 --reaction_id=INT-64 --repo=STRING --squirrel-girl-preview
    Delete an issue reaction -
    https://developer.github.com/v3/reactions/#delete-an-issue-reaction

  reactions delete-for-issue-comment --token=STRING --comment_id=INT-64 --reaction_id=INT-64 --repo=STRING --squirrel-girl-preview
    Delete an issue comment reaction -
    https://developer.github.com/v3/reactions/#delete-an-issue-comment-reaction

  reactions delete-for-pull-request-comment --token=STRING --comment_id=INT-64 --reaction_id=INT-64 --repo=STRING --squirrel-girl-preview
    Delete a pull request comment reaction -
    https://developer.github.com/v3/reactions/#delete-a-pull-request-comment-reaction

  reactions delete-for-team-discussion --token=STRING --discussion_number=INT-64 --org=STRING --reaction_id=INT-64 --squirrel-girl-preview --team_slug=STRING
    Delete team discussion reaction -
    https://developer.github.com/v3/reactions/#delete-team-discussion-reaction

  reactions delete-for-team-discussion-comment --token=STRING --comment_number=INT-64 --discussion_number=INT-64 --org=STRING --reaction_id=INT-64 --squirrel-girl-preview --team_slug=STRING
    Delete team discussion comment reaction -
    https://developer.github.com/v3/reactions/#delete-team-discussion-comment-reaction

  reactions delete-legacy --token=STRING --reaction_id=INT-64 --squirrel-girl-preview
    Delete a reaction (Legacy) -
    https://developer.github.com/v3/reactions/#delete-a-reaction-legacy

  reactions list-for-commit-comment --token=STRING --comment_id=INT-64 --repo=STRING --squirrel-girl-preview
    List reactions for a commit comment -
    https://developer.github.com/v3/reactions/#list-reactions-for-a-commit-comment

  reactions list-for-issue --token=STRING --issue_number=INT-64 --repo=STRING --squirrel-girl-preview
    List reactions for an issue -
    https://developer.github.com/v3/reactions/#list-reactions-for-an-issue

  reactions list-for-issue-comment --token=STRING --comment_id=INT-64 --repo=STRING --squirrel-girl-preview
    List reactions for an issue comment -
    https://developer.github.com/v3/reactions/#list-reactions-for-an-issue-comment

  reactions list-for-pull-request-review-comment --token=STRING --comment_id=INT-64 --repo=STRING --squirrel-girl-preview
    List reactions for a pull request review comment -
    https://developer.github.com/v3/reactions/#list-reactions-for-a-pull-request-review-comment

  reactions list-for-team-discussion-comment-in-org --token=STRING --comment_number=INT-64 --discussion_number=INT-64 --org=STRING --squirrel-girl-preview --team_slug=STRING
    List reactions for a team discussion comment -
    https://developer.github.com/v3/reactions/#list-reactions-for-a-team-discussion-comment

  reactions list-for-team-discussion-comment-legacy --token=STRING --comment_number=INT-64 --discussion_number=INT-64 --squirrel-girl-preview --team_id=INT-64
    List reactions for a team discussion comment (Legacy) -
    https://developer.github.com/v3/reactions/#list-reactions-for-a-team-discussion-comment-legacy

  reactions list-for-team-discussion-in-org --token=STRING --discussion_number=INT-64 --org=STRING --squirrel-girl-preview --team_slug=STRING
    List reactions for a team discussion -
    https://developer.github.com/v3/reactions/#list-reactions-for-a-team-discussion

  reactions list-for-team-discussion-legacy --token=STRING --discussion_number=INT-64 --squirrel-girl-preview --team_id=INT-64
    List reactions for a team discussion (Legacy) -
    https://developer.github.com/v3/reactions/#list-reactions-for-a-team-discussion-legacy

  repos accept-invitation --token=STRING --invitation_id=INT-64
    Accept a repository invitation -
    https://developer.github.com/v3/repos/invitations/#accept-a-repository-invitation

  repos add-collaborator --token=STRING --repo=STRING --username=STRING
    Add user as a collaborator -
    https://developer.github.com/v3/repos/collaborators/#add-user-as-a-collaborator

  repos add-deploy-key --token=STRING --key=STRING --repo=STRING
    Add a new deploy key -
    https://developer.github.com/v3/repos/keys/#add-a-new-deploy-key

  repos add-protected-branch-admin-enforcement --token=STRING --branch=STRING --repo=STRING
    Add admin enforcement of protected branch -
    https://developer.github.com/v3/repos/branches/#add-admin-enforcement-of-protected-branch

  repos add-protected-branch-app-restrictions --token=STRING --branch=STRING --repo=STRING
    Add app restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#add-app-restrictions-of-protected-branch

  repos add-protected-branch-required-signatures --token=STRING --branch=STRING --repo=STRING --zzzax-preview
    Add required signatures of protected branch -
    https://developer.github.com/v3/repos/branches/#add-required-signatures-of-protected-branch

  repos add-protected-branch-required-status-checks-contexts --token=STRING --branch=STRING --repo=STRING
    Add required status checks contexts of protected branch -
    https://developer.github.com/v3/repos/branches/#add-required-status-checks-contexts-of-protected-branch

  repos add-protected-branch-team-restrictions --token=STRING --branch=STRING --repo=STRING
    Add team restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#add-team-restrictions-of-protected-branch

  repos add-protected-branch-user-restrictions --token=STRING --branch=STRING --repo=STRING
    Add user restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#add-user-restrictions-of-protected-branch

  repos check-collaborator --token=STRING --repo=STRING --username=STRING
    Check if a user is a collaborator -
    https://developer.github.com/v3/repos/collaborators/#check-if-a-user-is-a-collaborator

  repos check-vulnerability-alerts --token=STRING --dorian-preview --repo=STRING
    Check if vulnerability alerts are enabled for a repository -
    https://developer.github.com/v3/repos/#check-if-vulnerability-alerts-are-enabled-for-a-repository

  repos compare-commits --token=STRING --base=STRING --head=STRING --repo=STRING
    Compare two commits -
    https://developer.github.com/v3/repos/commits/#compare-two-commits

  repos create-commit-comment --token=STRING --body=STRING --commit_sha=STRING --repo=STRING
    Create a commit comment -
    https://developer.github.com/v3/repos/comments/#create-a-commit-comment

  repos create-deployment --token=STRING --ref=STRING --repo=STRING
    Create a deployment -
    https://developer.github.com/v3/repos/deployments/#create-a-deployment

  repos create-deployment-status --token=STRING --deployment_id=INT-64 --repo=STRING --state=STRING
    Create a deployment status -
    https://developer.github.com/v3/repos/deployments/#create-a-deployment-status

  repos create-for-authenticated-user --token=STRING --name=STRING
    Create a repository for the authenticated user -
    https://developer.github.com/v3/repos/#create-a-repository-for-the-authenticated-user

  repos create-fork --token=STRING --repo=STRING
    Create a fork - https://developer.github.com/v3/repos/forks/#create-a-fork

  repos create-in-org --token=STRING --name=STRING --org=STRING
    Create an organization repository -
    https://developer.github.com/v3/repos/#create-an-organization-repository

  repos create-release --token=STRING --repo=STRING --tag_name=STRING
    Create a release -
    https://developer.github.com/v3/repos/releases/#create-a-release

  repos create-status --token=STRING --repo=STRING --sha=STRING --state=STRING
    Create a status -
    https://developer.github.com/v3/repos/statuses/#create-a-status

  repos create-using-template --token=STRING --baptiste-preview --name=STRING --template_owner=STRING --template_repo=STRING
    Create a repository using a template -
    https://developer.github.com/v3/repos/#create-a-repository-using-a-template

  repos decline-invitation --token=STRING --invitation_id=INT-64
    Decline a repository invitation -
    https://developer.github.com/v3/repos/invitations/#decline-a-repository-invitation

  repos delete --token=STRING --repo=STRING
    Delete a repository -
    https://developer.github.com/v3/repos/#delete-a-repository

  repos delete-commit-comment --token=STRING --comment_id=INT-64 --repo=STRING
    Delete a commit comment -
    https://developer.github.com/v3/repos/comments/#delete-a-commit-comment

  repos delete-deployment --token=STRING --deployment_id=INT-64 --repo=STRING
    Delete a deployment -
    https://developer.github.com/v3/repos/deployments/#delete-a-deployment

  repos delete-download --token=STRING --download_id=INT-64 --repo=STRING
    Delete a download -
    https://developer.github.com/v3/repos/downloads/#delete-a-download

  repos delete-hook --token=STRING --hook_id=INT-64 --repo=STRING
    Delete a hook - https://developer.github.com/v3/repos/hooks/#delete-a-hook

  repos delete-invitation --token=STRING --invitation_id=INT-64 --repo=STRING
    Delete a repository invitation -
    https://developer.github.com/v3/repos/invitations/#delete-a-repository-invitation

  repos delete-release --token=STRING --release_id=INT-64 --repo=STRING
    Delete a release -
    https://developer.github.com/v3/repos/releases/#delete-a-release

  repos delete-release-asset --token=STRING --asset_id=INT-64 --repo=STRING
    Delete a release asset -
    https://developer.github.com/v3/repos/releases/#delete-a-release-asset

  repos disable-automated-security-fixes --token=STRING --london-preview --repo=STRING
    Disable automated security fixes -
    https://developer.github.com/v3/repos/#disable-automated-security-fixes

  repos disable-pages-site --token=STRING --repo=STRING --switcheroo-preview
    Disable a Pages site -
    https://developer.github.com/v3/repos/pages/#disable-a-pages-site

  repos disable-vulnerability-alerts --token=STRING --dorian-preview --repo=STRING
    Disable vulnerability alerts -
    https://developer.github.com/v3/repos/#disable-vulnerability-alerts

  repos enable-automated-security-fixes --token=STRING --london-preview --repo=STRING
    Enable automated security fixes -
    https://developer.github.com/v3/repos/#enable-automated-security-fixes

  repos enable-vulnerability-alerts --token=STRING --dorian-preview --repo=STRING
    Enable vulnerability alerts -
    https://developer.github.com/v3/repos/#enable-vulnerability-alerts

  repos get --token=STRING --repo=STRING
    Get a repository - https://developer.github.com/v3/repos/#get-a-repository

  repos get-all-topics --token=STRING --mercy-preview --repo=STRING
    Get all repository topics -
    https://developer.github.com/v3/repos/#get-all-repository-topics

  repos get-apps-with-access-to-protected-branch --token=STRING --branch=STRING --repo=STRING
    Get apps with access to protected branch -
    https://developer.github.com/v3/repos/branches/#list-apps-with-access-to-protected-branch

  repos get-archive-link --token=STRING --archive_format=STRING --ref=STRING --repo=STRING
    Get archive link -
    https://developer.github.com/v3/repos/contents/#get-archive-link

  repos get-branch --token=STRING --branch=STRING --repo=STRING
    Get branch - https://developer.github.com/v3/repos/branches/#get-branch

  repos get-branch-protection --token=STRING --branch=STRING --repo=STRING
    Get branch protection -
    https://developer.github.com/v3/repos/branches/#get-branch-protection

  repos get-clones --token=STRING --repo=STRING
    Clones - https://developer.github.com/v3/repos/traffic/#clones

  repos get-code-frequency-stats --token=STRING --repo=STRING
    Get the number of additions and deletions per week -
    https://developer.github.com/v3/repos/statistics/#get-the-number-of-additions-and-deletions-per-week

  repos get-collaborator-permission-level --token=STRING --repo=STRING --username=STRING
    Review a user's permission level -
    https://developer.github.com/v3/repos/collaborators/#review-a-users-permission-level

  repos get-combined-status-for-ref --token=STRING --ref=STRING --repo=STRING
    Get the combined status for a specific ref -
    https://developer.github.com/v3/repos/statuses/#get-the-combined-status-for-a-specific-ref

  repos get-commit --token=STRING --ref=STRING --repo=STRING
    Get a single commit -
    https://developer.github.com/v3/repos/commits/#get-a-single-commit

  repos get-commit-activity-stats --token=STRING --repo=STRING
    Get the last year of commit activity data -
    https://developer.github.com/v3/repos/statistics/#get-the-last-year-of-commit-activity-data

  repos get-commit-comment --token=STRING --comment_id=INT-64 --repo=STRING
    Get a single commit comment -
    https://developer.github.com/v3/repos/comments/#get-a-single-commit-comment

  repos get-contents --token=STRING --path=STRING --repo=STRING
    Get contents - https://developer.github.com/v3/repos/contents/#get-contents

  repos get-contributors-stats --token=STRING --repo=STRING
    Get contributors list with additions, deletions, and commit counts -
    https://developer.github.com/v3/repos/statistics/#get-contributors-list-with-additions-deletions-and-commit-counts

  repos get-deploy-key --token=STRING --key_id=INT-64 --repo=STRING
    Get a deploy key -
    https://developer.github.com/v3/repos/keys/#get-a-deploy-key

  repos get-deployment --token=STRING --deployment_id=INT-64 --repo=STRING
    Get a single deployment -
    https://developer.github.com/v3/repos/deployments/#get-a-single-deployment

  repos get-deployment-status --token=STRING --deployment_id=INT-64 --repo=STRING --status_id=INT-64
    Get a single deployment status -
    https://developer.github.com/v3/repos/deployments/#get-a-single-deployment-status

  repos get-download --token=STRING --download_id=INT-64 --repo=STRING
    Get a single download -
    https://developer.github.com/v3/repos/downloads/#get-a-single-download

  repos get-hook --token=STRING --hook_id=INT-64 --repo=STRING
    Get single hook -
    https://developer.github.com/v3/repos/hooks/#get-single-hook

  repos get-latest-pages-build --token=STRING --repo=STRING
    Get latest Pages build -
    https://developer.github.com/v3/repos/pages/#get-latest-pages-build

  repos get-latest-release --token=STRING --repo=STRING
    Get the latest release -
    https://developer.github.com/v3/repos/releases/#get-the-latest-release

  repos get-pages --token=STRING --repo=STRING
    Get information about a Pages site -
    https://developer.github.com/v3/repos/pages/#get-information-about-a-pages-site

  repos get-pages-build --token=STRING --build_id=INT-64 --repo=STRING
    Get a specific Pages build -
    https://developer.github.com/v3/repos/pages/#get-a-specific-pages-build

  repos get-participation-stats --token=STRING --repo=STRING
    Get the weekly commit count for the repository owner and everyone else -
    https://developer.github.com/v3/repos/statistics/#get-the-weekly-commit-count-for-the-repository-owner-and-everyone-else

  repos get-protected-branch-admin-enforcement --token=STRING --branch=STRING --repo=STRING
    Get admin enforcement of protected branch -
    https://developer.github.com/v3/repos/branches/#get-admin-enforcement-of-protected-branch

  repos get-protected-branch-pull-request-review-enforcement --token=STRING --branch=STRING --repo=STRING
    Get pull request review enforcement of protected branch -
    https://developer.github.com/v3/repos/branches/#get-pull-request-review-enforcement-of-protected-branch

  repos get-protected-branch-required-signatures --token=STRING --branch=STRING --repo=STRING --zzzax-preview
    Get required signatures of protected branch -
    https://developer.github.com/v3/repos/branches/#get-required-signatures-of-protected-branch

  repos get-protected-branch-required-status-checks --token=STRING --branch=STRING --repo=STRING
    Get required status checks of protected branch -
    https://developer.github.com/v3/repos/branches/#get-required-status-checks-of-protected-branch

  repos get-protected-branch-restrictions --token=STRING --branch=STRING --repo=STRING
    Get restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#get-restrictions-of-protected-branch

  repos get-punch-card-stats --token=STRING --repo=STRING
    Get the number of commits per hour in each day -
    https://developer.github.com/v3/repos/statistics/#get-the-number-of-commits-per-hour-in-each-day

  repos get-readme --token=STRING --repo=STRING
    Get the README -
    https://developer.github.com/v3/repos/contents/#get-the-readme

  repos get-release --token=STRING --release_id=INT-64 --repo=STRING
    Get a single release -
    https://developer.github.com/v3/repos/releases/#get-a-single-release

  repos get-release-asset --token=STRING --asset_id=INT-64 --repo=STRING
    Get a single release asset -
    https://developer.github.com/v3/repos/releases/#get-a-single-release-asset

  repos get-release-by-tag --token=STRING --repo=STRING --tag=STRING
    Get a release by tag name -
    https://developer.github.com/v3/repos/releases/#get-a-release-by-tag-name

  repos get-teams-with-access-to-protected-branch --token=STRING --branch=STRING --repo=STRING
    Get teams with access to protected branch -
    https://developer.github.com/v3/repos/branches/#list-teams-with-access-to-protected-branch

  repos get-top-paths --token=STRING --repo=STRING
    List paths - https://developer.github.com/v3/repos/traffic/#list-paths

  repos get-top-referrers --token=STRING --repo=STRING
    List referrers -
    https://developer.github.com/v3/repos/traffic/#list-referrers

  repos get-users-with-access-to-protected-branch --token=STRING --branch=STRING --repo=STRING
    Get users with access to protected branch -
    https://developer.github.com/v3/repos/branches/#list-users-with-access-to-protected-branch

  repos get-views --token=STRING --repo=STRING
    Views - https://developer.github.com/v3/repos/traffic/#views

  repos list-assets-for-release --token=STRING --release_id=INT-64 --repo=STRING
    List assets for a release -
    https://developer.github.com/v3/repos/releases/#list-assets-for-a-release

  repos list-branches --token=STRING --repo=STRING
    List branches -
    https://developer.github.com/v3/repos/branches/#list-branches

  repos list-branches-for-head-commit --token=STRING --commit_sha=STRING --groot-preview --repo=STRING
    List branches for HEAD commit -
    https://developer.github.com/v3/repos/commits/#list-branches-for-head-commit

  repos list-collaborators --token=STRING --repo=STRING
    List collaborators -
    https://developer.github.com/v3/repos/collaborators/#list-collaborators

  repos list-comments-for-commit --token=STRING --commit_sha=STRING --repo=STRING
    List comments for a single commit -
    https://developer.github.com/v3/repos/comments/#list-comments-for-a-single-commit

  repos list-commit-comments --token=STRING --repo=STRING
    List commit comments for a repository -
    https://developer.github.com/v3/repos/comments/#list-commit-comments-for-a-repository

  repos list-commits --token=STRING --repo=STRING
    List commits on a repository -
    https://developer.github.com/v3/repos/commits/#list-commits-on-a-repository

  repos list-contributors --token=STRING --repo=STRING
    List contributors - https://developer.github.com/v3/repos/#list-contributors

  repos list-deploy-keys --token=STRING --repo=STRING
    List deploy keys -
    https://developer.github.com/v3/repos/keys/#list-deploy-keys

  repos list-deployment-statuses --token=STRING --deployment_id=INT-64 --repo=STRING
    List deployment statuses -
    https://developer.github.com/v3/repos/deployments/#list-deployment-statuses

  repos list-deployments --token=STRING --repo=STRING
    List deployments -
    https://developer.github.com/v3/repos/deployments/#list-deployments

  repos list-downloads --token=STRING --repo=STRING
    List downloads for a repository -
    https://developer.github.com/v3/repos/downloads/#list-downloads-for-a-repository

  repos list-for-authenticated-user --token=STRING
    List repositories for the authenticated user -
    https://developer.github.com/v3/repos/#list-repositories-for-the-authenticated-user

  repos list-for-org --token=STRING --org=STRING
    List organization repositories -
    https://developer.github.com/v3/repos/#list-organization-repositories

  repos list-for-user --token=STRING --username=STRING
    List repositories for a user -
    https://developer.github.com/v3/repos/#list-repositories-for-a-user

  repos list-forks --token=STRING --repo=STRING
    List forks - https://developer.github.com/v3/repos/forks/#list-forks

  repos list-hooks --token=STRING --repo=STRING
    List hooks - https://developer.github.com/v3/repos/hooks/#list-hooks

  repos list-invitations --token=STRING --repo=STRING
    List invitations for a repository -
    https://developer.github.com/v3/repos/invitations/#list-invitations-for-a-repository

  repos list-invitations-for-authenticated-user --token=STRING
    List a user's repository invitations -
    https://developer.github.com/v3/repos/invitations/#list-a-users-repository-invitations

  repos list-languages --token=STRING --repo=STRING
    List languages - https://developer.github.com/v3/repos/#list-languages

  repos list-pages-builds --token=STRING --repo=STRING
    List Pages builds -
    https://developer.github.com/v3/repos/pages/#list-pages-builds

  repos list-protected-branch-required-status-checks-contexts --token=STRING --branch=STRING --repo=STRING
    List required status checks contexts of protected branch -
    https://developer.github.com/v3/repos/branches/#list-required-status-checks-contexts-of-protected-branch

  repos list-public --token=STRING
    List public repositories -
    https://developer.github.com/v3/repos/#list-public-repositories

  repos list-pull-requests-associated-with-commit --token=STRING --commit_sha=STRING --groot-preview --repo=STRING
    List pull requests associated with commit -
    https://developer.github.com/v3/repos/commits/#list-pull-requests-associated-with-commit

  repos list-releases --token=STRING --repo=STRING
    List releases for a repository -
    https://developer.github.com/v3/repos/releases/#list-releases-for-a-repository

  repos list-statuses-for-ref --token=STRING --ref=STRING --repo=STRING
    List statuses for a specific ref -
    https://developer.github.com/v3/repos/statuses/#list-statuses-for-a-specific-ref

  repos list-tags --token=STRING --repo=STRING
    List tags - https://developer.github.com/v3/repos/#list-tags

  repos list-teams --token=STRING --repo=STRING
    List teams - https://developer.github.com/v3/repos/#list-teams

  repos merge --token=STRING --base=STRING --head=STRING --repo=STRING
    Perform a merge -
    https://developer.github.com/v3/repos/merging/#perform-a-merge

  repos ping-hook --token=STRING --hook_id=INT-64 --repo=STRING
    Ping a hook - https://developer.github.com/v3/repos/hooks/#ping-a-hook

  repos remove-branch-protection --token=STRING --branch=STRING --repo=STRING
    Remove branch protection -
    https://developer.github.com/v3/repos/branches/#remove-branch-protection

  repos remove-collaborator --token=STRING --repo=STRING --username=STRING
    Remove user as a collaborator -
    https://developer.github.com/v3/repos/collaborators/#remove-user-as-a-collaborator

  repos remove-deploy-key --token=STRING --key_id=INT-64 --repo=STRING
    Remove a deploy key -
    https://developer.github.com/v3/repos/keys/#remove-a-deploy-key

  repos remove-protected-branch-admin-enforcement --token=STRING --branch=STRING --repo=STRING
    Remove admin enforcement of protected branch -
    https://developer.github.com/v3/repos/branches/#remove-admin-enforcement-of-protected-branch

  repos remove-protected-branch-app-restrictions --token=STRING --branch=STRING --repo=STRING
    Remove app restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#remove-app-restrictions-of-protected-branch

  repos remove-protected-branch-pull-request-review-enforcement --token=STRING --branch=STRING --repo=STRING
    Remove pull request review enforcement of protected branch -
    https://developer.github.com/v3/repos/branches/#remove-pull-request-review-enforcement-of-protected-branch

  repos remove-protected-branch-required-signatures --token=STRING --branch=STRING --repo=STRING --zzzax-preview
    Remove required signatures of protected branch -
    https://developer.github.com/v3/repos/branches/#remove-required-signatures-of-protected-branch

  repos remove-protected-branch-required-status-checks --token=STRING --branch=STRING --repo=STRING
    Remove required status checks of protected branch -
    https://developer.github.com/v3/repos/branches/#remove-required-status-checks-of-protected-branch

  repos remove-protected-branch-required-status-checks-contexts --token=STRING --branch=STRING --repo=STRING
    Remove required status checks contexts of protected branch -
    https://developer.github.com/v3/repos/branches/#remove-required-status-checks-contexts-of-protected-branch

  repos remove-protected-branch-restrictions --token=STRING --branch=STRING --repo=STRING
    Remove restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#remove-restrictions-of-protected-branch

  repos remove-protected-branch-team-restrictions --token=STRING --branch=STRING --repo=STRING
    Remove team restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#remove-team-restrictions-of-protected-branch

  repos remove-protected-branch-user-restrictions --token=STRING --branch=STRING --repo=STRING
    Remove user restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#remove-user-restrictions-of-protected-branch

  repos replace-all-topics --token=STRING --mercy-preview --names=NAMES,... --repo=STRING
    Replace all repository topics -
    https://developer.github.com/v3/repos/#replace-all-repository-topics

  repos replace-protected-branch-app-restrictions --token=STRING --branch=STRING --repo=STRING
    Replace app restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#replace-app-restrictions-of-protected-branch

  repos replace-protected-branch-required-status-checks-contexts --token=STRING --branch=STRING --repo=STRING
    Replace required status checks contexts of protected branch -
    https://developer.github.com/v3/repos/branches/#replace-required-status-checks-contexts-of-protected-branch

  repos replace-protected-branch-team-restrictions --token=STRING --branch=STRING --repo=STRING
    Replace team restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#replace-team-restrictions-of-protected-branch

  repos replace-protected-branch-user-restrictions --token=STRING --branch=STRING --repo=STRING
    Replace user restrictions of protected branch -
    https://developer.github.com/v3/repos/branches/#replace-user-restrictions-of-protected-branch

  repos request-page-build --token=STRING --repo=STRING
    Request a page build -
    https://developer.github.com/v3/repos/pages/#request-a-page-build

  repos retrieve-community-profile-metrics --token=STRING --repo=STRING
    Retrieve community profile metrics -
    https://developer.github.com/v3/repos/community/#retrieve-community-profile-metrics

  repos test-push-hook --token=STRING --hook_id=INT-64 --repo=STRING
    Test a push hook -
    https://developer.github.com/v3/repos/hooks/#test-a-push-hook

  repos transfer --token=STRING --repo=STRING
    Transfer a repository -
    https://developer.github.com/v3/repos/#transfer-a-repository

  repos update --token=STRING --repo=STRING
    Update a repository -
    https://developer.github.com/v3/repos/#update-a-repository

  repos update-commit-comment --token=STRING --body=STRING --comment_id=INT-64 --repo=STRING
    Update a commit comment -
    https://developer.github.com/v3/repos/comments/#update-a-commit-comment

  repos update-information-about-pages-site --token=STRING --repo=STRING
    Update information about a Pages site -
    https://developer.github.com/v3/repos/pages/#update-information-about-a-pages-site

  repos update-invitation --token=STRING --invitation_id=INT-64 --repo=STRING
    Update a repository invitation -
    https://developer.github.com/v3/repos/invitations/#update-a-repository-invitation

  repos update-protected-branch-required-status-checks --token=STRING --branch=STRING --repo=STRING
    Update required status checks of protected branch -
    https://developer.github.com/v3/repos/branches/#update-required-status-checks-of-protected-branch

  repos update-release --token=STRING --release_id=INT-64 --repo=STRING
    Edit a release -
    https://developer.github.com/v3/repos/releases/#edit-a-release

  repos update-release-asset --token=STRING --asset_id=INT-64 --repo=STRING
    Edit a release asset -
    https://developer.github.com/v3/repos/releases/#edit-a-release-asset

  repos upload-release-asset --token=STRING --content-length=STRING --content-type=STRING --release_id=INT-64 --repo=STRING
    Upload a release asset -
    https://developer.github.com/v3/repos/releases/#upload-a-release-asset

  scim get-provisioning-details-for-user --token=STRING --org=STRING --scim_user_id=INT-64
    Get provisioning details for a single user -
    https://developer.github.com/v3/scim/#get-provisioning-details-for-a-single-user

  scim list-provisioned-identities --token=STRING --org=STRING
    Get a list of provisioned identities -
    https://developer.github.com/v3/scim/#get-a-list-of-provisioned-identities

  scim provision-and-invite-users --token=STRING --org=STRING
    Provision and invite users -
    https://developer.github.com/v3/scim/#provision-and-invite-users

  scim remove-user-from-org --token=STRING --org=STRING --scim_user_id=INT-64
    Remove a user from the organization -
    https://developer.github.com/v3/scim/#remove-a-user-from-the-organization

  scim replace-provisioned-user-information --token=STRING --org=STRING --scim_user_id=INT-64
    Replace a provisioned user's information -
    https://developer.github.com/v3/scim/#replace-a-provisioned-users-information

  scim update-user-attribute --token=STRING --org=STRING --scim_user_id=INT-64
    Update a user attribute -
    https://developer.github.com/v3/scim/#update-a-user-attribute

  search code --token=STRING --q=STRING
    Search code - https://developer.github.com/v3/search/#search-code

  search commits --token=STRING --cloak-preview --q=STRING
    Search commits - https://developer.github.com/v3/search/#search-commits

  search email-legacy --token=STRING --email=STRING
    Email search - https://developer.github.com/v3/search/legacy/#email-search

  search issues-and-pull-requests --token=STRING --q=STRING
    Search issues and pull requests -
    https://developer.github.com/v3/search/#search-issues-and-pull-requests

  search issues-legacy --token=STRING --keyword=STRING --owner=STRING --repository=STRING --state=STRING
    Search issues - https://developer.github.com/v3/search/legacy/#search-issues

  search labels --token=STRING --q=STRING --repository_id=INT-64
    Search labels - https://developer.github.com/v3/search/#search-labels

  search repos --token=STRING --q=STRING
    Search repositories -
    https://developer.github.com/v3/search/#search-repositories

  search repos-legacy --token=STRING --keyword=STRING
    Search repositories -
    https://developer.github.com/v3/search/legacy/#search-repositories

  search topics --token=STRING --q=STRING
    Search topics - https://developer.github.com/v3/search/#search-topics

  search users --token=STRING --q=STRING
    Search users - https://developer.github.com/v3/search/#search-users

  search users-legacy --token=STRING --keyword=STRING
    Search users - https://developer.github.com/v3/search/legacy/#search-users

  teams add-member-legacy --token=STRING --team_id=INT-64 --username=STRING
    Add team member (Legacy) -
    https://developer.github.com/v3/teams/members/#add-team-member-legacy

  teams add-or-update-membership-in-org --token=STRING --org=STRING --team_slug=STRING --username=STRING
    Add or update team membership -
    https://developer.github.com/v3/teams/members/#add-or-update-team-membership

  teams add-or-update-membership-legacy --token=STRING --team_id=INT-64 --username=STRING
    Add or update team membership (Legacy) -
    https://developer.github.com/v3/teams/members/#add-or-update-team-membership-legacy

  teams add-or-update-project-in-org --token=STRING --inertia-preview --org=STRING --project_id=INT-64 --team_slug=STRING
    Add or update team project -
    https://developer.github.com/v3/teams/#add-or-update-team-project

  teams add-or-update-project-legacy --token=STRING --inertia-preview --project_id=INT-64 --team_id=INT-64
    Add or update team project (Legacy) -
    https://developer.github.com/v3/teams/#add-or-update-team-project-legacy

  teams add-or-update-repo-in-org --token=STRING --org=STRING --repo=STRING --team_slug=STRING
    Add or update team repository -
    https://developer.github.com/v3/teams/#add-or-update-team-repository

  teams add-or-update-repo-legacy --token=STRING --repo=STRING --team_id=INT-64
    Add or update team repository (Legacy) -
    https://developer.github.com/v3/teams/#add-or-update-team-repository-legacy

  teams check-manages-repo-in-org --token=STRING --org=STRING --repo=STRING --team_slug=STRING
    Check if a team manages a repository -
    https://developer.github.com/v3/teams/#check-if-a-team-manages-a-repository

  teams check-manages-repo-legacy --token=STRING --repo=STRING --team_id=INT-64
    Check if a team manages a repository (Legacy) -
    https://developer.github.com/v3/teams/#check-if-a-team-manages-a-repository-legacy

  teams create --token=STRING --name=STRING --org=STRING
    Create team - https://developer.github.com/v3/teams/#create-team

  teams create-discussion-comment-in-org --token=STRING --body=STRING --discussion_number=INT-64 --org=STRING --team_slug=STRING
    Create a comment -
    https://developer.github.com/v3/teams/discussion_comments/#create-a-comment

  teams create-discussion-comment-legacy --token=STRING --body=STRING --discussion_number=INT-64 --team_id=INT-64
    Create a comment (Legacy) -
    https://developer.github.com/v3/teams/discussion_comments/#create-a-comment-legacy

  teams create-discussion-in-org --token=STRING --body=STRING --org=STRING --team_slug=STRING --title=STRING
    Create a discussion -
    https://developer.github.com/v3/teams/discussions/#create-a-discussion

  teams create-discussion-legacy --token=STRING --body=STRING --team_id=INT-64 --title=STRING
    Create a discussion (Legacy) -
    https://developer.github.com/v3/teams/discussions/#create-a-discussion-legacy

  teams delete-discussion-comment-in-org --token=STRING --comment_number=INT-64 --discussion_number=INT-64 --org=STRING --team_slug=STRING
    Delete a comment -
    https://developer.github.com/v3/teams/discussion_comments/#delete-a-comment

  teams delete-discussion-comment-legacy --token=STRING --comment_number=INT-64 --discussion_number=INT-64 --team_id=INT-64
    Delete a comment (Legacy) -
    https://developer.github.com/v3/teams/discussion_comments/#delete-a-comment-legacy

  teams delete-discussion-in-org --token=STRING --discussion_number=INT-64 --org=STRING --team_slug=STRING
    Delete a discussion -
    https://developer.github.com/v3/teams/discussions/#delete-a-discussion

  teams delete-discussion-legacy --token=STRING --discussion_number=INT-64 --team_id=INT-64
    Delete a discussion (Legacy) -
    https://developer.github.com/v3/teams/discussions/#delete-a-discussion-legacy

  teams delete-in-org --token=STRING --org=STRING --team_slug=STRING
    Delete team - https://developer.github.com/v3/teams/#delete-team

  teams delete-legacy --token=STRING --team_id=INT-64
    Delete team (Legacy) -
    https://developer.github.com/v3/teams/#delete-team-legacy

  teams get-by-name --token=STRING --org=STRING --team_slug=STRING
    Get team by name - https://developer.github.com/v3/teams/#get-team-by-name

  teams get-discussion-comment-in-org --token=STRING --comment_number=INT-64 --discussion_number=INT-64 --org=STRING --team_slug=STRING
    Get a single comment -
    https://developer.github.com/v3/teams/discussion_comments/#get-a-single-comment

  teams get-discussion-comment-legacy --token=STRING --comment_number=INT-64 --discussion_number=INT-64 --team_id=INT-64
    Get a single comment (Legacy) -
    https://developer.github.com/v3/teams/discussion_comments/#get-a-single-comment-legacy

  teams get-discussion-in-org --token=STRING --discussion_number=INT-64 --org=STRING --team_slug=STRING
    Get a single discussion -
    https://developer.github.com/v3/teams/discussions/#get-a-single-discussion

  teams get-discussion-legacy --token=STRING --discussion_number=INT-64 --team_id=INT-64
    Get a single discussion (Legacy) -
    https://developer.github.com/v3/teams/discussions/#get-a-single-discussion-legacy

  teams get-legacy --token=STRING --team_id=INT-64
    Get team (Legacy) - https://developer.github.com/v3/teams/#get-team-legacy

  teams get-member-legacy --token=STRING --team_id=INT-64 --username=STRING
    Get team member (Legacy) -
    https://developer.github.com/v3/teams/members/#get-team-member-legacy

  teams get-membership-in-org --token=STRING --org=STRING --team_slug=STRING --username=STRING
    Get team membership -
    https://developer.github.com/v3/teams/members/#get-team-membership

  teams get-membership-legacy --token=STRING --team_id=INT-64 --username=STRING
    Get team membership (Legacy) -
    https://developer.github.com/v3/teams/members/#get-team-membership-legacy

  teams list --token=STRING --org=STRING
    List teams - https://developer.github.com/v3/teams/#list-teams

  teams list-child-in-org --token=STRING --org=STRING --team_slug=STRING
    List child teams - https://developer.github.com/v3/teams/#list-child-teams

  teams list-child-legacy --token=STRING --team_id=INT-64
    List child teams (Legacy) -
    https://developer.github.com/v3/teams/#list-child-teams-legacy

  teams list-discussion-comments-in-org --token=STRING --discussion_number=INT-64 --org=STRING --team_slug=STRING
    List comments -
    https://developer.github.com/v3/teams/discussion_comments/#list-comments

  teams list-discussion-comments-legacy --token=STRING --discussion_number=INT-64 --team_id=INT-64
    List comments (Legacy) -
    https://developer.github.com/v3/teams/discussion_comments/#list-comments-legacy

  teams list-discussions-in-org --token=STRING --org=STRING --team_slug=STRING
    List discussions -
    https://developer.github.com/v3/teams/discussions/#list-discussions

  teams list-discussions-legacy --token=STRING --team_id=INT-64
    List discussions (Legacy) -
    https://developer.github.com/v3/teams/discussions/#list-discussions-legacy

  teams list-for-authenticated-user --token=STRING
    List user teams - https://developer.github.com/v3/teams/#list-user-teams

  teams list-id-p-groups-for-legacy --token=STRING --team_id=INT-64
    List IdP groups for a team (Legacy) -
    https://developer.github.com/v3/teams/team_sync/#list-idp-groups-for-a-team-legacy

  teams list-id-p-groups-for-org --token=STRING --org=STRING
    List IdP groups in an organization -
    https://developer.github.com/v3/teams/team_sync/#list-idp-groups-in-an-organization

  teams list-id-p-groups-in-org --token=STRING --org=STRING --team_slug=STRING
    List IdP groups for a team -
    https://developer.github.com/v3/teams/team_sync/#list-idp-groups-for-a-team

  teams list-members-in-org --token=STRING --org=STRING --team_slug=STRING
    List team members -
    https://developer.github.com/v3/teams/members/#list-team-members

  teams list-members-legacy --token=STRING --team_id=INT-64
    List team members (Legacy) -
    https://developer.github.com/v3/teams/members/#list-team-members-legacy

  teams list-pending-invitations-in-org --token=STRING --org=STRING --team_slug=STRING
    List pending team invitations -
    https://developer.github.com/v3/teams/members/#list-pending-team-invitations

  teams list-pending-invitations-legacy --token=STRING --team_id=INT-64
    List pending team invitations (Legacy) -
    https://developer.github.com/v3/teams/members/#list-pending-team-invitations-legacy

  teams list-projects-in-org --token=STRING --inertia-preview --org=STRING --team_slug=STRING
    List team projects -
    https://developer.github.com/v3/teams/#list-team-projects

  teams list-projects-legacy --token=STRING --inertia-preview --team_id=INT-64
    List team projects (Legacy) -
    https://developer.github.com/v3/teams/#list-team-projects-legacy

  teams list-repos-in-org --token=STRING --org=STRING --team_slug=STRING
    List team repos - https://developer.github.com/v3/teams/#list-team-repos

  teams list-repos-legacy --token=STRING --team_id=INT-64
    List team repos (Legacy) -
    https://developer.github.com/v3/teams/#list-team-repos-legacy

  teams remove-member-legacy --token=STRING --team_id=INT-64 --username=STRING
    Remove team member (Legacy) -
    https://developer.github.com/v3/teams/members/#remove-team-member-legacy

  teams remove-membership-in-org --token=STRING --org=STRING --team_slug=STRING --username=STRING
    Remove team membership -
    https://developer.github.com/v3/teams/members/#remove-team-membership

  teams remove-membership-legacy --token=STRING --team_id=INT-64 --username=STRING
    Remove team membership (Legacy) -
    https://developer.github.com/v3/teams/members/#remove-team-membership-legacy

  teams remove-project-in-org --token=STRING --org=STRING --project_id=INT-64 --team_slug=STRING
    Remove team project -
    https://developer.github.com/v3/teams/#remove-team-project

  teams remove-project-legacy --token=STRING --project_id=INT-64 --team_id=INT-64
    Remove team project (Legacy) -
    https://developer.github.com/v3/teams/#remove-team-project-legacy

  teams remove-repo-in-org --token=STRING --org=STRING --repo=STRING --team_slug=STRING
    Remove team repository -
    https://developer.github.com/v3/teams/#remove-team-repository

  teams remove-repo-legacy --token=STRING --repo=STRING --team_id=INT-64
    Remove team repository (Legacy) -
    https://developer.github.com/v3/teams/#remove-team-repository-legacy

  teams review-project-in-org --token=STRING --inertia-preview --org=STRING --project_id=INT-64 --team_slug=STRING
    Review a team project -
    https://developer.github.com/v3/teams/#review-a-team-project

  teams review-project-legacy --token=STRING --inertia-preview --project_id=INT-64 --team_id=INT-64
    Review a team project (Legacy) -
    https://developer.github.com/v3/teams/#review-a-team-project-legacy

  teams update-discussion-comment-in-org --token=STRING --body=STRING --comment_number=INT-64 --discussion_number=INT-64 --org=STRING --team_slug=STRING
    Edit a comment -
    https://developer.github.com/v3/teams/discussion_comments/#edit-a-comment

  teams update-discussion-comment-legacy --token=STRING --body=STRING --comment_number=INT-64 --discussion_number=INT-64 --team_id=INT-64
    Edit a comment (Legacy) -
    https://developer.github.com/v3/teams/discussion_comments/#edit-a-comment-legacy

  teams update-discussion-in-org --token=STRING --discussion_number=INT-64 --org=STRING --team_slug=STRING
    Edit a discussion -
    https://developer.github.com/v3/teams/discussions/#edit-a-discussion

  teams update-discussion-legacy --token=STRING --discussion_number=INT-64 --team_id=INT-64
    Edit a discussion (Legacy) -
    https://developer.github.com/v3/teams/discussions/#edit-a-discussion-legacy

  teams update-in-org --token=STRING --name=STRING --org=STRING --team_slug=STRING
    Edit team - https://developer.github.com/v3/teams/#edit-team

  teams update-legacy --token=STRING --name=STRING --team_id=INT-64
    Edit team (Legacy) - https://developer.github.com/v3/teams/#edit-team-legacy

  users add-emails --token=STRING --emails=EMAILS,...
    Add email address(es) -
    https://developer.github.com/v3/users/emails/#add-email-addresses

  users block --token=STRING --username=STRING
    Block a user - https://developer.github.com/v3/users/blocking/#block-a-user

  users check-blocked --token=STRING --username=STRING
    Check whether you've blocked a user -
    https://developer.github.com/v3/users/blocking/#check-whether-youve-blocked-a-user

  users check-following --token=STRING --username=STRING
    Check if you are following a user -
    https://developer.github.com/v3/users/followers/#check-if-you-are-following-a-user

  users check-following-for-user --token=STRING --target_user=STRING --username=STRING
    Check if one user follows another -
    https://developer.github.com/v3/users/followers/#check-if-one-user-follows-another

  users create-gpg-key --token=STRING
    Create a GPG key -
    https://developer.github.com/v3/users/gpg_keys/#create-a-gpg-key

  users create-public-key --token=STRING
    Create a public key -
    https://developer.github.com/v3/users/keys/#create-a-public-key

  users delete-emails --token=STRING --emails=EMAILS,...
    Delete email address(es) -
    https://developer.github.com/v3/users/emails/#delete-email-addresses

  users delete-gpg-key --token=STRING --gpg_key_id=INT-64
    Delete a GPG key -
    https://developer.github.com/v3/users/gpg_keys/#delete-a-gpg-key

  users delete-public-key --token=STRING --key_id=INT-64
    Delete a public key -
    https://developer.github.com/v3/users/keys/#delete-a-public-key

  users follow --token=STRING --username=STRING
    Follow a user -
    https://developer.github.com/v3/users/followers/#follow-a-user

  users get-authenticated --token=STRING
    Get the authenticated user -
    https://developer.github.com/v3/users/#get-the-authenticated-user

  users get-by-username --token=STRING --username=STRING
    Get a single user - https://developer.github.com/v3/users/#get-a-single-user

  users get-context-for-user --token=STRING --username=STRING
    Get contextual information about a user -
    https://developer.github.com/v3/users/#get-contextual-information-about-a-user

  users get-gpg-key --token=STRING --gpg_key_id=INT-64
    Get a single GPG key -
    https://developer.github.com/v3/users/gpg_keys/#get-a-single-gpg-key

  users get-public-key --token=STRING --key_id=INT-64
    Get a single public key -
    https://developer.github.com/v3/users/keys/#get-a-single-public-key

  users list --token=STRING
    Get all users - https://developer.github.com/v3/users/#get-all-users

  users list-blocked --token=STRING
    List blocked users -
    https://developer.github.com/v3/users/blocking/#list-blocked-users

  users list-emails --token=STRING
    List email addresses for a user -
    https://developer.github.com/v3/users/emails/#list-email-addresses-for-a-user

  users list-followed-by-authenticated --token=STRING
    List users followed by the authenticated user -
    https://developer.github.com/v3/users/followers/#list-users-followed-by-the-authenticated-user

  users list-followers-for-authenticated-user --token=STRING
    List followers of the authenticated user -
    https://developer.github.com/v3/users/followers/#list-followers-of-the-authenticated-user

  users list-followers-for-user --token=STRING --username=STRING
    List followers of a user -
    https://developer.github.com/v3/users/followers/#list-followers-of-a-user

  users list-following-for-user --token=STRING --username=STRING
    List users followed by another user -
    https://developer.github.com/v3/users/followers/#list-users-followed-by-another-user

  users list-gpg-keys --token=STRING
    List your GPG keys -
    https://developer.github.com/v3/users/gpg_keys/#list-your-gpg-keys

  users list-gpg-keys-for-user --token=STRING --username=STRING
    List GPG keys for a user -
    https://developer.github.com/v3/users/gpg_keys/#list-gpg-keys-for-a-user

  users list-public-emails --token=STRING
    List public email addresses for a user -
    https://developer.github.com/v3/users/emails/#list-public-email-addresses-for-a-user

  users list-public-keys --token=STRING
    List your public keys -
    https://developer.github.com/v3/users/keys/#list-your-public-keys

  users list-public-keys-for-user --token=STRING --username=STRING
    List public keys for a user -
    https://developer.github.com/v3/users/keys/#list-public-keys-for-a-user

  users toggle-primary-email-visibility --token=STRING --email=STRING --visibility=STRING
    Toggle primary email visibility -
    https://developer.github.com/v3/users/emails/#toggle-primary-email-visibility

  users unblock --token=STRING --username=STRING
    Unblock a user -
    https://developer.github.com/v3/users/blocking/#unblock-a-user

  users unfollow --token=STRING --username=STRING
    Unfollow a user -
    https://developer.github.com/v3/users/followers/#unfollow-a-user

  users update-authenticated --token=STRING
    Update the authenticated user -
    https://developer.github.com/v3/users/#update-the-authenticated-user

Run "octo <command> --help" for more information on a command.

```
<!--- END HELP OUTPUT --->
