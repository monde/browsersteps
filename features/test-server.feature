Feature: Test browsersteps functions

@remote
    Scenario: Open go.mod from the test file server
        Given I open the test server in a browser
        When I wait for 2 seconds
        And I should see the "go.mod" link text
        Then I wait for 2 seconds
