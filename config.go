package main

import (
	"fmt"
	"os"
)

type config struct {
	Domain        string
	DNSRecordType string
	APIKey        string
}

func getOptionalEnvVar(key string, def string) string {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}
	return def
}

func getRequiredEnvVar(key string) (string, error) {
	val, exists := os.LookupEnv(key)

	if exists {
		return val, nil
	}
	return "", fmt.Errorf("env var %s not found", key)
}

func loadConfigFromEnv() (config, error) {
	domain, err := getRequiredEnvVar("DOMAIN")
	if err != nil {
		return config{}, err
	}

	apiKey, err := getRequiredEnvVar("API_KEY")
	if err != nil {
		return config{}, err
	}

	return config{
		domain,
		getOptionalEnvVar("DNS_RECORD_TYPE", "A"),
		apiKey,
	}, nil
}
