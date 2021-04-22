package main

import (
	"github.com/jobbox-tech/recruiter-api/cmd"
)

// @title Recruiter API Documentation
// @version 2.0
// @description Recruiter API Documentation

// @contact.name API Support
// @contact.url http://jobbox.ai
// @contact.email hello@jobbox.ai

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
