package server

import (
	"os/signal"
	"syscall"
	"time"
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
		case syscall.SIGINT:
		case syscall.SIGTERM:
			log.Do.Info("Fast shutdown")
			break
		case syscall.SIGUSR1:
			log.Do.Info("Re-opening log files")
			break
		case syscall.SIGQUIT:
			log.Do.Info("Graceful shutdown")
			break
		case syscall.SIGHUP:
			log.Do.Info("Changing configuration")
			break
		default:
			log.Do.Info("Stopping server gracefully")
			//server.Stop()
		}
	}
}
