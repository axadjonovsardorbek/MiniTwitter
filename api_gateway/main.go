package main

import (
	"github.com/axadjonovsardorbek/MiniTwitter/api_gateway/config"
	"github.com/axadjonovsardorbek/MiniTwitter/api_gateway/app"

)

func main() {
	// Load configuration and start the API Gateway
	cfg := config.Load()

	// Set up runtime profiling if enabled
	app.Run(&cfg)
}
