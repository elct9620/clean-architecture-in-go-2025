// Package testability provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package testability

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

// CreateOrderRequest defines model for CreateOrderRequest.
type CreateOrderRequest struct {
	Id    string      `json:"id"`
	Items []OrderItem `json:"items"`
	Name  string      `json:"name"`
}

// CreateTokenRequest defines model for CreateTokenRequest.
type CreateTokenRequest struct {
	Data    string `json:"data"`
	Id      string `json:"id"`
	Version string `json:"version"`
}

// CreatedResponse defines model for CreatedResponse.
type CreatedResponse struct {
	Id string `json:"id"`
}

// OrderItem defines model for OrderItem.
type OrderItem struct {
	Name      string  `json:"name"`
	Quantity  float32 `json:"quantity"`
	UnitPrice float32 `json:"unit_price"`
}

// CreateOrderJSONRequestBody defines body for CreateOrder for application/json ContentType.
type CreateOrderJSONRequestBody = CreateOrderRequest

// CreateTokenJSONRequestBody defines body for CreateToken for application/json ContentType.
type CreateTokenJSONRequestBody = CreateTokenRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create a new order
	// (POST /testability/orders)
	CreateOrder(w http.ResponseWriter, r *http.Request)
	// Create a new token
	// (POST /testability/tokens)
	CreateToken(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Create a new order
// (POST /testability/orders)
func (_ Unimplemented) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create a new token
// (POST /testability/tokens)
func (_ Unimplemented) CreateToken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// CreateOrder operation middleware
func (siw *ServerInterfaceWrapper) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateOrder(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateToken operation middleware
func (siw *ServerInterfaceWrapper) CreateToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateToken(w, r)
	}))

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

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
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
		r.Post(options.BaseURL+"/testability/orders", wrapper.CreateOrder)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/testability/tokens", wrapper.CreateToken)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8yVX2/aPBTGv0rk973YJGgCtCpE6kXHxdSraRXSLiY0OckJnBb/wT5hQyjffbIdyFLS",
	"tRettBtIjJ/j5/z82BxYroRWEiRZlh6YzdcguH+cG+AEX0wB5h62FVhyo9ooDYYQ/Bws3Cf84kJvgKVs",
	"DOX15ayYDbPxFQwvZ0U5nJblZHh1WU6m15OyLLMpGzDaazfbkkG5YvWAIYEIBY8P/xsoWcr+i1t/cWMu",
	"9p7uCISTCpQoKsHS0akuN4bv3W+SC+gavFWkHvHcQj1gBrYVGihY+t311aiP3pYnicoeICdXPhBaqEeQ",
	"zxIqOPGuha/fZkk2Jj1f3dz0sngjpjswFpXsFtuNXtf7UTwIDTzffXEPVitp4b3Cce6uz0wbiTMbPSHQ",
	"7rsH2bbikpD2nemj3ozJSmRgnKiSSD+0wby7yij5q+5JW03WTgY6Zc8bdnKUpXIrEpLvar4BLqNbk6+R",
	"IKfKQIQy+qyiDwuwxDPcIO0//rG3KRtdJBeJ60FpkFwjS9nEDw2Y5rT2+GJq1bFymP2wViHtjjQnVPKu",
	"cB7aO4OFBsHSJ1V4ormSBNKruNYbzL0ufrAhpOF0v3T2e26lug40Qw69u3GSvPGKbc79agXY3KCmwDFM",
	"iTydyFZ5DtaW1cbvsq2E4GbfzuKRhJ9hrssgX1kXgPC+dIoOcXK3y4vE/R30rsQ7t9w/Q9zTeSVxahgd",
	"iYf3pS/ejB2erLJYN5tq3VGiNUQWzM4dyeN/S7Nv9bL+HQAA///nTxyxSgcAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
	res := make(map[string]func() ([]byte, error))
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
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
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
