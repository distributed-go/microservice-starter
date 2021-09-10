package main

import (
	"github.com/distributed-go/microservice-starter/cmd"
)

// @title Recruiter API Documentation
// @version 2.0
// @description Recruiter API Documentation

// @contact.name API Support
// @contact.url http://xyz.ai
// @contact.email hello@xyz.ai

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8001
// @BasePath /recruiter-api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	cmd.Execute()
}
