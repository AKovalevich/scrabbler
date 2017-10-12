package server

import (
	"os/signal"
	"net/http"
	"context"
	"time"
	"sync"
	"os"

	log "github.com/AKovalevich/scrabbler/log/logrus"
	"github.com/AKovalevich/scrabbler/config"
)

// Server is the reverse-proxy/load-balancer engine
type Server struct {
	mainConfiguration *config.ScrabblerConfiguration
	signals							chan os.Signal
	stopChan						chan bool
	mainHttpServer					*http.Server
	webUiHttpServer					*http.Server
}

func NewServer(config *config.ScrabblerConfiguration) Server {
	var server Server
	server.mainConfiguration = config

	log.Do.Info("Configure signals and listeners...")
	// Configure signals
	server.configureSignals()
	go server.listenSignals()

	log.Do.Info("Configure HTTP services...")
	// Configure main scrabbler HTTP server
	server.configureMainHttpServer()
	server.configureWebUiHttpServer()

	return server
}

func (server *Server) Serve() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		go server.runMainServer()
		go server.runWebUiServer()
	}()
	wg.Wait()
}

func (server *Server) Stop() {
	defer log.Do.Info("Server stopped")
	var entryPoints []*http.Server
	if server.mainConfiguration.WebUI {
		entryPoints = append(entryPoints, server.webUiHttpServer)
	}
	entryPoints = append(entryPoints, server.mainHttpServer)
	var wg sync.WaitGroup
	for _, v := range entryPoints {
		wg.Add(1)
		go func(srv *http.Server) {
			defer wg.Done()
			graceTimeOut := time.Duration(server.mainConfiguration.GraceTimeOut)
			ctx, cancel := context.WithTimeout(context.Background(), graceTimeOut)
			log.Do.Debugf("Waiting %s seconds before killing connections on %s...", graceTimeOut, "server")
			if err := srv.Shutdown(ctx); err != nil {
				log.Do.Debugf("Wait is over due to: %s", err)
				srv.Close()
			}
			cancel()
			log.Do.Debugf("Entry point is closed")
		}(v)
	}

	wg.Wait()
	server.Close()
}

func (server * Server) Close() {
	// Close Web UI HTTP server
	signal.Stop(server.signals)
	close(server.signals)
	os.Exit(1)
}
