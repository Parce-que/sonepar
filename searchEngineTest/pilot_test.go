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
		Name:                 "searchEngine",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	os.Exit(status)
	wd := getWebdriver()
}

func iAmOnANewSearchPage(engine string) error {
	// go to engine
	// Navigate to google <- to change to a variable
	if err := wd.Get("https://www.google.com/"); err != nil {
		panic(err) // TODO
	}

	return nil
}

func iLookForWords(word string) error {
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

	return nil
}
func iClickSearchButton() error {
	// Click the search button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, "#run")
	if err != nil {
		panic(err) // TODO
	}
	if err := btn.Click(); err != nil {
		panic(err) // TODO
	}

	return nil
}

func iShouldSeeWords(word string) error {
	// Wait for the program to finish running and get the output.
	outputDiv, err := wd.FindElement(selenium.ByCSSSelector, "#output")
	if err != nil {
		panic(err) // TODO
	}
	fmt.Println(outputDiv)
	// TODO : Check output and words in there
	return nil
}

// Settings the data first -- not needed here
func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(*godog.Scenario) {
	})

	ctx.Step(`^I am on a new pages of ""$`, iAmOnANewSearchPage) //TODO REGEX
	ctx.Step(`^I look for "words"$`, iLookForWords)
	ctx.Step(`^I I click the search button$`, iClickSearchButton)
	// ctx.Step(`^I should have a "200" http code$`, iShouldHaveHttpCode) <- selenium can't do that it seems
	ctx.Step(`^I should see the "words" in the resulting page$`, iShouldSeeWords)
}
