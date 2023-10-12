package id

import (
	"time"

	"github.com/disgoorg/snowflake/v2"
)

func Snowflake() string {
	return snowflake.New(time.Now()).String()
}
