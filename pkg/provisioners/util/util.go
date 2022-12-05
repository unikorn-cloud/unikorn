/*
Copyright 2022 EscherCloud.

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

package util

import (
	"context"
	"errors"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	ErrNamespaceLookup = errors.New("unable to lookup namespace")
)

func GetResourceNamespace(ctx context.Context, cli client.Client, label, name string) (*corev1.Namespace, error) {
	labelRequirement, err := labels.NewRequirement(label, selection.Equals, []string{name})
	if err != nil {
		return nil, err
	}

	selector := labels.Everything().Add(*labelRequirement)

	namespaces := &corev1.NamespaceList{}
	if err := cli.List(ctx, namespaces, &client.ListOptions{LabelSelector: selector}); err != nil {
		return nil, err
	}

	if len(namespaces.Items) != 1 {
		return nil, fmt.Errorf("%w: label %s, name %s", ErrNamespaceLookup, label, name)
	}

	return &namespaces.Items[0], nil
}