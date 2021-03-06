// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/openshift/hive/pkg/apis/hive/v1"
	scheme "github.com/openshift/hive/pkg/client/clientset-generated/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SyncIdentityProvidersGetter has a method to return a SyncIdentityProviderInterface.
// A group's client should implement this interface.
type SyncIdentityProvidersGetter interface {
	SyncIdentityProviders(namespace string) SyncIdentityProviderInterface
}

// SyncIdentityProviderInterface has methods to work with SyncIdentityProvider resources.
type SyncIdentityProviderInterface interface {
	Create(*v1.SyncIdentityProvider) (*v1.SyncIdentityProvider, error)
	Update(*v1.SyncIdentityProvider) (*v1.SyncIdentityProvider, error)
	UpdateStatus(*v1.SyncIdentityProvider) (*v1.SyncIdentityProvider, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.SyncIdentityProvider, error)
	List(opts metav1.ListOptions) (*v1.SyncIdentityProviderList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.SyncIdentityProvider, err error)
	SyncIdentityProviderExpansion
}

// syncIdentityProviders implements SyncIdentityProviderInterface
type syncIdentityProviders struct {
	client rest.Interface
	ns     string
}

// newSyncIdentityProviders returns a SyncIdentityProviders
func newSyncIdentityProviders(c *HiveV1Client, namespace string) *syncIdentityProviders {
	return &syncIdentityProviders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the syncIdentityProvider, and returns the corresponding syncIdentityProvider object, and an error if there is any.
func (c *syncIdentityProviders) Get(name string, options metav1.GetOptions) (result *v1.SyncIdentityProvider, err error) {
	result = &v1.SyncIdentityProvider{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("syncidentityproviders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SyncIdentityProviders that match those selectors.
func (c *syncIdentityProviders) List(opts metav1.ListOptions) (result *v1.SyncIdentityProviderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SyncIdentityProviderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("syncidentityproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested syncIdentityProviders.
func (c *syncIdentityProviders) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("syncidentityproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a syncIdentityProvider and creates it.  Returns the server's representation of the syncIdentityProvider, and an error, if there is any.
func (c *syncIdentityProviders) Create(syncIdentityProvider *v1.SyncIdentityProvider) (result *v1.SyncIdentityProvider, err error) {
	result = &v1.SyncIdentityProvider{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("syncidentityproviders").
		Body(syncIdentityProvider).
		Do().
		Into(result)
	return
}

// Update takes the representation of a syncIdentityProvider and updates it. Returns the server's representation of the syncIdentityProvider, and an error, if there is any.
func (c *syncIdentityProviders) Update(syncIdentityProvider *v1.SyncIdentityProvider) (result *v1.SyncIdentityProvider, err error) {
	result = &v1.SyncIdentityProvider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("syncidentityproviders").
		Name(syncIdentityProvider.Name).
		Body(syncIdentityProvider).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *syncIdentityProviders) UpdateStatus(syncIdentityProvider *v1.SyncIdentityProvider) (result *v1.SyncIdentityProvider, err error) {
	result = &v1.SyncIdentityProvider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("syncidentityproviders").
		Name(syncIdentityProvider.Name).
		SubResource("status").
		Body(syncIdentityProvider).
		Do().
		Into(result)
	return
}

// Delete takes name of the syncIdentityProvider and deletes it. Returns an error if one occurs.
func (c *syncIdentityProviders) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("syncidentityproviders").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *syncIdentityProviders) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("syncidentityproviders").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched syncIdentityProvider.
func (c *syncIdentityProviders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.SyncIdentityProvider, err error) {
	result = &v1.SyncIdentityProvider{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("syncidentityproviders").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
