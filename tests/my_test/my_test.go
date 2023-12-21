package my_test

import (
	"log/slog"
	"testing"
)

func TestSlog_1(t *testing.T) {
	slog.Default().Info("My test for slog info")
}
