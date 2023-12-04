package attrs

import (
	"log/slog"
	"time"
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
	f := make([]slog.Attr, 0, len(fields)/2)

	for i := 0; i < len(fields); i += 2 {
		key := fields[i]
		value := fields[i+1]

		switch v := value.(type) {
		case string:
			f = append(f, slog.String(key.(string), v))
		case time.Time:
			f = append(f, slog.Time(key.(string), v))
		case int:
			f = append(f, slog.Int(key.(string), v))
		case bool:
			f = append(f, slog.Bool(key.(string), v))
		default:
			f = append(f, slog.Any(key.(string), v))
		}
	}

	return f
}
