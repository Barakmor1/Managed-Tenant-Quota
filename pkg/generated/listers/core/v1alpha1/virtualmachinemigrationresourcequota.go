/*
Copyright 2023 The MTQ Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubevirt.io/managed-tenant-quota/staging/src/kubevirt.io/managed-tenant-quota-api/pkg/apis/core/v1alpha1"
)

// VirtualMachineMigrationResourceQuotaLister helps list VirtualMachineMigrationResourceQuotas.
// All objects returned here must be treated as read-only.
type VirtualMachineMigrationResourceQuotaLister interface {
	// List lists all VirtualMachineMigrationResourceQuotas in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineMigrationResourceQuota, err error)
	// VirtualMachineMigrationResourceQuotas returns an object that can list and get VirtualMachineMigrationResourceQuotas.
	VirtualMachineMigrationResourceQuotas(namespace string) VirtualMachineMigrationResourceQuotaNamespaceLister
	VirtualMachineMigrationResourceQuotaListerExpansion
}

// virtualMachineMigrationResourceQuotaLister implements the VirtualMachineMigrationResourceQuotaLister interface.
type virtualMachineMigrationResourceQuotaLister struct {
	indexer cache.Indexer
}

// NewVirtualMachineMigrationResourceQuotaLister returns a new VirtualMachineMigrationResourceQuotaLister.
func NewVirtualMachineMigrationResourceQuotaLister(indexer cache.Indexer) VirtualMachineMigrationResourceQuotaLister {
	return &virtualMachineMigrationResourceQuotaLister{indexer: indexer}
}

// List lists all VirtualMachineMigrationResourceQuotas in the indexer.
func (s *virtualMachineMigrationResourceQuotaLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineMigrationResourceQuota, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualMachineMigrationResourceQuota))
	})
	return ret, err
}

// VirtualMachineMigrationResourceQuotas returns an object that can list and get VirtualMachineMigrationResourceQuotas.
func (s *virtualMachineMigrationResourceQuotaLister) VirtualMachineMigrationResourceQuotas(namespace string) VirtualMachineMigrationResourceQuotaNamespaceLister {
	return virtualMachineMigrationResourceQuotaNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VirtualMachineMigrationResourceQuotaNamespaceLister helps list and get VirtualMachineMigrationResourceQuotas.
// All objects returned here must be treated as read-only.
type VirtualMachineMigrationResourceQuotaNamespaceLister interface {
	// List lists all VirtualMachineMigrationResourceQuotas in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineMigrationResourceQuota, err error)
	// Get retrieves the VirtualMachineMigrationResourceQuota from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.VirtualMachineMigrationResourceQuota, error)
	VirtualMachineMigrationResourceQuotaNamespaceListerExpansion
}

// virtualMachineMigrationResourceQuotaNamespaceLister implements the VirtualMachineMigrationResourceQuotaNamespaceLister
// interface.
type virtualMachineMigrationResourceQuotaNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VirtualMachineMigrationResourceQuotas in the indexer for a given namespace.
func (s virtualMachineMigrationResourceQuotaNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineMigrationResourceQuota, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualMachineMigrationResourceQuota))
	})
	return ret, err
}

// Get retrieves the VirtualMachineMigrationResourceQuota from the indexer for a given namespace and name.
func (s virtualMachineMigrationResourceQuotaNamespaceLister) Get(name string) (*v1alpha1.VirtualMachineMigrationResourceQuota, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("virtualmachinemigrationresourcequota"), name)
	}
	return obj.(*v1alpha1.VirtualMachineMigrationResourceQuota), nil
}
