package textclassifier

import (
	"fmt"
	"net/http"

	"github.com/AKovalevich/scrabbler/route"
)

func GetRoutes() []route.Route {
	return []route.Route{
		{
			Path: "/train",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "Welcome!\n")
			},
		},
	}
}
