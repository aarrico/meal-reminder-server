package supplements

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var controller = &controller{Repository: Repository{}}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Authentication",
		"POST",
		"/get-token",
		controller.GetToken,
	},
	Route{
		"GetSupplement",
		"GET",
		"/supplements/{id:[a-zA-Z0-9]*}",
		AuthenticationMiddleware(controller.GetSupplement),
	},
	Route{
		"AddSupplement",
		"GET",
		"/supplements",
		AuthenticationMiddleware(controller.AddSupplement),
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
