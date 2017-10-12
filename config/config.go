package config

import (
	"time"

	"github.com/BurntSushi/toml"
	"github.com/containous/flaeg"
	log "github.com/AKovalevich/scrabbler/log/logrus"
)


const (
	// DefaultGraceTimeout controls how long Scrabbler serves pending requests
	// prior to shutting down.
	DefaultGraceTimeout = 10 * time.Second

	// DefaultConfigFileName path to configuration file
	DefaultConfigPath = "/configuration.default.toml"
)

type ScrabblerConfiguration struct {
	// Main configuration
	Debug			bool	`toml:"debug" short:"d" description:"Enable debug mode" export:"true"`
	LogLevel		string	`toml:"log_level" short:"l" description:"Log level" export:"true"`
	ConfigFilePath	string	`toml:"config_file_path" short:"c" description:"Path to configuration directory, load configuration.toml file in a directory"`
	// Scrabbler server configuration
	ServerPort		int		`toml:"server_port" short:"sp" description:"Scrabbler web server port"`
	ServerHost		string	`toml:"server_host" short:"sd" description:"Scrabbler web server host"`
	// Web UI configuration
	WebUI			bool 	`toml:"web_ui" short:"w" description:"Run service with web UI"`
	WebUIPort		int		`toml:"web_ui_port" short:"wp" description:"Web UI port"`
	WebUIHost		string	`toml:"web_ui_host" short:"wh" description:"Web UI host"`
	// Shutdown configuration
	GraceTimeOut 	flaeg.Duration `grace_time_out:"g" description:"Duration to give active requests a chance to finish before Scrabbler stops"`
}

func NewScrabblerConfiguration() *ScrabblerConfiguration {
	return &ScrabblerConfiguration{}
}

func NewScrabblerDefaultConfiguration() *ScrabblerConfiguration {
	return &ScrabblerConfiguration{
		Debug: true,
		LogLevel: "info",
		ServerPort: 8787,
		ServerHost: "localhost",
		WebUI: true,
		WebUIPort: 8788,
		WebUIHost: "localhost",
		GraceTimeOut: flaeg.Duration(DefaultGraceTimeout),
	}
}

func (config *ScrabblerConfiguration) Reload() {
	var configFilePath string
	var tmpScreabblerConfiguration ScrabblerConfiguration

	if config.ConfigFilePath != "" {
		configFilePath = config.ConfigFilePath
	} else {
		configFilePath = DefaultConfigPath
	}

	if _, err := toml.DecodeFile(configFilePath, &tmpScreabblerConfiguration); err != nil {
		log.Do.Error(err)
	}


}
