// Code generated by go-github-cli/generator; DO NOT EDIT.

package services

type SearchCmd struct {
	Repos        SearchReposCmd        `cmd:"" help:"Search repositories"`
	Commits      SearchCommitsCmd      `cmd:"" help:"Search commits"`
	Code         SearchCodeCmd         `cmd:"" help:"Search code"`
	Issues       SearchIssuesCmd       `cmd:"" help:"Search issues"`
	Users        SearchUsersCmd        `cmd:"" help:"Search users"`
	Topics       SearchTopicsCmd       `cmd:"" help:"Search topics"`
	Labels       SearchLabelsCmd       `cmd:"" help:"Search labels"`
	IssuesLegacy SearchIssuesLegacyCmd `cmd:"" help:"Search issues"`
	ReposLegacy  SearchReposLegacyCmd  `cmd:"" help:"Search repositories"`
	UsersLegacy  SearchUsersLegacyCmd  `cmd:"" help:"Search users"`
	EmailLegacy  SearchEmailLegacyCmd  `cmd:"" help:"Email search"`
}

type SearchReposCmd struct {
	baseCmd
	Q       string `required:"" name:"q" help:"The search keywords, as well as any qualifiers."`
	Sort    string "name:\"sort\" help:\"The sort field. One of `stars`, `forks`, or `updated`.\""
	Order   string "name:\"order\" help:\"The sort order if `sort` parameter is provided. One of `asc` or `desc`.\""
	PerPage int64  `name:"per_page" help:"Results per page (max 100)"`
	Page    int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *SearchReposCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/search/repositories"
	c.updateURLQuery("q", c.Q)
	c.updateURLQuery("sort", c.Sort)
	c.updateURLQuery("order", c.Order)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type SearchCommitsCmd struct {
	baseCmd
	Q       string `required:"" name:"q" help:"The search terms."`
	Sort    string "name:\"sort\" help:\"The sort field. Can be `author-date` or `committer-date`.\""
	Order   string "name:\"order\" help:\"The sort order if `sort` parameter is provided. One of `asc` or `desc`.\""
	PerPage int64  `name:"per_page" help:"Results per page (max 100)"`
	Page    int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *SearchCommitsCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/search/commits"
	c.updateURLQuery("q", c.Q)
	c.updateURLQuery("sort", c.Sort)
	c.updateURLQuery("order", c.Order)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type SearchCodeCmd struct {
	baseCmd
	Q       string `required:"" name:"q" help:"The search terms."`
	Sort    string "name:\"sort\" help:\"The sort field. Can only be `indexed`, which indicates how recently a file has been indexed by the GitHub search infrastructure.\""
	Order   string "name:\"order\" help:\"The sort order if `sort` parameter is provided. One of `asc` or `desc`.\""
	PerPage int64  `name:"per_page" help:"Results per page (max 100)"`
	Page    int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *SearchCodeCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/search/code"
	c.updateURLQuery("q", c.Q)
	c.updateURLQuery("sort", c.Sort)
	c.updateURLQuery("order", c.Order)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type SearchIssuesCmd struct {
	baseCmd
	Q       string `required:"" name:"q" help:"The search terms."`
	Sort    string "name:\"sort\" help:\"The sort field. Can be `comments`, `created`, or `updated`.\""
	Order   string "name:\"order\" help:\"The sort order if `sort` parameter is provided. One of `asc` or `desc`.\""
	PerPage int64  `name:"per_page" help:"Results per page (max 100)"`
	Page    int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *SearchIssuesCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/search/issues"
	c.updateURLQuery("q", c.Q)
	c.updateURLQuery("sort", c.Sort)
	c.updateURLQuery("order", c.Order)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type SearchUsersCmd struct {
	baseCmd
	Q       string `required:"" name:"q" help:"The search terms."`
	Sort    string "name:\"sort\" help:\"The sort field. Can be `followers`, `repositories`, or `joined`.\""
	Order   string "name:\"order\" help:\"The sort order if `sort` parameter is provided. One of `asc` or `desc`.\""
	PerPage int64  `name:"per_page" help:"Results per page (max 100)"`
	Page    int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *SearchUsersCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/search/users"
	c.updateURLQuery("q", c.Q)
	c.updateURLQuery("sort", c.Sort)
	c.updateURLQuery("order", c.Order)
	c.updateURLQuery("per_page", c.PerPage)
	c.updateURLQuery("page", c.Page)
	return c.doRequest("GET")
}

type SearchTopicsCmd struct {
	baseCmd
	Q string `required:"" name:"q" help:"The search terms."`
}

func (c *SearchTopicsCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/search/topics"
	c.updateURLQuery("q", c.Q)
	return c.doRequest("GET")
}

type SearchLabelsCmd struct {
	baseCmd
	RepositoryId int64  `required:"" name:"repository_id" help:"The id of the repository."`
	Q            string `required:"" name:"q" help:"The search keywords."`
	Sort         string "name:\"sort\" help:\"The sort field. Can be one of `created` or `updated`.\""
	Order        string "name:\"order\" help:\"The sort order if the sort parameter is provided. Can be one of `asc` or `desc`.\""
}

func (c *SearchLabelsCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/search/labels"
	c.updateURLQuery("repository_id", c.RepositoryId)
	c.updateURLQuery("q", c.Q)
	c.updateURLQuery("sort", c.Sort)
	c.updateURLQuery("order", c.Order)
	return c.doRequest("GET")
}

type SearchIssuesLegacyCmd struct {
	baseCmd
	Owner      string `required:"" name:"owner"`
	Repository string `required:"" name:"repository"`
	State      string "required:\"\" name:\"state\" help:\"Indicates the state of the issues to return. Can be either `open` or `closed`.\""
	Keyword    string `required:"" name:"keyword" help:"The search term."`
}

func (c *SearchIssuesLegacyCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/legacy/issues/search/:owner/:repository/:state/:keyword"
	c.updateURLPath("owner", c.Owner)
	c.updateURLPath("repository", c.Repository)
	c.updateURLPath("state", c.State)
	c.updateURLPath("keyword", c.Keyword)
	return c.doRequest("GET")
}

type SearchReposLegacyCmd struct {
	baseCmd
	Keyword   string `required:"" name:"keyword" help:"The search term."`
	Language  string `name:"language" help:"Filter results by language."`
	StartPage string `name:"start_page" help:"The page number to fetch."`
	Sort      string "name:\"sort\" help:\"The sort field. One of `stars`, `forks`, or `updated`.\""
	Order     string "name:\"order\" help:\"The sort field. if `sort` param is provided. Can be either `asc` or `desc`.\""
}

func (c *SearchReposLegacyCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/legacy/repos/search/:keyword"
	c.updateURLPath("keyword", c.Keyword)
	c.updateURLQuery("language", c.Language)
	c.updateURLQuery("start_page", c.StartPage)
	c.updateURLQuery("sort", c.Sort)
	c.updateURLQuery("order", c.Order)
	return c.doRequest("GET")
}

type SearchUsersLegacyCmd struct {
	baseCmd
	Keyword   string `required:"" name:"keyword" help:"The search term."`
	StartPage string `name:"start_page" help:"The page number to fetch."`
	Sort      string "name:\"sort\" help:\"The sort field. One of `stars`, `forks`, or `updated`.\""
	Order     string "name:\"order\" help:\"The sort field. if `sort` param is provided. Can be either `asc` or `desc`.\""
}

func (c *SearchUsersLegacyCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/legacy/user/search/:keyword"
	c.updateURLPath("keyword", c.Keyword)
	c.updateURLQuery("start_page", c.StartPage)
	c.updateURLQuery("sort", c.Sort)
	c.updateURLQuery("order", c.Order)
	return c.doRequest("GET")
}

type SearchEmailLegacyCmd struct {
	baseCmd
	Email string `required:"" name:"email" help:"The email address."`
}

func (c *SearchEmailLegacyCmd) Run(isValueSetMap map[string]bool) error {
	c.isValueSetMap = isValueSetMap
	c.url.Path = "/legacy/user/email/:email"
	c.updateURLPath("email", c.Email)
	return c.doRequest("GET")
}
