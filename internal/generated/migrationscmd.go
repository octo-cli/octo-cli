// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import internal "github.com/octo-cli/octo-cli/internal"

type MigrationsCmd struct {
	CancelImport                      MigrationsCancelImportCmd                      `cmd:""`
	DeleteArchiveForAuthenticatedUser MigrationsDeleteArchiveForAuthenticatedUserCmd `cmd:""`
	DeleteArchiveForOrg               MigrationsDeleteArchiveForOrgCmd               `cmd:""`
	DownloadArchiveForOrg             MigrationsDownloadArchiveForOrgCmd             `cmd:""`
	GetArchiveForAuthenticatedUser    MigrationsGetArchiveForAuthenticatedUserCmd    `cmd:""`
	GetCommitAuthors                  MigrationsGetCommitAuthorsCmd                  `cmd:""`
	GetImportStatus                   MigrationsGetImportStatusCmd                   `cmd:""`
	GetLargeFiles                     MigrationsGetLargeFilesCmd                     `cmd:""`
	GetStatusForAuthenticatedUser     MigrationsGetStatusForAuthenticatedUserCmd     `cmd:""`
	GetStatusForOrg                   MigrationsGetStatusForOrgCmd                   `cmd:""`
	ListForAuthenticatedUser          MigrationsListForAuthenticatedUserCmd          `cmd:""`
	ListForOrg                        MigrationsListForOrgCmd                        `cmd:""`
	ListReposForOrg                   MigrationsListReposForOrgCmd                   `cmd:""`
	ListReposForUser                  MigrationsListReposForUserCmd                  `cmd:""`
	MapCommitAuthor                   MigrationsMapCommitAuthorCmd                   `cmd:""`
	SetLfsPreference                  MigrationsSetLfsPreferenceCmd                  `cmd:""`
	StartForAuthenticatedUser         MigrationsStartForAuthenticatedUserCmd         `cmd:""`
	StartForOrg                       MigrationsStartForOrgCmd                       `cmd:""`
	StartImport                       MigrationsStartImportCmd                       `cmd:""`
	UnlockRepoForAuthenticatedUser    MigrationsUnlockRepoForAuthenticatedUserCmd    `cmd:""`
	UnlockRepoForOrg                  MigrationsUnlockRepoForOrgCmd                  `cmd:""`
	UpdateImport                      MigrationsUpdateImportCmd                      `cmd:""`
}

type MigrationsCancelImportCmd struct {
	Repo string `name:"repo" required:"true"`
	internal.BaseCmd
}

func (c *MigrationsCancelImportCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/import")
	c.UpdateURLPath("repo", c.Repo)
	return c.DoRequest("DELETE")
}

type MigrationsDeleteArchiveForAuthenticatedUserCmd struct {
	Wyandotte   bool  `name:"wyandotte-preview" required:"true"`
	MigrationId int64 `name:"migration_id" required:"true"`
	internal.BaseCmd
}

func (c *MigrationsDeleteArchiveForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/migrations/{migration_id}/archive")
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("DELETE")
}

type MigrationsDeleteArchiveForOrgCmd struct {
	Wyandotte   bool   `name:"wyandotte-preview" required:"true"`
	Org         string `name:"org" required:"true"`
	MigrationId int64  `name:"migration_id" required:"true"`
	internal.BaseCmd
}

func (c *MigrationsDeleteArchiveForOrgCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/{org}/migrations/{migration_id}/archive")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("DELETE")
}

type MigrationsDownloadArchiveForOrgCmd struct {
	Wyandotte   bool   `name:"wyandotte-preview" required:"true"`
	Org         string `name:"org" required:"true"`
	MigrationId int64  `name:"migration_id" required:"true"`
	internal.BaseCmd
}

func (c *MigrationsDownloadArchiveForOrgCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/{org}/migrations/{migration_id}/archive")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("GET")
}

type MigrationsGetArchiveForAuthenticatedUserCmd struct {
	Wyandotte   bool  `name:"wyandotte-preview" required:"true"`
	MigrationId int64 `name:"migration_id" required:"true"`
	internal.BaseCmd
}

