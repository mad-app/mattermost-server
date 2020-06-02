// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/mad-app/mattermost-server/v5/model"

// ComplianceInterface is an autogenerated mock type for the ComplianceInterface type
type ComplianceInterface struct {
	mock.Mock
}

// RunComplianceJob provides a mock function with given fields: job
func (_m *ComplianceInterface) RunComplianceJob(job *model.Compliance) *model.AppError {
	ret := _m.Called(job)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.Compliance) *model.AppError); ok {
		r0 = rf(job)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// StartComplianceDailyJob provides a mock function with given fields:
func (_m *ComplianceInterface) StartComplianceDailyJob() {
	_m.Called()
}
