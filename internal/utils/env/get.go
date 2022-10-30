package env

import (
	"os"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}

	return value
}

func GetDefaultEnvDir(appEnv string) (string, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if appEnv == "production" {
		envDir := rootDir + "/.env.production"
		return envDir, nil
	}

	envDir := rootDir + "/.env"
	return envDir, nil
}
