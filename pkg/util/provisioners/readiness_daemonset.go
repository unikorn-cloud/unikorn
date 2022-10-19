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

package provisioners

import (
	"context"
	"errors"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var (
	ErrDaemonSetUnready = errors.New("daemonset readiness doesn't match desired")
)

type DaemonSetReady struct {
	// client is an intialized Kubernetes client.
	client kubernetes.Interface

	// namespace is the namespace a resource resides in.
	namespace string

	// name is the name of the resource.
	name string
}

// Ensure the ReadinessCheck interface is implemented.
var _ ReadinessCheck = &DaemonSetReady{}

// NewDaemonSetReady creates a new daemonset readiness check.
func NewDaemonSetReady(client kubernetes.Interface, namespace, name string) ReadinessCheck {
	return &DaemonSetReady{
		client:    client,
		namespace: namespace,
		name:      name,
	}
}

// Check implements the ReadinessCheck interface.
func (r *DaemonSetReady) Check() error {
	daemonset, err := r.client.AppsV1().DaemonSets(r.namespace).Get(context.TODO(), r.name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("daemonset get error: %w", err)
	}

	if daemonset.Status.NumberReady != daemonset.Status.DesiredNumberScheduled {
		return fmt.Errorf("%w: status mismatch", ErrDaemonSetUnready)
	}

	return nil

}