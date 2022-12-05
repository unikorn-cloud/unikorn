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

package readiness

import (
	"context"
	"errors"

	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	ErrResourceExists = errors.New("resource still exists")
)

type ResourceNotExists struct {
	client client.Client
	object client.Object
}

func NewResourceNotExists(client client.Client, object client.Object) *ResourceNotExists {
	return &ResourceNotExists{
		client: client,
		object: object,
	}
}

// Ensure the Check interface is implemented.
var _ Check = &Deployment{}

// Check implements the Check interface.
func (r *ResourceNotExists) Check(ctx context.Context) error {
	objectKey := client.ObjectKeyFromObject(r.object)

	// The Get call needs something to store the result into, it also needs
	// to be able to derive the GVK, so we just convert the object into an
	// unstructured resource.  Immutability is not guaranteed when the source
	// is also unstructured, you have been warned!
	var resource unstructured.Unstructured

	if err := r.client.Scheme().Convert(r.object, &resource, nil); err != nil {
		return err
	}

	if err := r.client.Get(ctx, objectKey, &resource); err != nil {
		if kerrors.IsNotFound(err) {
			return nil
		}

		return err
	}

	return ErrResourceExists
}