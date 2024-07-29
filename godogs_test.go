package main

import (
	"testing"

	"github.com/cucumber/godog"
)

var opts = godog.Options{}

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(setupHttpServer)

	ctx.Step(`^make a GET request to "([^"]*)"$`, makeAGETRequestTo)
	ctx.Step(`^make a POST request to "([^"]*)"$`, makeAPOSTRequestTo)
	ctx.Step(`^the response status code should be (\d+)$`, theResponseStatusCodeShouldBe)
	ctx.Step(`^the response JSON contains "([^"]*)" string$`, theResponseJSONContainsString)
	ctx.Step(`^the response JSON contains "([^"]*)" with value "([^"]*)"$`, theResponseJSONContainsWithValue)
	ctx.Step(`^the response JSON contains "([^"]*)" with value (\d+\.?\d?)$`, theResponseJSONContainsWithValueNumber)
	ctx.Step(`^the response body should be "([^"]*)"$`, theResponseBodyShouldBe)
}

func TestFeatures(t *testing.T) {
	o := opts
	o.TestingT = t

	status := godog.TestSuite{
		Name:                "Clean Architecture in Go 2025",
		ScenarioInitializer: InitializeScenario,
		Options:             &o,
	}.Run()

	if status != 0 {
		t.Errorf("godog failed with status: %d", status)
	}
}
