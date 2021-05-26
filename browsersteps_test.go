package browsersteps

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"strconv"
	"testing"
	"time"

	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/cucumber/messages-go/v10"
	"github.com/tebeka/selenium"
)

type TestHarness struct {
	bs *BrowserSteps
}

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "pretty", // "cucumber", "events", "junit", "pretty", "progress"
	Strict: true,
	Paths:  []string{"features"},
	//ShowStepDefinitions: true,
}

func init() {
	godog.BindCommandLineFlags("godog.", &opts) // godog v0.11.0 (latest)
}

func TestMain(m *testing.M) {
	th := &TestHarness{}
	status := godog.TestSuite{
		Name:                 "browsersteps",
		TestSuiteInitializer: th.InitializeTestSuite,
		ScenarioInitializer:  th.InitializeScenario,
		Options:              &opts,
	}.Run()

	os.Exit(status)
}

func (th *TestHarness) iShouldSeeInLinkText(text string) error {
	elem, err := th.bs.wd.FindElement(selenium.ByLinkText, text)
	if err != nil {
		return err
	}
	err = elem.Click()

	return nil
}

func (th *TestHarness) iWaitFor(amount int, unit string) error {
	var u time.Duration
	switch unit {
	case "millisecond":
		u = time.Millisecond
	case "milliseconds":
		u = time.Millisecond
	case "second":
		u = time.Second
	case "seconds":
		u = time.Second
	default:
		u = time.Second
	}
	time.Sleep(u * time.Duration(amount))
	return nil
}

func (th *TestHarness) iOpenTheTestServerInABrowser() error {
	return th.bs.GetWebDriver().Get(th.bs.URL.String())
}

func (th *TestHarness) InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {})
	ctx.AfterSuite(func() {})
}

func (th *TestHarness) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I open the test server in a browser$`, th.iOpenTheTestServerInABrowser)
	ctx.Step(`^I should see the "([^"]*)" link text$`, th.iShouldSeeInLinkText)
	ctx.Step(`^I wait for (\d+) (milliseconds|millisecond|seconds|second)$`, th.iWaitFor)

	debug := os.Getenv("DEBUG")
	if debug != "" {
		val, err := strconv.ParseBool(debug)
		if err == nil {
			selenium.SetDebug(val)
		}
	}

	capabilities := selenium.Capabilities{"browserName": "chrome"}
	capEnv := os.Getenv("SELENIUM_CAPABILITIES")
	if capEnv != "" {
		err := json.Unmarshal([]byte(capEnv), &capabilities)
		if err != nil {
			log.Panic(err)
		}
	}

	bs := NewBrowserSteps(ctx, capabilities, os.Getenv("SELENIUM_URL"))
	if th.bs == nil {
		th.bs = bs
	}

	var server *httptest.Server
	ctx.BeforeScenario(func(sc *messages.Pickle) {
		server = httptest.NewUnstartedServer(http.FileServer(http.Dir(".")))
		var err error
		server.Listener, err = net.Listen("tcp4", "127.0.0.1:")
		if err != nil {
			log.Fatal(err)
		}
		server.Start()
		u, err := url.Parse(server.URL)
		if err != nil {
			log.Panic(err.Error())
		}
		bs.SetBaseURL(u)
	})

	ctx.AfterScenario(func(sc *messages.Pickle, err error) {
		if server != nil {
			server.Close()
			server = nil
		}
	})
}
