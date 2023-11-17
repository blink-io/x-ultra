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

	EnvSuffix            = ".env"
	EnvSuffixFormat      = ".env.%s"
	EnvLocalSuffix       = ".env.local"
	EnvLocalSuffixFormat = ".env.%s.local"
)

// Flow load multiple files based on the following orders (from lower to higher)
// .env
// .env.production
// .env.test
// .env.development
// .env.local
// .env.production.local
// .env.test.local
// .env.development.local
// supported env: development|test|production
// See https://github.com/bkeepers/dotenv#what-other-env-files-can-i-use
func Flow(paths ...string) error {
	if len(paths) > 0 {
		envs := []string{EnvProduction, EnvTest, EnvDevelopment}
		for _, path := range paths {
			// Load .env first
			_ = godotenv.Overload(filepath.Join(path, EnvSuffix))
			// Load .env.(production|test|development)
			for _, env := range envs {
				_ = godotenv.Overload(filepath.Join(path, fmt.Sprintf(EnvSuffixFormat, env)))
			}
			// Load .env.local
			_ = godotenv.Overload(filepath.Join(path, EnvLocalSuffix))
			// Load .env.local.(production|test|development)
			for _, env := range envs {
				_ = godotenv.Overload(filepath.Join(path, fmt.Sprintf(EnvLocalSuffixFormat, env)))
			}
		}
	}
	return nil
}
