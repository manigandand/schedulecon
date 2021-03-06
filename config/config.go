package config

import (
	"fmt"
	"os"
)

const (
	// EnvDev const represents dev environment
	EnvDev = "dev"
	// EnvStaging const represents staging environment
	EnvStaging = "staging"
	// EnvProduction const represents production environment
	EnvProduction = "production"
)

// Env holds the current environment
var (
	Env  string
	Port string
)

// Initialize ...
func Initialize() {
	GetAllEnv()
}

// GetAllEnv should get all the env configs required for the app.
func GetAllEnv() {
	// API Configs
	mustEnv("ENV", &Env, EnvDev)
	mustEnv("PORT", &Port, "8080")
}

// mustEnv get the env variable with the name 'key' and store it in 'value'
func mustEnv(key string, value *string, defaultVal string) {
	if *value = os.Getenv(key); *value == "" {
		*value = defaultVal
		fmt.Printf("%s env variable not set, using default value.\n", key)
	}
}
