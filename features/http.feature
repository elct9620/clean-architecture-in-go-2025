Feature: HTTP Server
  Scenario: Server is available
    When make a GET request to "/"
    Then the response status code should be 404
