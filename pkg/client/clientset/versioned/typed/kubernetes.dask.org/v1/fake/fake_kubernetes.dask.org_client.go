// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/bstadlbauer/dask-k8s-operator-go-client/pkg/client/clientset/versioned/typed/kubernetes.dask.org/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKubernetesV1 struct {
	*testing.Fake
}

func (c *FakeKubernetesV1) DaskClusters(namespace string) v1.DaskClusterInterface {
	return &FakeDaskClusters{c, namespace}
}

func (c *FakeKubernetesV1) DaskJobs(namespace string) v1.DaskJobInterface {
	return &FakeDaskJobs{c, namespace}
}

func (c *FakeKubernetesV1) DaskWorkerGroups(namespace string) v1.DaskWorkerGroupInterface {
	return &FakeDaskWorkerGroups{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKubernetesV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
