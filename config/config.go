package config

import (
	"time"

	//"github.com/BurntSushi/toml"
	"github.com/containous/flaeg"
)


const (
	// DefaultGraceTimeout controls how long Scrabbler serves pending requests
	// prior to shutting down.
	DefaultGraceTimeout = 10 * time.Second
)

type ScrabblerConfiguration struct {
	// Main configuration
	Debug			bool	`short:"d" description:"Enable debug mode" export:"true"`
	LogLevel		string	`short:"l" description:"Log level" export:"true"`
	ConfigFileDir	string	`short:"c" description:"Path to configuration directory, load configuration.toml file in a directory"`
	// Scrabbler server configuration
	ServerPort		int		`short:"sp" description:"Scrabbler web server port"`
	ServerHost		string	`short:"sd" description:"Scrabbler web server host"`
	// Web UI configuration
	WebUI			bool 	`short:"w" description:"Run service with web UI"`
	WebUIPort		int		`short:"wp" description:"Web UI port"`
	WebUIHost		string	`short:"wh" description:"Web UI host"`
	// Shutdown configuration
	GraceTimeOut 	flaeg.Duration `short:"g" description:"Duration to give active requests a chance to finish before Scrabbler stops"`
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

}
