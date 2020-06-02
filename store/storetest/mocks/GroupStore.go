// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mad-app/mattermost-server/v5/model"
	mock "github.com/stretchr/testify/mock"
)

// GroupStore is an autogenerated mock type for the GroupStore type
type GroupStore struct {
	mock.Mock
}

// AdminRoleGroupsForSyncableMember provides a mock function with given fields: userID, syncableID, syncableType
func (_m *GroupStore) AdminRoleGroupsForSyncableMember(userID string, syncableID string, syncableType model.GroupSyncableType) ([]string, *model.AppError) {
	ret := _m.Called(userID, syncableID, syncableType)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, string, model.GroupSyncableType) []string); ok {
		r0 = rf(userID, syncableID, syncableType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string, model.GroupSyncableType) *model.AppError); ok {
		r1 = rf(userID, syncableID, syncableType)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// ChannelMembersMinusGroupMembers provides a mock function with given fields: channelID, groupIDs, page, perPage
func (_m *GroupStore) ChannelMembersMinusGroupMembers(channelID string, groupIDs []string, page int, perPage int) ([]*model.UserWithGroups, *model.AppError) {
	ret := _m.Called(channelID, groupIDs, page, perPage)

	var r0 []*model.UserWithGroups
	if rf, ok := ret.Get(0).(func(string, []string, int, int) []*model.UserWithGroups); ok {
		r0 = rf(channelID, groupIDs, page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.UserWithGroups)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, []string, int, int) *model.AppError); ok {
		r1 = rf(channelID, groupIDs, page, perPage)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// ChannelMembersToAdd provides a mock function with given fields: since, channelID
func (_m *GroupStore) ChannelMembersToAdd(since int64, channelID *string) ([]*model.UserChannelIDPair, *model.AppError) {
	ret := _m.Called(since, channelID)

	var r0 []*model.UserChannelIDPair
	if rf, ok := ret.Get(0).(func(int64, *string) []*model.UserChannelIDPair); ok {
		r0 = rf(since, channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.UserChannelIDPair)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(int64, *string) *model.AppError); ok {
		r1 = rf(since, channelID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// ChannelMembersToRemove provides a mock function with given fields: channelID
func (_m *GroupStore) ChannelMembersToRemove(channelID *string) ([]*model.ChannelMember, *model.AppError) {
	ret := _m.Called(channelID)

	var r0 []*model.ChannelMember
	if rf, ok := ret.Get(0).(func(*string) []*model.ChannelMember); ok {
		r0 = rf(channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ChannelMember)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*string) *model.AppError); ok {
		r1 = rf(channelID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// CountChannelMembersMinusGroupMembers provides a mock function with given fields: channelID, groupIDs
func (_m *GroupStore) CountChannelMembersMinusGroupMembers(channelID string, groupIDs []string) (int64, *model.AppError) {
	ret := _m.Called(channelID, groupIDs)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, []string) int64); ok {
		r0 = rf(channelID, groupIDs)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, []string) *model.AppError); ok {
		r1 = rf(channelID, groupIDs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// CountGroupsByChannel provides a mock function with given fields: channelId, opts
func (_m *GroupStore) CountGroupsByChannel(channelId string, opts model.GroupSearchOpts) (int64, *model.AppError) {
	ret := _m.Called(channelId, opts)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, model.GroupSearchOpts) int64); ok {
		r0 = rf(channelId, opts)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, model.GroupSearchOpts) *model.AppError); ok {
		r1 = rf(channelId, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// CountGroupsByTeam provides a mock function with given fields: teamId, opts
func (_m *GroupStore) CountGroupsByTeam(teamId string, opts model.GroupSearchOpts) (int64, *model.AppError) {
	ret := _m.Called(teamId, opts)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, model.GroupSearchOpts) int64); ok {
		r0 = rf(teamId, opts)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, model.GroupSearchOpts) *model.AppError); ok {
		r1 = rf(teamId, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// CountTeamMembersMinusGroupMembers provides a mock function with given fields: teamID, groupIDs
func (_m *GroupStore) CountTeamMembersMinusGroupMembers(teamID string, groupIDs []string) (int64, *model.AppError) {
	ret := _m.Called(teamID, groupIDs)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, []string) int64); ok {
		r0 = rf(teamID, groupIDs)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, []string) *model.AppError); ok {
		r1 = rf(teamID, groupIDs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// Create provides a mock function with given fields: group
func (_m *GroupStore) Create(group *model.Group) (*model.Group, *model.AppError) {
	ret := _m.Called(group)

	var r0 *model.Group
	if rf, ok := ret.Get(0).(func(*model.Group) *model.Group); ok {
		r0 = rf(group)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Group)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*model.Group) *model.AppError); ok {
		r1 = rf(group)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// CreateGroupSyncable provides a mock function with given fields: groupSyncable
func (_m *GroupStore) CreateGroupSyncable(groupSyncable *model.GroupSyncable) (*model.GroupSyncable, *model.AppError) {
	ret := _m.Called(groupSyncable)

	var r0 *model.GroupSyncable
	if rf, ok := ret.Get(0).(func(*model.GroupSyncable) *model.GroupSyncable); ok {
		r0 = rf(groupSyncable)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GroupSyncable)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*model.GroupSyncable) *model.AppError); ok {
		r1 = rf(groupSyncable)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// Delete provides a mock function with given fields: groupID
func (_m *GroupStore) Delete(groupID string) (*model.Group, *model.AppError) {
	ret := _m.Called(groupID)

	var r0 *model.Group
	if rf, ok := ret.Get(0).(func(string) *model.Group); ok {
		r0 = rf(groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Group)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(groupID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// DeleteGroupSyncable provides a mock function with given fields: groupID, syncableID, syncableType
func (_m *GroupStore) DeleteGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) (*model.GroupSyncable, *model.AppError) {
	ret := _m.Called(groupID, syncableID, syncableType)

	var r0 *model.GroupSyncable
	if rf, ok := ret.Get(0).(func(string, string, model.GroupSyncableType) *model.GroupSyncable); ok {
		r0 = rf(groupID, syncableID, syncableType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GroupSyncable)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string, model.GroupSyncableType) *model.AppError); ok {
		r1 = rf(groupID, syncableID, syncableType)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// DeleteMember provides a mock function with given fields: groupID, userID
func (_m *GroupStore) DeleteMember(groupID string, userID string) (*model.GroupMember, *model.AppError) {
	ret := _m.Called(groupID, userID)

	var r0 *model.GroupMember
	if rf, ok := ret.Get(0).(func(string, string) *model.GroupMember); ok {
		r0 = rf(groupID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GroupMember)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string) *model.AppError); ok {
		r1 = rf(groupID, userID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// Get provides a mock function with given fields: groupID
func (_m *GroupStore) Get(groupID string) (*model.Group, *model.AppError) {
	ret := _m.Called(groupID)

	var r0 *model.Group
	if rf, ok := ret.Get(0).(func(string) *model.Group); ok {
		r0 = rf(groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Group)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(groupID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetAllBySource provides a mock function with given fields: groupSource
func (_m *GroupStore) GetAllBySource(groupSource model.GroupSource) ([]*model.Group, *model.AppError) {
	ret := _m.Called(groupSource)

	var r0 []*model.Group
	if rf, ok := ret.Get(0).(func(model.GroupSource) []*model.Group); ok {
		r0 = rf(groupSource)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Group)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(model.GroupSource) *model.AppError); ok {
		r1 = rf(groupSource)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetAllGroupSyncablesByGroupId provides a mock function with given fields: groupID, syncableType
func (_m *GroupStore) GetAllGroupSyncablesByGroupId(groupID string, syncableType model.GroupSyncableType) ([]*model.GroupSyncable, *model.AppError) {
	ret := _m.Called(groupID, syncableType)

	var r0 []*model.GroupSyncable
	if rf, ok := ret.Get(0).(func(string, model.GroupSyncableType) []*model.GroupSyncable); ok {
		r0 = rf(groupID, syncableType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.GroupSyncable)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, model.GroupSyncableType) *model.AppError); ok {
		r1 = rf(groupID, syncableType)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetByIDs provides a mock function with given fields: groupIDs
func (_m *GroupStore) GetByIDs(groupIDs []string) ([]*model.Group, *model.AppError) {
	ret := _m.Called(groupIDs)

	var r0 []*model.Group
	if rf, ok := ret.Get(0).(func([]string) []*model.Group); ok {
		r0 = rf(groupIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Group)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func([]string) *model.AppError); ok {
		r1 = rf(groupIDs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: name
func (_m *GroupStore) GetByName(name string) (*model.Group, *model.AppError) {
	ret := _m.Called(name)

	var r0 *model.Group
	if rf, ok := ret.Get(0).(func(string) *model.Group); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Group)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetByRemoteID provides a mock function with given fields: remoteID, groupSource
func (_m *GroupStore) GetByRemoteID(remoteID string, groupSource model.GroupSource) (*model.Group, *model.AppError) {
	ret := _m.Called(remoteID, groupSource)

	var r0 *model.Group
	if rf, ok := ret.Get(0).(func(string, model.GroupSource) *model.Group); ok {
		r0 = rf(remoteID, groupSource)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Group)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, model.GroupSource) *model.AppError); ok {
		r1 = rf(remoteID, groupSource)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetByUser provides a mock function with given fields: userId
func (_m *GroupStore) GetByUser(userId string) ([]*model.Group, *model.AppError) {
	ret := _m.Called(userId)

	var r0 []*model.Group
	if rf, ok := ret.Get(0).(func(string) []*model.Group); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Group)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(userId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetGroupSyncable provides a mock function with given fields: groupID, syncableID, syncableType
func (_m *GroupStore) GetGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) (*model.GroupSyncable, *model.AppError) {
	ret := _m.Called(groupID, syncableID, syncableType)

	var r0 *model.GroupSyncable
	if rf, ok := ret.Get(0).(func(string, string, model.GroupSyncableType) *model.GroupSyncable); ok {
		r0 = rf(groupID, syncableID, syncableType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GroupSyncable)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string, model.GroupSyncableType) *model.AppError); ok {
		r1 = rf(groupID, syncableID, syncableType)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetGroups provides a mock function with given fields: page, perPage, opts
func (_m *GroupStore) GetGroups(page int, perPage int, opts model.GroupSearchOpts) ([]*model.Group, *model.AppError) {
	ret := _m.Called(page, perPage, opts)

	var r0 []*model.Group
	if rf, ok := ret.Get(0).(func(int, int, model.GroupSearchOpts) []*model.Group); ok {
		r0 = rf(page, perPage, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Group)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(int, int, model.GroupSearchOpts) *model.AppError); ok {
		r1 = rf(page, perPage, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetGroupsByChannel provides a mock function with given fields: channelId, opts
func (_m *GroupStore) GetGroupsByChannel(channelId string, opts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, *model.AppError) {
	ret := _m.Called(channelId, opts)

	var r0 []*model.GroupWithSchemeAdmin
	if rf, ok := ret.Get(0).(func(string, model.GroupSearchOpts) []*model.GroupWithSchemeAdmin); ok {
		r0 = rf(channelId, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.GroupWithSchemeAdmin)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, model.GroupSearchOpts) *model.AppError); ok {
		r1 = rf(channelId, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetGroupsByTeam provides a mock function with given fields: teamId, opts
func (_m *GroupStore) GetGroupsByTeam(teamId string, opts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, *model.AppError) {
	ret := _m.Called(teamId, opts)

	var r0 []*model.GroupWithSchemeAdmin
	if rf, ok := ret.Get(0).(func(string, model.GroupSearchOpts) []*model.GroupWithSchemeAdmin); ok {
		r0 = rf(teamId, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.GroupWithSchemeAdmin)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, model.GroupSearchOpts) *model.AppError); ok {
		r1 = rf(teamId, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetMemberCount provides a mock function with given fields: groupID
func (_m *GroupStore) GetMemberCount(groupID string) (int64, *model.AppError) {
	ret := _m.Called(groupID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(groupID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(groupID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetMemberUsers provides a mock function with given fields: groupID
func (_m *GroupStore) GetMemberUsers(groupID string) ([]*model.User, *model.AppError) {
	ret := _m.Called(groupID)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string) []*model.User); ok {
		r0 = rf(groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(groupID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetMemberUsersPage provides a mock function with given fields: groupID, page, perPage
func (_m *GroupStore) GetMemberUsersPage(groupID string, page int, perPage int) ([]*model.User, *model.AppError) {
	ret := _m.Called(groupID, page, perPage)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, int, int) []*model.User); ok {
		r0 = rf(groupID, page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, int, int) *model.AppError); ok {
		r1 = rf(groupID, page, perPage)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// PermanentDeleteMembersByUser provides a mock function with given fields: userId
func (_m *GroupStore) PermanentDeleteMembersByUser(userId string) *model.AppError {
	ret := _m.Called(userId)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string) *model.AppError); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// PermittedSyncableAdmins provides a mock function with given fields: syncableID, syncableType
func (_m *GroupStore) PermittedSyncableAdmins(syncableID string, syncableType model.GroupSyncableType) ([]string, *model.AppError) {
	ret := _m.Called(syncableID, syncableType)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, model.GroupSyncableType) []string); ok {
		r0 = rf(syncableID, syncableType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, model.GroupSyncableType) *model.AppError); ok {
		r1 = rf(syncableID, syncableType)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// TeamMembersMinusGroupMembers provides a mock function with given fields: teamID, groupIDs, page, perPage
func (_m *GroupStore) TeamMembersMinusGroupMembers(teamID string, groupIDs []string, page int, perPage int) ([]*model.UserWithGroups, *model.AppError) {
	ret := _m.Called(teamID, groupIDs, page, perPage)

	var r0 []*model.UserWithGroups
	if rf, ok := ret.Get(0).(func(string, []string, int, int) []*model.UserWithGroups); ok {
		r0 = rf(teamID, groupIDs, page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.UserWithGroups)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, []string, int, int) *model.AppError); ok {
		r1 = rf(teamID, groupIDs, page, perPage)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// TeamMembersToAdd provides a mock function with given fields: since, teamID
func (_m *GroupStore) TeamMembersToAdd(since int64, teamID *string) ([]*model.UserTeamIDPair, *model.AppError) {
	ret := _m.Called(since, teamID)

	var r0 []*model.UserTeamIDPair
	if rf, ok := ret.Get(0).(func(int64, *string) []*model.UserTeamIDPair); ok {
		r0 = rf(since, teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.UserTeamIDPair)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(int64, *string) *model.AppError); ok {
		r1 = rf(since, teamID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// TeamMembersToRemove provides a mock function with given fields: teamID
func (_m *GroupStore) TeamMembersToRemove(teamID *string) ([]*model.TeamMember, *model.AppError) {
	ret := _m.Called(teamID)

	var r0 []*model.TeamMember
	if rf, ok := ret.Get(0).(func(*string) []*model.TeamMember); ok {
		r0 = rf(teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TeamMember)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*string) *model.AppError); ok {
		r1 = rf(teamID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: group
func (_m *GroupStore) Update(group *model.Group) (*model.Group, *model.AppError) {
	ret := _m.Called(group)

	var r0 *model.Group
	if rf, ok := ret.Get(0).(func(*model.Group) *model.Group); ok {
		r0 = rf(group)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Group)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*model.Group) *model.AppError); ok {
		r1 = rf(group)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// UpdateGroupSyncable provides a mock function with given fields: groupSyncable
func (_m *GroupStore) UpdateGroupSyncable(groupSyncable *model.GroupSyncable) (*model.GroupSyncable, *model.AppError) {
	ret := _m.Called(groupSyncable)

	var r0 *model.GroupSyncable
	if rf, ok := ret.Get(0).(func(*model.GroupSyncable) *model.GroupSyncable); ok {
		r0 = rf(groupSyncable)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GroupSyncable)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*model.GroupSyncable) *model.AppError); ok {
		r1 = rf(groupSyncable)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// UpsertMember provides a mock function with given fields: groupID, userID
func (_m *GroupStore) UpsertMember(groupID string, userID string) (*model.GroupMember, *model.AppError) {
	ret := _m.Called(groupID, userID)

	var r0 *model.GroupMember
	if rf, ok := ret.Get(0).(func(string, string) *model.GroupMember); ok {
		r0 = rf(groupID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GroupMember)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string) *model.AppError); ok {
		r1 = rf(groupID, userID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}
