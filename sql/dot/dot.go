package dot

import (
	"embed"

	"github.com/qustavo/dotsql"
)

func LoadFromEmbed(ef embed.FS, name string) (*dotsql.DotSql, error) {
	fs, err := ef.Open(name)
	if err != nil {
		return nil, err
	}
	return dotsql.Load(fs)
}
