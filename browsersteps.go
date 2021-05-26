package browsersteps

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github.com/tebeka/selenium"
)

/*BrowserSteps represents a WebDriver context to run the Scenarios*/
type BrowserSteps struct {
	wd             selenium.WebDriver
	Capabilities   selenium.Capabilities
	DefaultURL     string
	URL            *url.URL
	ScreenshotPath string
}

/*SetBaseURL sets the absolute URL used to complete relative URLs*/
func (b *BrowserSteps) SetBaseURL(url *url.URL) error {
	if !url.IsAbs() {
		return errors.New("BaseURL must be absolute")
	}
	b.URL = url
	return nil
}

//BeforeScenario is executed before each scenario
func (b *BrowserSteps) BeforeScenario(sc *messages.Pickle) {
	var err error
	b.wd, err = selenium.NewRemote(b.Capabilities, b.DefaultURL)
	if err != nil {
		log.Panic(err)
	}
}

//AfterScenario is executed after each scenario
func (b *BrowserSteps) AfterScenario(sc *messages.Pickle, err error) {
	if err != nil && b.ScreenshotPath != "" {
		filename := fmt.Sprintf("FAILED STEP - %d.png", time.Now().Unix())

		buff, err := b.GetWebDriver().Screenshot()
		if err != nil {
			fmt.Printf("Error %+v\n", err)
		}

		if _, err := os.Stat(b.ScreenshotPath); os.IsNotExist(err) {
			os.MkdirAll(b.ScreenshotPath, 0755)
		}
		pathname := filepath.Join(b.ScreenshotPath, filename)
		ioutil.WriteFile(pathname, buff, 0644)
	}
	b.GetWebDriver().Quit()
}

//NewBrowserSteps starts a new BrowserSteps instance.
func NewBrowserSteps(ctx *godog.ScenarioContext, cap selenium.Capabilities, defaultURL string) *BrowserSteps {
	bs := &BrowserSteps{Capabilities: cap, DefaultURL: defaultURL, ScreenshotPath: os.Getenv("SCREENSHOT_PATH")}

	ctx.BeforeScenario(bs.BeforeScenario)
	ctx.AfterScenario(bs.AfterScenario)

	return bs
}
