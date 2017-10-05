package scrabbler

import (
	"runtime"
	"fmt"
	"os"

	"github.com/AKovalevich/scrabbler/config"
	"github.com/AKovalevich/scrabbler/log"
	"github.com/containous/flaeg"
)

func Run(args []string) int {
	runtime.GOMAXPROCS(runtime.NumCPU())

	scrabblerConfiguration := config.NewScrabblerConfiguration()
	traefikPointersConfiguration := config.NewScrabblerDefaultConfiguration()

	scrabblerCmd := &flaeg.Command{
		Name:					"scrabbler",
		Description:			`scrabbler text classification`,
		Config:					scrabblerConfiguration,
		DefaultPointersConfig:	traefikPointersConfiguration,
		Run: func() error {
			start(scrabblerConfiguration)
			return nil
		},
	}

	healthCheckCmd := &flaeg.Command{
		Name:					"healthcheck",
		Description:			`Calls scrabbler /ping to check health (web provider must be enabled)`,
		Config:					struct{}{},
		DefaultPointersConfig:	struct{}{},
		Run: func() error {
			fmt.Print("OK")
			os.Exit(0)
			return nil
		},
		Metadata: map[string]string{
			"parseAllSources": "true",
		},
	}

	f := flaeg.New(scrabblerCmd, args)
	f.AddCommand(healthCheckCmd)
	f.AddCommand(newVersionCmd())

	//run test
	if err := f.Run(); err != nil {
		log.Error("Running error", log.String("error", err.Error()))
	}
	return 1
}

// start scrabbler application
func start(config *config.ScrabblerConfiguration) {
	//log.Info("Scrabbler started",
	//	log.String("Log level", config.LogLevel),
	//	log.Bool("Debug", config.Debug),
	//)
}
