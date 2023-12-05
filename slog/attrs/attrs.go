package attrs

import (
	"log/slog"
	"time"

	"github.com/blink-io/x/cast"
)

func Options(ops ...Option) []slog.Attr {
	opt := applyOptions(ops...)
	fields := opt.fields
	attrs := makeInternal(opt, fields)
	return attrs
}

func Make(fields ...any) []slog.Attr {
	opt := applyOptions()
	attrs := makeInternal(opt, fields...)
	return attrs
}

func makeInternal(opt *options, fields ...any) []slog.Attr {
	var p = 0
	if len(fields)%2 == 1 {
		p = 1
	}
	f := make([]slog.Attr, 0, len(fields)/2+p)

	for k, v := range opt.extras {
		f = append(f, slog.Any(k, v))
	}

	for i := 0; i < len(fields); i += 2 {
		key := cast.ToString(fields[i])
		value := fields[i+1]

		switch v := value.(type) {
		case string:
			f = append(f, slog.String(key, v))
		case time.Time:
			f = append(f, slog.Time(key, v))
		case time.Duration:
			f = append(f, slog.Duration(key, v))
		case int:
			f = append(f, slog.Int(key, v))
		case bool:
			f = append(f, slog.Bool(key, v))
		default:
			f = append(f, slog.Any(key, v))
		}
	}

	return f
}
