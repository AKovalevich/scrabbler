package route

import (
	"net/http"
)

type Route struct {
	Handler func(http.ResponseWriter, *http.Request)
	Path string
}
