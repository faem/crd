/*
Copyright The Kubernetes Authors.

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

package fake

import (
	v1alpha1 "crd/pkg/apis/crd.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCustomPods implements CustomPodInterface
type FakeCustomPods struct {
	Fake *FakeCrdV1alpha1
	ns   string
}

var custompodsResource = schema.GroupVersionResource{Group: "crd.com", Version: "v1alpha1", Resource: "custompods"}

var custompodsKind = schema.GroupVersionKind{Group: "crd.com", Version: "v1alpha1", Kind: "CustomPod"}

// Get takes name of the customPod, and returns the corresponding customPod object, and an error if there is any.
func (c *FakeCustomPods) Get(name string, options v1.GetOptions) (result *v1alpha1.CustomPod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(custompodsResource, c.ns, name), &v1alpha1.CustomPod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomPod), err
}

// List takes label and field selectors, and returns the list of CustomPods that match those selectors.
func (c *FakeCustomPods) List(opts v1.ListOptions) (result *v1alpha1.CustomPodList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(custompodsResource, custompodsKind, c.ns, opts), &v1alpha1.CustomPodList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CustomPodList{ListMeta: obj.(*v1alpha1.CustomPodList).ListMeta}
	for _, item := range obj.(*v1alpha1.CustomPodList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested customPods.
func (c *FakeCustomPods) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(custompodsResource, c.ns, opts))

}

// Create takes the representation of a customPod and creates it.  Returns the server's representation of the customPod, and an error, if there is any.
func (c *FakeCustomPods) Create(customPod *v1alpha1.CustomPod) (result *v1alpha1.CustomPod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(custompodsResource, c.ns, customPod), &v1alpha1.CustomPod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomPod), err
}

// Update takes the representation of a customPod and updates it. Returns the server's representation of the customPod, and an error, if there is any.
func (c *FakeCustomPods) Update(customPod *v1alpha1.CustomPod) (result *v1alpha1.CustomPod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(custompodsResource, c.ns, customPod), &v1alpha1.CustomPod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomPod), err
}

// Delete takes name of the customPod and deletes it. Returns an error if one occurs.
func (c *FakeCustomPods) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(custompodsResource, c.ns, name), &v1alpha1.CustomPod{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCustomPods) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(custompodsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CustomPodList{})
	return err
}

// Patch applies the patch and returns the patched customPod.
func (c *FakeCustomPods) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CustomPod, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(custompodsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CustomPod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomPod), err
}
