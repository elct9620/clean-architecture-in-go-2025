package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/cucumber/godog"
	"github.com/elct9620/clean-architecture-in-go-2025/internal/api/rest"
	jmespath "github.com/jmespath/go-jmespath"
)

type httpCtx struct{}
type resCtx struct{}

func setupHttpServer(ctx context.Context, gc *godog.Scenario) (context.Context, error) {
	server, err := initializeTest()
	if err != nil {
		return ctx, err
	}

	return context.WithValue(ctx, httpCtx{}, server), nil
}

func getHttpServer(ctx context.Context) (*rest.Server, error) {
	server, ok := ctx.Value(httpCtx{}).(*rest.Server)
	if !ok {
		return nil, errors.New("http server not found in context")
	}

	return server, nil
}

func getResponse(ctx context.Context) (*httptest.ResponseRecorder, error) {
	res, ok := ctx.Value(resCtx{}).(*httptest.ResponseRecorder)
	if !ok {
		return nil, errors.New("response not found in context")
	}

	return res, nil
}

func searchJsonPath(ctx context.Context, path string) (any, error) {
	res, err := getResponse(ctx)
	if err != nil {
		return nil, err
	}

	var body any
	if err := json.Unmarshal(res.Body.Bytes(), &body); err != nil {
		return nil, err
	}

	return jmespath.Search(path, body)
}

func makeAGETRequestTo(ctx context.Context, path string) (context.Context, error) {
	server, err := getHttpServer(ctx)
	if err != nil {
		return ctx, err
	}

	req := httptest.NewRequest(http.MethodGet, path, nil)
	res := httptest.NewRecorder()

	server.ServeHTTP(res, req)

	return context.WithValue(ctx, resCtx{}, res), nil
}

func makeAPOSTRequestTo(ctx context.Context, path string, doc *godog.DocString) (context.Context, error) {
	server, err := getHttpServer(ctx)
	if err != nil {
		return ctx, err
	}

	req := httptest.NewRequest(http.MethodPost, path, bytes.NewBufferString(doc.Content))
	req.Header.Add("Content-Type", "application/json")
	res := httptest.NewRecorder()

	server.ServeHTTP(res, req)

	return context.WithValue(ctx, resCtx{}, res), nil
}

func theResponseStatusCodeShouldBe(ctx context.Context, code int) error {
	res, err := getResponse(ctx)
	if err != nil {
		return err
	}

	if res.Code != code {
		return errors.New("response status code not match")
	}

	return nil
}

func theResponseJSONContainsString(ctx context.Context, path string) error {
	raw, err := searchJsonPath(ctx, path)
	if err != nil {
		return fmt.Errorf("failed to search json path: %w", err)
	}

	if _, ok := raw.(string); !ok {
		return fmt.Errorf("expected string, got %T", raw)
	}

	return nil
}

func theResponseJSONContainsWithValue(ctx context.Context, path, expected string) error {
	raw, err := searchJsonPath(ctx, path)
	if err != nil {
		return fmt.Errorf("failed to search json path: %w", err)
	}

	actual, ok := raw.(string)
	if !ok {
		return fmt.Errorf("expected string, got %T", raw)
	}

	if actual != expected {
		return fmt.Errorf("expected %s, got %v", expected, actual)
	}

	return nil
}

func theResponseJSONContainsWithValueNumber(ctx context.Context, path string, expected float64) error {
	raw, err := searchJsonPath(ctx, path)
	if err != nil {
		return fmt.Errorf("failed to search json path: %w", err)
	}

	actual, ok := raw.(float64)
	if !ok {
		return fmt.Errorf("expected int, got %T", raw)
	}

	if actual != expected {
		return fmt.Errorf("expected %f, got %v", expected, actual)
	}

	return nil
}
