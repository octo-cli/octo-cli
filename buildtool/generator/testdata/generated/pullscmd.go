// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import "github.com/octo-cli/octo-cli/internal"

type PullsCmd struct {
	List                 PullsListCmd                 `cmd:"" help:"List pull requests - https://developer.github.com/v3/pulls/#list-pull-requests"`
	Get                  PullsGetCmd                  `cmd:"" help:"Get a single pull request - https://developer.github.com/v3/pulls/#get-a-single-pull-request"`
	Create               PullsCreateCmd               `cmd:"" help:"Create a pull request - https://developer.github.com/v3/pulls/#create-a-pull-request"`
	CreateFromIssue      PullsCreateFromIssueCmd      `cmd:"" help:"Create a Pull Request from an Issue - https://developer.github.com/v3/pulls/#create-a-pull-request"`
	Update               PullsUpdateCmd               `cmd:"" help:"Update a pull request - https://developer.github.com/v3/pulls/#update-a-pull-request"`
	ListCommits          PullsListCommitsCmd          `cmd:"" help:"List commits on a pull request - https://developer.github.com/v3/pulls/#list-commits-on-a-pull-request"`
	ListFiles            PullsListFilesCmd            `cmd:"" help:"List pull requests files - https://developer.github.com/v3/pulls/#list-pull-requests-files"`
	CheckIfMerged        PullsCheckIfMergedCmd        `cmd:"" help:"Get if a pull request has been merged - https://developer.github.com/v3/pulls/#get-if-a-pull-request-has-been-merged"`
	Merge                PullsMergeCmd                `cmd:"" help:"Merge a pull request (Merge Button) - https://developer.github.com/v3/pulls/#merge-a-pull-request-merge-button"`
	ListReviews          PullsListReviewsCmd          `cmd:"" help:"List reviews on a pull request - https://developer.github.com/v3/pulls/reviews/#list-reviews-on-a-pull-request"`
	GetReview            PullsGetReviewCmd            `cmd:"" help:"Get a single review - https://developer.github.com/v3/pulls/reviews/#get-a-single-review"`
	DeletePendingReview  PullsDeletePendingReviewCmd  `cmd:"" help:"Delete a pending review - https://developer.github.com/v3/pulls/reviews/#delete-a-pending-review"`
	GetCommentsForReview PullsGetCommentsForReviewCmd `cmd:"" help:"Get comments for a single review - https://developer.github.com/v3/pulls/reviews/#get-comments-for-a-single-review"`
	SubmitReview         PullsSubmitReviewCmd         `cmd:"" help:"Submit a pull request review - https://developer.github.com/v3/pulls/reviews/#submit-a-pull-request-review"`
	DismissReview        PullsDismissReviewCmd        `cmd:"" help:"Dismiss a pull request review - https://developer.github.com/v3/pulls/reviews/#dismiss-a-pull-request-review"`
	ListComments         PullsListCommentsCmd         `cmd:"" help:"List comments on a pull request - https://developer.github.com/v3/pulls/comments/#list-comments-on-a-pull-request"`
	ListCommentsForRepo  PullsListCommentsForRepoCmd  `cmd:"" help:"List comments in a repository - https://developer.github.com/v3/pulls/comments/#list-comments-in-a-repository"`
	GetComment           PullsGetCommentCmd           `cmd:"" help:"Get a single comment - https://developer.github.com/v3/pulls/comments/#get-a-single-comment"`
	CreateComment        PullsCreateCommentCmd        `cmd:"" help:"Create a comment - https://developer.github.com/v3/pulls/comments/#create-a-comment"`
	CreateCommentReply   PullsCreateCommentReplyCmd   `cmd:"" help:"Create a comment reply - https://developer.github.com/v3/pulls/comments/#create-a-comment"`
	EditComment          PullsEditCommentCmd          `cmd:"" help:"Edit a comment - https://developer.github.com/v3/pulls/comments/#edit-a-comment"`
	DeleteComment        PullsDeleteCommentCmd        `cmd:"" help:"Delete a comment - https://developer.github.com/v3/pulls/comments/#delete-a-comment"`
	ListReviewRequests   PullsListReviewRequestsCmd   `cmd:"" help:"List review requests - https://developer.github.com/v3/pulls/review_requests/#list-review-requests"`
	CreateReviewRequest  PullsCreateReviewRequestCmd  `cmd:"" help:"Create a review request - https://developer.github.com/v3/pulls/review_requests/#create-a-review-request"`
	DeleteReviewRequest  PullsDeleteReviewRequestCmd  `cmd:"" help:"Delete a review request - https://developer.github.com/v3/pulls/review_requests/#delete-a-review-request"`
}

