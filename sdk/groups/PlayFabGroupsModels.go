package groups

import "time"
                    
// AcceptGroupApplicationRequest accepts an outstanding invitation to to join a group if the invited entity is not blocked by the group. Nothing is
// returned in the case of success.
type AcceptGroupApplicationRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity optional. Type of the entity to accept as. If specified, must be the same entity as the claimant or an entity that is a
// child of the claimant entity. Defaults to the claimant entity.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
}

// AcceptGroupInvitationRequest accepts an outstanding invitation to join the group if the invited entity is not blocked by the group. Only the invited
// entity or a parent in its chain (e.g. title) may accept the invitation on the invited entity's behalf. Nothing is
// returned in the case of success.
type AcceptGroupInvitationRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
}

// AddMembersRequest adds members to a group or role. Existing members of the group will added to roles within the group, but if the user is
// not already a member of the group, only title claimants may add them to the group, and others must use the group
// application or invite system to add new members to a group. Returns nothing if successful.
type AddMembersRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
    // Members list of entities to add to the group. Only entities of type title_player_account and character may be added to groups.
    Members []EntityKeyModel `json:"Members,omitempty"`
    // RoleId optional: The ID of the existing role to add the entities to. If this is not specified, the default member role for the
// group will be used. Role IDs must be between 1 and 64 characters long.
    RoleId string `json:"RoleId,omitempty"`
}

// ApplyToGroupRequest creates an application to join a group. Calling this while a group application already exists will return the same
// application instead of an error and will not refresh the time before the application expires. By default, if the entity
// has an invitation to join the group outstanding, this will accept the invitation to join the group instead and return an
// error indicating such, rather than creating a duplicate application to join that will need to be cleaned up later.
// Returns information about the application or an error indicating an invitation was accepted instead.
type ApplyToGroupRequestModel struct {
    // AutoAcceptOutstandingInvite optional, default true. Automatically accept an outstanding invitation if one exists instead of creating an application
    AutoAcceptOutstandingInvite bool `json:"AutoAcceptOutstandingInvite"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
}

// ApplyToGroupResponse describes an application to join a group
type ApplyToGroupResponseModel struct {
    // Entity type of entity that requested membership
    Entity *EntityWithLineageModel `json:"Entity,omitempty"`
    // Expires when the application to join will expire and be deleted
    Expires time.Time `json:"Expires,omitempty"`
    // Group iD of the group that the entity requesting membership to
    Group *EntityKeyModel `json:"Group,omitempty"`
}

// BlockEntityRequest blocks a list of entities from joining a group. Blocked entities may not create new applications to join, be invited to
// join, accept an invitation, or have an application accepted. Failure due to being blocked does not clean up existing
// applications or invitations to the group. No data is returned in the case of success.
type BlockEntityRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
}

// ChangeMemberRoleRequest changes the role membership of a list of entities from one role to another in in a single operation. The destination
// role must already exist. This is equivalent to adding the entities to the destination role and removing from the origin
// role. Returns nothing if successful.
type ChangeMemberRoleRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // DestinationRoleId the ID of the role that the entities will become a member of. This must be an existing role. Role IDs must be between 1
// and 64 characters long.
    DestinationRoleId string `json:"DestinationRoleId,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
    // Members list of entities to move between roles in the group. All entities in this list must be members of the group and origin
// role.
    Members []EntityKeyModel `json:"Members,omitempty"`
    // OriginRoleId the ID of the role that the entities currently are a member of. Role IDs must be between 1 and 64 characters long.
    OriginRoleId string `json:"OriginRoleId,omitempty"`
}

// CreateGroupRequest creates a new group, as well as administration and member roles, based off of a title's group template. Returns
// information about the group that was created.
type CreateGroupRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // GroupName the name of the group. This is unique at the title level by default.
    GroupName string `json:"GroupName,omitempty"`
}

