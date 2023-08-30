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
// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/openkruise/kruise/apis/apps/v1alpha1"
	v1beta1 "github.com/openkruise/kruise/apis/apps/v1beta1"
	policyv1alpha1 "github.com/openkruise/kruise/apis/policy/v1alpha1"
	policyv1beta1 "github.com/openkruise/kruise/apis/policy/v1beta1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=apps.kruise.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("advancedcronjobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().AdvancedCronJobs().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("broadcastjobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().BroadcastJobs().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("clonesets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().CloneSets().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("containerrecreaterequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().ContainerRecreateRequests().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("daemonsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().DaemonSets().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("ephemeraljobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().EphemeralJobs().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("imagelistpulljobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().ImageListPullJobs().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("imagepulljobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().ImagePullJobs().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("nodeimages"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().NodeImages().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("nodepodprobes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().NodePodProbes().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("persistentpodstates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().PersistentPodStates().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("podprobemarkers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().PodProbeMarkers().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("resourcedistributions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().ResourceDistributions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("sidecarsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().SidecarSets().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("statefulsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().StatefulSets().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("uniteddeployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().UnitedDeployments().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("workloadspreads"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().WorkloadSpreads().Informer()}, nil

		// Group=apps.kruise.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("advancedcronjobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().AdvancedCronJobs().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("broadcastjobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().BroadcastJobs().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("clonesets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().CloneSets().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("containerrecreaterequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().ContainerRecreateRequests().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("daemonsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().DaemonSets().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("ephemeraljobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().EphemeralJobs().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("imagelistpulljobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().ImageListPullJobs().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("imagepulljobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().ImagePullJobs().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("nodeimages"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().NodeImages().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("nodepodprobes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().NodePodProbes().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("persistentpodstates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().PersistentPodStates().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("podprobemarkers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().PodProbeMarkers().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("resourcedistributions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().ResourceDistributions().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("sidecarsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().SidecarSets().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("statefulsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().StatefulSets().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("uniteddeployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().UnitedDeployments().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("workloadspreads"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1beta1().WorkloadSpreads().Informer()}, nil

		// Group=policy.kruise.io, Version=v1alpha1
	case policyv1alpha1.SchemeGroupVersion.WithResource("podunavailablebudgets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Policy().V1alpha1().PodUnavailableBudgets().Informer()}, nil

		// Group=policy.kruise.io, Version=v1beta1
	case policyv1beta1.SchemeGroupVersion.WithResource("podunavailablebudgets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Policy().V1beta1().PodUnavailableBudgets().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
