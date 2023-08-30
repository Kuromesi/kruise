/*
Copyright 2021 The Kruise Authors.

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

package v1beta1

import (
	v1beta1 "github.com/openkruise/kruise/apis/apps/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PodProbeMarkerLister helps list PodProbeMarkers.
// All objects returned here must be treated as read-only.
type PodProbeMarkerLister interface {
	// List lists all PodProbeMarkers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.PodProbeMarker, err error)
	// PodProbeMarkers returns an object that can list and get PodProbeMarkers.
	PodProbeMarkers(namespace string) PodProbeMarkerNamespaceLister
	PodProbeMarkerListerExpansion
}

// podProbeMarkerLister implements the PodProbeMarkerLister interface.
type podProbeMarkerLister struct {
	indexer cache.Indexer
}

// NewPodProbeMarkerLister returns a new PodProbeMarkerLister.
func NewPodProbeMarkerLister(indexer cache.Indexer) PodProbeMarkerLister {
	return &podProbeMarkerLister{indexer: indexer}
}

// List lists all PodProbeMarkers in the indexer.
func (s *podProbeMarkerLister) List(selector labels.Selector) (ret []*v1beta1.PodProbeMarker, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.PodProbeMarker))
	})
	return ret, err
}

// PodProbeMarkers returns an object that can list and get PodProbeMarkers.
func (s *podProbeMarkerLister) PodProbeMarkers(namespace string) PodProbeMarkerNamespaceLister {
	return podProbeMarkerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PodProbeMarkerNamespaceLister helps list and get PodProbeMarkers.
// All objects returned here must be treated as read-only.
type PodProbeMarkerNamespaceLister interface {
	// List lists all PodProbeMarkers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.PodProbeMarker, err error)
	// Get retrieves the PodProbeMarker from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.PodProbeMarker, error)
	PodProbeMarkerNamespaceListerExpansion
}

// podProbeMarkerNamespaceLister implements the PodProbeMarkerNamespaceLister
// interface.
type podProbeMarkerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PodProbeMarkers in the indexer for a given namespace.
func (s podProbeMarkerNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.PodProbeMarker, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.PodProbeMarker))
	})
	return ret, err
}

// Get retrieves the PodProbeMarker from the indexer for a given namespace and name.
func (s podProbeMarkerNamespaceLister) Get(name string) (*v1beta1.PodProbeMarker, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("podprobemarker"), name)
	}
	return obj.(*v1beta1.PodProbeMarker), nil
}
