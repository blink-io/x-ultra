package zap

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"strings"
	"text/template"
	"time"

	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// QueryHookOptions logging options
type QueryHookOptions struct {
	LogSlow         time.Duration
	Logger          *zap.Logger
	QueryLevel      zapcore.Level
	SlowLevel       zapcore.Level
	ErrorLevel      zapcore.Level
	MessageTemplate string
	ErrorTemplate   string
}

// hook wraps query hook
type hook struct {
	opts            QueryHookOptions
	errorTemplate   *template.Template
	messageTemplate *template.Template
}

// LogEntryVars variables made available to template
type LogEntryVars struct {
	Timestamp time.Time
	Query     string
	Operation string
	Duration  time.Duration
	Error     error
}

// New returns new instance
func New(opts QueryHookOptions) bun.QueryHook {
	h := new(hook)

	if opts.ErrorTemplate == "" {
		opts.ErrorTemplate = "{{.Operation}}[{{.Duration}}]: {{.Query}}: {{.Error}}"
	}
	if opts.MessageTemplate == "" {
		opts.MessageTemplate = "{{.Operation}}[{{.Duration}}]: {{.Query}}"
	}
	h.opts = opts
	errorTemplate, err := template.New("ErrorTemplate").Parse(h.opts.ErrorTemplate)
	if err != nil {
		panic(err)
	}
	messageTemplate, err := template.New("MessageTemplate").Parse(h.opts.MessageTemplate)
	if err != nil {
		panic(err)
	}

	h.errorTemplate = errorTemplate
	h.messageTemplate = messageTemplate
	return h
}

// BeforeQuery does nothing tbh
func (h *hook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

// AfterQuery convert a bun QueryEvent into a logrus message
func (h *hook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	var level zapcore.Level
	var isError bool
	var msg bytes.Buffer

	now := time.Now()
	dur := now.Sub(event.StartTime)

	switch {
	case event.Err == nil, errors.Is(event.Err, sql.ErrNoRows):
		isError = false
		if h.opts.LogSlow > 0 && dur >= h.opts.LogSlow {
			level = h.opts.SlowLevel
		} else {
			level = h.opts.QueryLevel
		}
	default:
		isError = true
		level = h.opts.ErrorLevel
	}
	if level == 0 {
		return
	}

	args := &LogEntryVars{
		Timestamp: now,
		Query:     event.Query,
		Operation: eventOperation(event),
		Duration:  dur,
		Error:     event.Err,
	}

	if isError {
		if err := h.errorTemplate.Execute(&msg, args); err != nil {
			panic(err)
		}
	} else {
		if err := h.messageTemplate.Execute(&msg, args); err != nil {
			panic(err)
		}
	}

	switch level {
	case zapcore.DebugLevel:
		h.opts.Logger.Debug(msg.String())
	case zapcore.InfoLevel:
		h.opts.Logger.Info(msg.String())
	case zapcore.WarnLevel:
		h.opts.Logger.Warn(msg.String())
	case zapcore.ErrorLevel:
		h.opts.Logger.Error(msg.String())
	case zapcore.FatalLevel:
		h.opts.Logger.Fatal(msg.String())
	case zapcore.PanicLevel:
		h.opts.Logger.Panic(msg.String())
	default:
		//panic(fmt.Errorf("unsupported level: %v", level))
		h.opts.Logger.Info(msg.String())
	}
}

// taken from bun
func eventOperation(event *bun.QueryEvent) string {
	switch event.IQuery.(type) {
	case *bun.SelectQuery:
		return "SELECT"
	case *bun.InsertQuery:
		return "INSERT"
	case *bun.UpdateQuery:
		return "UPDATE"
	case *bun.DeleteQuery:
		return "DELETE"
	case *bun.CreateTableQuery:
		return "CREATE TABLE"
	case *bun.DropTableQuery:
		return "DROP TABLE"
	default:
		return queryOperation(event.Query)
	}
}

// taken from bun
func queryOperation(name string) string {
	if idx := strings.Index(name, " "); idx > 0 {
		name = name[:idx]
	}
	if len(name) > 16 {
		name = name[:16]
	}
	return name
}
