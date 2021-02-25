// Code generated by MockGen. DO NOT EDIT.
// Source: inventory_client.go

// Package inventory_client is a generated GoMock package.
package inventory_client

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/openshift/assisted-service/models"
	logrus "github.com/sirupsen/logrus"
)

// MockInventoryClient is a mock of InventoryClient interface
type MockInventoryClient struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryClientMockRecorder
}

// MockInventoryClientMockRecorder is the mock recorder for MockInventoryClient
type MockInventoryClientMockRecorder struct {
	mock *MockInventoryClient
}

// NewMockInventoryClient creates a new mock instance
func NewMockInventoryClient(ctrl *gomock.Controller) *MockInventoryClient {
	mock := &MockInventoryClient{ctrl: ctrl}
	mock.recorder = &MockInventoryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInventoryClient) EXPECT() *MockInventoryClientMockRecorder {
	return m.recorder
}

// DownloadFile mocks base method
func (m *MockInventoryClient) DownloadFile(ctx context.Context, filename, dest string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadFile", ctx, filename, dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// DownloadFile indicates an expected call of DownloadFile
func (mr *MockInventoryClientMockRecorder) DownloadFile(ctx, filename, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadFile", reflect.TypeOf((*MockInventoryClient)(nil).DownloadFile), ctx, filename, dest)
}

// DownloadHostIgnition mocks base method
func (m *MockInventoryClient) DownloadHostIgnition(ctx context.Context, hostID, dest string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadHostIgnition", ctx, hostID, dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// DownloadHostIgnition indicates an expected call of DownloadHostIgnition
func (mr *MockInventoryClientMockRecorder) DownloadHostIgnition(ctx, hostID, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadHostIgnition", reflect.TypeOf((*MockInventoryClient)(nil).DownloadHostIgnition), ctx, hostID, dest)
}

// UpdateHostInstallProgress mocks base method
func (m *MockInventoryClient) UpdateHostInstallProgress(ctx context.Context, hostId string, newStage models.HostStage, info string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHostInstallProgress", ctx, hostId, newStage, info)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHostInstallProgress indicates an expected call of UpdateHostInstallProgress
func (mr *MockInventoryClientMockRecorder) UpdateHostInstallProgress(ctx, hostId, newStage, info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHostInstallProgress", reflect.TypeOf((*MockInventoryClient)(nil).UpdateHostInstallProgress), ctx, hostId, newStage, info)
}

// GetEnabledHostsNamesHosts mocks base method
func (m *MockInventoryClient) GetEnabledHostsNamesHosts(ctx context.Context, log logrus.FieldLogger) (map[string]HostData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnabledHostsNamesHosts", ctx, log)
	ret0, _ := ret[0].(map[string]HostData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnabledHostsNamesHosts indicates an expected call of GetEnabledHostsNamesHosts
func (mr *MockInventoryClientMockRecorder) GetEnabledHostsNamesHosts(ctx, log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnabledHostsNamesHosts", reflect.TypeOf((*MockInventoryClient)(nil).GetEnabledHostsNamesHosts), ctx, log)
}

// UploadIngressCa mocks base method
func (m *MockInventoryClient) UploadIngressCa(ctx context.Context, ingressCA, clusterId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadIngressCa", ctx, ingressCA, clusterId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadIngressCa indicates an expected call of UploadIngressCa
func (mr *MockInventoryClientMockRecorder) UploadIngressCa(ctx, ingressCA, clusterId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadIngressCa", reflect.TypeOf((*MockInventoryClient)(nil).UploadIngressCa), ctx, ingressCA, clusterId)
}

// GetCluster mocks base method
func (m *MockInventoryClient) GetCluster(ctx context.Context) (*models.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCluster", ctx)
	ret0, _ := ret[0].(*models.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCluster indicates an expected call of GetCluster
func (mr *MockInventoryClientMockRecorder) GetCluster(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockInventoryClient)(nil).GetCluster), ctx)
}

// CompleteInstallation mocks base method
func (m *MockInventoryClient) CompleteInstallation(ctx context.Context, clusterId string, isSuccess bool, errorInfo string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteInstallation", ctx, clusterId, isSuccess, errorInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteInstallation indicates an expected call of CompleteInstallation
func (mr *MockInventoryClientMockRecorder) CompleteInstallation(ctx, clusterId, isSuccess, errorInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteInstallation", reflect.TypeOf((*MockInventoryClient)(nil).CompleteInstallation), ctx, clusterId, isSuccess, errorInfo)
}

// GetHosts mocks base method
func (m *MockInventoryClient) GetHosts(ctx context.Context, log logrus.FieldLogger, skippedStatuses []string) (map[string]HostData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHosts", ctx, log, skippedStatuses)
	ret0, _ := ret[0].(map[string]HostData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHosts indicates an expected call of GetHosts
func (mr *MockInventoryClientMockRecorder) GetHosts(ctx, log, skippedStatuses interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHosts", reflect.TypeOf((*MockInventoryClient)(nil).GetHosts), ctx, log, skippedStatuses)
}

// UploadLogs mocks base method
func (m *MockInventoryClient) UploadLogs(ctx context.Context, clusterId string, logsType models.LogsType, upfile io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadLogs", ctx, clusterId, logsType, upfile)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadLogs indicates an expected call of UploadLogs
func (mr *MockInventoryClientMockRecorder) UploadLogs(ctx, clusterId, logsType, upfile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadLogs", reflect.TypeOf((*MockInventoryClient)(nil).UploadLogs), ctx, clusterId, logsType, upfile)
}

// ClusterLogProgressReport mocks base method
func (m *MockInventoryClient) ClusterLogProgressReport(ctx context.Context, clusterId string, progress models.LogsState) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClusterLogProgressReport", ctx, clusterId, progress)
}

// ClusterLogProgressReport indicates an expected call of ClusterLogProgressReport
func (mr *MockInventoryClientMockRecorder) ClusterLogProgressReport(ctx, clusterId, progress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterLogProgressReport", reflect.TypeOf((*MockInventoryClient)(nil).ClusterLogProgressReport), ctx, clusterId, progress)
}

// HostLogProgressReport mocks base method
func (m *MockInventoryClient) HostLogProgressReport(ctx context.Context, clusterId, hostId string, progress models.LogsState) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HostLogProgressReport", ctx, clusterId, hostId, progress)
}

// HostLogProgressReport indicates an expected call of HostLogProgressReport
func (mr *MockInventoryClientMockRecorder) HostLogProgressReport(ctx, clusterId, hostId, progress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostLogProgressReport", reflect.TypeOf((*MockInventoryClient)(nil).HostLogProgressReport), ctx, clusterId, hostId, progress)
}

// UpdateClusterInstallProgress mocks base method
func (m *MockInventoryClient) UpdateClusterInstallProgress(ctx context.Context, clusterId, progress string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClusterInstallProgress", ctx, clusterId, progress)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClusterInstallProgress indicates an expected call of UpdateClusterInstallProgress
func (mr *MockInventoryClientMockRecorder) UpdateClusterInstallProgress(ctx, clusterId, progress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClusterInstallProgress", reflect.TypeOf((*MockInventoryClient)(nil).UpdateClusterInstallProgress), ctx, clusterId, progress)
}
