package config

import (
	"os"
	"strconv"
	"time"
)

var (
	// Port to be listened by application
	Port string
	// Timeout string
	Timeout string
	// SlackToken string
	SlackToken string
	// SlackSigningSecret string
	SlackSigningSecret string
	// SlackChannel string
	SlackChannel string
	// SlackSlashCommand string
	SlackSlashCommand string
	// CACertificate string
	CACertificate string
	// TelegramToken string
	TelegramToken string
	// TelegramURL string
	TelegramURL string
	// SensuAPIToken string
	SensuAPIToken string
	// SensuAPI string
	SensuAPI string
	// SensuAPIScheme string
	SensuAPIScheme string
	// DebugSensuRequests string
	DebugSensuRequests string
	// BotTimeout time.Duration
	BotTimeout time.Duration
)

// GetEnv func return a default value if missing an Environment variable
func GetEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func init() {
	Port = GetEnv("SENSUBOT_PORT", "9090")
	Timeout = GetEnv("SENSUBOT_TIMEOUT", "15")
	DebugSensuRequests = GetEnv("SENSUBOT_DEBUG_SENSU_REQUESTS", "false")
	SlackSlashCommand = GetEnv("SENSUBOT_SLASH_COMMAND", "/sensubot")
	SlackToken = GetEnv("SENSUBOT_SLACK_TOKEN", "disabled")
	SlackSigningSecret = GetEnv("SENSUBOT_SLACK_SIGNING_SECRET", "disabled")
	SlackChannel = GetEnv("SENSUBOT_SLACK_CHANNEL", "disabled")
	CACertificate = GetEnv("SENSUBOT_CA_CERTIFICATE", "Absent")
	TelegramToken = GetEnv("SENSUBOT_TELEGRAM_TOKEN", "disabled")
	TelegramURL = GetEnv("SENSUBOT_TELEGRAM_URL", "https://api.telegram.org/bot")
	SensuAPIScheme = GetEnv("SENSUBOT_API_SCHEME", "http")
	SensuAPIToken = GetEnv("SENSUBOT_API_TOKEN", "Absent")
	SensuAPI = GetEnv("SENSUBOT_API_URL", "Absent")

	tmpTimeout, err := strconv.Atoi(Timeout)
	if err != nil {
		tmpTimeout = 15
	}
	BotTimeout = time.Duration(tmpTimeout)
}
