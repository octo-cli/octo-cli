# octo-cli

`octo-cli` is a cli client for GitHub's REST API.  It is generated
by inspecting https://octokit.github.io/routes/.

See ./generator for more on how it is generated.

## Work In Progress

 octo-cli is a work in progress.  Use it with
caution.

## Usage

You can set the environment variable GITHUB_TOKEN instead of using a `--token` flag to
avoid putting credentials on the command line.


```
Usage: octo-cli <command>

Flags:
  --help    Show context-sensitive help.

Commands:
  activity list-public-events --token=STRING
    List public events

  activity list-repo-events --token=STRING --owner=STRING --repo=STRING
    List repository events

  activity list-public-events-for-repo-network --token=STRING --owner=STRING --repo=STRING
    List public events for a network of repositories

  activity list-public-events-for-org --token=STRING --org=STRING
    List public events for an organization

  activity list-received-events-for-user --token=STRING --username=STRING
    List events that a user has received

  activity list-received-public-events-for-user --token=STRING --username=STRING
    List public events that a user has received

  activity list-events-for-user --token=STRING --username=STRING
    List events performed by a user

  activity list-public-events-for-user --token=STRING --username=STRING
    List public events performed by a user

  activity list-events-for-org --token=STRING --username=STRING --org=STRING
    List events for an organization

  activity list-feeds --token=STRING
    List feeds

  activity list-notifications --token=STRING
    List your notifications

  activity list-notifications-for-repo --token=STRING --owner=STRING --repo=STRING
    List your notifications in a repository

  activity mark-as-read --token=STRING
    Mark as read

  activity mark-notifications-as-read-for-repo --token=STRING --owner=STRING --repo=STRING
    Mark notifications as read in a repository

  activity get-thread --token=STRING --thread_id=INT-64
    View a single thread

  activity mark-thread-as-read --token=STRING --thread_id=INT-64
    Mark a thread as read

  activity get-thread-subscription --token=STRING --thread_id=INT-64
    Get a thread subscription

  activity set-thread-subscription --token=STRING --thread_id=INT-64
    Set a thread subscription

  activity delete-thread-subscription --token=STRING --thread_id=INT-64
    Delete a thread subscription

  activity list-stargazers-for-repo --token=STRING --owner=STRING --repo=STRING
    List Stargazers

  activity list-repos-starred-by-user --token=STRING --username=STRING
    List repositories being starred by a user

  activity list-repos-starred-by-authenticated-user --token=STRING
    List repositories being starred by the authenticated user

  activity check-starring-repo --token=STRING --owner=STRING --repo=STRING
    Check if you are starring a repository

  activity star-repo --token=STRING --owner=STRING --repo=STRING
    Star a repository

  activity unstar-repo --token=STRING --owner=STRING --repo=STRING
    Unstar a repository

  activity list-watchers-for-repo --token=STRING --owner=STRING --repo=STRING
    List watchers

  activity list-repos-watched-by-user --token=STRING --username=STRING
    List repositories being watched by a user

  activity list-watched-repos-for-authenticated-user --token=STRING
    List repositories being watched by the authenticated user

  activity get-repo-subscription --token=STRING --owner=STRING --repo=STRING
    Get a Repository Subscription

  activity set-repo-subscription --token=STRING --owner=STRING --repo=STRING
    Set a Repository Subscription

  activity delete-repo-subscription --token=STRING --owner=STRING --repo=STRING
    Delete a Repository Subscription

  activity check-watching-repo-legacy --token=STRING --owner=STRING --repo=STRING
    Check if you are watching a repository (LEGACY)

  activity watch-repo-legacy --token=STRING --owner=STRING --repo=STRING
    Watch a repository (LEGACY)

  activity stop-watching-repo-legacy --token=STRING --owner=STRING --repo=STRING
    Stop watching a repository (LEGACY)

  apps get-by-slug --token=STRING --app_slug=STRING
    Get a single GitHub App

  apps get-authenticated --token=STRING --machine-man-preview
    Get the authenticated GitHub App

  apps list-installations --token=STRING --machine-man-preview
    Find installations

  apps get-installation --token=STRING --machine-man-preview --installation_id=INT-64
    Get a single installation

  apps list-installations-for-authenticated-user --token=STRING --machine-man-preview
    List installations for user

  apps create-installation-token --token=STRING --machine-man-preview --installation_id=INT-64
    Create a new installation token

  apps find-org-installation --token=STRING --machine-man-preview --org=STRING
    Find organization installation

  apps find-repo-installation --token=STRING --machine-man-preview --owner=STRING --repo=STRING
    Find repository installation

  apps find-user-installation --token=STRING --machine-man-preview --username=STRING
    Find user installation

  apps create-from-manifest --token=STRING --fury-preview --code=STRING
    Create a GitHub App from a manifest

  apps list-repos --token=STRING
    List repositories

  apps list-installation-repos-for-authenticated-user --token=STRING --machine-man-preview --installation_id=INT-64
    List repositories accessible to the user for an installation

  apps add-repo-to-installation --token=STRING --machine-man-preview --installation_id=INT-64 --repository_id=INT-64
    Add repository to installation

  apps remove-repo-from-installation --token=STRING --machine-man-preview --installation_id=INT-64 --repository_id=INT-64
    Remove repository from installation

  apps list-plans --token=STRING
    List all plans for your Marketplace listing

  apps list-plans-stubbed --token=STRING
    List all plans for your Marketplace listing (stubbed)

  apps list-accounts-user-or-org-on-plan --token=STRING --plan_id=INT-64
    List all GitHub accounts (user or organization) on a specific plan

  apps list-accounts-user-or-org-on-plan-stubbed --token=STRING --plan_id=INT-64
    List all GitHub accounts (user or organization) on a specific plan (stubbed)

  apps check-account-is-associated-with-any --token=STRING --account_id=INT-64
    Check if a GitHub account is associated with any Marketplace listing

  apps check-account-is-associated-with-any-stubbed --token=STRING --account_id=INT-64
    Check if a GitHub account is associated with any Marketplace listing
    (stubbed)

  apps list-marketplace-purchases-for-authenticated-user --token=STRING
    Get a user's Marketplace purchases

  apps list-marketplace-purchases-for-authenticated-user-stubbed --token=STRING
    Get a user's Marketplace purchases (stubbed)

  checks list-for-ref --token=STRING --antiope-preview --owner=STRING --repo=STRING --ref=STRING
    List check runs for a specific ref

  checks list-for-suite --token=STRING --antiope-preview --owner=STRING --repo=STRING --check_suite_id=INT-64
    List check runs in a check suite

  checks get --token=STRING --antiope-preview --owner=STRING --repo=STRING --check_run_id=INT-64
    Get a single check run

  checks list-annotations --token=STRING --antiope-preview --owner=STRING --repo=STRING --check_run_id=INT-64
    List annotations for a check run

  checks get-suite --token=STRING --antiope-preview --owner=STRING --repo=STRING --check_suite_id=INT-64
    Get a single check suite

  checks list-suites-for-ref --token=STRING --antiope-preview --owner=STRING --repo=STRING --ref=STRING
    List check suites for a specific ref

  checks create-suite --token=STRING --antiope-preview --owner=STRING --repo=STRING --head_sha=STRING
    Create a check suite

  checks rerequest-suite --token=STRING --antiope-preview --owner=STRING --repo=STRING --check_suite_id=INT-64
    Rerequest check suite

  codes-of-conduct list-conduct-codes --token=STRING
    List all codes of conduct

  codes-of-conduct get-conduct-code --token=STRING --key=STRING
    Get an individual code of conduct

  codes-of-conduct get-for-repo --token=STRING --owner=STRING --repo=STRING
    Get the contents of a repository's code of conduct

  emojis get --token=STRING
    Get

  gists list-public-for-user --token=STRING --username=STRING
    List public gists for the specified user

  gists list --token=STRING
    List the authenticated user's gists or if called anonymously, this will
    return all public gists

  gists list-public --token=STRING
    List all public gists

  gists list-starred --token=STRING
    List starred gists

  gists get --token=STRING --gist_id=STRING
    Get a single gist

  gists get-revision --token=STRING --gist_id=STRING --sha=STRING
    Get a specific revision of a gist

  gists list-commits --token=STRING --gist_id=STRING
    List gist commits

  gists star --token=STRING --gist_id=STRING
    Star a gist

  gists unstar --token=STRING --gist_id=STRING
    Unstar a gist

  gists check-is-starred --token=STRING --gist_id=STRING
    Check if a gist is starred

  gists fork --token=STRING --gist_id=STRING
    Fork a gist

  gists list-forks --token=STRING --gist_id=STRING
    List gist forks

  gists delete --token=STRING --gist_id=STRING
    Delete a gist

  gists list-comments --token=STRING --gist_id=STRING
    List comments on a gist

  gists get-comment --token=STRING --gist_id=STRING --comment_id=INT-64
    Get a single comment

  gists create-comment --token=STRING --gist_id=STRING --body=STRING
    Create a comment

  gists edit-comment --token=STRING --gist_id=STRING --comment_id=INT-64 --body=STRING
    Edit a comment

  gists delete-comment --token=STRING --gist_id=STRING --comment_id=INT-64
    Delete a comment

  git get-blob --token=STRING --owner=STRING --repo=STRING --file_sha=STRING
    Get a blob

  git create-blob --token=STRING --owner=STRING --repo=STRING --content=STRING
    Create a blob

  git get-commit --token=STRING --owner=STRING --repo=STRING --commit_sha=STRING
    Get a commit

  git get-ref --token=STRING --owner=STRING --repo=STRING --ref=STRING
    Get a reference

  git list-refs --token=STRING --owner=STRING --repo=STRING
    Get all references

  git create-ref --token=STRING --owner=STRING --repo=STRING --ref=STRING --sha=STRING
    Create a reference

  git update-ref --token=STRING --owner=STRING --repo=STRING --ref=STRING --sha=STRING
    Update a reference

  git delete-ref --token=STRING --owner=STRING --repo=STRING --ref=STRING
    Delete a reference

  git get-tag --token=STRING --owner=STRING --repo=STRING --tag_sha=STRING
    Get a tag

  git get-tree --token=STRING --owner=STRING --repo=STRING --tree_sha=STRING
    Get a tree

  gitignore list-templates --token=STRING
    Listing available templates

  gitignore get-template --token=STRING --name=STRING
    Get a single template

  issues list --token=STRING
    List all issues assigned to the authenticated user across all visible
    repositories including owned repositories, member repositories, and
    organization repositories

  issues list-for-authenticated-user --token=STRING
    List all issues across owned and member repositories assigned to the
    authenticated user

  issues list-for-org --token=STRING --org=STRING
    List all issues for a given organization assigned to the authenticated user

  issues list-for-repo --token=STRING --owner=STRING --repo=STRING
    List issues for a repository

  issues get --token=STRING --owner=STRING --repo=STRING --number=INT-64
    Get a single issue

  issues create --token=STRING --owner=STRING --repo=STRING --title=STRING
    Create an issue

  issues edit --token=STRING --owner=STRING --repo=STRING --number=INT-64
    Edit an issue

  issues lock --token=STRING --owner=STRING --repo=STRING --number=INT-64
    Lock an issue

  issues unlock --token=STRING --owner=STRING --repo=STRING --number=INT-64
    Unlock an issue

  issues list-assignees --token=STRING --owner=STRING --repo=STRING
    List assignees

  issues check-assignee --token=STRING --owner=STRING --repo=STRING --assignee=STRING
    Check assignee

  issues add-assignees --token=STRING --owner=STRING --repo=STRING --number=INT-64
    Add assignees to an issue

  issues remove-assignees --token=STRING --owner=STRING --repo=STRING --number=INT-64
    Remove assignees from an issue

  issues list-comments --token=STRING --owner=STRING --repo=STRING --number=INT-64
    List comments on an issue

  issues list-comments-for-repo --token=STRING --owner=STRING --repo=STRING
    List comments in a repository

  issues get-comment --token=STRING --owner=STRING --repo=STRING --comment_id=INT-64
    Get a single comment

  issues create-comment --token=STRING --owner=STRING --repo=STRING --number=INT-64 --body=STRING
    Create a comment

  issues edit-comment --token=STRING --owner=STRING --repo=STRING --comment_id=INT-64 --body=STRING
    Edit a comment

  issues delete-comment --token=STRING --owner=STRING --repo=STRING --comment_id=INT-64
    Delete a comment

  issues list-events --token=STRING --owner=STRING --repo=STRING --number=INT-64
    List events for an issue

  issues list-events-for-repo --token=STRING --owner=STRING --repo=STRING
    List events for a repository

  issues get-event --token=STRING --owner=STRING --repo=STRING --event_id=INT-64
    Get a single event

  issues list-labels-for-repo --token=STRING --owner=STRING --repo=STRING
    List all labels for this repository

  issues get-label --token=STRING --owner=STRING --repo=STRING --name=STRING
    Get a single label

  issues create-label --token=STRING --owner=STRING --repo=STRING --name=STRING --color=STRING
    Create a label

  issues update-label --token=STRING --owner=STRING --repo=STRING --current_name=STRING
    Update a label

  issues delete-label --token=STRING --owner=STRING --repo=STRING --name=STRING
    Delete a label

  issues list-labels-on-issue --token=STRING --owner=STRING --repo=STRING --number=INT-64
    List labels on an issue

  issues add-labels --token=STRING --owner=STRING --repo=STRING --number=INT-64 --labels=LABELS,...
    Add labels to an issue

  issues remove-label --token=STRING --owner=STRING --repo=STRING --number=INT-64 --name=STRING
    Remove a label from an issue

  issues replace-labels --token=STRING --owner=STRING --repo=STRING --number=INT-64 --labels=LABELS,...
    Replace all labels for an issue

  issues remove-labels --token=STRING --owner=STRING --repo=STRING --number=INT-64
    Remove all labels from an issue

  issues list-labels-for-milestone --token=STRING --owner=STRING --repo=STRING --number=INT-64
    Get labels for every issue in a milestone

  issues list-milestones-for-repo --token=STRING --owner=STRING --repo=STRING
    List milestones for a repository

  issues get-milestone --token=STRING --owner=STRING --repo=STRING --number=INT-64
    Get a single milestone

  issues create-milestone --token=STRING --owner=STRING --repo=STRING --title=STRING
    Create a milestone

  issues update-milestone --token=STRING --owner=STRING --repo=STRING --number=INT-64
    Update a milestone

  issues delete-milestone --token=STRING --owner=STRING --repo=STRING --number=INT-64
    Delete a milestone

  issues list-events-for-timeline --token=STRING --mockingbird-preview --owner=STRING --repo=STRING --number=INT-64
    List events for an issue

  licenses list --token=STRING
    List all licenses

  licenses get --token=STRING --license=STRING
    Get an individual license

  licenses get-for-repo --token=STRING --owner=STRING --repo=STRING
    Get the contents of a repository's license

  markdown render --token=STRING --text=STRING
    Render an arbitrary Markdown document

  markdown render-raw --token=STRING --data=STRING
    Render a Markdown document in raw mode

  meta get --token=STRING
    Get

  migrations start-for-org --token=STRING --org=STRING --repositories=REPOSITORIES,...
    Start an organization migration

  migrations list-for-org --token=STRING --org=STRING
    Get a list of organization migrations

  migrations get-status-for-org --token=STRING --org=STRING --migration_id=INT-64
    Get the status of an organization migration

  migrations get-archive-for-org --token=STRING --org=STRING --migration_id=INT-64
    Download an organization migration archive

  migrations delete-archive-for-org --token=STRING --org=STRING --migration_id=INT-64
    Delete an organization migration archive

  migrations unlock-repo-for-org --token=STRING --org=STRING --migration_id=INT-64 --repo_name=STRING
    Unlock an organization repository

  migrations start-import --token=STRING --barred-rock-preview --owner=STRING --repo=STRING --vcs_url=STRING
    Start an import

  migrations get-import-progress --token=STRING --barred-rock-preview --owner=STRING --repo=STRING
    Get import progress

  migrations update-import --token=STRING --barred-rock-preview --owner=STRING --repo=STRING
    Update existing import

  migrations get-commit-authors --token=STRING --barred-rock-preview --owner=STRING --repo=STRING
    Get commit authors

  migrations map-commit-author --token=STRING --barred-rock-preview --owner=STRING --repo=STRING --author_id=INT-64
    Map a commit author

  migrations set-lfs-preference --token=STRING --barred-rock-preview --owner=STRING --repo=STRING --use_lfs=STRING
    Set Git LFS preference

  migrations get-large-files --token=STRING --barred-rock-preview --owner=STRING --repo=STRING
    Get large files

  migrations cancel-import --token=STRING --barred-rock-preview --owner=STRING --repo=STRING
    Cancel an import

  migrations start-for-authenticated-user --token=STRING --repositories=REPOSITORIES,...
    Start a user migration

  migrations list-for-authenticated-user --token=STRING
    Get a list of user migrations

  migrations get-status-for-authenticated-user --token=STRING --migration_id=INT-64
    Get the status of a user migration

  migrations get-archive-for-authenticated-user --token=STRING --migration_id=INT-64
    Download a user migration archive

  migrations delete-archive-for-authenticated-user --token=STRING --migration_id=INT-64
    Delete a user migration archive

  migrations unlock-repo-for-authenticated-user --token=STRING --migration_id=INT-64 --repo_name=STRING
    Unlock a user repository

  oauth-authorizations list-grants --token=STRING
    List your grants

  oauth-authorizations get-grant --token=STRING --grant_id=INT-64
    Get a single grant

  oauth-authorizations delete-grant --token=STRING --grant_id=INT-64
    Delete a grant

  oauth-authorizations list-authorizations --token=STRING
    List your authorizations

  oauth-authorizations get-authorization --token=STRING --authorization_id=INT-64
    Get a single authorization

  oauth-authorizations create-authorization --token=STRING --note=STRING
    Create a new authorization

  oauth-authorizations get-or-create-authorization-for-app --token=STRING --client_id=STRING --client_secret=STRING
    Get-or-create an authorization for a specific app

  oauth-authorizations get-or-create-authorization-for-app-fingerprint --token=STRING --client_id=STRING --fingerprint=STRING --client_secret=STRING
    Get-or-create an authorization for a specific app and fingerprint

  oauth-authorizations update-authorization --token=STRING --authorization_id=INT-64
    Update an existing authorization

  oauth-authorizations delete-authorization --token=STRING --authorization_id=INT-64
    Delete an authorization

  oauth-authorizations check-authorization --token=STRING --client_id=STRING --access_token=STRING
    Check an authorization

  oauth-authorizations reset-authorization --token=STRING --client_id=STRING --access_token=STRING
    Reset an authorization

  oauth-authorizations revoke-authorization-for-application --token=STRING --client_id=STRING --access_token=STRING
    Revoke an authorization for an application

  oauth-authorizations revoke-grant-for-application --token=STRING --client_id=STRING --access_token=STRING
    Revoke a grant for an application

  orgs list-for-current-user --token=STRING
    List your organizations

  orgs list --token=STRING
    List all organizations

  orgs list-for-user --token=STRING --username=STRING
    List user organizations

  orgs get --token=STRING --org=STRING
    Get an organization

  orgs edit --token=STRING --org=STRING
    Edit an organization

  orgs list-blocked-users --token=STRING --org=STRING
    List blocked users

  orgs check-blocked-user --token=STRING --org=STRING --username=STRING
    Check whether a user is blocked from an organization

  orgs block-user --token=STRING --org=STRING --username=STRING
    Block a user

  orgs unblock-user --token=STRING --org=STRING --username=STRING
    Unblock a user

  orgs list-members --token=STRING --org=STRING
    Members list

  orgs check-membership --token=STRING --org=STRING --username=STRING
    Check membership

  orgs remove-member --token=STRING --org=STRING --username=STRING
    Remove a member

  orgs list-public-members --token=STRING --org=STRING
    Public members list

  orgs check-public-membership --token=STRING --org=STRING --username=STRING
    Check public membership

  orgs publicize-membership --token=STRING --org=STRING --username=STRING
    Publicize a user's membership

  orgs conceal-membership --token=STRING --org=STRING --username=STRING
    Conceal a user's membership

  orgs get-membership-for-user --token=STRING --org=STRING --username=STRING
    Get organization membership

  orgs add-or-update-membership --token=STRING --org=STRING --username=STRING
    Add or update organization membership

  orgs remove-membership --token=STRING --org=STRING --username=STRING
    Remove organization membership

  orgs list-invitation-teams --token=STRING --org=STRING --invitation_id=INT-64
    List organization invitation teams

  orgs list-pending-invitations --token=STRING --org=STRING
    List pending organization invitations

  orgs create-invitation --token=STRING --org=STRING
    Create organization invitation

  orgs list-memberships --token=STRING
    List your organization memberships

  orgs get-membership --token=STRING --org=STRING
    Get your organization membership

  orgs edit-membership --token=STRING --org=STRING --state=STRING
    Edit your organization membership

  orgs list-outside-collaborators --token=STRING --org=STRING
    List outside collaborators

  orgs remove-outside-collaborator --token=STRING --org=STRING --username=STRING
    Remove outside collaborator

  orgs convert-member-to-outside-collaborator --token=STRING --org=STRING --username=STRING
    Convert member to outside collaborator

  orgs list-hooks --token=STRING --org=STRING
    List hooks

  orgs get-hook --token=STRING --org=STRING --hook_id=INT-64
    Get single hook

  orgs ping-hook --token=STRING --org=STRING --hook_id=INT-64
    Ping a hook

  orgs delete-hook --token=STRING --org=STRING --hook_id=INT-64
    Delete a hook

  projects list-for-repo --token=STRING --inertia-preview --owner=STRING --repo=STRING
    List repository projects

  projects list-for-org --token=STRING --inertia-preview --org=STRING
    List organization projects

  projects get --token=STRING --inertia-preview --project_id=INT-64
    Get a project

  projects create-for-repo --token=STRING --inertia-preview --owner=STRING --repo=STRING --name=STRING
    Create a repository project

  projects create-for-org --token=STRING --inertia-preview --org=STRING --name=STRING
    Create an organization project

  projects update --token=STRING --inertia-preview --project_id=INT-64
    Update a project

  projects delete --token=STRING --inertia-preview --project_id=INT-64
    Delete a project

  projects list-cards --token=STRING --inertia-preview --column_id=INT-64
    List project cards

  projects get-card --token=STRING --inertia-preview --card_id=INT-64
    Get a project card

  projects create-card --token=STRING --inertia-preview --column_id=INT-64
    Create a project card

  projects update-card --token=STRING --inertia-preview --card_id=INT-64
    Update a project card

  projects delete-card --token=STRING --inertia-preview --card_id=INT-64
    Delete a project card

  projects move-card --token=STRING --inertia-preview --card_id=INT-64 --position=STRING
    Move a project card

  projects list-collaborators --token=STRING --inertia-preview --project_id=INT-64
    List collaborators

  projects review-user-permission-level --token=STRING --inertia-preview --project_id=INT-64 --username=STRING
    Review a user's permission level

  projects add-collaborator --token=STRING --inertia-preview --project_id=INT-64 --username=STRING
    Add user as a collaborator

  projects remove-collaborator --token=STRING --inertia-preview --project_id=INT-64 --username=STRING
    Remove user as a collaborator

  projects list-columns --token=STRING --inertia-preview --project_id=INT-64
    List project columns

  projects get-column --token=STRING --inertia-preview --column_id=INT-64
    Get a project column

  projects create-column --token=STRING --inertia-preview --project_id=INT-64 --name=STRING
    Create a project column

  projects update-column --token=STRING --inertia-preview --column_id=INT-64 --name=STRING
    Update a project column

  projects delete-column --token=STRING --inertia-preview --column_id=INT-64
    Delete a project column

  projects move-column --token=STRING --inertia-preview --column_id=INT-64 --position=STRING
    Move a project column

  pulls list --token=STRING --owner=STRING --repo=STRING
    List pull requests

  pulls get --token=STRING --owner=STRING --repo=STRING --number=INT-64
    Get a single pull request

  pulls create --token=STRING --owner=STRING --repo=STRING --title=STRING --head=STRING --base=STRING
    Create a pull request

  pulls create-from-issue --token=STRING --owner=STRING --repo=STRING --issue=INT-64 --head=STRING --base=STRING
    Create a Pull Request from an Issue

  pulls update --token=STRING --owner=STRING --repo=STRING --number=INT-64
    Update a pull request

  pulls list-commits --token=STRING --owner=STRING --repo=STRING --number=INT-64
    List commits on a pull request

  pulls list-files --token=STRING --owner=STRING --repo=STRING --number=INT-64
    List pull requests files

  pulls check-if-merged --token=STRING --owner=STRING --repo=STRING --number=INT-64
    Get if a pull request has been merged

  pulls merge --token=STRING --owner=STRING --repo=STRING --number=INT-64
    Merge a pull request (Merge Button)

  pulls list-reviews --token=STRING --owner=STRING --repo=STRING --number=INT-64
    List reviews on a pull request

  pulls get-review --token=STRING --owner=STRING --repo=STRING --number=INT-64 --review_id=INT-64
    Get a single review

  pulls delete-pending-review --token=STRING --owner=STRING --repo=STRING --number=INT-64 --review_id=INT-64
    Delete a pending review

  pulls get-comments-for-review --token=STRING --owner=STRING --repo=STRING --number=INT-64 --review_id=INT-64
    Get comments for a single review

  pulls submit-review --token=STRING --owner=STRING --repo=STRING --number=INT-64 --review_id=INT-64 --event=STRING
    Submit a pull request review

  pulls dismiss-review --token=STRING --owner=STRING --repo=STRING --number=INT-64 --review_id=INT-64 --message=STRING
    Dismiss a pull request review

  pulls list-comments --token=STRING --owner=STRING --repo=STRING --number=INT-64
    List comments on a pull request

  pulls list-comments-for-repo --token=STRING --owner=STRING --repo=STRING
    List comments in a repository

  pulls get-comment --token=STRING --owner=STRING --repo=STRING --comment_id=INT-64
    Get a single comment

  pulls create-comment --token=STRING --owner=STRING --repo=STRING --number=INT-64 --body=STRING --commit_id=STRING --path=STRING --position=INT-64
    Create a comment

  pulls create-comment-reply --token=STRING --owner=STRING --repo=STRING --number=INT-64 --body=STRING --in_reply_to=INT-64
    Create a comment reply

  pulls edit-comment --token=STRING --owner=STRING --repo=STRING --comment_id=INT-64 --body=STRING
    Edit a comment

  pulls delete-comment --token=STRING --owner=STRING --repo=STRING --comment_id=INT-64
    Delete a comment

  pulls list-review-requests --token=STRING --owner=STRING --repo=STRING --number=INT-64
    List review requests

  pulls create-review-request --token=STRING --owner=STRING --repo=STRING --number=INT-64
    Create a review request

  pulls delete-review-request --token=STRING --owner=STRING --repo=STRING --number=INT-64
    Delete a review request

  rate-limit get --token=STRING
    Get your current rate limit status

  reactions list-for-commit-comment --token=STRING --squirrel-girl-preview --owner=STRING --repo=STRING --comment_id=INT-64
    List reactions for a commit comment

  reactions create-for-commit-comment --token=STRING --squirrel-girl-preview --owner=STRING --repo=STRING --comment_id=INT-64 --content=STRING
    Create reaction for a commit comment

  reactions list-for-issue --token=STRING --squirrel-girl-preview --owner=STRING --repo=STRING --number=INT-64
    List reactions for an issue

  reactions create-for-issue --token=STRING --squirrel-girl-preview --owner=STRING --repo=STRING --number=INT-64 --content=STRING
    Create reaction for an issue

  reactions list-for-issue-comment --token=STRING --squirrel-girl-preview --owner=STRING --repo=STRING --comment_id=INT-64
    List reactions for an issue comment

  reactions create-for-issue-comment --token=STRING --squirrel-girl-preview --owner=STRING --repo=STRING --comment_id=INT-64 --content=STRING
    Create reaction for an issue comment

  reactions list-for-pull-request-review-comment --token=STRING --squirrel-girl-preview --owner=STRING --repo=STRING --comment_id=INT-64
    List reactions for a pull request review comment

  reactions create-for-pull-request-review-comment --token=STRING --squirrel-girl-preview --owner=STRING --repo=STRING --comment_id=INT-64 --content=STRING
    Create reaction for a pull request review comment

  reactions list-for-team-discussion --token=STRING --echo-preview --squirrel-girl-preview --team_id=INT-64 --discussion_number=INT-64
    List reactions for a team discussion

  reactions create-for-team-discussion --token=STRING --echo-preview --squirrel-girl-preview --team_id=INT-64 --discussion_number=INT-64 --content=STRING
    Create reaction for a team discussion

  reactions list-for-team-discussion-comment --token=STRING --echo-preview --squirrel-girl-preview --team_id=INT-64 --discussion_number=INT-64 --comment_number=INT-64
    List reactions for a team discussion comment

  reactions create-for-team-discussion-comment --token=STRING --echo-preview --squirrel-girl-preview --team_id=INT-64 --discussion_number=INT-64 --comment_number=INT-64 --content=STRING
    Create reaction for a team discussion comment

  reactions delete --token=STRING --echo-preview --squirrel-girl-preview --reaction_id=INT-64
    Delete a reaction

  repos list --token=STRING
    List your repositories

  repos list-for-user --token=STRING --username=STRING
    List user repositories

  repos list-for-org --token=STRING --org=STRING
    List organization repositories

  repos list-public --token=STRING
    List all public repositories

  repos create-for-authenticated-user --token=STRING --name=STRING
    Create a new repository for the authenticated user

  repos create-in-org --token=STRING --org=STRING --name=STRING
    Create a new repository in this organization

  repos get --token=STRING --owner=STRING --repo=STRING
    Get

  repos edit --token=STRING --owner=STRING --repo=STRING --name=STRING
    Edit

  repos list-topics --token=STRING --owner=STRING --repo=STRING
    List all topics for a repository

  repos replace-topics --token=STRING --owner=STRING --repo=STRING --names=NAMES,...
    Replace all topics for a repository

  repos list-contributors --token=STRING --owner=STRING --repo=STRING
    List contributors

  repos list-languages --token=STRING --owner=STRING --repo=STRING
    List languages

  repos list-teams --token=STRING --owner=STRING --repo=STRING
    List teams

  repos list-tags --token=STRING --owner=STRING --repo=STRING
    List tags

  repos delete --token=STRING --owner=STRING --repo=STRING
    Delete a repository

  repos transfer --token=STRING --nightshade-preview --owner=STRING --repo=STRING
    Transfer a repository

  repos list-branches --token=STRING --owner=STRING --repo=STRING
    List branches

  repos get-branch --token=STRING --owner=STRING --repo=STRING --branch=STRING
    Get branch

  repos get-branch-protection --token=STRING --owner=STRING --repo=STRING --branch=STRING
    Get branch protection

  repos remove-branch-protection --token=STRING --owner=STRING --repo=STRING --branch=STRING
    Remove branch protection

  repos get-protected-branch-required-status-checks --token=STRING --owner=STRING --repo=STRING --branch=STRING
    Get required status checks of protected branch

  repos update-protected-branch-required-status-checks --token=STRING --owner=STRING --repo=STRING --branch=STRING
    Update required status checks of protected branch

  repos remove-protected-branch-required-status-checks --token=STRING --owner=STRING --repo=STRING --branch=STRING
    Remove required status checks of protected branch

  repos list-protected-branch-required-status-checks-contexts --token=STRING --owner=STRING --repo=STRING --branch=STRING
    List required status checks contexts of protected branch

  repos replace-protected-branch-required-status-checks-contexts --token=STRING --owner=STRING --repo=STRING --branch=STRING --contexts=CONTEXTS,...
    Replace required status checks contexts of protected branch

  repos add-protected-branch-required-status-checks-contexts --token=STRING --owner=STRING --repo=STRING --branch=STRING --contexts=CONTEXTS,...
    Add required status checks contexts of protected branch

  repos remove-protected-branch-required-status-checks-contexts --token=STRING --owner=STRING --repo=STRING --branch=STRING --contexts=CONTEXTS,...
    Remove required status checks contexts of protected branch

  repos get-protected-branch-pull-request-review-enforcement --token=STRING --owner=STRING --repo=STRING --branch=STRING
    Get pull request review enforcement of protected branch

  repos remove-protected-branch-pull-request-review-enforcement --token=STRING --owner=STRING --repo=STRING --branch=STRING
    Remove pull request review enforcement of protected branch

  repos get-protected-branch-required-signatures --token=STRING --zzzax-preview --owner=STRING --repo=STRING --branch=STRING
    Get required signatures of protected branch

  repos add-protected-branch-required-signatures --token=STRING --zzzax-preview --owner=STRING --repo=STRING --branch=STRING
    Add required signatures of protected branch

  repos remove-protected-branch-required-signatures --token=STRING --zzzax-preview --owner=STRING --repo=STRING --branch=STRING
    Remove required signatures of protected branch

  repos get-protected-branch-admin-enforcement --token=STRING --owner=STRING --repo=STRING --branch=STRING
    Get admin enforcement of protected branch

  repos add-protected-branch-admin-enforcement --token=STRING --owner=STRING --repo=STRING --branch=STRING
    Add admin enforcement of protected branch

  repos remove-protected-branch-admin-enforcement --token=STRING --owner=STRING --repo=STRING --branch=STRING
    Remove admin enforcement of protected branch

  repos get-protected-branch-restrictions --token=STRING --owner=STRING --repo=STRING --branch=STRING
    Get restrictions of protected branch

  repos remove-protected-branch-restrictions --token=STRING --owner=STRING --repo=STRING --branch=STRING
    Remove restrictions of protected branch

  repos list-protected-branch-team-restrictions --token=STRING --owner=STRING --repo=STRING --branch=STRING
    List team restrictions of protected branch

  repos replace-protected-branch-team-restrictions --token=STRING --owner=STRING --repo=STRING --branch=STRING --teams=TEAMS,...
    Replace team restrictions of protected branch

  repos add-protected-branch-team-restrictions --token=STRING --owner=STRING --repo=STRING --branch=STRING --teams=TEAMS,...
    Add team restrictions of protected branch

  repos remove-protected-branch-team-restrictions --token=STRING --owner=STRING --repo=STRING --branch=STRING --teams=TEAMS,...
    Remove team restrictions of protected branch

  repos list-protected-branch-user-restrictions --token=STRING --owner=STRING --repo=STRING --branch=STRING
    List user restrictions of protected branch

  repos replace-protected-branch-user-restrictions --token=STRING --owner=STRING --repo=STRING --branch=STRING --users=USERS,...
    Replace user restrictions of protected branch

  repos add-protected-branch-user-restrictions --token=STRING --owner=STRING --repo=STRING --branch=STRING --users=USERS,...
    Add user restrictions of protected branch

  repos remove-protected-branch-user-restrictions --token=STRING --owner=STRING --repo=STRING --branch=STRING --users=USERS,...
    Remove user restrictions of protected branch

  repos list-collaborators --token=STRING --owner=STRING --repo=STRING
    List collaborators

  repos check-collaborator --token=STRING --owner=STRING --repo=STRING --username=STRING
    Check if a user is a collaborator

  repos get-collaborator-permission-level --token=STRING --owner=STRING --repo=STRING --username=STRING
    Review a user's permission level

  repos add-collaborator --token=STRING --owner=STRING --repo=STRING --username=STRING
    Add user as a collaborator

  repos remove-collaborator --token=STRING --owner=STRING --repo=STRING --username=STRING
    Remove user as a collaborator

  repos list-commit-comments --token=STRING --owner=STRING --repo=STRING
    List commit comments for a repository

  repos list-comments-for-commit --token=STRING --owner=STRING --repo=STRING --ref=STRING
    List comments for a single commit

  repos create-commit-comment --token=STRING --owner=STRING --repo=STRING --sha=STRING --body=STRING
    Create a commit comment

  repos get-commit-comment --token=STRING --owner=STRING --repo=STRING --comment_id=INT-64
    Get a single commit comment

  repos update-commit-comment --token=STRING --owner=STRING --repo=STRING --comment_id=INT-64 --body=STRING
    Update a commit comment

  repos delete-commit-comment --token=STRING --owner=STRING --repo=STRING --comment_id=INT-64
    Delete a commit comment

  repos list-commits --token=STRING --owner=STRING --repo=STRING
    List commits on a repository

  repos get-commit --token=STRING --owner=STRING --repo=STRING --sha=STRING
    Get a single commit

  repos get-commit-ref-sha --token=STRING --owner=STRING --repo=STRING --ref=STRING
    Get the SHA-1 of a commit reference

  repos compare-commits --token=STRING --owner=STRING --repo=STRING --base=STRING --head=STRING
    Compare two commits

  repos retrieve-community-profile-metrics --token=STRING --owner=STRING --repo=STRING
    Retrieve community profile metrics

  repos get-readme --token=STRING --owner=STRING --repo=STRING
    Get the README

  repos get-contents --token=STRING --owner=STRING --repo=STRING --path=STRING
    Get contents

  repos get-archive-link --token=STRING --owner=STRING --repo=STRING --archive_format=STRING --ref=STRING
    Get archive link

  repos list-deploy-keys --token=STRING --owner=STRING --repo=STRING
    List deploy keys

  repos get-deploy-key --token=STRING --owner=STRING --repo=STRING --key_id=INT-64
    Get a deploy key

  repos add-deploy-key --token=STRING --owner=STRING --repo=STRING --key=STRING
    Add a new deploy key

  repos remove-deploy-key --token=STRING --owner=STRING --repo=STRING --key_id=INT-64
    Remove a deploy key

  repos list-deployments --token=STRING --owner=STRING --repo=STRING
    List deployments

  repos get-deployment --token=STRING --owner=STRING --repo=STRING --deployment_id=INT-64
    Get a single deployment

  repos create-deployment --token=STRING --owner=STRING --repo=STRING --ref=STRING
    Create a deployment

  repos list-deployment-statuses --token=STRING --owner=STRING --repo=STRING --deployment_id=INT-64
    List deployment statuses

  repos get-deployment-status --token=STRING --owner=STRING --repo=STRING --deployment_id=INT-64 --status_id=INT-64
    Get a single deployment status

  repos create-deployment-status --token=STRING --owner=STRING --repo=STRING --deployment_id=INT-64 --state=STRING
    Create a deployment status

  repos list-downloads --token=STRING --owner=STRING --repo=STRING
    List downloads for a repository

  repos get-download --token=STRING --owner=STRING --repo=STRING --download_id=INT-64
    Get a single download

  repos delete-download --token=STRING --owner=STRING --repo=STRING --download_id=INT-64
    Delete a download

  repos list-forks --token=STRING --owner=STRING --repo=STRING
    List forks

  repos create-fork --token=STRING --owner=STRING --repo=STRING
    Create a fork

  repos list-invitations --token=STRING --owner=STRING --repo=STRING
    List invitations for a repository

  repos delete-invitation --token=STRING --owner=STRING --repo=STRING --invitation_id=INT-64
    Delete a repository invitation

  repos update-invitation --token=STRING --owner=STRING --repo=STRING --invitation_id=INT-64
    Update a repository invitation

  repos list-invitations-for-authenticated-user --token=STRING
    List a user's repository invitations

  repos accept-invitation --token=STRING --invitation_id=INT-64
    Accept a repository invitation

  repos decline-invitation --token=STRING --invitation_id=INT-64
    Decline a repository invitation

  repos merge --token=STRING --owner=STRING --repo=STRING --base=STRING --head=STRING
    Perform a merge

  repos get-pages --token=STRING --mister-fantastic-preview --owner=STRING --repo=STRING
    Get information about a Pages site

  repos update-information-about-pages-site --token=STRING --mister-fantastic-preview --owner=STRING --repo=STRING
    Update information about a Pages site

  repos request-page-build --token=STRING --mister-fantastic-preview --owner=STRING --repo=STRING
    Request a page build

  repos list-pages-builds --token=STRING --owner=STRING --repo=STRING
    List Pages builds

  repos get-latest-pages-build --token=STRING --owner=STRING --repo=STRING
    Get latest Pages build

  repos get-pages-build --token=STRING --owner=STRING --repo=STRING --build_id=INT-64
    Get a specific Pages build

  repos list-releases --token=STRING --owner=STRING --repo=STRING
    List releases for a repository

  repos get-release --token=STRING --owner=STRING --repo=STRING --release_id=INT-64
    Get a single release

  repos get-latest-release --token=STRING --owner=STRING --repo=STRING
    Get the latest release

  repos get-release-by-tag --token=STRING --owner=STRING --repo=STRING --tag=STRING
    Get a release by tag name

  repos create-release --token=STRING --owner=STRING --repo=STRING --tag_name=STRING
    Create a release

  repos edit-release --token=STRING --owner=STRING --repo=STRING --release_id=INT-64
    Edit a release

  repos delete-release --token=STRING --owner=STRING --repo=STRING --release_id=INT-64
    Delete a release

  repos list-assets-for-release --token=STRING --owner=STRING --repo=STRING --release_id=INT-64
    List assets for a release

  repos get-release-asset --token=STRING --owner=STRING --repo=STRING --asset_id=INT-64
    Get a single release asset

  repos edit-release-asset --token=STRING --owner=STRING --repo=STRING --asset_id=INT-64
    Edit a release asset

  repos delete-release-asset --token=STRING --owner=STRING --repo=STRING --asset_id=INT-64
    Delete a release asset

  repos get-contributors-stats --token=STRING --owner=STRING --repo=STRING
    Get contributors list with additions, deletions, and commit counts

  repos get-commit-activity-stats --token=STRING --owner=STRING --repo=STRING
    Get the last year of commit activity data

  repos get-code-frequency-stats --token=STRING --owner=STRING --repo=STRING
    Get the number of additions and deletions per week

  repos get-participation-stats --token=STRING --owner=STRING --repo=STRING
    Get the weekly commit count for the repository owner and everyone else

  repos get-punch-card-stats --token=STRING --owner=STRING --repo=STRING
    Get the number of commits per hour in each day

  repos create-status --token=STRING --owner=STRING --repo=STRING --sha=STRING --state=STRING
    Create a status

  repos list-statuses-for-ref --token=STRING --owner=STRING --repo=STRING --ref=STRING
    List statuses for a specific ref

  repos get-combined-status-for-ref --token=STRING --owner=STRING --repo=STRING --ref=STRING
    Get the combined status for a specific ref

  repos get-top-referrers --token=STRING --owner=STRING --repo=STRING
    List referrers

  repos get-top-paths --token=STRING --owner=STRING --repo=STRING
    List paths

  repos get-views --token=STRING --owner=STRING --repo=STRING
    Views

  repos get-clones --token=STRING --owner=STRING --repo=STRING
    Clones

  repos list-hooks --token=STRING --owner=STRING --repo=STRING
    List hooks

  repos get-hook --token=STRING --owner=STRING --repo=STRING --hook_id=INT-64
    Get single hook

  repos test-push-hook --token=STRING --owner=STRING --repo=STRING --hook_id=INT-64
    Test a push hook

  repos ping-hook --token=STRING --owner=STRING --repo=STRING --hook_id=INT-64
    Ping a hook

  repos delete-hook --token=STRING --owner=STRING --repo=STRING --hook_id=INT-64
    Delete a hook

  scim get-provisioned-identities-list --token=STRING --organization_id=INT-64
    Get a list of provisioned identities

  scim get-provisioning-details-for-user --token=STRING --organization_id=INT-64 --external_identity_guid=STRING
    Get provisioning details for a single user

  scim provision-invite-users --token=STRING --organization_id=INT-64
    Provision and invite users

  scim update-provisioned-org-membership --token=STRING --organization_id=INT-64 --external_identity_guid=STRING
    Update a provisioned organization membership

  scim update-user-attribute --token=STRING --organization_id=INT-64 --external_identity_guid=STRING
    Update a user attribute

  scim remove-user-from-org --token=STRING --organization_id=INT-64 --external_identity_guid=STRING
    Remove a user from the organization

  search repos --token=STRING --q=STRING
    Search repositories

  search commits --token=STRING --cloak-preview --q=STRING
    Search commits

  search code --token=STRING --q=STRING
    Search code

  search issues --token=STRING --q=STRING
    Search issues

  search users --token=STRING --q=STRING
    Search users

  search topics --token=STRING --q=STRING
    Search topics

  search labels --token=STRING --repository_id=INT-64 --q=STRING
    Search labels

  search issues-legacy --token=STRING --owner=STRING --repository=STRING --state=STRING --keyword=STRING
    Search issues

  search repos-legacy --token=STRING --keyword=STRING
    Search repositories

  search users-legacy --token=STRING --keyword=STRING
    Search users

  search email-legacy --token=STRING --email=STRING
    Email search

  teams list --token=STRING --org=STRING
    List teams

  teams get --token=STRING --team_id=INT-64
    Get team

  teams create --token=STRING --org=STRING --name=STRING
    Create team

  teams edit --token=STRING --team_id=INT-64 --name=STRING
    Edit team

  teams delete --token=STRING --team_id=INT-64
    Delete team

  teams list-child --token=STRING --hellcat-preview --team_id=INT-64
    List child teams

  teams list-repos --token=STRING --team_id=INT-64
    List team repos

  teams check-manages-repo --token=STRING --team_id=INT-64 --owner=STRING --repo=STRING
    Check if a team manages a repository

  teams add-or-update-repo --token=STRING --team_id=INT-64 --owner=STRING --repo=STRING
    Add or update team repository

  teams remove-repo --token=STRING --team_id=INT-64 --owner=STRING --repo=STRING
    Remove team repository

  teams list-for-authenticated-user --token=STRING
    List user teams

  teams list-projects --token=STRING --inertia-preview --team_id=INT-64
    List team projects

  teams review-project --token=STRING --inertia-preview --team_id=INT-64 --project_id=INT-64
    Review a team project

  teams add-or-update-project --token=STRING --inertia-preview --team_id=INT-64 --project_id=INT-64
    Add or update team project

  teams remove-project --token=STRING --team_id=INT-64 --project_id=INT-64
    Remove team project

  teams list-discussions --token=STRING --echo-preview --team_id=INT-64
    List discussions

  teams get-discussion --token=STRING --echo-preview --team_id=INT-64 --discussion_number=INT-64
    Get a single discussion

  teams create-discussion --token=STRING --echo-preview --team_id=INT-64 --title=STRING --body=STRING
    Create a discussion

  teams edit-discussion --token=STRING --echo-preview --team_id=INT-64 --discussion_number=INT-64
    Edit a discussion

  teams delete-discussion --token=STRING --echo-preview --team_id=INT-64 --discussion_number=INT-64
    Delete a discussion

  teams list-discussion-comments --token=STRING --echo-preview --team_id=INT-64 --discussion_number=INT-64
    List comments

  teams get-discussion-comment --token=STRING --echo-preview --team_id=INT-64 --discussion_number=INT-64 --comment_number=INT-64
    Get a single comment

  teams create-discussion-comment --token=STRING --echo-preview --team_id=INT-64 --discussion_number=INT-64 --body=STRING
    Create a comment

  teams edit-discussion-comment --token=STRING --echo-preview --team_id=INT-64 --discussion_number=INT-64 --comment_number=INT-64 --body=STRING
    Edit a comment

  teams delete-discussion-comment --token=STRING --echo-preview --team_id=INT-64 --discussion_number=INT-64 --comment_number=INT-64
    Delete a comment

  teams list-members --token=STRING --team_id=INT-64
    List team members

  teams get-member --token=STRING --team_id=INT-64 --username=STRING
    Get team member

  teams add-member --token=STRING --team_id=INT-64 --username=STRING
    Add team member

  teams remove-member --token=STRING --team_id=INT-64 --username=STRING
    Remove team member

  teams get-membership --token=STRING --team_id=INT-64 --username=STRING
    Get team membership

  teams add-or-update-membership --token=STRING --team_id=INT-64 --username=STRING
    Add or update team membership

  teams remove-membership --token=STRING --team_id=INT-64 --username=STRING
    Remove team membership

  teams list-pending-invitations --token=STRING --team_id=INT-64
    List pending team invitations

  users get-by-username --token=STRING --username=STRING
    Get a single user

  users get-authenticated --token=STRING
    Get the authenticated user

  users update-authenticated --token=STRING
    Update the authenticated user

  users get-context-for-user --token=STRING --hagar-preview --username=STRING
    Get contextual information about a user

  users list --token=STRING
    Get all users

  users list-blocked --token=STRING
    List blocked users

  users check-blocked --token=STRING --username=STRING
    Check whether you've blocked a user

  users block --token=STRING --username=STRING
    Block a user

  users unblock --token=STRING --username=STRING
    Unblock a user

  users list-emails --token=STRING
    List email addresses for a user

  users list-public-emails --token=STRING
    List public email addresses for a user

  users add-emails --token=STRING --emails=EMAILS,...
    Add email address(es)

  users delete-emails --token=STRING --emails=EMAILS,...
    Delete email address(es)

  users toggle-primary-email-visibility --token=STRING --email=STRING --visibility=STRING
    Toggle primary email visibility

  users list-followers-for-user --token=STRING --username=STRING
    List a user's followers

  users list-followers-for-authenticated-user --token=STRING
    List the authenticated user's followers

  users list-following-for-user --token=STRING --username=STRING
    List who a user is following

  users list-following-for-authenticated-user --token=STRING
    List who the authenticated user is following

  users check-following --token=STRING --username=STRING
    Check if you are following a user

  users check-following-for-user --token=STRING --username=STRING --target_user=STRING
    Check if one user follows another

  users follow --token=STRING --username=STRING
    Follow a user

  users unfollow --token=STRING --username=STRING
    Unfollow a user

  users list-public-keys-for-user --token=STRING --username=STRING
    List public keys for a user

  users list-public-keys --token=STRING
    List your public keys

  users get-public-key --token=STRING --key_id=INT-64
    Get a single public key

  users create-public-key --token=STRING
    Create a public key

  users delete-public-key --token=STRING --key_id=INT-64
    Delete a public key

  users list-gpg-keys-for-user --token=STRING --username=STRING
    List GPG keys for a user

  users list-gpg-keys --token=STRING
    List your GPG keys

  users get-gpg-key --token=STRING --gpg_key_id=INT-64
    Get a single GPG key

  users create-gpg-key --token=STRING
    Create a GPG key

  users delete-gpg-key --token=STRING --gpg_key_id=INT-64
    Delete a GPG key

Run "octo-cli <command> --help" for more information on a command.
```

