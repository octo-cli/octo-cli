// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import "github.com/octo-cli/octo-cli/internal"

type OrgsCmd struct {
	AddOrUpdateMembership              OrgsAddOrUpdateMembershipCmd              `cmd:""`
	BlockUser                          OrgsBlockUserCmd                          `cmd:""`
	CheckBlockedUser                   OrgsCheckBlockedUserCmd                   `cmd:""`
	CheckMembership                    OrgsCheckMembershipCmd                    `cmd:""`
	CheckPublicMembership              OrgsCheckPublicMembershipCmd              `cmd:""`
	ConcealMembership                  OrgsConcealMembershipCmd                  `cmd:""`
	ConvertMemberToOutsideCollaborator OrgsConvertMemberToOutsideCollaboratorCmd `cmd:""`
	CreateInvitation                   OrgsCreateInvitationCmd                   `cmd:""`
	DeleteHook                         OrgsDeleteHookCmd                         `cmd:""`
	Get                                OrgsGetCmd                                `cmd:""`
	GetHook                            OrgsGetHookCmd                            `cmd:""`
	GetMembershipForAuthenticatedUser  OrgsGetMembershipForAuthenticatedUserCmd  `cmd:""`
	GetMembershipForUser               OrgsGetMembershipForUserCmd               `cmd:""`
	List                               OrgsListCmd                               `cmd:""`
	ListBlockedUsers                   OrgsListBlockedUsersCmd                   `cmd:""`
	ListCredentialAuthorizations       OrgsListCredentialAuthorizationsCmd       `cmd:""`
	ListForAuthenticatedUser           OrgsListForAuthenticatedUserCmd           `cmd:""`
	ListForUser                        OrgsListForUserCmd                        `cmd:""`
	ListHooks                          OrgsListHooksCmd                          `cmd:""`
	ListInstallations                  OrgsListInstallationsCmd                  `cmd:""`
	ListInvitationTeams                OrgsListInvitationTeamsCmd                `cmd:""`
	ListMembers                        OrgsListMembersCmd                        `cmd:""`
	ListMemberships                    OrgsListMembershipsCmd                    `cmd:""`
	ListOutsideCollaborators           OrgsListOutsideCollaboratorsCmd           `cmd:""`
	ListPendingInvitations             OrgsListPendingInvitationsCmd             `cmd:""`
	ListPublicMembers                  OrgsListPublicMembersCmd                  `cmd:""`
	PingHook                           OrgsPingHookCmd                           `cmd:""`
	PublicizeMembership                OrgsPublicizeMembershipCmd                `cmd:""`
	RemoveCredentialAuthorization      OrgsRemoveCredentialAuthorizationCmd      `cmd:""`
	RemoveMember                       OrgsRemoveMemberCmd                       `cmd:""`
	RemoveMembership                   OrgsRemoveMembershipCmd                   `cmd:""`
	RemoveOutsideCollaborator          OrgsRemoveOutsideCollaboratorCmd          `cmd:""`
	UnblockUser                        OrgsUnblockUserCmd                        `cmd:""`
	Update                             OrgsUpdateCmd                             `cmd:""`
	UpdateMembership                   OrgsUpdateMembershipCmd                   `cmd:""`
}

type OrgsAddOrUpdateMembershipCmd struct {
	Org      string `required:"" name:"org"`
	Role     string `name:"role"`
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *OrgsAddOrUpdateMembershipCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/memberships/:username")
	c.UpdateURLPath("org", c.Org)
	c.UpdateBody("role", c.Role)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("PUT")
}

type OrgsBlockUserCmd struct {
	Org      string `required:"" name:"org"`
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *OrgsBlockUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/blocks/:username")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("PUT")
}

type OrgsCheckBlockedUserCmd struct {
	Org      string `required:"" name:"org"`
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *OrgsCheckBlockedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/blocks/:username")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("GET")
}

