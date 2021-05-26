# BrowserSteps

[![Build Status](https://travis-ci.org/monde/browsersteps.svg?branch=master)](https://travis-ci.org/monde/browsersteps)


This package provides Cucumber steps for Browser automation.

## Installation

    go get github.com/monde/browsersteps

## Usage

1. Create cucumber features in your project's `/features` folder.
1. Use this repository `example/example_tests.go` file as `main_test.go` in your project.
1. Execute Selenium Server.
1. Run `godog` or `go test`.

## Testing (this project/module)

Examples assumes local Selenium server is running and chromedriver is installed
your milage may vary.

OSX brew installation.

* `brew install selenium-server-standalone`
* `brew install chromedriver`

Run Selenium in one shell.

```
$ selenium-server -port 4444
```

Run the tests in another shell.

```
$ SELENIUM_URL="http://127.0.0.1:4444/wd/hub" go test -v
```

## Acknowledgements

* [tebeka/selenium](https://github.com/tebeka/selenium) project for Selenium client for Golang.
* [cucumber/godog](http://github.com/cucumber/godog) project for Cucumber implementation for Golang.
* [Browserstack](http://browserstack.com) for providing support and automation to this project.

## History

The `monde/browsersteps` package is cleaned up for go modules. It is a fork of
[`tjipbv/jumba-browsersteps`](https://github.com/llonchj/browsersteps) which is
a fork of [`llonchj/browsersteps`](https://github.com/llonchj/browsersteps).
