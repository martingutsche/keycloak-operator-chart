// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	scheme "github.com/keycloak/keycloak-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KeycloaksGetter has a method to return a KeycloakInterface.
// A group's client should implement this interface.
type KeycloaksGetter interface {
	Keycloaks(namespace string) KeycloakInterface
}

// KeycloakInterface has methods to work with Keycloak resources.
type KeycloakInterface interface {
	Create(ctx context.Context, keycloak *v1alpha1.Keycloak, opts v1.CreateOptions) (*v1alpha1.Keycloak, error)
	Update(ctx context.Context, keycloak *v1alpha1.Keycloak, opts v1.UpdateOptions) (*v1alpha1.Keycloak, error)
	UpdateStatus(ctx context.Context, keycloak *v1alpha1.Keycloak, opts v1.UpdateOptions) (*v1alpha1.Keycloak, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Keycloak, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.KeycloakList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Keycloak, err error)
	KeycloakExpansion
}

// keycloaks implements KeycloakInterface
type keycloaks struct {
	client rest.Interface
	ns     string
}

// newKeycloaks returns a Keycloaks
func newKeycloaks(c *KeycloakV1alpha1Client, namespace string) *keycloaks {
	return &keycloaks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the keycloak, and returns the corresponding keycloak object, and an error if there is any.
func (c *keycloaks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Keycloak, err error) {
	result = &v1alpha1.Keycloak{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("keycloaks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Keycloaks that match those selectors.
func (c *keycloaks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KeycloakList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KeycloakList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("keycloaks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested keycloaks.
func (c *keycloaks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("keycloaks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a keycloak and creates it.  Returns the server's representation of the keycloak, and an error, if there is any.
func (c *keycloaks) Create(ctx context.Context, keycloak *v1alpha1.Keycloak, opts v1.CreateOptions) (result *v1alpha1.Keycloak, err error) {
	result = &v1alpha1.Keycloak{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("keycloaks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(keycloak).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a keycloak and updates it. Returns the server's representation of the keycloak, and an error, if there is any.
func (c *keycloaks) Update(ctx context.Context, keycloak *v1alpha1.Keycloak, opts v1.UpdateOptions) (result *v1alpha1.Keycloak, err error) {
	result = &v1alpha1.Keycloak{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("keycloaks").
		Name(keycloak.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(keycloak).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *keycloaks) UpdateStatus(ctx context.Context, keycloak *v1alpha1.Keycloak, opts v1.UpdateOptions) (result *v1alpha1.Keycloak, err error) {
	result = &v1alpha1.Keycloak{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("keycloaks").
		Name(keycloak.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(keycloak).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the keycloak and deletes it. Returns an error if one occurs.
func (c *keycloaks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("keycloaks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *keycloaks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("keycloaks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched keycloak.
func (c *keycloaks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Keycloak, err error) {
	result = &v1alpha1.Keycloak{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("keycloaks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
