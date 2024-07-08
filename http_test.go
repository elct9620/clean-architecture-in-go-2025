package main

import (
	"context"
	"errors"
	"net/http/httptest"

	"github.com/cucumber/godog"
	"github.com/elct9620/clean-architecture-in-go-2025/internal/api/rest"
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

func makeAGETRequestTo(ctx context.Context, path string) (context.Context, error) {
	server, err := getHttpServer(ctx)
	if err != nil {
		return ctx, err
	}

	req := httptest.NewRequest("GET", path, nil)
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