// CreateGroupResponse 
type CreateGroupResponseModel struct {
    // AdminRoleId the ID of the administrator role for the group.
    AdminRoleId string `json:"AdminRoleId,omitempty"`
    // Created the server date and time the group was created.
    Created time.Time `json:"Created,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
    // GroupName the name of the group.
    GroupName string `json:"GroupName,omitempty"`
    // MemberRoleId the ID of the default member role for the group.
    MemberRoleId string `json:"MemberRoleId,omitempty"`
    // ProfileVersion the current version of the profile, can be used for concurrency control during updates.
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
    // Roles the list of roles and names that belong to the group.
    Roles map[string]string `json:"Roles,omitempty"`
}

// CreateGroupRoleRequest creates a new role within an existing group, with no members. Both the role ID and role name must be unique within the
// group, but the name can be the same as the ID. The role ID is set at creation and cannot be changed. Returns information
// about the role that was created.
type CreateGroupRoleRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
    // RoleId the ID of the role. This must be unique within the group and cannot be changed. Role IDs must be between 1 and 64
// characters long.
    RoleId string `json:"RoleId,omitempty"`
    // RoleName the name of the role. This must be unique within the group and can be changed later. Role names must be between 1 and
// 100 characters long
    RoleName string `json:"RoleName,omitempty"`
}

// CreateGroupRoleResponse 
type CreateGroupRoleResponseModel struct {
    // ProfileVersion the current version of the group profile, can be used for concurrency control during updates.
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
    // RoleId iD for the role
    RoleId string `json:"RoleId,omitempty"`
    // RoleName the name of the role
    RoleName string `json:"RoleName,omitempty"`
}

// DeleteGroupRequest deletes a group and all roles, invitations, join requests, and blocks associated with it. Permission to delete is only
// required the group itself to execute this action. The group and data cannot be cannot be recovered once removed, but any
// abuse reports about the group will remain. No data is returned in the case of success.
type DeleteGroupRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Group iD of the group or role to remove
    Group* EntityKeyModel `json:"Group,omitempty"`
}

// DeleteRoleRequest returns information about the role
type DeleteRoleRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
    // RoleId the ID of the role to delete. Role IDs must be between 1 and 64 characters long.
    RoleId string `json:"RoleId,omitempty"`
}

// EmptyResponse 
type EmptyResponseModel struct {
}

// EntityKey combined entity type and ID structure which uniquely identifies a single entity.
type EntityKeyModel struct {
    // Id unique ID of the entity.
    Id string `json:"Id,omitempty"`
    // Type entity type. See https://docs.microsoft.com/gaming/playfab/features/data/entities/available-built-in-entity-types
    Type string `json:"Type,omitempty"`
}

// EntityMemberRole 
type EntityMemberRoleModel struct {
    // Members the list of members in the role
    Members []EntityWithLineageModel `json:"Members,omitempty"`
    // RoleId the ID of the role.
    RoleId string `json:"RoleId,omitempty"`
    // RoleName the name of the role
    RoleName string `json:"RoleName,omitempty"`
}

// EntityWithLineage entity wrapper class that contains the entity key and the entities that make up the lineage of the entity.
type EntityWithLineageModel struct {
    // Key the entity key for the specified entity
    Key *EntityKeyModel `json:"Key,omitempty"`
    // Lineage dictionary of entity keys for related entities. Dictionary key is entity type.
    Lineage map[string]EntityKeyModel `json:"Lineage,omitempty"`
}

// GetGroupRequest returns the ID, name, role list and other non-membership related information about a group.
type GetGroupRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Group the identifier of the group
    Group *EntityKeyModel `json:"Group,omitempty"`
    // GroupName the full name of the group
    GroupName string `json:"GroupName,omitempty"`
}

// GetGroupResponse 
type GetGroupResponseModel struct {
    // AdminRoleId the ID of the administrator role for the group.
    AdminRoleId string `json:"AdminRoleId,omitempty"`
    // Created the server date and time the group was created.
    Created time.Time `json:"Created,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
    // GroupName the name of the group.
    GroupName string `json:"GroupName,omitempty"`
    // MemberRoleId the ID of the default member role for the group.
    MemberRoleId string `json:"MemberRoleId,omitempty"`
    // ProfileVersion the current version of the profile, can be used for concurrency control during updates.
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
    // Roles the list of roles and names that belong to the group.
    Roles map[string]string `json:"Roles,omitempty"`
}

