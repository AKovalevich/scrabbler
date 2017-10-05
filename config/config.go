package config

type ScrabblerConfiguration struct {
	Debug		bool	`short:"d" description:"Enable debug mode" export:"true"`
	LogLevel	string	`short:"l" description:"Log level" export:"true"`
	Port		int		`short:"p" description:"Scrabbler web server port"`
}

func NewScrabblerConfiguration() *ScrabblerConfiguration {
	return &ScrabblerConfiguration{}
}

func NewScrabblerDefaultConfiguration() *ScrabblerConfiguration {
	return &ScrabblerConfiguration{
		Debug: false,
		LogLevel: "info",
	}
}
