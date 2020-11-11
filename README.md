## Requiriment:
Create a web application that takes in a number and returns to the user the highest prime number lower than the input number. For example an input of 55 would return 53. Treat this like a real production application that you would be proud of.

You will be graded on:
-- Performance of algorithm
-- User experience
-- Test coverage
-- Ease of deployment and best practices
-- Git commit history
-- Documentation
-- Code readability and maintainability
-- Anything else you would pay attention to in real production system

Extra credit:
-- Make your service more performant
-- Containerize your application for a one line build experience
-- Use Golang (1.11 or above with module support) and/or typescript
-- Host and run your application on a cloud provider

## Run Project:
    # Requriments: install GoLang version more than 1.11 and nodejs. If you use docker please make sure docker installed
    # how to run:
        - 1> via command lines:
            . cd ./stably_test
            . go run main.go
            . cd ./stably_test/static/stably_dashboard
            . npm run serve
        - 2> via docker:
            . cd ./stably_test
            . make linux
            . make docker
            . cd ./stably_test/static/stably_dashboard
            . make docker
            . back to ./stably_test and docker-compose -u -d
        - 3> run it:
            . goto: http://localhost:8080 for running it