package discordrus

import (
	"github.com/Sirupsen/logrus"
	"os"
	"testing"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.DebugLevel)

	logrus.AddHook(&Hook{
		// Use environment variable for security reasons
		WebhookURL: os.Getenv("DISCORDRUS_WEBHOOK_URL"),
		// Set minimum level to DebugLevel to receive all log entries
		MinLevel:   logrus.DebugLevel,
		Opts: &Opts{
			Username: "Test Username",
			Author: "", // Setting this to a non-empty string adds the author text to the message header
			DisableTimestamp: false, // Setting this to true will disable timestamps from appearing in the footer
			TimestampFormat: "Jan 2 15:04:05.00000", // The timestamp takes this format; if it is unset, it will take logrus' default format
			Asynchronous: false, // If set to true, the HTTP request will be made in a goroutine
			EnableCustomColors: true, // If set to true, the below CustomLevelColors will apply
			CustomLevelColors: &LevelColors{
				Debug: 10170623,
				Info:  3581519,
				Warn:  14327864,
				Error: 13631488,
				Panic: 13631488,
				Fatal: 13631488,
			},
			DisableInlineFields: false, // If set to true, fields will not appear in columns ("inline")
		},
	})
}

// TestHookIntegration is an integration test to ensure that log entries do send
func TestHookIntegration(t *testing.T) {
	logrus.WithFields(logrus.Fields{"String": "hi", "Integer": 2, "Boolean": false}).Debug("Check this out! Awesome, right?")
	logrus.WithFields(logrus.Fields{"String": "hi", "Integer": 2, "Boolean": false}).Info("Check this out! Awesome, right?")
	logrus.WithFields(logrus.Fields{"String": "hi", "Integer": 2, "Boolean": false}).Warn("Check this out! Awesome, right?")
	logrus.WithFields(logrus.Fields{"String": "hi", "Integer": 2, "Boolean": false}).Error("Check this out! Awesome, right?")
}
