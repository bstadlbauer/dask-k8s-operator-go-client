// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	kubernetesdaskorgv1 "github.com/bstadlbauer/dask-k8s-operator-go-client/pkg/apis/kubernetes.dask.org/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDaskJobs implements DaskJobInterface
type FakeDaskJobs struct {
	Fake *FakeKubernetesV1
	ns   string
}

var daskjobsResource = schema.GroupVersionResource{Group: "kubernetes.dask.org", Version: "v1", Resource: "daskjobs"}

var daskjobsKind = schema.GroupVersionKind{Group: "kubernetes.dask.org", Version: "v1", Kind: "DaskJob"}

// Get takes name of the daskJob, and returns the corresponding daskJob object, and an error if there is any.
func (c *FakeDaskJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *kubernetesdaskorgv1.DaskJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(daskjobsResource, c.ns, name), &kubernetesdaskorgv1.DaskJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubernetesdaskorgv1.DaskJob), err
}

// List takes label and field selectors, and returns the list of DaskJobs that match those selectors.
func (c *FakeDaskJobs) List(ctx context.Context, opts v1.ListOptions) (result *kubernetesdaskorgv1.DaskJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(daskjobsResource, daskjobsKind, c.ns, opts), &kubernetesdaskorgv1.DaskJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubernetesdaskorgv1.DaskJobList{ListMeta: obj.(*kubernetesdaskorgv1.DaskJobList).ListMeta}
	for _, item := range obj.(*kubernetesdaskorgv1.DaskJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested daskJobs.
func (c *FakeDaskJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(daskjobsResource, c.ns, opts))

}

// Create takes the representation of a daskJob and creates it.  Returns the server's representation of the daskJob, and an error, if there is any.
func (c *FakeDaskJobs) Create(ctx context.Context, daskJob *kubernetesdaskorgv1.DaskJob, opts v1.CreateOptions) (result *kubernetesdaskorgv1.DaskJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(daskjobsResource, c.ns, daskJob), &kubernetesdaskorgv1.DaskJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubernetesdaskorgv1.DaskJob), err
}

// Update takes the representation of a daskJob and updates it. Returns the server's representation of the daskJob, and an error, if there is any.
func (c *FakeDaskJobs) Update(ctx context.Context, daskJob *kubernetesdaskorgv1.DaskJob, opts v1.UpdateOptions) (result *kubernetesdaskorgv1.DaskJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(daskjobsResource, c.ns, daskJob), &kubernetesdaskorgv1.DaskJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubernetesdaskorgv1.DaskJob), err
}

// Delete takes name of the daskJob and deletes it. Returns an error if one occurs.
func (c *FakeDaskJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(daskjobsResource, c.ns, name, opts), &kubernetesdaskorgv1.DaskJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDaskJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(daskjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &kubernetesdaskorgv1.DaskJobList{})
	return err
}

// Patch applies the patch and returns the patched daskJob.
func (c *FakeDaskJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kubernetesdaskorgv1.DaskJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(daskjobsResource, c.ns, name, pt, data, subresources...), &kubernetesdaskorgv1.DaskJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubernetesdaskorgv1.DaskJob), err
}
