// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package generated

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /api/v1/applicationbundles/cluster)
	GetApiV1ApplicationbundlesCluster(w http.ResponseWriter, r *http.Request)

	// (GET /api/v1/applicationbundles/controlPlane)
	GetApiV1ApplicationbundlesControlPlane(w http.ResponseWriter, r *http.Request)

	// (GET /api/v1/applications)
	GetApiV1Applications(w http.ResponseWriter, r *http.Request)

	// (GET /api/v1/controlplanes)
	GetApiV1Controlplanes(w http.ResponseWriter, r *http.Request)

	// (POST /api/v1/controlplanes)
	PostApiV1Controlplanes(w http.ResponseWriter, r *http.Request)

	// (DELETE /api/v1/controlplanes/{controlPlaneName})
	DeleteApiV1ControlplanesControlPlaneName(w http.ResponseWriter, r *http.Request, controlPlaneName ControlPlaneNameParameter)

	// (GET /api/v1/controlplanes/{controlPlaneName})
	GetApiV1ControlplanesControlPlaneName(w http.ResponseWriter, r *http.Request, controlPlaneName ControlPlaneNameParameter)

	// (PUT /api/v1/controlplanes/{controlPlaneName})
	PutApiV1ControlplanesControlPlaneName(w http.ResponseWriter, r *http.Request, controlPlaneName ControlPlaneNameParameter)

	// (GET /api/v1/controlplanes/{controlPlaneName}/clusters)
	GetApiV1ControlplanesControlPlaneNameClusters(w http.ResponseWriter, r *http.Request, controlPlaneName ControlPlaneNameParameter)

	// (POST /api/v1/controlplanes/{controlPlaneName}/clusters)
	PostApiV1ControlplanesControlPlaneNameClusters(w http.ResponseWriter, r *http.Request, controlPlaneName ControlPlaneNameParameter)

	// (DELETE /api/v1/controlplanes/{controlPlaneName}/clusters/{clusterName})
	DeleteApiV1ControlplanesControlPlaneNameClustersClusterName(w http.ResponseWriter, r *http.Request, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter)

	// (GET /api/v1/controlplanes/{controlPlaneName}/clusters/{clusterName})
	GetApiV1ControlplanesControlPlaneNameClustersClusterName(w http.ResponseWriter, r *http.Request, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter)

	// (PUT /api/v1/controlplanes/{controlPlaneName}/clusters/{clusterName})
	PutApiV1ControlplanesControlPlaneNameClustersClusterName(w http.ResponseWriter, r *http.Request, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter)

	// (GET /api/v1/controlplanes/{controlPlaneName}/clusters/{clusterName}/kubeconfig)
	GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfig(w http.ResponseWriter, r *http.Request, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter)

	// (DELETE /api/v1/project)
	DeleteApiV1Project(w http.ResponseWriter, r *http.Request)

	// (POST /api/v1/project)
	PostApiV1Project(w http.ResponseWriter, r *http.Request)

	// (GET /api/v1/regions)
	GetApiV1Regions(w http.ResponseWriter, r *http.Request)

	// (GET /api/v1/regions/{regionName}/availability-zones/block-storage)
	GetApiV1RegionsRegionNameAvailabilityZonesBlockStorage(w http.ResponseWriter, r *http.Request, regionName RegionNameParameter)

	// (GET /api/v1/regions/{regionName}/availability-zones/compute)
	GetApiV1RegionsRegionNameAvailabilityZonesCompute(w http.ResponseWriter, r *http.Request, regionName RegionNameParameter)

	// (GET /api/v1/regions/{regionName}/external-networks)
	GetApiV1RegionsRegionNameExternalNetworks(w http.ResponseWriter, r *http.Request, regionName RegionNameParameter)

	// (GET /api/v1/regions/{regionName}/flavors)
	GetApiV1RegionsRegionNameFlavors(w http.ResponseWriter, r *http.Request, regionName RegionNameParameter)

	// (GET /api/v1/regions/{regionName}/images)
	GetApiV1RegionsRegionNameImages(w http.ResponseWriter, r *http.Request, regionName RegionNameParameter)

	// (GET /api/v1/regions/{regionName}/key-pairs)
	GetApiV1RegionsRegionNameKeyPairs(w http.ResponseWriter, r *http.Request, regionName RegionNameParameter)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetApiV1ApplicationbundlesCluster operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1ApplicationbundlesCluster(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1ApplicationbundlesCluster(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1ApplicationbundlesControlPlane operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1ApplicationbundlesControlPlane(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1ApplicationbundlesControlPlane(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1Applications operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1Applications(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1Applications(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1Controlplanes operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1Controlplanes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1Controlplanes(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostApiV1Controlplanes operation middleware
func (siw *ServerInterfaceWrapper) PostApiV1Controlplanes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostApiV1Controlplanes(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteApiV1ControlplanesControlPlaneName operation middleware
func (siw *ServerInterfaceWrapper) DeleteApiV1ControlplanesControlPlaneName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "controlPlaneName" -------------
	var controlPlaneName ControlPlaneNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, chi.URLParam(r, "controlPlaneName"), &controlPlaneName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "controlPlaneName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteApiV1ControlplanesControlPlaneName(w, r, controlPlaneName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1ControlplanesControlPlaneName operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1ControlplanesControlPlaneName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "controlPlaneName" -------------
	var controlPlaneName ControlPlaneNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, chi.URLParam(r, "controlPlaneName"), &controlPlaneName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "controlPlaneName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1ControlplanesControlPlaneName(w, r, controlPlaneName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutApiV1ControlplanesControlPlaneName operation middleware
func (siw *ServerInterfaceWrapper) PutApiV1ControlplanesControlPlaneName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "controlPlaneName" -------------
	var controlPlaneName ControlPlaneNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, chi.URLParam(r, "controlPlaneName"), &controlPlaneName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "controlPlaneName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutApiV1ControlplanesControlPlaneName(w, r, controlPlaneName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1ControlplanesControlPlaneNameClusters operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1ControlplanesControlPlaneNameClusters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "controlPlaneName" -------------
	var controlPlaneName ControlPlaneNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, chi.URLParam(r, "controlPlaneName"), &controlPlaneName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "controlPlaneName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1ControlplanesControlPlaneNameClusters(w, r, controlPlaneName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostApiV1ControlplanesControlPlaneNameClusters operation middleware
func (siw *ServerInterfaceWrapper) PostApiV1ControlplanesControlPlaneNameClusters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "controlPlaneName" -------------
	var controlPlaneName ControlPlaneNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, chi.URLParam(r, "controlPlaneName"), &controlPlaneName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "controlPlaneName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostApiV1ControlplanesControlPlaneNameClusters(w, r, controlPlaneName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteApiV1ControlplanesControlPlaneNameClustersClusterName operation middleware
func (siw *ServerInterfaceWrapper) DeleteApiV1ControlplanesControlPlaneNameClustersClusterName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "controlPlaneName" -------------
	var controlPlaneName ControlPlaneNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, chi.URLParam(r, "controlPlaneName"), &controlPlaneName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "controlPlaneName", Err: err})
		return
	}

	// ------------- Path parameter "clusterName" -------------
	var clusterName ClusterNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, chi.URLParam(r, "clusterName"), &clusterName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "clusterName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteApiV1ControlplanesControlPlaneNameClustersClusterName(w, r, controlPlaneName, clusterName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1ControlplanesControlPlaneNameClustersClusterName operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1ControlplanesControlPlaneNameClustersClusterName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "controlPlaneName" -------------
	var controlPlaneName ControlPlaneNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, chi.URLParam(r, "controlPlaneName"), &controlPlaneName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "controlPlaneName", Err: err})
		return
	}

	// ------------- Path parameter "clusterName" -------------
	var clusterName ClusterNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, chi.URLParam(r, "clusterName"), &clusterName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "clusterName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1ControlplanesControlPlaneNameClustersClusterName(w, r, controlPlaneName, clusterName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutApiV1ControlplanesControlPlaneNameClustersClusterName operation middleware
func (siw *ServerInterfaceWrapper) PutApiV1ControlplanesControlPlaneNameClustersClusterName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "controlPlaneName" -------------
	var controlPlaneName ControlPlaneNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, chi.URLParam(r, "controlPlaneName"), &controlPlaneName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "controlPlaneName", Err: err})
		return
	}

	// ------------- Path parameter "clusterName" -------------
	var clusterName ClusterNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, chi.URLParam(r, "clusterName"), &clusterName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "clusterName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutApiV1ControlplanesControlPlaneNameClustersClusterName(w, r, controlPlaneName, clusterName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfig operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfig(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "controlPlaneName" -------------
	var controlPlaneName ControlPlaneNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, chi.URLParam(r, "controlPlaneName"), &controlPlaneName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "controlPlaneName", Err: err})
		return
	}

	// ------------- Path parameter "clusterName" -------------
	var clusterName ClusterNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, chi.URLParam(r, "clusterName"), &clusterName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "clusterName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfig(w, r, controlPlaneName, clusterName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteApiV1Project operation middleware
