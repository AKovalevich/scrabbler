package server

import (
	"os/signal"
	"syscall"
	"time"

	log "github.com/AKovalevich/scrabbler/log/logrus"

)

func (server *Server) configureSignals() {
	signal.Notify(server.signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
}

func (server *Server) listenSignals() {
	for {
		sig := <-server.signals
		switch sig {
		case syscall.SIGUSR1:
			log.Do.Infof("Signal: ", sig)
		default:
			log.Do.Infof("I have to go... %+v", sig)
			reqAcceptGraceTimeOut := time.Duration(3)
			if reqAcceptGraceTimeOut > 0 {
				log.Do.Infof("Waiting %s for incoming requests to cease", 3)
				time.Sleep(3)
			}
			log.Do.Info("Stopping server gracefully")
			//server.Stop()
		}
	}
}