// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePrometheusRules implements PrometheusRuleInterface
type FakePrometheusRules struct {
	Fake *FakeMonitoringV1
	ns   string
}

var prometheusrulesResource = schema.GroupVersionResource{Group: "monitoring.coreos.com", Version: "v1", Resource: "prometheusrules"}

var prometheusrulesKind = schema.GroupVersionKind{Group: "monitoring.coreos.com", Version: "v1", Kind: "PrometheusRule"}

// Get takes name of the prometheusRule, and returns the corresponding prometheusRule object, and an error if there is any.
func (c *FakePrometheusRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *monitoringv1.PrometheusRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(prometheusrulesResource, c.ns, name), &monitoringv1.PrometheusRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.PrometheusRule), err
}

// List takes label and field selectors, and returns the list of PrometheusRules that match those selectors.
func (c *FakePrometheusRules) List(ctx context.Context, opts v1.ListOptions) (result *monitoringv1.PrometheusRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(prometheusrulesResource, prometheusrulesKind, c.ns, opts), &monitoringv1.PrometheusRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &monitoringv1.PrometheusRuleList{ListMeta: obj.(*monitoringv1.PrometheusRuleList).ListMeta}
	for _, item := range obj.(*monitoringv1.PrometheusRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested prometheusRules.
func (c *FakePrometheusRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(prometheusrulesResource, c.ns, opts))

}

// Create takes the representation of a prometheusRule and creates it.  Returns the server's representation of the prometheusRule, and an error, if there is any.
func (c *FakePrometheusRules) Create(ctx context.Context, prometheusRule *monitoringv1.PrometheusRule, opts v1.CreateOptions) (result *monitoringv1.PrometheusRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(prometheusrulesResource, c.ns, prometheusRule), &monitoringv1.PrometheusRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.PrometheusRule), err
}

// Update takes the representation of a prometheusRule and updates it. Returns the server's representation of the prometheusRule, and an error, if there is any.
func (c *FakePrometheusRules) Update(ctx context.Context, prometheusRule *monitoringv1.PrometheusRule, opts v1.UpdateOptions) (result *monitoringv1.PrometheusRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(prometheusrulesResource, c.ns, prometheusRule), &monitoringv1.PrometheusRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.PrometheusRule), err
}

// Delete takes name of the prometheusRule and deletes it. Returns an error if one occurs.
func (c *FakePrometheusRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(prometheusrulesResource, c.ns, name), &monitoringv1.PrometheusRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrometheusRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(prometheusrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &monitoringv1.PrometheusRuleList{})
	return err
}

// Patch applies the patch and returns the patched prometheusRule.
func (c *FakePrometheusRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *monitoringv1.PrometheusRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(prometheusrulesResource, c.ns, name, pt, data, subresources...), &monitoringv1.PrometheusRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.PrometheusRule), err
}
