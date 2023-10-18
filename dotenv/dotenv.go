package dotenv

import (
	"fmt"
	"path/filepath"

	"github.com/joho/godotenv"
)

const (
	EnvDevelopment = "development"
	EnvTest        = "test"
	EnvProduction  = "production"
)

// Flow load multiple files based on the following order
// .env.development.local
// .env.local
// .env.development
// .env
// supported env: development|test|production
// See https://github.com/bkeepers/dotenv#what-other-env-files-can-i-use
func Flow(env string, paths ...string) error {
	if env == EnvDevelopment ||
		env == EnvTest ||
		env == EnvProduction {
		if paths == nil {
			paths = []string{""}
		}
		for _, path := range paths {
			envFile := filepath.Join(path, ".env")
			customEnvFile := filepath.Join(path, fmt.Sprintf(".env.%s", env))
			envLocalFile := filepath.Join(path, ".env.local")
			customEnvLocalFile := filepath.Join(path, fmt.Sprintf(".env.%s.local", env))
			_ = godotenv.Overload(envFile)
			_ = godotenv.Overload(customEnvFile)
			_ = godotenv.Overload(envLocalFile)
			_ = godotenv.Overload(customEnvLocalFile)
		}
	}
	return nil
}
