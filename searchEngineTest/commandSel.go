package selenium_test

import (
	"fmt"
	"os"

	"github.com/tebeka/selenium"
)

// TODO add Variable depending from the gherkin test

func Example() {
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	const (
		// To change to dockers path
		seleniumPath    = "vendor/selenium-server-standalone-3.4.jar"
		geckoDriverPath = "vendor/geckodriver-v0.18.0-linux64"
		port            = 8080
	)
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // TODO : Change to proper error handling here and for every PANIC
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to google <- to change to a variable
	if err := wd.Get("https://www.google.com/"); err != nil {
		panic(err) // TODO
	}

	// Get a reference to the text box containing code.
	elem, err := wd.FindElement(selenium.ByCSSSelector, "#code")
	if err != nil {
		panic(err) // TODO
	}

	// Enter variable in text box.
	err = elem.SendKeys(`oiseau`)
	if err != nil {
		panic(err) // TODO
	}

	// Click the search button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, "#run")
	if err != nil {
		panic(err) // TODO
	}
	if err := btn.Click(); err != nil {
		panic(err) // TODO
	}

	// Wait for the program to finish running and get the output.
	outputDiv, err := wd.FindElement(selenium.ByCSSSelector, "#output")
	if err != nil {
		panic(err) // TODO
	}
	fmt.Println(outputDiv)
	// TODO : Check output and words in there
}
