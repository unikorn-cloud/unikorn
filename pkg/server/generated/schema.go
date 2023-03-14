// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"//94ZGEAjooAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
