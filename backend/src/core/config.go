package core

import (
	"os"
	"time"
)

type AppConfig struct {
	AppName               string
	Version               string
	ZincSearchAPIURL      string
	ZincSearchAuthUser    string
	ZincSearchAuthPass    string
	EmailOutputDir        string
	HttpPort              string
	ServerTearDownTimeOut time.Duration
}

func BuildConfigFromEnv() *AppConfig {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "Email Indexer"
	}
	version := os.Getenv("VERSION")
	if version == "" {
		version = "1.0"
	}
	zincSearchAPIURL := os.Getenv("ZINC_SEARCH_API_URL")
	if zincSearchAPIURL == "" {
		zincSearchAPIURL = "https://api.zinc.io/v1/search"
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "3000"
	}
	authUser := os.Getenv("ZINC_SEARCH_AUTH_USER")
	authPass := os.Getenv("ZINC_SEARCH_AUTH_PASS")

	return &AppConfig{
		AppName:               appName,
		Version:               version,
		ZincSearchAPIURL:      zincSearchAPIURL,
		EmailOutputDir:        "./enron_mail_20110402/maildir",
		ZincSearchAuthUser:    authUser,
		ZincSearchAuthPass:    authPass,
		HttpPort:              httpPort,
		ServerTearDownTimeOut: 10 * time.Second,
	}
}
