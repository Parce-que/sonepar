package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	flag "github.com/spf13/pflag"
	"github.com/tebeka/selenium"
)

var opts = godog.Options{Output: colors.Colored(os.Stdout)}
var wd = getWebdriver()

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opts.Paths = flag.Args()

	status := godog.TestSuite{
		Name:                 "searchEngine",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	os.Exit(status)

}

func iAmOnANewPagesOf(wd selenium.WebDriver, sEngine string) error {
	// go to engine
	// Navigate to google <- to change to a variable
	if err := wd.Get("https://www.google.com/"); err != nil {
		panic(err) // TODO
	}

	return godog.ErrPending
}

func iLookFor(wd selenium.WebDriver, word string) error {
	// Get a reference to the text field containing search.
	elem, err := wd.FindElement(selenium.ByCSSSelector, "#search")
	if err != nil {
		panic(err) // TODO
	}

	// Need to clear ? Not sure
	// Enter variable in text box.
	err = elem.SendKeys(`oiseau`)
	if err != nil {
		panic(err) // TODO
	}
	// Enter variable in text box.
	err = elem.SendKeys(word)
	if err != nil {
		panic(err) // TODO
	}

	return godog.ErrPending
}

func iClickTheSearchButton(wd selenium.WebDriver) error {
	// Click the search button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, "#run")
	if err != nil {
		panic(err) // TODO
	}
	if err := btn.Click(); err != nil {
		panic(err) // TODO
	}

	return godog.ErrPending
}

func iShouldSeeTheInTheResultingPage(wd selenium.WebDriver, word string) error {
	// Wait for the program to finish running and get the output.

	outputDiv, err := wd.FindElement(selenium.ByCSSSelector, "#output")
	if err != nil {
		panic(err) // TODO
	}
	fmt.Println(outputDiv)
	// TODO : Check output and words in there
	return godog.ErrPending
}

// Settings the data first -- not needed here
func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I am on a new pages of "([^"]*)"$`, iAmOnANewPagesOf)
	ctx.Step(`^I click the search button$`, iClickTheSearchButton)
	ctx.Step(`^I look for "([^"]*)"$`, iLookFor)
	ctx.Step(`^I should see the "([^"]*)" in the resulting page$`, iShouldSeeTheInTheResultingPage)
}