func (siw *ServerInterfaceWrapper) DeleteApiV1Project(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteApiV1Project(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostApiV1Project operation middleware
func (siw *ServerInterfaceWrapper) PostApiV1Project(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostApiV1Project(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1Regions operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1Regions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1Regions(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1RegionsRegionNameAvailabilityZonesBlockStorage operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1RegionsRegionNameAvailabilityZonesBlockStorage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "regionName" -------------
	var regionName RegionNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "regionName", runtime.ParamLocationPath, chi.URLParam(r, "regionName"), &regionName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "regionName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1RegionsRegionNameAvailabilityZonesBlockStorage(w, r, regionName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1RegionsRegionNameAvailabilityZonesCompute operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1RegionsRegionNameAvailabilityZonesCompute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "regionName" -------------
	var regionName RegionNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "regionName", runtime.ParamLocationPath, chi.URLParam(r, "regionName"), &regionName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "regionName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1RegionsRegionNameAvailabilityZonesCompute(w, r, regionName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1RegionsRegionNameExternalNetworks operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1RegionsRegionNameExternalNetworks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "regionName" -------------
	var regionName RegionNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "regionName", runtime.ParamLocationPath, chi.URLParam(r, "regionName"), &regionName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "regionName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1RegionsRegionNameExternalNetworks(w, r, regionName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1RegionsRegionNameFlavors operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1RegionsRegionNameFlavors(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "regionName" -------------
	var regionName RegionNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "regionName", runtime.ParamLocationPath, chi.URLParam(r, "regionName"), &regionName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "regionName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1RegionsRegionNameFlavors(w, r, regionName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1RegionsRegionNameImages operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1RegionsRegionNameImages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "regionName" -------------
	var regionName RegionNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "regionName", runtime.ParamLocationPath, chi.URLParam(r, "regionName"), &regionName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "regionName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1RegionsRegionNameImages(w, r, regionName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiV1RegionsRegionNameKeyPairs operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1RegionsRegionNameKeyPairs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "regionName" -------------
	var regionName RegionNameParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "regionName", runtime.ParamLocationPath, chi.URLParam(r, "regionName"), &regionName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "regionName", Err: err})
		return
	}

	ctx = context.WithValue(ctx, Oauth2AuthenticationScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiV1RegionsRegionNameKeyPairs(w, r, regionName)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/applicationbundles/cluster", wrapper.GetApiV1ApplicationbundlesCluster)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/applicationbundles/controlPlane", wrapper.GetApiV1ApplicationbundlesControlPlane)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/applications", wrapper.GetApiV1Applications)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/controlplanes", wrapper.GetApiV1Controlplanes)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/v1/controlplanes", wrapper.PostApiV1Controlplanes)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/v1/controlplanes/{controlPlaneName}", wrapper.DeleteApiV1ControlplanesControlPlaneName)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/controlplanes/{controlPlaneName}", wrapper.GetApiV1ControlplanesControlPlaneName)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/api/v1/controlplanes/{controlPlaneName}", wrapper.PutApiV1ControlplanesControlPlaneName)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/controlplanes/{controlPlaneName}/clusters", wrapper.GetApiV1ControlplanesControlPlaneNameClusters)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/v1/controlplanes/{controlPlaneName}/clusters", wrapper.PostApiV1ControlplanesControlPlaneNameClusters)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/v1/controlplanes/{controlPlaneName}/clusters/{clusterName}", wrapper.DeleteApiV1ControlplanesControlPlaneNameClustersClusterName)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/controlplanes/{controlPlaneName}/clusters/{clusterName}", wrapper.GetApiV1ControlplanesControlPlaneNameClustersClusterName)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/api/v1/controlplanes/{controlPlaneName}/clusters/{clusterName}", wrapper.PutApiV1ControlplanesControlPlaneNameClustersClusterName)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/controlplanes/{controlPlaneName}/clusters/{clusterName}/kubeconfig", wrapper.GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfig)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/v1/project", wrapper.DeleteApiV1Project)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/v1/project", wrapper.PostApiV1Project)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/regions", wrapper.GetApiV1Regions)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/regions/{regionName}/availability-zones/block-storage", wrapper.GetApiV1RegionsRegionNameAvailabilityZonesBlockStorage)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/regions/{regionName}/availability-zones/compute", wrapper.GetApiV1RegionsRegionNameAvailabilityZonesCompute)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/regions/{regionName}/external-networks", wrapper.GetApiV1RegionsRegionNameExternalNetworks)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/regions/{regionName}/flavors", wrapper.GetApiV1RegionsRegionNameFlavors)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/regions/{regionName}/images", wrapper.GetApiV1RegionsRegionNameImages)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/regions/{regionName}/key-pairs", wrapper.GetApiV1RegionsRegionNameKeyPairs)
	})

	return r
}
