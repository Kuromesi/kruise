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
// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/openkruise/kruise/apis/apps/v1beta1"
	scheme "github.com/openkruise/kruise/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BroadcastJobsGetter has a method to return a BroadcastJobInterface.
// A group's client should implement this interface.
type BroadcastJobsGetter interface {
	BroadcastJobs(namespace string) BroadcastJobInterface
}

// BroadcastJobInterface has methods to work with BroadcastJob resources.
type BroadcastJobInterface interface {
	Create(ctx context.Context, broadcastJob *v1beta1.BroadcastJob, opts v1.CreateOptions) (*v1beta1.BroadcastJob, error)
	Update(ctx context.Context, broadcastJob *v1beta1.BroadcastJob, opts v1.UpdateOptions) (*v1beta1.BroadcastJob, error)
	UpdateStatus(ctx context.Context, broadcastJob *v1beta1.BroadcastJob, opts v1.UpdateOptions) (*v1beta1.BroadcastJob, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.BroadcastJob, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.BroadcastJobList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BroadcastJob, err error)
	BroadcastJobExpansion
}

// broadcastJobs implements BroadcastJobInterface
type broadcastJobs struct {
	client rest.Interface
	ns     string
}

// newBroadcastJobs returns a BroadcastJobs
func newBroadcastJobs(c *AppsV1beta1Client, namespace string) *broadcastJobs {
	return &broadcastJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the broadcastJob, and returns the corresponding broadcastJob object, and an error if there is any.
func (c *broadcastJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BroadcastJob, err error) {
	result = &v1beta1.BroadcastJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("broadcastjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BroadcastJobs that match those selectors.
func (c *broadcastJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BroadcastJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.BroadcastJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("broadcastjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested broadcastJobs.
func (c *broadcastJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("broadcastjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a broadcastJob and creates it.  Returns the server's representation of the broadcastJob, and an error, if there is any.
func (c *broadcastJobs) Create(ctx context.Context, broadcastJob *v1beta1.BroadcastJob, opts v1.CreateOptions) (result *v1beta1.BroadcastJob, err error) {
	result = &v1beta1.BroadcastJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("broadcastjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(broadcastJob).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a broadcastJob and updates it. Returns the server's representation of the broadcastJob, and an error, if there is any.
func (c *broadcastJobs) Update(ctx context.Context, broadcastJob *v1beta1.BroadcastJob, opts v1.UpdateOptions) (result *v1beta1.BroadcastJob, err error) {
	result = &v1beta1.BroadcastJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("broadcastjobs").
		Name(broadcastJob.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(broadcastJob).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *broadcastJobs) UpdateStatus(ctx context.Context, broadcastJob *v1beta1.BroadcastJob, opts v1.UpdateOptions) (result *v1beta1.BroadcastJob, err error) {
	result = &v1beta1.BroadcastJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("broadcastjobs").
		Name(broadcastJob.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(broadcastJob).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the broadcastJob and deletes it. Returns an error if one occurs.
func (c *broadcastJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("broadcastjobs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *broadcastJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("broadcastjobs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched broadcastJob.
func (c *broadcastJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BroadcastJob, err error) {
	result = &v1beta1.BroadcastJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("broadcastjobs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