func (c *MigrationsGetArchiveForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/migrations/{migration_id}/archive")
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("GET")
}

type MigrationsGetCommitAuthorsCmd struct {
	Repo  string `name:"repo" required:"true"`
	Since string `name:"since"`
	internal.BaseCmd
}

func (c *MigrationsGetCommitAuthorsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/import/authors")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLQuery("since", c.Since)
	return c.DoRequest("GET")
}

type MigrationsGetImportStatusCmd struct {
	Repo string `name:"repo" required:"true"`
	internal.BaseCmd
}

func (c *MigrationsGetImportStatusCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/import")
	c.UpdateURLPath("repo", c.Repo)
	return c.DoRequest("GET")
}

type MigrationsGetLargeFilesCmd struct {
	Repo string `name:"repo" required:"true"`
	internal.BaseCmd
}

func (c *MigrationsGetLargeFilesCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/import/large_files")
	c.UpdateURLPath("repo", c.Repo)
	return c.DoRequest("GET")
}

type MigrationsGetStatusForAuthenticatedUserCmd struct {
	Wyandotte   bool     `name:"wyandotte-preview" required:"true"`
	MigrationId int64    `name:"migration_id" required:"true"`
	Exclude     []string `name:"exclude"`
	internal.BaseCmd
}

func (c *MigrationsGetStatusForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/migrations/{migration_id}")
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdateURLQuery("exclude", c.Exclude)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("GET")
}

type MigrationsGetStatusForOrgCmd struct {
	Wyandotte   bool   `name:"wyandotte-preview" required:"true"`
	Org         string `name:"org" required:"true"`
	MigrationId int64  `name:"migration_id" required:"true"`
	internal.BaseCmd
}

func (c *MigrationsGetStatusForOrgCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/{org}/migrations/{migration_id}")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("GET")
}

type MigrationsListForAuthenticatedUserCmd struct {
	Wyandotte bool  `name:"wyandotte-preview" required:"true"`
	Page      int64 `name:"page"`
	PerPage   int64 `name:"per_page"`
	internal.BaseCmd
}

func (c *MigrationsListForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/migrations")
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("GET")
}

type MigrationsListForOrgCmd struct {
	Wyandotte bool   `name:"wyandotte-preview" required:"true"`
	Org       string `name:"org" required:"true"`
	Page      int64  `name:"page"`
	PerPage   int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *MigrationsListForOrgCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/{org}/migrations")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("GET")
}

type MigrationsListReposForOrgCmd struct {
	Wyandotte   bool   `name:"wyandotte-preview" required:"true"`
	Org         string `name:"org" required:"true"`
	MigrationId int64  `name:"migration_id" required:"true"`
	Page        int64  `name:"page"`
	PerPage     int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *MigrationsListReposForOrgCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/{org}/migrations/{migration_id}/repositories")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("GET")
}

type MigrationsListReposForUserCmd struct {
	Wyandotte   bool  `name:"wyandotte-preview" required:"true"`
	MigrationId int64 `name:"migration_id" required:"true"`
	Page        int64 `name:"page"`
	PerPage     int64 `name:"per_page"`
	internal.BaseCmd
}

func (c *MigrationsListReposForUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/migrations/{migration_id}/repositories")
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("GET")
}

type MigrationsMapCommitAuthorCmd struct {
	Repo     string `name:"repo" required:"true"`
	AuthorId int64  `name:"author_id" required:"true"`
	Email    string `name:"email"`
	Name     string `name:"name"`
	RemoteId string `name:"remote_id"`
	internal.BaseCmd
}

func (c *MigrationsMapCommitAuthorCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/import/authors/{author_id}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("author_id", c.AuthorId)
	c.UpdateBody("email", c.Email)
	c.UpdateBody("name", c.Name)
	c.UpdateBody("remote_id", c.RemoteId)
	return c.DoRequest("PATCH")
}

type MigrationsSetLfsPreferenceCmd struct {
	Repo   string `name:"repo" required:"true"`
	UseLfs string `name:"use_lfs" required:"true"`
	internal.BaseCmd
}

func (c *MigrationsSetLfsPreferenceCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/import/lfs")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("use_lfs", c.UseLfs)
	return c.DoRequest("PATCH")
}

