// Code generated by mockery v2.2.2. DO NOT EDIT.

package mocks

import (
	config "github.com/aws-controllers-k8s/runtime/pkg/config"
	logr "github.com/go-logr/logr"

	metrics "github.com/aws-controllers-k8s/runtime/pkg/metrics"

	mock "github.com/stretchr/testify/mock"

	session "github.com/aws/aws-sdk-go/aws/session"

	types "github.com/aws-controllers-k8s/runtime/pkg/types"

	v1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
)

// AWSResourceManagerFactory is an autogenerated mock type for the AWSResourceManagerFactory type
type AWSResourceManagerFactory struct {
	mock.Mock
}

// ManagerFor provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5, _a6
func (_m *AWSResourceManagerFactory) ManagerFor(_a0 config.Config, _a1 logr.Logger, _a2 *metrics.Metrics, _a3 types.AWSResourceReconciler, _a4 *session.Session, _a5 v1alpha1.AWSAccountID, _a6 v1alpha1.AWSRegion) (types.AWSResourceManager, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5, _a6)

	var r0 types.AWSResourceManager
	if rf, ok := ret.Get(0).(func(config.Config, logr.Logger, *metrics.Metrics, types.AWSResourceReconciler, *session.Session, v1alpha1.AWSAccountID, v1alpha1.AWSRegion) types.AWSResourceManager); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.AWSResourceManager)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(config.Config, logr.Logger, *metrics.Metrics, types.AWSResourceReconciler, *session.Session, v1alpha1.AWSAccountID, v1alpha1.AWSRegion) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceDescriptor provides a mock function with given fields:
func (_m *AWSResourceManagerFactory) ResourceDescriptor() types.AWSResourceDescriptor {
	ret := _m.Called()

	var r0 types.AWSResourceDescriptor
	if rf, ok := ret.Get(0).(func() types.AWSResourceDescriptor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.AWSResourceDescriptor)
		}
	}

	return r0
}