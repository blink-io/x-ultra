package slog

import (
	"fmt"
	"log/slog"
	"testing"

	xslog "github.com/blink-io/x/slog/ext"
	"github.com/stretchr/testify/require"
)

func TestLoggerInfoExpected(t *testing.T) {
	checkMessages(t, slog.LevelDebug, nil, slog.LevelInfo, []string{
		"hello",
		"s1s21 2 3s34s56",
		"hello world",
		"",
		"foo",
		"foo bar",
		"s1 s2 1 2 3 s3 4 s5 6",
		"hello",
		"s1s21 2 3s34s56",
		"hello world",
		"",
		"foo",
		"foo bar",
		"s1 s2 1 2 3 s3 4 s5 6",
	}, func(logger *Logger) {
		logger.Info("hello")
		logger.Info("s1", "s2", 1, 2, 3, "s3", 4, "s5", 6)
		logger.Infof("%s world", "hello")
		logger.Infoln()
		logger.Infoln("foo")
		logger.Infoln("foo", "bar")
		logger.Infoln("s1", "s2", 1, 2, 3, "s3", 4, "s5", 6)
		logger.Print("hello")
		logger.Print("s1", "s2", 1, 2, 3, "s3", 4, "s5", 6)
		logger.Printf("%s world", "hello")
		logger.Println()
		logger.Println("foo")
		logger.Println("foo", "bar")
		logger.Println("s1", "s2", 1, 2, 3, "s3", 4, "s5", 6)
	})
}

func TestLoggerDebugExpected(t *testing.T) {
	checkMessages(t, slog.LevelDebug, []Option{WithDebug()}, slog.LevelDebug, []string{
		"hello",
		"s1s21 2 3s34s56",
		"hello world",
		"",
		"foo",
		"foo bar",
		"s1 s2 1 2 3 s3 4 s5 6",
	}, func(logger *Logger) {
		logger.Print("hello")
		logger.Print("s1", "s2", 1, 2, 3, "s3", 4, "s5", 6)
		logger.Printf("%s world", "hello")
		logger.Println()
		logger.Println("foo")
		logger.Println("foo", "bar")
		logger.Println("s1", "s2", 1, 2, 3, "s3", 4, "s5", 6)
	})
}

func TestLoggerDebugSuppressed(t *testing.T) {
	checkMessages(t, slog.LevelInfo, []Option{WithDebug()}, slog.LevelDebug, nil, func(logger *Logger) {
		logger.Print("hello")
		logger.Printf("%s world", "hello")
		logger.Println()
		logger.Println("foo")
		logger.Println("foo", "bar")
	})
}

func TestLoggerWarningExpected(t *testing.T) {
	checkMessages(t, slog.LevelDebug, nil, slog.LevelWarn, []string{
		"hello",
		"s1s21 2 3s34s56",
		"hello world",
		"",
		"foo",
		"foo bar",
		"s1 s2 1 2 3 s3 4 s5 6",
	}, func(logger *Logger) {
		logger.Warning("hello")
		logger.Warning("s1", "s2", 1, 2, 3, "s3", 4, "s5", 6)
		logger.Warningf("%s world", "hello")
		logger.Warningln()
		logger.Warningln("foo")
		logger.Warningln("foo", "bar")
		logger.Warningln("s1", "s2", 1, 2, 3, "s3", 4, "s5", 6)
	})
}

func TestLoggerErrorExpected(t *testing.T) {
	checkMessages(t, slog.LevelDebug, nil, slog.LevelError, []string{
		"hello",
		"s1s21 2 3s34s56",
		"hello world",
		"",
		"foo",
		"foo bar",
		"s1 s2 1 2 3 s3 4 s5 6",
	}, func(logger *Logger) {
		logger.Error("hello")
		logger.Error("s1", "s2", 1, 2, 3, "s3", 4, "s5", 6)
		logger.Errorf("%s world", "hello")
		logger.Errorln()
		logger.Errorln("foo")
		logger.Errorln("foo", "bar")
		logger.Errorln("s1", "s2", 1, 2, 3, "s3", 4, "s5", 6)
	})
}

