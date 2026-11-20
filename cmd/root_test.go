package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/devsy-org/devsy/pkg/log"
	"github.com/spf13/cobra"
)

const (
	testInfoLevel  = "info"
	testErrorLevel = "error"
	testDebugLevel = "debug"
	testTraceLevel = "trace"
	testWarnLevel  = "warn"
)

func TestProviderLogging(t *testing.T) {
	if os.Getenv("PROVIDER_LOG_TEST_CHILD") == "1" {
		serveLoggingProbe()
		return
	}

	for _, tc := range []struct {
		name   string
		level  string
		format string
		want   []string
	}{
		{name: "standalone", want: []string{testInfoLevel, testWarnLevel, testErrorLevel}},
		{name: testErrorLevel, level: testErrorLevel, want: []string{testErrorLevel}},
		{name: testWarnLevel, level: testWarnLevel, want: []string{testWarnLevel, testErrorLevel}},
		{
			name: testInfoLevel, level: testInfoLevel,
			want: []string{testInfoLevel, testWarnLevel, testErrorLevel},
		},
		{
			name: "uppercase", level: "DEBUG", format: "JSON",
			want: []string{testDebugLevel, testInfoLevel, testWarnLevel, testErrorLevel},
		},
		{
			name: testDebugLevel, level: testDebugLevel,
			want: []string{testDebugLevel, testInfoLevel, testWarnLevel, testErrorLevel},
		},
		{
			name: testTraceLevel, level: testTraceLevel, format: "auto",
			want: []string{testTraceLevel, testDebugLevel, testInfoLevel, testWarnLevel, testErrorLevel},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr, err := runLoggingProbe(t, tc.level, tc.format)
			if err != nil {
				t.Fatalf("probe failed: %v; stderr: %s", err, stderr)
			}
			if stdout != "protocol-output" {
				t.Fatalf("protocol output = %q", stdout)
			}
			if strings.Contains(stderr, "provider-secret-canary-846297") {
				t.Fatal("diagnostics exposed secret")
			}
			levels := assertDiagnosticLevels(t, stderr)
			if fmt.Sprint(levels) != fmt.Sprint(tc.want) {
				t.Fatalf("levels = %v, want %v", levels, tc.want)
			}
		})
	}
}

func TestProviderTextDiagnostics(t *testing.T) {
	for _, format := range []string{"text", "plain", "logfmt"} {
		t.Run(format, func(t *testing.T) {
			stdout, stderr, err := runLoggingProbe(t, testWarnLevel, format)
			if err != nil || stdout != "protocol-output" ||
				!strings.Contains(stderr, "warn-canary") ||
				strings.Contains(stderr, "info-canary") ||
				strings.HasPrefix(stderr, "{") {
				t.Fatalf("stdout = %q, stderr = %q, err = %v", stdout, stderr, err)
			}
		})
	}
}

func TestProviderRejectsInvalidLogging(t *testing.T) {
	for _, tc := range []struct{ level, format string }{
		{level: "provider-secret-canary-846297"},
		{format: "provider-secret-canary-846297"},
		{level: " info"},
		{format: " json"},
	} {
		t.Run("invalid-"+tc.level+tc.format, func(t *testing.T) {
			stdout, stderr, err := runLoggingProbe(t, tc.level, tc.format)
			if err == nil || stdout != "" || !strings.Contains(stderr, "invalid DEVSY_LOG_") ||
				strings.Contains(stderr, "provider-secret-canary-846297") {
				t.Fatalf("stdout = %q, stderr = %q, err = %v", stdout, stderr, err)
			}
		})
	}
}

func runLoggingProbe(t *testing.T, level, format string) (string, string, error) {
	t.Helper()
	t.Setenv("PROVIDER_LOG_TEST_CHILD", "1")
	t.Setenv("PROVIDER_TEST_TOKEN", "provider-secret-canary-846297")
	t.Setenv("DEVSY_LOG_LEVEL", level)
	t.Setenv("DEVSY_LOG_FORMAT", format)
	cmd := newLoggingProbeCommand(t)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func assertDiagnosticLevels(t *testing.T, stderr string) []string {
	t.Helper()
	var levels []string
	for line := range strings.SplitSeq(strings.TrimSpace(stderr), "\n") {
		var record map[string]any
		if err := json.Unmarshal([]byte(line), &record); err != nil {
			t.Fatalf("invalid JSON diagnostic %q: %v", line, err)
		}
		level, ok := record["level"].(string)
		msg, hasMessage := record["msg"].(string)
		if !ok || !hasMessage || record["ts"] == nil {
			t.Fatalf("missing diagnostic fields: %v", record)
		}
		if !strings.Contains(msg, level+"-canary") {
			t.Fatalf("unexpected diagnostic: %v", record)
		}
		levels = append(levels, level)
	}
	return levels
}

func newLoggingProbeCommand(t *testing.T) *exec.Cmd {
	t.Helper()
	executable, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	return &exec.Cmd{
		Path: executable,
		Args: []string{executable, "-test.run=^TestProviderLogging$"},
	}
}

func serveLoggingProbe() {
	root := NewRootCmd()
	root.AddCommand(
		&cobra.Command{Use: "probe", RunE: func(_ *cobra.Command, _ []string) error {
			if _, err := fmt.Fprint(os.Stdout, "protocol-output"); err != nil {
				return err
			}
			log.Trace("trace-canary")
			log.Debug("debug-canary")
			log.Info("info-canary " + os.Getenv("PROVIDER_TEST_TOKEN"))
			log.Warn("warn-canary")
			log.Error("error-canary")
			return nil
		}},
	)
	root.SetArgs([]string{"probe"})
	if err := root.Execute(); err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}
	os.Exit(0)
}
