/*
Copyright 2022-2023 EscherCloud.

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

package application_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	unikornv1 "github.com/eschercloudai/unikorn/pkg/apis/unikorn/v1alpha1"
	"github.com/eschercloudai/unikorn/pkg/cd"
	"github.com/eschercloudai/unikorn/pkg/cd/mock"
	"github.com/eschercloudai/unikorn/pkg/provisioners"
	"github.com/eschercloudai/unikorn/pkg/provisioners/application"
	mockprovisioners "github.com/eschercloudai/unikorn/pkg/provisioners/mock"
	clientutil "github.com/eschercloudai/unikorn/pkg/util/client"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func toPointer[T any](t T) *T {
	return &t
}

const (
	resourceBundleName = "bundle-x.y.z"
	resourceBundleKind = unikornv1.ApplicationBundleResourceKindControlPlane
)

// mutuallyExclusiveResource defines an abstract type that is able to uniquely
// identify itself with a set of labels.
type mutuallyExclusiveResource labels.Set

func (m mutuallyExclusiveResource) ApplicationBundleKind() unikornv1.ApplicationBundleResourceKind {
	return resourceBundleKind
}

func (m mutuallyExclusiveResource) ApplicationBundleName() string {
	return resourceBundleName
}

func (m mutuallyExclusiveResource) ResourceLabels() (labels.Set, error) {
	return labels.Set(m), nil
}

// testContext provides a common framework for test execution.
type testContext struct {
	client client.Client
}

func mustNewTestContext(t *testing.T, objects ...client.Object) *testContext {
	t.Helper()

	scheme, err := clientutil.NewScheme()
	if err != nil {
		t.Fatal(err)
	}

	tc := &testContext{
		client: fake.NewClientBuilder().WithScheme(scheme).WithStatusSubresource(&unikornv1.Project{}).WithObjects(objects...).Build(),
	}

	return tc
}

const (
	applicationName         = "test"
	overrideApplicationName = "testinate"
	repo                    = "foo"
	chart                   = "bar"
	version                 = "baz"
)

func newBundle(applications ...*unikornv1.HelmApplication) *unikornv1.ApplicationBundle {
	apps := make([]unikornv1.ApplicationBundleApplication, 0, len(applications))

	for _, application := range applications {
		apps = append(apps, unikornv1.ApplicationBundleApplication{
			Name: toPointer(application.Name),
			Reference: &unikornv1.ApplicationReference{
				Kind: toPointer(unikornv1.ApplicationReferenceKindHelm),
				Name: toPointer(application.Name),
			},
		})
	}

	bundle := &unikornv1.ApplicationBundle{
		ObjectMeta: metav1.ObjectMeta{
			Name: resourceBundleName,
		},
		Spec: unikornv1.ApplicationBundleSpec{
			Kind:         toPointer(resourceBundleKind),
			Applications: apps,
		},
	}

	return bundle
}

// TestApplicationCreateHelm tests that given the requested input the provisioner
// creates a CD Application, and the fields are populated as expected.
func TestApplicationCreateHelm(t *testing.T) {
	t.Parallel()

	app := &unikornv1.HelmApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name: applicationName,
		},
		Spec: unikornv1.HelmApplicationSpec{
			Repo:    toPointer(repo),
			Chart:   toPointer(chart),
			Version: toPointer(version),
		},
	}

	tc := mustNewTestContext(t, app, newBundle(app))
	ctx := context.Background()

	c := gomock.NewController(t)
	defer c.Finish()

	driverAppID := &cd.ResourceIdentifier{
		Name: applicationName,
	}

	driverApp := &cd.HelmApplication{
		Repo:    repo,
		Chart:   chart,
		Version: version,
	}

	driver := mock.NewMockDriver(c)
	driver.EXPECT().Client().Return(tc.client)
	driver.EXPECT().CreateOrUpdateHelmApplication(ctx, driverAppID, driverApp).Return(provisioners.ErrYield)

	owner := &mutuallyExclusiveResource{}
	provisioner := application.New(driver, applicationName, owner)

	assert.ErrorIs(t, provisioner.Provision(ctx), provisioners.ErrYield)
}

// TestApplicationCreateHelmExtended tests that given the requested input the provisioner
// creates an ArgoCD Application, and the fields are populated as expected.
func TestApplicationCreateHelmExtended(t *testing.T) {
	t.Parallel()

	release := "epic"
	parameter := "foo"
	value := "bah"
	idLabel1 := "cow"
	idLabel1Value := "moo"
	remoteClusterName := "bar"
	remoteClusterLabel1 := "dog"
	remoteClusterLabel1Value := "woof"
	remoteClusterLabel2 := "cat"
	remoteClusterLabel2value := "meow"

	app := &unikornv1.HelmApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name: applicationName,
		},
		Spec: unikornv1.HelmApplicationSpec{
			Repo:    toPointer(repo),
			Chart:   toPointer(chart),
			Version: toPointer(version),
			Release: toPointer(release),
			Parameters: []unikornv1.HelmApplicationSpecParameter{
				{
					Name:  toPointer(parameter),
					Value: toPointer(value),
				},
			},
			CreateNamespace: toPointer(true),
			ServerSideApply: toPointer(true),
		},
	}

	tc := mustNewTestContext(t, app, newBundle(app))
	ctx := context.Background()

	c := gomock.NewController(t)
	defer c.Finish()

	remoteID := &cd.ResourceIdentifier{
		Name: remoteClusterName,
		Labels: []cd.ResourceIdentifierLabel{
			{
				Name:  remoteClusterLabel1,
				Value: remoteClusterLabel1Value,
			},
			{
				Name:  remoteClusterLabel2,
				Value: remoteClusterLabel2value,
			},
		},
	}

	r := mockprovisioners.NewMockRemoteCluster(c)
	r.EXPECT().ID().Return(remoteID)

	driverAppID := &cd.ResourceIdentifier{
		Name: overrideApplicationName,
		Labels: []cd.ResourceIdentifierLabel{
			{
				Name:  idLabel1,
				Value: idLabel1Value,
			},
		},
	}

	driverApp := &cd.HelmApplication{
		Repo:    repo,
		Chart:   chart,
		Version: version,
		Release: release,
		Parameters: []cd.HelmApplicationParameter{
			{
				Name:  parameter,
				Value: value,
			},
		},
		Cluster:         remoteID,
		CreateNamespace: true,
		ServerSideApply: true,
	}

	driver := mock.NewMockDriver(c)
	driver.EXPECT().Client().Return(tc.client)
	driver.EXPECT().CreateOrUpdateHelmApplication(ctx, driverAppID, driverApp).Return(provisioners.ErrYield)

	owner := &mutuallyExclusiveResource{
		idLabel1: idLabel1Value,
	}
	provisioner := application.New(driver, applicationName, owner).OnRemote(r).WithApplicationName(overrideApplicationName)

	assert.ErrorIs(t, provisioner.Provision(ctx), provisioners.ErrYield)
}

// TestApplicationCreateGit tests that given the requested input the provisioner
// creates an ArgoCD Application, and the fields are populated as expected.
func TestApplicationCreateGit(t *testing.T) {
	t.Parallel()

	path := "bar"

	app := &unikornv1.HelmApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name: applicationName,
		},
		Spec: unikornv1.HelmApplicationSpec{
			Repo:    toPointer(repo),
			Path:    toPointer(path),
			Version: toPointer(version),
		},
	}

	tc := mustNewTestContext(t, app, newBundle(app))
	ctx := context.Background()

	c := gomock.NewController(t)
	defer c.Finish()

	driverAppID := &cd.ResourceIdentifier{
		Name: applicationName,
	}

	driverApp := &cd.HelmApplication{
		Repo:    repo,
		Path:    path,
		Version: version,
	}

	driver := mock.NewMockDriver(c)
	driver.EXPECT().Client().Return(tc.client)
	driver.EXPECT().CreateOrUpdateHelmApplication(ctx, driverAppID, driverApp).Return(provisioners.ErrYield)

	owner := &mutuallyExclusiveResource{}
	provisioner := application.New(driver, applicationName, owner)

	assert.ErrorIs(t, provisioner.Provision(ctx), provisioners.ErrYield)
}

const (
	mutatorRelease                  = "sentinel"
	mutatorParameter                = "foo"
	mutatorValue                    = "bar"
	mutatorIgnoreDifferencesGroup   = "hippes"
	mutatorIgnoreDifferencesKind    = "treeHugger"
	mutatorIgnoreDifferencesPointer = "arrow"
)

//nolint:gochecknoglobals
var mutatorValues = map[string]string{
	mutatorParameter: mutatorValue,
}

// mutator does just that allows modifications of the application.
type mutator struct{}

var _ application.ReleaseNamer = &mutator{}
var _ application.Paramterizer = &mutator{}
var _ application.ValuesGenerator = &mutator{}
var _ application.Customizer = &mutator{}

func (m *mutator) ReleaseName() string {
	return "sentinel"
}

func (m *mutator) Parameters(version *string) (map[string]string, error) {
	p := map[string]string{
		mutatorParameter: mutatorValue,
	}

	return p, nil
}

func (m *mutator) Values(version *string) (interface{}, error) {
	return mutatorValues, nil
}

func (m *mutator) Customize(version *string) ([]cd.HelmApplicationField, error) {
	differences := []cd.HelmApplicationField{
		{
			Group: mutatorIgnoreDifferencesGroup,
			Kind:  mutatorIgnoreDifferencesKind,
			JSONPointers: []string{
				mutatorIgnoreDifferencesPointer,
			},
		},
	}

	return differences, nil
}

// TestApplicationCreateMutate tests that given the requested input the provisioner
// creates an ArgoCD Application, and the fields are populated as expected.
func TestApplicationCreateMutate(t *testing.T) {
	t.Parallel()

	namespace := "gerbils"

	app := &unikornv1.HelmApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name: applicationName,
		},
		Spec: unikornv1.HelmApplicationSpec{
			Repo:    toPointer(repo),
			Chart:   toPointer(chart),
			Version: toPointer(version),
		},
	}

	tc := mustNewTestContext(t, app, newBundle(app))
	ctx := context.Background()

	c := gomock.NewController(t)
	defer c.Finish()

	driverAppID := &cd.ResourceIdentifier{
		Name: applicationName,
	}

	driverApp := &cd.HelmApplication{
		Repo:      repo,
		Chart:     chart,
		Version:   version,
		Release:   mutatorRelease,
		Namespace: namespace,
		Parameters: []cd.HelmApplicationParameter{
			{
				Name:  mutatorParameter,
				Value: mutatorValue,
			},
		},
		Values: mutatorValues,
		IgnoreDifferences: []cd.HelmApplicationField{
			{
				Group: mutatorIgnoreDifferencesGroup,
				Kind:  mutatorIgnoreDifferencesKind,
				JSONPointers: []string{
					mutatorIgnoreDifferencesPointer,
				},
			},
		},
	}

	driver := mock.NewMockDriver(c)
	driver.EXPECT().Client().Return(tc.client)
	driver.EXPECT().CreateOrUpdateHelmApplication(ctx, driverAppID, driverApp).Return(provisioners.ErrYield)

	owner := &mutuallyExclusiveResource{}
	provisioner := application.New(driver, applicationName, owner).WithGenerator(&mutator{}).InNamespace(namespace)

	assert.ErrorIs(t, provisioner.Provision(ctx), provisioners.ErrYield)
}

// TestApplicationDeleteNotFound tests the provisioner returns nil when an application
// doesn't exist.
func TestApplicationDeleteNotFound(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	c := gomock.NewController(t)
	defer c.Finish()

	driverAppID := &cd.ResourceIdentifier{
		Name: applicationName,
	}

	driver := mock.NewMockDriver(c)
	driver.EXPECT().DeleteHelmApplication(ctx, driverAppID, false).Return(provisioners.ErrYield)

	owner := &mutuallyExclusiveResource{}
	provisioner := application.New(driver, applicationName, owner)

	assert.ErrorIs(t, provisioner.Deprovision(ctx), provisioners.ErrYield)
}