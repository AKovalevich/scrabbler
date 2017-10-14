package server

import (
	"os/signal"
	"syscall"
	"os"

	log "github.com/AKovalevich/scrabbler/log/logrus"
)

func (server *Server) configureSignals() {
	server.signals = make(chan os.Signal, 1)
	signal.Notify(server.signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
}

func (server *Server) listenSignals() {
	for {
		sig := <-server.signals
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			log.Do.Info("Graceful shutdown...")
			server.Stop()
		case syscall.SIGUSR1:
			log.Do.Info("Re-opening log files")
		case syscall.SIGHUP:
			server.mainConfiguration.Reload()
		}
	}
}
