package main

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/spf13/pflag"
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

func TestMain(m *testing.M) {
	pflag.Parse()
	opts.Paths = pflag.Args()

	status := godog.TestSuite{
		Name:                "Clean Architecture in Go 2025",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	// Optional: Run `testing` package's logic besides godog.
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
