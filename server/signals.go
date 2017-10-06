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

		default:
			log.Do.Infof("I have to go... %+v", sig)
			reqAcceptGraceTimeOut := time.Duration(server.globalConfiguration.LifeCycle.RequestAcceptGraceTimeout)
			if reqAcceptGraceTimeOut > 0 {
				log.Do.Infof("Waiting %s for incoming requests to cease", reqAcceptGraceTimeOut)
				time.Sleep(reqAcceptGraceTimeOut)
			}
			log.Do.Info("Stopping server gracefully")
			server.Stop()
		}
	}
}