// GroupApplication describes an application to join a group
type GroupApplicationModel struct {
    // Entity type of entity that requested membership
    Entity *EntityWithLineageModel `json:"Entity,omitempty"`
    // Expires when the application to join will expire and be deleted
    Expires time.Time `json:"Expires,omitempty"`
    // Group iD of the group that the entity requesting membership to
    Group *EntityKeyModel `json:"Group,omitempty"`
}

// GroupBlock describes an entity that is blocked from joining a group.
type GroupBlockModel struct {
    // Entity the entity that is blocked
    Entity *EntityWithLineageModel `json:"Entity,omitempty"`
    // Group iD of the group that the entity is blocked from
    Group* EntityKeyModel `json:"Group,omitempty"`
}

// GroupInvitation describes an invitation to a group.
type GroupInvitationModel struct {
    // Expires when the invitation will expire and be deleted
    Expires time.Time `json:"Expires,omitempty"`
    // Group the group that the entity invited to
    Group *EntityKeyModel `json:"Group,omitempty"`
    // InvitedByEntity the entity that created the invitation
    InvitedByEntity *EntityWithLineageModel `json:"InvitedByEntity,omitempty"`
    // InvitedEntity the entity that is invited
    InvitedEntity *EntityWithLineageModel `json:"InvitedEntity,omitempty"`
    // RoleId iD of the role in the group to assign the user to.
    RoleId string `json:"RoleId,omitempty"`
}

// GroupRole describes a group role
type GroupRoleModel struct {
    // RoleId iD for the role
    RoleId string `json:"RoleId,omitempty"`
    // RoleName the name of the role
    RoleName string `json:"RoleName,omitempty"`
}

// GroupWithRoles describes a group and the roles that it contains
type GroupWithRolesModel struct {
    // Group iD for the group
    Group *EntityKeyModel `json:"Group,omitempty"`
    // GroupName the name of the group
    GroupName string `json:"GroupName,omitempty"`
    // ProfileVersion the current version of the profile, can be used for concurrency control during updates.
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
    // Roles the list of roles within the group
    Roles []GroupRoleModel `json:"Roles,omitempty"`
}