func TestLoggerFatalExpected(t *testing.T) {
	checkMessages(t, slog.LevelDebug, nil, xslog.LevelFatal, []string{
		"hello",
		"s1s21 2 3s34s56",
		"hello world",
		"",
		"foo",
		"foo bar",
		"s1 s2 1 2 3 s3 4 s5 6",
	}, func(logger *Logger) {
		logger.Fatal("hello")
		logger.Fatal("s1", "s2", 1, 2, 3, "s3", 4, "s5", 6)
		logger.Fatalf("%s world", "hello")
		logger.Fatalln()
		logger.Fatalln("foo")
		logger.Fatalln("foo", "bar")
		logger.Fatalln("s1", "s2", 1, 2, 3, "s3", 4, "s5", 6)
	})
}

func TestLoggerV(t *testing.T) {
	tests := []struct {
		slogLevel    slog.Level
		grpcEnabled  []int
		grpcDisabled []int
	}{
		{
			slogLevel:    slog.LevelDebug,
			grpcEnabled:  []int{grpcLvlInfo, grpcLvlWarn, grpcLvlError, grpcLvlFatal},
			grpcDisabled: []int{}, // everything is enabled, nothing is disabled
		},
		{
			slogLevel:    slog.LevelInfo,
			grpcEnabled:  []int{grpcLvlInfo, grpcLvlWarn, grpcLvlError, grpcLvlFatal},
			grpcDisabled: []int{}, // everything is enabled, nothing is disabled
		},
		{
			slogLevel:    slog.LevelWarn,
			grpcEnabled:  []int{grpcLvlWarn, grpcLvlError, grpcLvlFatal},
			grpcDisabled: []int{grpcLvlInfo},
		},
		{
			slogLevel:    slog.LevelError,
			grpcEnabled:  []int{grpcLvlError, grpcLvlFatal},
			grpcDisabled: []int{grpcLvlInfo, grpcLvlWarn},
		},
		//{
		//	slogLevel:    zapcore.DPanicLevel,
		//	grpcEnabled:  []int{grpcLvlFatal},
		//	grpcDisabled: []int{grpcLvlInfo, grpcLvlWarn, grpcLvlError},
		//},
		//{
		//	slogLevel:    zapcore.PanicLevel,
		//	grpcEnabled:  []int{grpcLvlFatal},
		//	grpcDisabled: []int{grpcLvlInfo, grpcLvlWarn, grpcLvlError},
		//},
		{
			slogLevel:    xslog.LevelFatal,
			grpcEnabled:  []int{grpcLvlFatal},
			grpcDisabled: []int{grpcLvlInfo, grpcLvlWarn, grpcLvlError},
		},
	}
	for _, tst := range tests {
		for _, grpcLvl := range tst.grpcEnabled {
			t.Run(fmt.Sprintf("enabled %s %d", tst.slogLevel, grpcLvl), func(t *testing.T) {
				checkLevel(t, tst.slogLevel, true, func(logger *Logger) bool {
					return logger.V(grpcLvl)
				})
			})
		}
		for _, grpcLvl := range tst.grpcDisabled {
			t.Run(fmt.Sprintf("disabled %s %d", tst.slogLevel, grpcLvl), func(t *testing.T) {
				checkLevel(t, tst.slogLevel, false, func(logger *Logger) bool {
					return logger.V(grpcLvl)
				})
			})
		}
	}
}

func TestGRPCSlog_1(t *testing.T) {
	gl := NewLogger(slog.Default())
	gl.Infof("Hello, %s", "世界")
}

func checkLevel(
	t testing.TB,
	minimalLevel slog.Level,
	expectedBool bool,
	f func(*Logger) bool,
) {
	withLogger(minimalLevel, nil, func(logger *Logger) {
		actualBool := f(logger)
		if expectedBool {
			require.True(t, actualBool)
		} else {
			require.False(t, actualBool)
		}
	})
}

func checkMessages(
	t testing.TB,
	minimalLevel slog.Level,
	opts []Option,
	expectedLevel slog.Level,
	expectedMessages []string,
	f func(*Logger),
) {
	//if expectedLevel == LevelFatal {
	//	expectedLevel = slog.LevelWarn
	//}
	//withLogger(enab, opts, func(logger *Logger) {
	//	f(logger)
	//	logEntries := observedLogs.All()
	//	require.Equal(t, len(expectedMessages), len(logEntries))
	//	for i, logEntry := range logEntries {
	//		require.Equal(t, expectedLevel, logEntry.Level)
	//		require.Equal(t, expectedMessages[i], logEntry.Message)
	//	}
	//})
}

func withLogger(
	minimalLevel slog.Level,
	opts []Option,
	f func(*Logger),
) {
	f(NewLogger(slog.Default(), append(opts, withWarn())...))
}
