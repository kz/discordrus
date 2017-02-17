package discordrus

import (
	"github.com/Sirupsen/logrus"
	"encoding/json"
	"strings"
	"net/http"
	"bytes"
)

// Opts contains the options available for the hook
type Opts struct {
	// Username replaces the default username of the webhook bot for the sent message only if set (default: none)
	Username string
	// Author adds an author field if set (default: none)
	Author string
	// Asynchronous specifies whether the HTTP request should be made in a goroutine without return (default: false)
	Asynchronous bool
	// DisableInlineFields causes fields to be displayed one per line as opposed to being inline (i.e., in columns) (default: false)
	DisableInlineFields bool
	// EnableCustomColors specifies whether CustomLevelColors should be used instead of DefaultLevelColors (default: true)
	EnableCustomColors bool
	// CustomLevelColors is a LevelColors struct which replaces DefaultLevelColors if EnableCustomColors is set to true (default: none)
	CustomLevelColors LevelColors
	// DisableTimestamp specifies whether the timestamp in the footer should be disabled (default: false)
	DisableTimestamp bool
	// TimestampFormat specifies a custom format for the footer
	TimestampFormat string
}

// Hook is a hook to send logs to Discord
type Hook struct {
	// WebhookURL is the full Discord webhook URL
	WebhookURL string
	// MinLevel is the minimum priority level to enable logging for
	MinLevel logrus.Level
	// Opts contains the options available for the hook
	Opts *Opts
}

func (hook *Hook) Fire(entry *logrus.Entry) error {
	// Parse the entry to a Discord webhook object in JSON form
	webhookObject, err := hook.parseToJson(entry)
	if err != nil {
		return err
	}

	// Send the JSON data to the webhook URL, via goroutine if required
	if hook.Opts.Asynchronous {
		go hook.send(webhookObject)
		return nil
	}
	err = hook.send(webhookObject)
	if err != nil {
		return err
	}

	return nil
}

func (hook *Hook) Levels() []logrus.Level {
	return LevelThreshold(hook.MinLevel)
}

func (hook *Hook) parseToJson(entry *logrus.Entry) ([]byte, error) {
	// Create struct mapping to Discord webhook object
	var data = map[string]interface{}{
		"embeds": []map[string]interface{}{},
	}
	var embed = map[string]interface{}{
		"title":       strings.ToUpper(entry.Level.String()),
		"description": entry.Message,
	}
	var fields = []map[string]interface{}{}

	// Add username to data
	if hook.Opts.Username != "" {
		data["username"] = hook.Opts.Username
	}

	// Add color to embed
	if hook.Opts.EnableCustomColors {
		embed["color"] = hook.Opts.CustomLevelColors.LevelColor(entry.Level)
	} else {
		embed["color"] = DefaultLevelColors.LevelColor(entry.Level)
	}

	// Add author to embed
	if hook.Opts.Author != "" {
		embed["author"] = map[string]interface{}{"name": hook.Opts.Author}
	}

	// Add footer to embed
	if !hook.Opts.DisableTimestamp {
		timestamp := ""
		if hook.Opts.TimestampFormat != "" {
			timestamp = entry.Time.Format(hook.Opts.TimestampFormat)
		} else {
			timestamp = entry.Time.String()
		}
		embed["footer"] = map[string]interface{}{
			"text": timestamp,
		}
	}

	// Add fields to embed
	for name, value := range entry.Data {
		var embedField = map[string]interface{}{
			"name":   name,
			"value":  value,
			"inline": !hook.Opts.DisableInlineFields,
		}
		fields = append(fields, embedField)
	}

	// Merge fields and embed into data
	embed["fields"] = fields
	data["embeds"] = []map[string]interface{}{embed}

	return json.Marshal(data)
}

func (hook *Hook) send(webhookObject []byte) error {
	_, err := http.Post(hook.WebhookURL, "application/json; charset=utf-8", bytes.NewBuffer(webhookObject))
	if err != nil {
		return err
	}
	return nil
}
