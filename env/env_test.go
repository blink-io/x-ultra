package env

import (
	"fmt"
	"testing"

	"github.com/caarlos0/env/v10"
	"github.com/stretchr/testify/require"
)

func TestEnv_1(t *testing.T) {
	type config struct {
		Home      string `env:"HOME"`
		Shell     string `env:"SHELL"`
		User      string `env:"USER"`
		LocalTime string `env:"LC_TIME"`
		LocalName string `env:"LC_NAME"`
	}

	cfg := new(config)
	require.NoError(t, env.Parse(cfg))

	fmt.Println("config: ", cfg)
}
