// Code generated by MockGen. DO NOT EDIT.
// Source: chgm.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	cloudtrail "github.com/aws/aws-sdk-go/service/cloudtrail"
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	gomock "github.com/golang/mock/gomock"
	ocm "github.com/openshift/configuration-anomaly-detection/pkg/ocm"
	pagerduty "github.com/openshift/configuration-anomaly-detection/pkg/pagerduty"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// AddNote mocks base method.
func (m *MockService) AddNote(notes string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNote", notes)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNote indicates an expected call of AddNote.
func (mr *MockServiceMockRecorder) AddNote(notes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNote", reflect.TypeOf((*MockService)(nil).AddNote), notes)
}

// CreateNewAlert mocks base method.
func (m *MockService) CreateNewAlert(newAlert pagerduty.NewAlert, serviceID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewAlert", newAlert, serviceID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNewAlert indicates an expected call of CreateNewAlert.
func (mr *MockServiceMockRecorder) CreateNewAlert(newAlert, serviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewAlert", reflect.TypeOf((*MockService)(nil).CreateNewAlert), newAlert, serviceID)
}

// DeleteLimitedSupportReasons mocks base method.
func (m *MockService) DeleteLimitedSupportReasons(ls ocm.LimitedSupportReason, clusterID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLimitedSupportReasons", ls, clusterID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLimitedSupportReasons indicates an expected call of DeleteLimitedSupportReasons.
func (mr *MockServiceMockRecorder) DeleteLimitedSupportReasons(ls, clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLimitedSupportReasons", reflect.TypeOf((*MockService)(nil).DeleteLimitedSupportReasons), ls, clusterID)
}

// GetServiceID mocks base method.
func (m *MockService) GetServiceID() string {
	m.ctrl.T.Helper()
<<<<<<< HEAD
	ret := m.ctrl.Call(m, "DeleteCHGMLimitedSupportReason", clusterID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCHGMLimitedSupportReason indicates an expected call of DeleteCHGMLimitedSupportReason.
func (mr *MockServiceMockRecorder) DeleteCHGMLimitedSupportReason(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCHGMLimitedSupportReason", reflect.TypeOf((*MockService)(nil).DeleteCHGMLimitedSupportReason), clusterID)
}

// ExtractServiceIDFromBytes mocks base method.
func (m *MockService) ExtractServiceIDFromBytes(data []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractServiceIDFromBytes", data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExtractServiceIDFromBytes indicates an expected call of ExtractServiceIDFromBytes.
func (mr *MockServiceMockRecorder) ExtractServiceIDFromBytes(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractServiceIDFromBytes", reflect.TypeOf((*MockService)(nil).ExtractServiceIDFromBytes), data)
}

// GetClusterDeployment mocks base method.
func (m *MockService) GetClusterDeployment(clusterID string) (*v10.ClusterDeployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterDeployment", clusterID)
	ret0, _ := ret[0].(*v10.ClusterDeployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterDeployment indicates an expected call of GetClusterDeployment.
func (mr *MockServiceMockRecorder) GetClusterDeployment(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterDeployment", reflect.TypeOf((*MockService)(nil).GetClusterDeployment), clusterID)
}

// GetClusterInfo mocks base method.
func (m *MockService) GetClusterInfo(identifier string) (*v1.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterInfo", identifier)
	ret0, _ := ret[0].(*v1.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterInfo indicates an expected call of GetClusterInfo.
func (mr *MockServiceMockRecorder) GetClusterInfo(identifier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterInfo", reflect.TypeOf((*MockService)(nil).GetClusterInfo), identifier)
}

// GetClusterMachinePools mocks base method.
func (m *MockService) GetClusterMachinePools(clusterID string) ([]*v1.MachinePool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterMachinePools", clusterID)
	ret0, _ := ret[0].([]*v1.MachinePool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterMachinePools indicates an expected call of GetClusterMachinePools.
func (mr *MockServiceMockRecorder) GetClusterMachinePools(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterMachinePools", reflect.TypeOf((*MockService)(nil).GetClusterMachinePools), clusterID)
}

// GetEscalationPolicy mocks base method.
func (m *MockService) GetEscalationPolicy() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEscalationPolicy")
=======
	ret := m.ctrl.Call(m, "GetServiceID")
>>>>>>> c657fa3 (Generate mock)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetServiceID indicates an expected call of GetServiceID.
func (mr *MockServiceMockRecorder) GetServiceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceID", reflect.TypeOf((*MockService)(nil).GetServiceID))
}

// LimitedSupportReasonsExist mocks base method.
func (m *MockService) LimitedSupportReasonsExist(clusterID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LimitedSupportReasonsExist", clusterID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LimitedSupportReasonsExist indicates an expected call of LimitedSupportReasonsExist.
func (mr *MockServiceMockRecorder) LimitedSupportReasonsExist(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LimitedSupportReasonsExist", reflect.TypeOf((*MockService)(nil).LimitedSupportReasonsExist), clusterID)
}

// ListNonRunningInstances mocks base method.
func (m *MockService) ListNonRunningInstances(infraID string) ([]*ec2.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNonRunningInstances", infraID)
	ret0, _ := ret[0].([]*ec2.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNonRunningInstances indicates an expected call of ListNonRunningInstances.
func (mr *MockServiceMockRecorder) ListNonRunningInstances(infraID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNonRunningInstances", reflect.TypeOf((*MockService)(nil).ListNonRunningInstances), infraID)
}

// ListRunningInstances mocks base method.
func (m *MockService) ListRunningInstances(infraID string) ([]*ec2.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRunningInstances", infraID)
	ret0, _ := ret[0].([]*ec2.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRunningInstances indicates an expected call of ListRunningInstances.
func (mr *MockServiceMockRecorder) ListRunningInstances(infraID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRunningInstances", reflect.TypeOf((*MockService)(nil).ListRunningInstances), infraID)
}

// PollInstanceStopEventsFor mocks base method.
func (m *MockService) PollInstanceStopEventsFor(instances []*ec2.Instance, retryTimes int) ([]*cloudtrail.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PollInstanceStopEventsFor", instances, retryTimes)
	ret0, _ := ret[0].([]*cloudtrail.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PollInstanceStopEventsFor indicates an expected call of PollInstanceStopEventsFor.
func (mr *MockServiceMockRecorder) PollInstanceStopEventsFor(instances, retryTimes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollInstanceStopEventsFor", reflect.TypeOf((*MockService)(nil).PollInstanceStopEventsFor), instances, retryTimes)
}

// PostLimitedSupportReason mocks base method.
func (m *MockService) PostLimitedSupportReason(limitedSupportReason ocm.LimitedSupportReason, clusterID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostLimitedSupportReason", limitedSupportReason, clusterID)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostLimitedSupportReason indicates an expected call of PostLimitedSupportReason.
func (mr *MockServiceMockRecorder) PostLimitedSupportReason(limitedSupportReason, clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostLimitedSupportReason", reflect.TypeOf((*MockService)(nil).PostLimitedSupportReason), limitedSupportReason, clusterID)
}

// SilenceAlert mocks base method.
func (m *MockService) SilenceAlert(notes string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SilenceAlert", notes)
	ret0, _ := ret[0].(error)
	return ret0
}

// SilenceAlert indicates an expected call of SilenceAlert.
func (mr *MockServiceMockRecorder) SilenceAlert(notes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SilenceAlert", reflect.TypeOf((*MockService)(nil).SilenceAlert), notes)
}

// UnrelatedLimitedSupportExists mocks base method.
func (m *MockService) UnrelatedLimitedSupportExists(ls ocm.LimitedSupportReason, clusterID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnrelatedLimitedSupportExists", ls, clusterID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnrelatedLimitedSupportExists indicates an expected call of UnrelatedLimitedSupportExists.
func (mr *MockServiceMockRecorder) UnrelatedLimitedSupportExists(ls, clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnrelatedLimitedSupportExists", reflect.TypeOf((*MockService)(nil).UnrelatedLimitedSupportExists), ls, clusterID)
}

// UpdateAndEscalateAlert mocks base method.
func (m *MockService) UpdateAndEscalateAlert(notes string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAndEscalateAlert", notes)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAndEscalateAlert indicates an expected call of UpdateAndEscalateAlert.
func (mr *MockServiceMockRecorder) UpdateAndEscalateAlert(notes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAndEscalateAlert", reflect.TypeOf((*MockService)(nil).UpdateAndEscalateAlert), notes)
}
