package config

import (
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
	// DefaultIntegrationName string
	DefaultIntegrationName string
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
	// TelegramBotName string
	TelegramBotName string
	// SensuAPIToken string
	SensuAPIToken string
	// SensuUser string
	SensuUser string
	// SensuPassword string
	SensuPassword string
	// SensuAPI string
	SensuAPI string
	// SensuAPIScheme string
	SensuAPIScheme string
	// BotTimeout int
	BotTimeout int
	// GoogleChatProjectID string
	GoogleChatProjectID string
	// GoogleChatSAPath string
	GoogleChatSAPath string
	// BlockedVerbs []string splited by comma
	BlockedVerbs []string
	// BlockedResources []string splited by comma
	BlockedResources []string
	// SlackAdminIDList []string splited by comma
	SlackAdminIDList []string
	// TelegramAdminIDList []string splited by comma
	TelegramAdminIDList []string
	// GoogleChatAdminList []string splited by comma
	GoogleChatAdminList []string
	// DeleteAllowedResources []string splited by comma
	DeleteAllowedResources []string
	// AlertManagerEndpoints []string
	AlertManagerEndpoints []string
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
	_ = viper.BindEnv("DefaultIntegrationName", "SENSUBOT_DEFAULT_INTEGRATION_NAME")
	viper.SetDefault("DefaultIntegrationName", "sensu")
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
	_ = viper.BindEnv("TelegramBotName", "SENSUBOT_TELEGRAM_NAME")
	viper.SetDefault("TelegramBotName", "@sensubot")
	_ = viper.BindEnv("SensuAPIScheme", "SENSUBOT_API_SCHEME")
	viper.SetDefault("SensuAPIScheme", "http")
	_ = viper.BindEnv("SensuAPIToken", "SENSUBOT_API_TOKEN")
	viper.SetDefault("SensuAPIToken", "Absent")
	_ = viper.BindEnv("SensuUser", "SENSUBOT_USER")
	viper.SetDefault("SensuUser", "Absent")
	_ = viper.BindEnv("SensuPassword", "SENSUBOT_PASSWORD")
	viper.SetDefault("SensuPassword", "Absent")
	_ = viper.BindEnv("SensuAPI", "SENSUBOT_API_URL")
	viper.SetDefault("SensuAPI", "Absent")
	_ = viper.BindEnv("GoogleChatProjectID", "SENSUBOT_GCHAT_PROJECTID")
	viper.SetDefault("GoogleChatProjectID", "Absent")
	_ = viper.BindEnv("GoogleChatSAPath", "SENSUBOT_GCHAT_SA_PATH")
	viper.SetDefault("GoogleChatSAPath", "Absent")
	_ = viper.BindEnv("BlockedVerbs", "SENSUBOT_BLOCKED_VERBS")
	viper.SetDefault("BlockedVerbs", []string{})
	_ = viper.BindEnv("BlockedResources", "SENSUBOT_BLOCKED_RESOURCES")
	viper.SetDefault("BlockedResources", []string{})
	_ = viper.BindEnv("SlackAdminListID", "SENSUBOT_SLACK_ADMIN_ID_LIST")
	viper.SetDefault("SlackAdminListID", []string{})
	_ = viper.BindEnv("TelegramAdminIDList", "SENSUBOT_TELEGRAM_ADMIN_ID_LIST")
	viper.SetDefault("TelegramAdminIDList", []string{})
	_ = viper.BindEnv("GoogleChatAdminList", "SENSUBOT_GCHAT_ADMIN_LIST")
	viper.SetDefault("GoogleChatAdminList", []string{})
	_ = viper.BindEnv("DeleteAllowedResources", "SENSUBOT_DELETE_ALLOWED_RESOURCES")
	viper.SetDefault("DeleteAllowedResources", []string{"events", "checks"})
	_ = viper.BindEnv("AlertManagerEndpoints", "SENSUBOT_ALERTMANAGER_ENDPOINTS")
	viper.SetDefault("AlertManagerEndpoints", []string{})

	_ = viper.Unmarshal(&Values)
}