type PullsListCmd struct {
	internal.BaseCmd
	Symmetra  bool   "name:\"symmetra-preview\" help:\"**Note:** You can now use emoji in label names, add descriptions to labels, and search for labels in a repository. See the [blog post](/changes/2018-02-22-label-description-search-preview) for full details. To access these features and receive payloads with this data during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.symmetra-preview+json\n\n```\""
	Owner     string `name:"owner"`
	Repo      string `required:"" name:"repo"`
	State     string "name:\"state\" help:\"Either `open`, `closed`, or `all` to filter by state.\""
	Head      string "name:\"head\" help:\"Filter pulls by head user and branch name in the format of `user:ref-name`. Example: `github:new-script-format`.\""
	Base      string "name:\"base\" help:\"Filter pulls by base branch name. Example: `gh-pages`.\""
	Sort      string "name:\"sort\" help:\"What to sort results by. Can be either `created`, `updated`, `popularity` (comment count) or `long-running` (age, filtering by pulls updated in the last month).\""
	Direction string "name:\"direction\" help:\"The direction of the sort. Can be either `asc` or `desc`.\""
	PerPage   int64  `name:"per_page" help:"Results per page (max 100)"`
	Page      int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *PullsListCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls")
	c.UpdatePreview("symmetra", c.Symmetra)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("state", c.State)
	c.UpdateURLQuery("head", c.Head)
	c.UpdateURLQuery("base", c.Base)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("direction", c.Direction)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type PullsGetCmd struct {
	internal.BaseCmd
	Symmetra bool   "name:\"symmetra-preview\" help:\"**Note:** You can now use emoji in label names, add descriptions to labels, and search for labels in a repository. See the [blog post](/changes/2018-02-22-label-description-search-preview) for full details. To access these features and receive payloads with this data during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.symmetra-preview+json\n\n```\""
	Owner    string `name:"owner"`
	Repo     string `required:"" name:"repo"`
	Number   int64  `required:"" name:"number"`
}

func (c *PullsGetCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/:number")
	c.UpdatePreview("symmetra", c.Symmetra)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	return c.DoRequest("GET")
}

type PullsCreateCmd struct {
	internal.BaseCmd
	Symmetra            bool   "name:\"symmetra-preview\" help:\"**Note:** You can now use emoji in label names, add descriptions to labels, and search for labels in a repository. See the [blog post](/changes/2018-02-22-label-description-search-preview) for full details. To access these features and receive payloads with this data during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.symmetra-preview+json\n\n```\""
	Owner               string `name:"owner"`
	Repo                string `required:"" name:"repo"`
	Title               string `required:"" name:"title" help:"The title of the pull request."`
	Head                string "required:\"\" name:\"head\" help:\"The name of the branch where your changes are implemented. For cross-repository pull requests in the same network, namespace `head` with a user like this: `username:branch`.\""
	Base                string `required:"" name:"base" help:"The name of the branch you want the changes pulled into. This should be an existing branch on the current repository. You cannot submit a pull request to one repository that requests a merge to a base of another repository."`
	Body                string `name:"body" help:"The contents of the pull request."`
	MaintainerCanModify bool   `name:"maintainer_can_modify" help:"Indicates whether [maintainers can modify](https://help.github.com/articles/allowing-changes-to-a-pull-request-branch-created-from-a-fork/) the pull request."`
}

func (c *PullsCreateCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls")
	c.UpdatePreview("symmetra", c.Symmetra)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("title", c.Title)
	c.UpdateBody("head", c.Head)
	c.UpdateBody("base", c.Base)
	c.UpdateBody("body", c.Body)
	c.UpdateBody("maintainer_can_modify", c.MaintainerCanModify)
	return c.DoRequest("POST")
}

type PullsCreateFromIssueCmd struct {
	internal.BaseCmd
	Symmetra            bool   "name:\"symmetra-preview\" help:\"**Note:** You can now use emoji in label names, add descriptions to labels, and search for labels in a repository. See the [blog post](/changes/2018-02-22-label-description-search-preview) for full details. To access these features and receive payloads with this data during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.symmetra-preview+json\n\n```\""
	Owner               string `name:"owner"`
	Repo                string `required:"" name:"repo"`
	Issue               int64  `required:"" name:"issue" help:"The issue number in this repository to turn into a Pull Request."`
	Head                string "required:\"\" name:\"head\" help:\"The name of the branch where your changes are implemented. For cross-repository pull requests in the same network, namespace `head` with a user like this: `username:branch`.\""
	Base                string `required:"" name:"base" help:"The name of the branch you want the changes pulled into. This should be an existing branch on the current repository. You cannot submit a pull request to one repository that requests a merge to a base of another repository."`
	MaintainerCanModify bool   `name:"maintainer_can_modify" help:"Indicates whether [maintainers can modify](https://help.github.com/articles/allowing-changes-to-a-pull-request-branch-created-from-a-fork/) the pull request."`
}

func (c *PullsCreateFromIssueCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls")
	c.UpdatePreview("symmetra", c.Symmetra)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("issue", c.Issue)
	c.UpdateBody("head", c.Head)
	c.UpdateBody("base", c.Base)
	c.UpdateBody("maintainer_can_modify", c.MaintainerCanModify)
	return c.DoRequest("POST")
}

type PullsUpdateCmd struct {
	internal.BaseCmd
	Symmetra            bool   "name:\"symmetra-preview\" help:\"**Note:** You can now use emoji in label names, add descriptions to labels, and search for labels in a repository. See the [blog post](/changes/2018-02-22-label-description-search-preview) for full details. To access these features and receive payloads with this data during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.symmetra-preview+json\n\n```\""
	Owner               string `name:"owner"`
	Repo                string `required:"" name:"repo"`
	Number              int64  `required:"" name:"number"`
	Title               string `name:"title" help:"The title of the pull request."`
	Body                string `name:"body" help:"The contents of the pull request."`
	State               string "name:\"state\" help:\"State of this Pull Request. Either `open` or `closed`.\""
	Base                string `name:"base" help:"The name of the branch you want your changes pulled into. This should be an existing branch on the current repository. You cannot update the base branch on a pull request to point to another repository."`
	MaintainerCanModify bool   `name:"maintainer_can_modify" help:"Indicates whether [maintainers can modify](https://help.github.com/articles/allowing-changes-to-a-pull-request-branch-created-from-a-fork/) the pull request."`
}

func (c *PullsUpdateCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/:number")
	c.UpdatePreview("symmetra", c.Symmetra)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	c.UpdateBody("title", c.Title)
	c.UpdateBody("body", c.Body)
	c.UpdateBody("state", c.State)
	c.UpdateBody("base", c.Base)
	c.UpdateBody("maintainer_can_modify", c.MaintainerCanModify)
	return c.DoRequest("PATCH")
}

type PullsListCommitsCmd struct {
	internal.BaseCmd
	Owner   string `name:"owner"`
	Repo    string `required:"" name:"repo"`
	Number  int64  `required:"" name:"number"`
	PerPage int64  `name:"per_page" help:"Results per page (max 100)"`
	Page    int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *PullsListCommitsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/:number/commits")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type PullsListFilesCmd struct {
	internal.BaseCmd
	Owner   string `name:"owner"`
	Repo    string `required:"" name:"repo"`
	Number  int64  `required:"" name:"number"`
	PerPage int64  `name:"per_page" help:"Results per page (max 100)"`
	Page    int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *PullsListFilesCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/:number/files")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type PullsCheckIfMergedCmd struct {
	internal.BaseCmd
	Owner  string `name:"owner"`
	Repo   string `required:"" name:"repo"`
	Number int64  `required:"" name:"number"`
}

func (c *PullsCheckIfMergedCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/:number/merge")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	return c.DoRequest("GET")
}

type PullsMergeCmd struct {
	internal.BaseCmd
	Owner         string `name:"owner"`
	Repo          string `required:"" name:"repo"`
	Number        int64  `required:"" name:"number"`
	CommitTitle   string `name:"commit_title" help:"Title for the automatic commit message."`
	CommitMessage string `name:"commit_message" help:"Extra detail to append to automatic commit message."`
	Sha           string `name:"sha" help:"SHA that pull request head must match to allow merge."`
	MergeMethod   string "name:\"merge_method\" help:\"Merge method to use. Possible values are `merge`, `squash` or `rebase`. Default is `merge`.\""
}

func (c *PullsMergeCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/:number/merge")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	c.UpdateBody("commit_title", c.CommitTitle)
	c.UpdateBody("commit_message", c.CommitMessage)
	c.UpdateBody("sha", c.Sha)
	c.UpdateBody("merge_method", c.MergeMethod)
	return c.DoRequest("PUT")
}

type PullsListReviewsCmd struct {
	internal.BaseCmd
	Owner   string `name:"owner"`
	Repo    string `required:"" name:"repo"`
	Number  int64  `required:"" name:"number"`
	PerPage int64  `name:"per_page" help:"Results per page (max 100)"`
	Page    int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *PullsListReviewsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/:number/reviews")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type PullsGetReviewCmd struct {
	internal.BaseCmd
	Owner    string `name:"owner"`
	Repo     string `required:"" name:"repo"`
	Number   int64  `required:"" name:"number"`
	ReviewId int64  `required:"" name:"review_id"`
}

func (c *PullsGetReviewCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/:number/reviews/:review_id")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	c.UpdateURLPath("review_id", c.ReviewId)
	return c.DoRequest("GET")
}

type PullsDeletePendingReviewCmd struct {
	internal.BaseCmd
	Owner    string `name:"owner"`
	Repo     string `required:"" name:"repo"`
	Number   int64  `required:"" name:"number"`
	ReviewId int64  `required:"" name:"review_id"`
}

func (c *PullsDeletePendingReviewCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/:number/reviews/:review_id")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	c.UpdateURLPath("review_id", c.ReviewId)
	return c.DoRequest("DELETE")
}

type PullsGetCommentsForReviewCmd struct {
	internal.BaseCmd
	Owner    string `name:"owner"`
	Repo     string `required:"" name:"repo"`
	Number   int64  `required:"" name:"number"`
	ReviewId int64  `required:"" name:"review_id"`
	PerPage  int64  `name:"per_page" help:"Results per page (max 100)"`
	Page     int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *PullsGetCommentsForReviewCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/:number/reviews/:review_id/comments")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	c.UpdateURLPath("review_id", c.ReviewId)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type PullsSubmitReviewCmd struct {
	internal.BaseCmd
	Owner    string `name:"owner"`
	Repo     string `required:"" name:"repo"`
	Number   int64  `required:"" name:"number"`
	ReviewId int64  `required:"" name:"review_id"`
	Body     string `name:"body" help:"The body text of the pull request review"`
	Event    string "required:\"\" name:\"event\" help:\"The review action you want to perform. The review actions include: `APPROVE`, `REQUEST_CHANGES`, or `COMMENT`. When you leave this blank, the API returns _HTTP 422 (Unrecognizable entity)_ and sets the review action state to `PENDING`, which means you will need to re-submit the pull request review using a review action.\""
}

func (c *PullsSubmitReviewCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/:number/reviews/:review_id/events")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	c.UpdateURLPath("review_id", c.ReviewId)
	c.UpdateBody("body", c.Body)
	c.UpdateBody("event", c.Event)
	return c.DoRequest("POST")
}

type PullsDismissReviewCmd struct {
	internal.BaseCmd
	Owner    string `name:"owner"`
	Repo     string `required:"" name:"repo"`
	Number   int64  `required:"" name:"number"`
	ReviewId int64  `required:"" name:"review_id"`
	Message  string `required:"" name:"message" help:"The message for the pull request review dismissal"`
}

func (c *PullsDismissReviewCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/:number/reviews/:review_id/dismissals")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	c.UpdateURLPath("review_id", c.ReviewId)
	c.UpdateBody("message", c.Message)
	return c.DoRequest("PUT")
}

type PullsListCommentsCmd struct {
	internal.BaseCmd
	SquirrelGirl bool   "name:\"squirrel-girl-preview\" help:\"An additional `reactions` object in the review comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](/changes/2016-05-12-reactions-api-preview) for full details.\n\nTo access the API you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.squirrel-girl-preview\n\n```\n\nThe `reactions` key will have the following payload where `url` can be used to construct the API location for [listing and creating](/v3/reactions) reactions.\""
	Owner        string `name:"owner"`
	Repo         string `required:"" name:"repo"`
	Number       int64  `required:"" name:"number"`
	Sort         string "name:\"sort\" help:\"Can be either `created` or `updated` comments.\""
	Direction    string "name:\"direction\" help:\"Can be either `asc` or `desc`. Ignored without `sort` parameter.\""
	Since        string "name:\"since\" help:\"This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Only returns comments `updated` at or after this time.\""
	PerPage      int64  `name:"per_page" help:"Results per page (max 100)"`
	Page         int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *PullsListCommentsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/:number/comments")
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("direction", c.Direction)
	c.UpdateURLQuery("since", c.Since)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type PullsListCommentsForRepoCmd struct {
	internal.BaseCmd
	SquirrelGirl bool   "name:\"squirrel-girl-preview\" help:\"An additional `reactions` object in the review comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](/changes/2016-05-12-reactions-api-preview) for full details.\n\nTo access the API you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.squirrel-girl-preview\n\n```\n\nThe `reactions` key will have the following payload where `url` can be used to construct the API location for [listing and creating](/v3/reactions) reactions.\""
	Owner        string `name:"owner"`
	Repo         string `required:"" name:"repo"`
	Sort         string "name:\"sort\" help:\"Can be either `created` or `updated` comments.\""
	Direction    string "name:\"direction\" help:\"Can be either `asc` or `desc`. Ignored without `sort` parameter.\""
	Since        string "name:\"since\" help:\"This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Only returns comments `updated` at or after this time.\""
	PerPage      int64  `name:"per_page" help:"Results per page (max 100)"`
	Page         int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *PullsListCommentsForRepoCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/comments")
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("direction", c.Direction)
	c.UpdateURLQuery("since", c.Since)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type PullsGetCommentCmd struct {
	internal.BaseCmd
	SquirrelGirl bool   "name:\"squirrel-girl-preview\" help:\"An additional `reactions` object in the review comment payload is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](/changes/2016-05-12-reactions-api-preview) for full details.\n\nTo access the API you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.squirrel-girl-preview\n\n```\n\nThe `reactions` key will have the following payload where `url` can be used to construct the API location for [listing and creating](/v3/reactions) reactions.\""
	Owner        string `name:"owner"`
	Repo         string `required:"" name:"repo"`
	CommentId    int64  `required:"" name:"comment_id"`
}

func (c *PullsGetCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/comments/:comment_id")
	c.UpdatePreview("squirrel-girl", c.SquirrelGirl)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("comment_id", c.CommentId)
	return c.DoRequest("GET")
}

type PullsCreateCommentCmd struct {
	internal.BaseCmd
	Owner    string `name:"owner"`
	Repo     string `required:"" name:"repo"`
	Number   int64  `required:"" name:"number"`
	Body     string `required:"" name:"body" help:"The text of the comment."`
	CommitId string "required:\"\" name:\"commit_id\" help:\"The SHA of the commit needing a comment. Not using the latest commit SHA may render your comment outdated if a subsequent commit modifies the line you specify as the `position`.\""
	Path     string `required:"" name:"path" help:"The relative path to the file that necessitates a comment."`
	Position int64  `required:"" name:"position" help:"The position in the diff where you want to add a review comment. Note this value is not the same as the line number in the file. For help finding the position value, read the note below."`
}

func (c *PullsCreateCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/:number/comments")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	c.UpdateBody("body", c.Body)
	c.UpdateBody("commit_id", c.CommitId)
	c.UpdateBody("path", c.Path)
	c.UpdateBody("position", c.Position)
	return c.DoRequest("POST")
}

type PullsCreateCommentReplyCmd struct {
	internal.BaseCmd
	Owner     string `name:"owner"`
	Repo      string `required:"" name:"repo"`
	Number    int64  `required:"" name:"number"`
	Body      string `required:"" name:"body" help:"The text of the comment."`
	InReplyTo int64  `required:"" name:"in_reply_to" help:"The comment ID to reply to. **Note**: This must be the ID of a _top-level comment_, not a reply to that comment. Replies to replies are not supported."`
}

func (c *PullsCreateCommentReplyCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/:number/comments")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	c.UpdateBody("body", c.Body)
	c.UpdateBody("in_reply_to", c.InReplyTo)
	return c.DoRequest("POST")
}

type PullsEditCommentCmd struct {
	internal.BaseCmd
	Owner     string `name:"owner"`
	Repo      string `required:"" name:"repo"`
	CommentId int64  `required:"" name:"comment_id"`
	Body      string `required:"" name:"body" help:"The text of the comment."`
}

func (c *PullsEditCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/comments/:comment_id")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("comment_id", c.CommentId)
	c.UpdateBody("body", c.Body)
	return c.DoRequest("PATCH")
}

type PullsDeleteCommentCmd struct {
	internal.BaseCmd
	Owner     string `name:"owner"`
	Repo      string `required:"" name:"repo"`
	CommentId int64  `required:"" name:"comment_id"`
}

func (c *PullsDeleteCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/comments/:comment_id")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("comment_id", c.CommentId)
	return c.DoRequest("DELETE")
}

type PullsListReviewRequestsCmd struct {
	internal.BaseCmd
	Owner   string `name:"owner"`
	Repo    string `required:"" name:"repo"`
	Number  int64  `required:"" name:"number"`
	PerPage int64  `name:"per_page" help:"Results per page (max 100)"`
	Page    int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *PullsListReviewRequestsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/:number/requested_reviewers")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type PullsCreateReviewRequestCmd struct {
	internal.BaseCmd
	Symmetra      bool     "name:\"symmetra-preview\" help:\"**Note:** You can now use emoji in label names, add descriptions to labels, and search for labels in a repository. See the [blog post](/changes/2018-02-22-label-description-search-preview) for full details. To access these features and receive payloads with this data during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\napplication/vnd.github.symmetra-preview+json\n\n```\""
	Owner         string   `name:"owner"`
	Repo          string   `required:"" name:"repo"`
	Number        int64    `required:"" name:"number"`
	Reviewers     []string "name:\"reviewers\" help:\"An array of user `login`s that will be requested.\""
	TeamReviewers []string "name:\"team_reviewers\" help:\"An array of team `slug`s that will be requested.\""
}

func (c *PullsCreateReviewRequestCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/:number/requested_reviewers")
	c.UpdatePreview("symmetra", c.Symmetra)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	c.UpdateBody("reviewers", c.Reviewers)
	c.UpdateBody("team_reviewers", c.TeamReviewers)
	return c.DoRequest("POST")
}

type PullsDeleteReviewRequestCmd struct {
	internal.BaseCmd
	Owner         string   `name:"owner"`
	Repo          string   `required:"" name:"repo"`
	Number        int64    `required:"" name:"number"`
	Reviewers     []string "name:\"reviewers\" help:\"An array of user `login`s that will be removed.\""
	TeamReviewers []string "name:\"team_reviewers\" help:\"An array of team `slug`s that will be removed.\""
}

func (c *PullsDeleteReviewRequestCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/pulls/:number/requested_reviewers")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("number", c.Number)
	c.UpdateBody("reviewers", c.Reviewers)
	c.UpdateBody("team_reviewers", c.TeamReviewers)
	return c.DoRequest("DELETE")
}