// InviteToGroupRequest invites a player to join a group, if they are not blocked by the group. An optional role can be provided to
// automatically assign the player to the role if they accept the invitation. By default, if the entity has an application
// to the group outstanding, this will accept the application instead and return an error indicating such, rather than
// creating a duplicate invitation to join that will need to be cleaned up later. Returns information about the new
// invitation or an error indicating an existing application to join was accepted.
type InviteToGroupRequestModel struct {
    // AutoAcceptOutstandingApplication optional, default true. Automatically accept an application if one exists instead of creating an invitation
    AutoAcceptOutstandingApplication bool `json:"AutoAcceptOutstandingApplication"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
    // RoleId optional. ID of an existing a role in the group to assign the user to. The group's default member role is used if this
// is not specified. Role IDs must be between 1 and 64 characters long.
    RoleId string `json:"RoleId,omitempty"`
}

// InviteToGroupResponse describes an invitation to a group.
type InviteToGroupResponseModel struct {
    // Expires when the invitation will expire and be deleted
    Expires time.Time `json:"Expires,omitempty"`
    // Group the group that the entity invited to
    Group *EntityKeyModel `json:"Group,omitempty"`
    // InvitedByEntity the entity that created the invitation
    InvitedByEntity *EntityWithLineageModel `json:"InvitedByEntity,omitempty"`
    // InvitedEntity the entity that is invited
    InvitedEntity *EntityWithLineageModel `json:"InvitedEntity,omitempty"`
    // RoleId iD of the role in the group to assign the user to.
    RoleId string `json:"RoleId,omitempty"`
}

// IsMemberRequest checks to see if an entity is a member of a group or role within the group. A result indicating if the entity is a
// member of the group is returned, or a permission error if the caller does not have permission to read the group's member
// list.
type IsMemberRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
    // RoleId optional: ID of the role to check membership of. Defaults to any role (that is, check to see if the entity is a member
// of the group in any capacity) if not specified.
    RoleId string `json:"RoleId,omitempty"`
}

// IsMemberResponse 
type IsMemberResponseModel struct {
    // IsMember a value indicating whether or not the entity is a member.
    IsMember bool `json:"IsMember"`
}

// ListGroupApplicationsRequest lists all outstanding requests to join a group. Returns a list of all requests to join, as well as when the request will
// expire. To get the group applications for a specific entity, use ListMembershipOpportunities.
type ListGroupApplicationsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
}

// ListGroupApplicationsResponse 
type ListGroupApplicationsResponseModel struct {
    // Applications the requested list of applications to the group.
    Applications []GroupApplicationModel `json:"Applications,omitempty"`
}

// ListGroupBlocksRequest lists all entities blocked from joining a group. A list of blocked entities is returned
type ListGroupBlocksRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
}

// ListGroupBlocksResponse 
type ListGroupBlocksResponseModel struct {
    // BlockedEntities the requested list blocked entities.
    BlockedEntities []GroupBlockModel `json:"BlockedEntities,omitempty"`
}

// ListGroupInvitationsRequest lists all outstanding invitations for a group. Returns a list of entities that have been invited, as well as when the
// invitation will expire. To get the group invitations for a specific entity, use ListMembershipOpportunities.
type ListGroupInvitationsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
}

// ListGroupInvitationsResponse 
type ListGroupInvitationsResponseModel struct {
    // Invitations the requested list of group invitations.
    Invitations []GroupInvitationModel `json:"Invitations,omitempty"`
}

// ListGroupMembersRequest gets a list of members and the roles they belong to within the group. If the caller does not have permission to view the
// role, and the member is in no other role, the member is not displayed. Returns a list of entities that are members of
// the group.
type ListGroupMembersRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Group iD of the group to list the members and roles for
    Group* EntityKeyModel `json:"Group,omitempty"`
}

// ListGroupMembersResponse 
type ListGroupMembersResponseModel struct {
    // Members the requested list of roles and member entity IDs.
    Members []EntityMemberRoleModel `json:"Members,omitempty"`
}

// ListMembershipOpportunitiesRequest lists all outstanding group applications and invitations for an entity. Anyone may call this for any entity, but data
// will only be returned for the entity or a parent of that entity. To list invitations or applications for a group to
// check if a player is trying to join, use ListGroupInvitations and ListGroupApplications.
type ListMembershipOpportunitiesRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// ListMembershipOpportunitiesResponse 
type ListMembershipOpportunitiesResponseModel struct {
    // Applications the requested list of group applications.
    Applications []GroupApplicationModel `json:"Applications,omitempty"`
    // Invitations the requested list of group invitations.
    Invitations []GroupInvitationModel `json:"Invitations,omitempty"`
}

// ListMembershipRequest lists the groups and roles that an entity is a part of, checking to see if group and role metadata and memberships
// should be visible to the caller. If the entity is not in any roles that are visible to the caller, the group is not
// returned in the results, even if the caller otherwise has permission to see that the entity is a member of that group.
type ListMembershipRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// ListMembershipResponse 
type ListMembershipResponseModel struct {
    // Groups the list of groups
    Groups []GroupWithRolesModel `json:"Groups,omitempty"`
}

// OperationTypes 
type OperationTypes string
  
const (
     OperationTypesCreated OperationTypes = "Created"
     OperationTypesUpdated OperationTypes = "Updated"
     OperationTypesDeleted OperationTypes = "Deleted"
     OperationTypesNone OperationTypes = "None"
      )
// RemoveGroupApplicationRequest removes an existing application to join the group. This is used for both rejection of an application as well as
// withdrawing an application. The applying entity or a parent in its chain (e.g. title) may withdraw the application, and
// any caller with appropriate access in the group may reject an application. No data is returned in the case of success.
type RemoveGroupApplicationRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
}

// RemoveGroupInvitationRequest removes an existing invitation to join the group. This is used for both rejection of an invitation as well as rescinding
// an invitation. The invited entity or a parent in its chain (e.g. title) may reject the invitation by calling this
// method, and any caller with appropriate access in the group may rescind an invitation. No data is returned in the case
// of success.
type RemoveGroupInvitationRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
}

// RemoveMembersRequest removes members from a group. A member can always remove themselves from a group, regardless of permissions. Returns
// nothing if successful.
type RemoveMembersRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
    // Members list of entities to remove
    Members []EntityKeyModel `json:"Members,omitempty"`
    // RoleId the ID of the role to remove the entities from.
    RoleId string `json:"RoleId,omitempty"`
}

// UnblockEntityRequest unblocks a list of entities from joining a group. No data is returned in the case of success.
type UnblockEntityRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
}

// UpdateGroupRequest updates data about a group, such as the name or default member role. Returns information about whether the update was
// successful. Only title claimants may modify the administration role for a group.
type UpdateGroupRequestModel struct {
    // AdminRoleId optional: the ID of an existing role to set as the new administrator role for the group
    AdminRoleId string `json:"AdminRoleId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ExpectedProfileVersion optional field used for concurrency control. By specifying the previously returned value of ProfileVersion from the
// GetGroup API, you can ensure that the group data update will only be performed if the group has not been updated by any
// other clients since the version you last loaded.
    ExpectedProfileVersion int32 `json:"ExpectedProfileVersion,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
    // GroupName optional: the new name of the group
    GroupName string `json:"GroupName,omitempty"`
    // MemberRoleId optional: the ID of an existing role to set as the new member role for the group
    MemberRoleId string `json:"MemberRoleId,omitempty"`
}

// UpdateGroupResponse 
type UpdateGroupResponseModel struct {
    // OperationReason optional reason to explain why the operation was the result that it was.
    OperationReason string `json:"OperationReason,omitempty"`
    // ProfileVersion new version of the group data.
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
    // SetResult indicates which operation was completed, either Created, Updated, Deleted or None.
    SetResult OperationTypes `json:"SetResult,omitempty"`
}

// UpdateGroupRoleRequest updates the role name. Returns information about whether the update was successful.
type UpdateGroupRoleRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ExpectedProfileVersion optional field used for concurrency control. By specifying the previously returned value of ProfileVersion from the
// GetGroup API, you can ensure that the group data update will only be performed if the group has not been updated by any
// other clients since the version you last loaded.
    ExpectedProfileVersion int32 `json:"ExpectedProfileVersion,omitempty"`
    // Group the identifier of the group
    Group* EntityKeyModel `json:"Group,omitempty"`
    // RoleId iD of the role to update. Role IDs must be between 1 and 64 characters long.
    RoleId string `json:"RoleId,omitempty"`
    // RoleName the new name of the role
    RoleName string `json:"RoleName,omitempty"`
}

// UpdateGroupRoleResponse 
type UpdateGroupRoleResponseModel struct {
    // OperationReason optional reason to explain why the operation was the result that it was.
    OperationReason string `json:"OperationReason,omitempty"`
    // ProfileVersion new version of the role data.
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
    // SetResult indicates which operation was completed, either Created, Updated, Deleted or None.
    SetResult OperationTypes `json:"SetResult,omitempty"`
}
