package discordrus

import (
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

// TestAllLevels ensures that logrus' AllLevels has not changed
func TestAllLevels(t *testing.T) {
	// AllLevels is a slice of all the supported Logrus levels
	var AllLevels = []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	}

	if !reflect.DeepEqual(AllLevels, logrus.AllLevels) {
		t.Error("discordrus' AllLevels is not the same as logrus' AllLevels")
	}
}

// TestLevelThreshold ensures that the slice returned contains the correct level
func TestLevelThreshold(t *testing.T) {
	var expectedTraceThreshold = []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	}
	var expectedDebugThreshold = []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
	var expectedPanicThreshold = []logrus.Level{
		logrus.PanicLevel,
	}
	var expectedErrorThreshold = []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}

	// Test extreme boundaries
	traceThreshold := LevelThreshold(logrus.TraceLevel)
	if !reflect.DeepEqual(traceThreshold, expectedTraceThreshold) {
		t.Error("Trace threshold does not match expected slice")
	}

	debugThreshold := LevelThreshold(logrus.DebugLevel)
	if !reflect.DeepEqual(debugThreshold, expectedDebugThreshold) {
		t.Error("Debug threshold does not match expected slice")
	}

	panicThreshold := LevelThreshold(logrus.PanicLevel)
	if !reflect.DeepEqual(panicThreshold, expectedPanicThreshold) {
		t.Error("Panic threshold does not match expected slice")
	}

	// Test within boundaries
	errorThreshold := LevelThreshold(logrus.ErrorLevel)
	if !reflect.DeepEqual(errorThreshold, expectedErrorThreshold) {
		t.Error("Error threshold does not match expected slice")
	}
}

// TestLevelColor ensures LevelColor is able to return the respect color for both default and custom LevelColors
func TestLevelColor(t *testing.T) {
	// Set up custom LevelColors
	customLevelColors := LevelColors{
		Trace: 1,
		Debug: 2,
		Info:  3,
		Warn:  4,
		Error: 5,
		Panic: 6,
		Fatal: 7,
	}

	defaultColors := []struct {
		expected int
		input    logrus.Level
	}{
		{
			input:    logrus.TraceLevel,
			expected: DefaultLevelColors.Trace,
		},
		{
			input:    logrus.DebugLevel,
			expected: DefaultLevelColors.Debug,
		},
		{
			input:    logrus.InfoLevel,
			expected: DefaultLevelColors.Info,
		},
		{
			input:    logrus.WarnLevel,
			expected: DefaultLevelColors.Warn,
		},
		{
			input:    logrus.ErrorLevel,
			expected: DefaultLevelColors.Error,
		},
		{
			input:    logrus.PanicLevel,
			expected: DefaultLevelColors.Panic,
		},
		{
			input:    logrus.FatalLevel,
			expected: DefaultLevelColors.Fatal,
		},
	}

	for _, c := range defaultColors {
		if c.expected != DefaultLevelColors.LevelColor(c.input) {
			t.Errorf("Error color for default: %s is not as expected", c.input)
		}
	}

	customColors := []struct {
		expected int
		input    logrus.Level
	}{
		{
			expected: customLevelColors.Trace,
			input:    logrus.TraceLevel,
		},
		{
			expected: customLevelColors.Debug,
			input:    logrus.DebugLevel,
		},
		{
			expected: customLevelColors.Info,
			input:    logrus.InfoLevel,
		},
		{
			expected: customLevelColors.Warn,
			input:    logrus.WarnLevel,
		},
		{
			expected: customLevelColors.Error,
			input:    logrus.ErrorLevel,
		},
		{
			expected: customLevelColors.Panic,
			input:    logrus.PanicLevel,
		},
		{
			expected: customLevelColors.Fatal,
			input:    logrus.FatalLevel,
		},
	}

	for _, c := range customColors {
		if c.expected != customLevelColors.LevelColor(c.input) {
			t.Errorf("Panic color for custom: %s is not as expected", c.input)
		}
	}

}
