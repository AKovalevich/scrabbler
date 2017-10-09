package config

type ScrabblerConfiguration struct {
	// Main configuration
	Debug		bool	`short:"d" description:"Enable debug mode" export:"true"`
	LogLevel	string	`short:"l" description:"Log level" export:"true"`
	// Scrabbler server configuration
	ServerPort	int		`short:"sp" description:"Scrabbler web server port"`
	ServerHost	string	`short:"sd" description:"Scrabbler web server host"`
	// Web UI configuration
	WebUI		bool 	`short:"w" description:"Run service with web UI"`
	WebUIPort	int		`short:"wp" description:"Web UI port"`
	WebUIHost	string	`short:"wh" description:"Web UI host"`
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
		WebUIPort: 8788,
		WebUIHost: "localhost",
	}
}
