package server


import (
	"fmt"

	log "github.com/AKovalevich/scrabbler/log/logrus"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"os"
	"github.com/cdipaolo/goml/cluster"
	"net/http"
)

// Server is the reverse-proxy/load-balancer engine
type Server struct {
	configurationChan             chan types.ConfigMessage
	configurationValidatedChan    chan types.ConfigMessage
	signals                       chan os.Signal
	stopChan                      chan bool
	providers                     []provider.Provider
	currentConfigurations         safe.Safe
	globalConfiguration           configuration.GlobalConfiguration
	accessLoggerMiddleware        *accesslog.LogHandler
	routinesPool                  *safe.Pool
	leadership                    *cluster.Leadership
	defaultForwardingRoundTripper http.RoundTripper
	metricsRegistry               metrics.Registry
}

func Serve() {
	router := fasthttprouter.New()
	router.GET("/", func (ctx *fasthttp.RequestCtx) {
		fmt.Fprint(ctx, "Welcome!\n")
	})
	log.Do.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
