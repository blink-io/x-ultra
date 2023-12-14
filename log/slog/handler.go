package slog

import (
	"context"
	"log/slog"

	"github.com/blink-io/x/log"

	slogcommon "github.com/samber/slog-common"
)

var LogLevels = map[slog.Level]log.Level{
	slog.LevelDebug: log.LevelDebug,
	slog.LevelInfo:  log.LevelInfo,
	slog.LevelWarn:  log.LevelWarn,
	slog.LevelError: log.LevelError,
}

type Option struct {
	// log level (default: debug)
	Level slog.Leveler

	// optional: zap logger (default: log.DefaultLogger)
	Logger log.Logger

	// optional: customize json payload builder
	Converter Converter

	// optional: see slog.HandlerOptions
	AddSource bool

	ReplaceAttr func(groups []string, a slog.Attr) slog.Attr
}

func (o Option) NewHandler() slog.Handler {
	if o.Level == nil {
		o.Level = slog.LevelDebug
	}

	if o.Logger == nil {
		o.Logger = log.DefaultLogger
	}

	return &Handler{
		option: o,
		attrs:  []slog.Attr{},
		groups: []string{},
	}
}

var _ slog.Handler = (*Handler)(nil)

type Handler struct {
	option Option
	attrs  []slog.Attr
	groups []string
}

func (h *Handler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.option.Level.Level()
}

func (h *Handler) Handle(ctx context.Context, record slog.Record) error {
	converter := DefaultConverter
	if h.option.Converter != nil {
		converter = h.option.Converter
	}

	level := LogLevels[record.Level]
	fields := converter(h.option.AddSource, h.option.ReplaceAttr, h.attrs, h.groups, &record)

	h.option.Logger.Log(level, record.Message, fields)

	return nil
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{
		option: h.option,
		attrs:  slogcommon.AppendAttrsToGroup(h.groups, h.attrs, attrs...),
		groups: h.groups,
	}
}

func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{
		option: h.option,
		attrs:  h.attrs,
		groups: append(h.groups, name),
	}
}