type MigrationsStartForAuthenticatedUserCmd struct {
	Exclude            []string `name:"exclude"`
	ExcludeAttachments bool     `name:"exclude_attachments"`
	LockRepositories   bool     `name:"lock_repositories"`
	Repositories       []string `name:"repositories" required:"true"`
	internal.BaseCmd
}

func (c *MigrationsStartForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/migrations")
	c.UpdateBody("exclude", c.Exclude)
	c.UpdateBody("exclude_attachments", c.ExcludeAttachments)
	c.UpdateBody("lock_repositories", c.LockRepositories)
	c.UpdateBody("repositories", c.Repositories)
	return c.DoRequest("POST")
}

type MigrationsStartForOrgCmd struct {
	Org                string   `name:"org" required:"true"`
	Exclude            []string `name:"exclude"`
	ExcludeAttachments bool     `name:"exclude_attachments"`
	LockRepositories   bool     `name:"lock_repositories"`
	Repositories       []string `name:"repositories" required:"true"`
	internal.BaseCmd
}

func (c *MigrationsStartForOrgCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/{org}/migrations")
	c.UpdateURLPath("org", c.Org)
	c.UpdateBody("exclude", c.Exclude)
	c.UpdateBody("exclude_attachments", c.ExcludeAttachments)
	c.UpdateBody("lock_repositories", c.LockRepositories)
	c.UpdateBody("repositories", c.Repositories)
	return c.DoRequest("POST")
}

type MigrationsStartImportCmd struct {
	Repo        string `name:"repo" required:"true"`
	TfvcProject string `name:"tfvc_project"`
	Vcs         string `name:"vcs"`
	VcsPassword string `name:"vcs_password"`
	VcsUsername string `name:"vcs_username"`
	VcsUrl      string `name:"vcs_url" required:"true"`
	internal.BaseCmd
}

func (c *MigrationsStartImportCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/import")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("tfvc_project", c.TfvcProject)
	c.UpdateBody("vcs", c.Vcs)
	c.UpdateBody("vcs_password", c.VcsPassword)
	c.UpdateBody("vcs_url", c.VcsUrl)
	c.UpdateBody("vcs_username", c.VcsUsername)
	return c.DoRequest("PUT")
}

type MigrationsUnlockRepoForAuthenticatedUserCmd struct {
	Wyandotte   bool   `name:"wyandotte-preview" required:"true"`
	MigrationId int64  `name:"migration_id" required:"true"`
	RepoName    string `name:"repo_name" required:"true"`
	internal.BaseCmd
}

func (c *MigrationsUnlockRepoForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/migrations/{migration_id}/repos/{repo_name}/lock")
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdateURLPath("repo_name", c.RepoName)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("DELETE")
}

type MigrationsUnlockRepoForOrgCmd struct {
	Wyandotte   bool   `name:"wyandotte-preview" required:"true"`
	Org         string `name:"org" required:"true"`
	MigrationId int64  `name:"migration_id" required:"true"`
	RepoName    string `name:"repo_name" required:"true"`
	internal.BaseCmd
}

func (c *MigrationsUnlockRepoForOrgCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/{org}/migrations/{migration_id}/repos/{repo_name}/lock")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("migration_id", c.MigrationId)
	c.UpdateURLPath("repo_name", c.RepoName)
	c.UpdatePreview("wyandotte", c.Wyandotte)
	return c.DoRequest("DELETE")
}

type MigrationsUpdateImportCmd struct {
	Repo        string `name:"repo" required:"true"`
	TfvcProject string `name:"tfvc_project"`
	Vcs         string `name:"vcs"`
	VcsPassword string `name:"vcs_password"`
	VcsUsername string `name:"vcs_username"`
	internal.BaseCmd
}

func (c *MigrationsUpdateImportCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/import")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("tfvc_project", c.TfvcProject)
	c.UpdateBody("vcs", c.Vcs)
	c.UpdateBody("vcs_password", c.VcsPassword)
	c.UpdateBody("vcs_username", c.VcsUsername)
	return c.DoRequest("PATCH")
}
