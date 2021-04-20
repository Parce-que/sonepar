package main

import (
	"fmt"
	"os"

	"github.com/tebeka/selenium"
)

// TODO add Variable depending from the gherkin test

func getWebdriver() WebDriver {
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	const (
		// To change to dockers path
		seleniumPath    = "selenium server"
		geckoDriverPath = "geckodriver-v0.18.0-linux64"
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
	return wd
}
