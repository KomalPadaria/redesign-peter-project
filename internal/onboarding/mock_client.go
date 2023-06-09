// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package onboarding is a generated GoMock package.
package onboarding

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// UpdateOnboardingStatus mocks base method.
func (m *MockClient) UpdateOnboardingStatus(ctx context.Context, onboardingStep int, companyUuid, updated_by uuid.UUID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateOnboardingStatus", ctx, onboardingStep, companyUuid, updated_by)
}

// UpdateOnboardingStatus indicates an expected call of UpdateOnboardingStatus.
func (mr *MockClientMockRecorder) UpdateOnboardingStatus(ctx, onboardingStep, companyUuid, updated_by interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOnboardingStatus", reflect.TypeOf((*MockClient)(nil).UpdateOnboardingStatus), ctx, onboardingStep, companyUuid, updated_by)
}
