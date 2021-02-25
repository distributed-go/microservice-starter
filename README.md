# recruiter-api

# Go Restful API Boilerplate

[![GoDoc Badge]][GoDoc] [![GoReportCard Badge]][GoReportCard]

Easily extendible RESTful API boilerplate aiming to follow idiomatic go and best practice.

The goal of this boiler is to have a solid and structured foundation to build upon on.

Any feedback and pull requests are welcome and highly appreciated. Feel free to open issues just for comments and discussions.

## Features
The following feature set is a minimal selection of typical Web API requirements:

- Configuration using [viper](https://github.com/spf13/viper)
- CLI features using [cobra](https://github.com/spf13/cobra)
- PostgreSQL support including migrations using [go-pg](https://github.com/go-pg/pg)
- Structured logging with [Logrus](https://github.com/sirupsen/logrus)
- Routing with [chi router](https://github.com/go-chi/chi) and middleware
- JWT Authentication using [jwt-go](https://github.com/dgrijalva/jwt-go) with example passwordless email authentication
- Request data validation using [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
- HTML emails with [gomail](https://github.com/go-gomail/gomail)

## Start Application
- Clone this repository
- Create a postgres database and set environment variables for your database accordingly if not using same as default
- Run the application to see available commands: ```go run main.go```
- First initialize the database running all migrations found in ./database/migrate at once with command *migrate*: ```go run main.go migrate```
- Run the application with command *serve*: ```go run main.go serve```

## API Routes
chechout api folder for swagger API documentation

### Testing
Package auth/pwdless contains example api tests using a mocked database.