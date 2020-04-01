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

// ClusterDeploymentsGetter has a method to return a ClusterDeploymentInterface.
// A group's client should implement this interface.
type ClusterDeploymentsGetter interface {
	ClusterDeployments(namespace string) ClusterDeploymentInterface
}

// ClusterDeploymentInterface has methods to work with ClusterDeployment resources.
type ClusterDeploymentInterface interface {
	Create(*v1.ClusterDeployment) (*v1.ClusterDeployment, error)
	Update(*v1.ClusterDeployment) (*v1.ClusterDeployment, error)
	UpdateStatus(*v1.ClusterDeployment) (*v1.ClusterDeployment, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ClusterDeployment, error)
	List(opts metav1.ListOptions) (*v1.ClusterDeploymentList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterDeployment, err error)
	ClusterDeploymentExpansion
}

// clusterDeployments implements ClusterDeploymentInterface
type clusterDeployments struct {
	client rest.Interface
	ns     string
}

// newClusterDeployments returns a ClusterDeployments
func newClusterDeployments(c *HiveV1Client, namespace string) *clusterDeployments {
	return &clusterDeployments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterDeployment, and returns the corresponding clusterDeployment object, and an error if there is any.
func (c *clusterDeployments) Get(name string, options metav1.GetOptions) (result *v1.ClusterDeployment, err error) {
	result = &v1.ClusterDeployment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterdeployments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterDeployments that match those selectors.
func (c *clusterDeployments) List(opts metav1.ListOptions) (result *v1.ClusterDeploymentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ClusterDeploymentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterdeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterDeployments.
func (c *clusterDeployments) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusterdeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a clusterDeployment and creates it.  Returns the server's representation of the clusterDeployment, and an error, if there is any.
func (c *clusterDeployments) Create(clusterDeployment *v1.ClusterDeployment) (result *v1.ClusterDeployment, err error) {
	result = &v1.ClusterDeployment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusterdeployments").
		Body(clusterDeployment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clusterDeployment and updates it. Returns the server's representation of the clusterDeployment, and an error, if there is any.
func (c *clusterDeployments) Update(clusterDeployment *v1.ClusterDeployment) (result *v1.ClusterDeployment, err error) {
	result = &v1.ClusterDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterdeployments").
		Name(clusterDeployment.Name).
		Body(clusterDeployment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *clusterDeployments) UpdateStatus(clusterDeployment *v1.ClusterDeployment) (result *v1.ClusterDeployment, err error) {
	result = &v1.ClusterDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterdeployments").
		Name(clusterDeployment.Name).
		SubResource("status").
		Body(clusterDeployment).
		Do().
		Into(result)
	return
}

// Delete takes name of the clusterDeployment and deletes it. Returns an error if one occurs.
func (c *clusterDeployments) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterdeployments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterDeployments) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterdeployments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clusterDeployment.
func (c *clusterDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterDeployment, err error) {
	result = &v1.ClusterDeployment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clusterdeployments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
