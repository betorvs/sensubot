package config

import (
	"time"

	"github.com/spf13/viper"
)

//Values stores the current configuration values
var Values Config

//Config contains the application's configuration values. Add here your own variables and bind it on init() function
type Config struct {
	//Port contains the port in which the application listens
	Port string
	//AppName for displaying in Monitoring
	AppName string
	//LogLevel - DEBUG or INFO or WARNING or ERROR or PANIC or FATAL
	LogLevel string
	//TestRun state if the current execution is a test execution
	TestRun bool
	//UsePrometheus to enable prometheus metrics endpoint
	UsePrometheus bool
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
}

func init() {
	_ = viper.BindEnv("TestRun", "TESTRUN")
	viper.SetDefault("TestRun", false)
	_ = viper.BindEnv("UsePrometheus", "USEPROMETHEUS")
	viper.SetDefault("UsePrometheus", false)
	_ = viper.BindEnv("Port", "PORT")
	viper.SetDefault("Port", "9090")
	_ = viper.BindEnv("AppName", "APP_NAME")
	viper.SetDefault("AppName", "sensubot")
	_ = viper.BindEnv("LogLevel", "LOG_LEVEL")
	viper.SetDefault("LogLevel", "INFO")
	_ = viper.BindEnv("BotTimeout", "SENSUBOT_TIMEOUT")
	viper.SetDefault("BotTimeout", 15)
	_ = viper.BindEnv("SlackSlashCommand", "SENSUBOT_SLASH_COMMAND")
	viper.SetDefault("SlackSlashCommand", "/sensubot")
	_ = viper.BindEnv("SlackToken", "SENSUBOT_SLACK_TOKEN")
	viper.SetDefault("SlackToken", "disabled")
	_ = viper.BindEnv("SlackSigningSecret", "SENSUBOT_SLACK_SIGNING_SECRET")
	viper.SetDefault("SlackSigningSecret", "disabled")
	_ = viper.BindEnv("SlackChannel", "SENSUBOT_SLACK_CHANNEL")
	viper.SetDefault("SlackChannel", "disabled")
	_ = viper.BindEnv("CACertificate", "SENSUBOT_CA_CERTIFICATE")
	viper.SetDefault("CACertificate", "Absent")
	_ = viper.BindEnv("TelegramToken", "SENSUBOT_TELEGRAM_TOKEN")
	viper.SetDefault("TelegramToken", "disabled")
	_ = viper.BindEnv("TelegramURL", "SENSUBOT_TELEGRAM_URL")
	viper.SetDefault("TelegramURL", "https://api.telegram.org/bot")
	_ = viper.BindEnv("SensuAPIScheme", "SENSUBOT_API_SCHEME")
	viper.SetDefault("SensuAPIScheme", "http")
	_ = viper.BindEnv("SensuAPIToken", "SENSUBOT_API_TOKEN")
	viper.SetDefault("SensuAPIToken", "Absent")
	_ = viper.BindEnv("SensuAPI", "SENSUBOT_API_URL")
	viper.SetDefault("SensuAPI", "Absent")

	_ = viper.Unmarshal(&Values)
}
