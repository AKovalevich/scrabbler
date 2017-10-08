package scrabbler

import (
	"runtime"
	"fmt"
	"os"

	log "github.com/AKovalevich/scrabbler/log/logrus"
	"github.com/AKovalevich/scrabbler/config"
	"github.com/AKovalevich/scrabbler/server"
	"github.com/containous/flaeg"
	"sync"
)

func Run(args []string) int {
	runtime.GOMAXPROCS(runtime.NumCPU())

	scrabblerConfiguration := config.NewScrabblerConfiguration()
	scrabblerPointersConfiguration := config.NewScrabblerDefaultConfiguration()

	scrabblerCmd := &flaeg.Command{
		Name:					"scrabbler",
		Description:			`scrabbler text classification`,
		Config:					scrabblerConfiguration,
		DefaultPointersConfig:	scrabblerPointersConfiguration,
		Run: func() error {
			start(scrabblerConfiguration)
			return nil
		},
	}

	healthCheckCmd := &flaeg.Command{
		Name:					"healthcheck",
		Description:			`Calls scrabbler /ping to check health`,
		Config:					struct{}{},
		DefaultPointersConfig:	struct{}{},
		Run: func() error {
			fmt.Print("OK")
			os.Exit(0)
			return nil
		},
		Metadata: map[string]string {
			"parseAllSources": "true",
		},
	}

	f := flaeg.New(scrabblerCmd, args)
	f.AddCommand(healthCheckCmd)
	f.AddCommand(newVersionCmd())

	if err := f.Run(); err != nil {
		log.Do.Error("Running error: ", err.Error())
	}
	return 1
}

// Start scrabbler application
func start(config *config.ScrabblerConfiguration) {
	log.Do.Infof("Scrabbler started")
	log.Do.Infof("PID: %d\n", os.Getpid())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server.Serve()
	}()
	wg.Wait()
}
