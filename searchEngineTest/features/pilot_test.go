package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	flag "github.com/spf13/pflag"
)

var opts = godog.Options{Output: colors.Colored(os.Stdout)}

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opts.Paths = flag.Args()

	status := godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	os.Exit(status)
}

func iAmOnANewSearchPage(engine string) error {
	// go to engine
	return nil
}

func iLookForWords(word string) error {
	if Godogs < num {
		return fmt.Errorf()
	}
	Godogs -= num
	return nil
}

func iShouldHaveHttpCode( ) error {
	if  {
		return fmt.Errorf()
	}
	return nil
}

func iShouldSeeWords(word string) error {
	Godogs = available
	return nil
}


func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() { Godogs = 0 })
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(*godog.Scenario) {
		Godogs = 0 // clean the state before every scenario
	})

	ctx.Step(`^I am on a new pages of ""$`, iAmOnANewSearchPage) //TODO REGEX
	ctx.Step(`^I look for "words"$`, iLookForWords)
	ctx.Step(`^I should have a "200" http code$`, iShouldHaveHttpCode)
	ctx.Step(`^I should see the "words" in the resulting page$`, iShouldSeeWords)
}