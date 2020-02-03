// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	cache "github.com/argoproj/argo-cd/engine/pkg/utils/kube/cache"

	kube "github.com/argoproj/argo-cd/engine/pkg/utils/kube"

	mock "github.com/stretchr/testify/mock"

	schema "k8s.io/apimachinery/pkg/runtime/schema"

	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	v1alpha1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
)

// LiveStateCache is an autogenerated mock type for the LiveStateCache type
type LiveStateCache struct {
	mock.Mock
}

// GetClustersInfo provides a mock function with given fields:
func (_m *LiveStateCache) GetClustersInfo() []cache.ClusterInfo {
	ret := _m.Called()

	var r0 []cache.ClusterInfo
	if rf, ok := ret.Get(0).(func() []cache.ClusterInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cache.ClusterInfo)
		}
	}

	return r0
}

// GetManagedLiveObjs provides a mock function with given fields: a, targetObjs
func (_m *LiveStateCache) GetManagedLiveObjs(a *v1alpha1.Application, targetObjs []*unstructured.Unstructured) (map[kube.ResourceKey]*unstructured.Unstructured, error) {
	ret := _m.Called(a, targetObjs)

	var r0 map[kube.ResourceKey]*unstructured.Unstructured
	if rf, ok := ret.Get(0).(func(*v1alpha1.Application, []*unstructured.Unstructured) map[kube.ResourceKey]*unstructured.Unstructured); ok {
		r0 = rf(a, targetObjs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[kube.ResourceKey]*unstructured.Unstructured)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1alpha1.Application, []*unstructured.Unstructured) error); ok {
		r1 = rf(a, targetObjs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNamespaceTopLevelResources provides a mock function with given fields: server, namespace
func (_m *LiveStateCache) GetNamespaceTopLevelResources(server string, namespace string) (map[kube.ResourceKey]v1alpha1.ResourceNode, error) {
	ret := _m.Called(server, namespace)

	var r0 map[kube.ResourceKey]v1alpha1.ResourceNode
	if rf, ok := ret.Get(0).(func(string, string) map[kube.ResourceKey]v1alpha1.ResourceNode); ok {
		r0 = rf(server, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[kube.ResourceKey]v1alpha1.ResourceNode)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(server, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServerVersion provides a mock function with given fields: serverURL
func (_m *LiveStateCache) GetServerVersion(serverURL string) (string, error) {
	ret := _m.Called(serverURL)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(serverURL)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(serverURL)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsNamespaced provides a mock function with given fields: server, gk
func (_m *LiveStateCache) IsNamespaced(server string, gk schema.GroupKind) (bool, error) {
	ret := _m.Called(server, gk)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, schema.GroupKind) bool); ok {
		r0 = rf(server, gk)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, schema.GroupKind) error); ok {
		r1 = rf(server, gk)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IterateHierarchy provides a mock function with given fields: server, key, action
func (_m *LiveStateCache) IterateHierarchy(server string, key kube.ResourceKey, action func(v1alpha1.ResourceNode, string)) error {
	ret := _m.Called(server, key, action)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, kube.ResourceKey, func(v1alpha1.ResourceNode, string)) error); ok {
		r0 = rf(server, key, action)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Run provides a mock function with given fields: ctx
func (_m *LiveStateCache) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