type OrgsCheckMembershipCmd struct {
	Org      string `required:"" name:"org"`
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *OrgsCheckMembershipCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/members/:username")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("GET")
}

type OrgsCheckPublicMembershipCmd struct {
	Org      string `required:"" name:"org"`
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *OrgsCheckPublicMembershipCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/public_members/:username")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("GET")
}

type OrgsConcealMembershipCmd struct {
	Org      string `required:"" name:"org"`
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *OrgsConcealMembershipCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/public_members/:username")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("DELETE")
}

type OrgsConvertMemberToOutsideCollaboratorCmd struct {
	Org      string `required:"" name:"org"`
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *OrgsConvertMemberToOutsideCollaboratorCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/outside_collaborators/:username")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("PUT")
}

type OrgsCreateInvitationCmd struct {
	Email     string  `name:"email"`
	InviteeId int64   `name:"invitee_id"`
	Org       string  `required:"" name:"org"`
	Role      string  `name:"role"`
	TeamIds   []int64 `name:"team_ids"`
	internal.BaseCmd
}

func (c *OrgsCreateInvitationCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/invitations")
	c.UpdateBody("email", c.Email)
	c.UpdateBody("invitee_id", c.InviteeId)
	c.UpdateURLPath("org", c.Org)
	c.UpdateBody("role", c.Role)
	c.UpdateBody("team_ids", c.TeamIds)
	return c.DoRequest("POST")
}

type OrgsDeleteHookCmd struct {
	HookId int64  `required:"" name:"hook_id"`
	Org    string `required:"" name:"org"`
	internal.BaseCmd
}

func (c *OrgsDeleteHookCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/hooks/:hook_id")
	c.UpdateURLPath("hook_id", c.HookId)
	c.UpdateURLPath("org", c.Org)
	return c.DoRequest("DELETE")
}

type OrgsGetCmd struct {
	Org    string `required:"" name:"org"`
	Surtur bool   "name:\"surtur-preview\" help:\"New repository creation permissions are available to preview. You can now use `members_can_create_public_repositories`, `members_can_create_private_repositories`, and `members_can_create_internal_repositories`. You can only allow members to create internal repositories if your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+. These parameters provide more granular permissions to configure the type of repositories organization members can create.\n\nTo access these new parameters during the preview period, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.surtur-preview+json\n```\""
	internal.BaseCmd
}

func (c *OrgsGetCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org")
	c.UpdateURLPath("org", c.Org)
	c.UpdatePreview("surtur", c.Surtur)
	return c.DoRequest("GET")
}

type OrgsGetHookCmd struct {
	HookId int64  `required:"" name:"hook_id"`
	Org    string `required:"" name:"org"`
	internal.BaseCmd
}

func (c *OrgsGetHookCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/hooks/:hook_id")
	c.UpdateURLPath("hook_id", c.HookId)
	c.UpdateURLPath("org", c.Org)
	return c.DoRequest("GET")
}

type OrgsGetMembershipForAuthenticatedUserCmd struct {
	Org string `required:"" name:"org"`
	internal.BaseCmd
}

func (c *OrgsGetMembershipForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/memberships/orgs/:org")
	c.UpdateURLPath("org", c.Org)
	return c.DoRequest("GET")
}

type OrgsGetMembershipForUserCmd struct {
	Org      string `required:"" name:"org"`
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *OrgsGetMembershipForUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/memberships/:username")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("GET")
}

type OrgsListBlockedUsersCmd struct {
	Org string `required:"" name:"org"`
	internal.BaseCmd
}

func (c *OrgsListBlockedUsersCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/blocks")
	c.UpdateURLPath("org", c.Org)
	return c.DoRequest("GET")
}

type OrgsListCmd struct {
	Since int64 `name:"since"`
	internal.BaseCmd
}

func (c *OrgsListCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/organizations")
	c.UpdateURLQuery("since", c.Since)
	return c.DoRequest("GET")
}

type OrgsListCredentialAuthorizationsCmd struct {
	Org string `required:"" name:"org"`
	internal.BaseCmd
}

func (c *OrgsListCredentialAuthorizationsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/credential-authorizations")
	c.UpdateURLPath("org", c.Org)
	return c.DoRequest("GET")
}

type OrgsListForAuthenticatedUserCmd struct {
	Page    int64 `name:"page"`
	PerPage int64 `name:"per_page"`
	internal.BaseCmd
}

func (c *OrgsListForAuthenticatedUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/orgs")
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type OrgsListForUserCmd struct {
	Page     int64  `name:"page"`
	PerPage  int64  `name:"per_page"`
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *OrgsListForUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/users/:username/orgs")
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("GET")
}

type OrgsListHooksCmd struct {
	Org     string `required:"" name:"org"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *OrgsListHooksCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/hooks")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type OrgsListInstallationsCmd struct {
	MachineMan bool   "name:\"machine-man-preview\" required:\"\" help:\"To access the API with your GitHub App, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` Header for your requests.\n\n```\napplication/vnd.github.machine-man-preview+json\n```\""
	Org        string `required:"" name:"org"`
	Page       int64  `name:"page"`
	PerPage    int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *OrgsListInstallationsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/installations")
	c.UpdatePreview("machine-man", c.MachineMan)
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type OrgsListInvitationTeamsCmd struct {
	InvitationId int64  `required:"" name:"invitation_id"`
	Org          string `required:"" name:"org"`
	Page         int64  `name:"page"`
	PerPage      int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *OrgsListInvitationTeamsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/invitations/:invitation_id/teams")
	c.UpdateURLPath("invitation_id", c.InvitationId)
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type OrgsListMembersCmd struct {
	Filter  string `name:"filter"`
	Org     string `required:"" name:"org"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	Role    string `name:"role"`
	internal.BaseCmd
}

func (c *OrgsListMembersCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/members")
	c.UpdateURLQuery("filter", c.Filter)
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("role", c.Role)
	return c.DoRequest("GET")
}

type OrgsListMembershipsCmd struct {
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	State   string `name:"state"`
	internal.BaseCmd
}

func (c *OrgsListMembershipsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/memberships/orgs")
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("state", c.State)
	return c.DoRequest("GET")
}

type OrgsListOutsideCollaboratorsCmd struct {
	Filter  string `name:"filter"`
	Org     string `required:"" name:"org"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *OrgsListOutsideCollaboratorsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/outside_collaborators")
	c.UpdateURLQuery("filter", c.Filter)
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type OrgsListPendingInvitationsCmd struct {
	Org     string `required:"" name:"org"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *OrgsListPendingInvitationsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/invitations")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type OrgsListPublicMembersCmd struct {
	Org     string `required:"" name:"org"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *OrgsListPublicMembersCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/public_members")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	return c.DoRequest("GET")
}

type OrgsPingHookCmd struct {
	HookId int64  `required:"" name:"hook_id"`
	Org    string `required:"" name:"org"`
	internal.BaseCmd
}

func (c *OrgsPingHookCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/hooks/:hook_id/pings")
	c.UpdateURLPath("hook_id", c.HookId)
	c.UpdateURLPath("org", c.Org)
	return c.DoRequest("POST")
}

type OrgsPublicizeMembershipCmd struct {
	Org      string `required:"" name:"org"`
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *OrgsPublicizeMembershipCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/public_members/:username")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("PUT")
}

type OrgsRemoveCredentialAuthorizationCmd struct {
	CredentialId int64  `required:"" name:"credential_id"`
	Org          string `required:"" name:"org"`
	internal.BaseCmd
}

func (c *OrgsRemoveCredentialAuthorizationCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/credential-authorizations/:credential_id")
	c.UpdateURLPath("credential_id", c.CredentialId)
	c.UpdateURLPath("org", c.Org)
	return c.DoRequest("DELETE")
}

type OrgsRemoveMemberCmd struct {
	Org      string `required:"" name:"org"`
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *OrgsRemoveMemberCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/members/:username")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("DELETE")
}

type OrgsRemoveMembershipCmd struct {
	Org      string `required:"" name:"org"`
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *OrgsRemoveMembershipCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/memberships/:username")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("DELETE")
}

type OrgsRemoveOutsideCollaboratorCmd struct {
	Org      string `required:"" name:"org"`
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *OrgsRemoveOutsideCollaboratorCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/outside_collaborators/:username")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("DELETE")
}

type OrgsUnblockUserCmd struct {
	Org      string `required:"" name:"org"`
	Username string `required:"" name:"username"`
	internal.BaseCmd
}

func (c *OrgsUnblockUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org/blocks/:username")
	c.UpdateURLPath("org", c.Org)
	c.UpdateURLPath("username", c.Username)
	return c.DoRequest("DELETE")
}

type OrgsUpdateCmd struct {
	BillingEmail                         string `name:"billing_email"`
	Company                              string `name:"company"`
	DefaultRepositoryPermission          string `name:"default_repository_permission"`
	Description                          string `name:"description"`
	Email                                string `name:"email"`
	HasOrganizationProjects              bool   `name:"has_organization_projects"`
	HasRepositoryProjects                bool   `name:"has_repository_projects"`
	Location                             string `name:"location"`
	MembersAllowedRepositoryCreationType string `name:"members_allowed_repository_creation_type"`
	MembersCanCreateInternalRepositories bool   `name:"members_can_create_internal_repositories"`
	MembersCanCreatePrivateRepositories  bool   `name:"members_can_create_private_repositories"`
	MembersCanCreatePublicRepositories   bool   `name:"members_can_create_public_repositories"`
	MembersCanCreateRepositories         bool   `name:"members_can_create_repositories"`
	Name                                 string `name:"name"`
	Org                                  string `required:"" name:"org"`
	Surtur                               bool   "name:\"surtur-preview\" help:\"New repository creation permissions are available to preview. You can now use `members_can_create_public_repositories`, `members_can_create_private_repositories`, and `members_can_create_internal_repositories`. You can only allow members to create internal repositories if your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+. These parameters provide more granular permissions to configure the type of repositories organization members can create.\n\nTo access these new parameters during the preview period, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.surtur-preview+json\n```\""
	internal.BaseCmd
}

func (c *OrgsUpdateCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/orgs/:org")
	c.UpdateBody("billing_email", c.BillingEmail)
	c.UpdateBody("company", c.Company)
	c.UpdateBody("default_repository_permission", c.DefaultRepositoryPermission)
	c.UpdateBody("description", c.Description)
	c.UpdateBody("email", c.Email)
	c.UpdateBody("has_organization_projects", c.HasOrganizationProjects)
	c.UpdateBody("has_repository_projects", c.HasRepositoryProjects)
	c.UpdateBody("location", c.Location)
	c.UpdateBody("members_allowed_repository_creation_type", c.MembersAllowedRepositoryCreationType)
	c.UpdateBody("members_can_create_internal_repositories", c.MembersCanCreateInternalRepositories)
	c.UpdateBody("members_can_create_private_repositories", c.MembersCanCreatePrivateRepositories)
	c.UpdateBody("members_can_create_public_repositories", c.MembersCanCreatePublicRepositories)
	c.UpdateBody("members_can_create_repositories", c.MembersCanCreateRepositories)
	c.UpdateBody("name", c.Name)
	c.UpdateURLPath("org", c.Org)
	c.UpdatePreview("surtur", c.Surtur)
	return c.DoRequest("PATCH")
}

type OrgsUpdateMembershipCmd struct {
	Org   string `required:"" name:"org"`
	State string `required:"" name:"state"`
	internal.BaseCmd
}

func (c *OrgsUpdateMembershipCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/user/memberships/orgs/:org")
	c.UpdateURLPath("org", c.Org)
	c.UpdateBody("state", c.State)
	return c.DoRequest("PATCH")
}
