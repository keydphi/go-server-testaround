/*
 * message exchange
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/keydphi/MsgSrvc/0.1/",
		Index,
	},

	Route{
		"GetMessage",
		strings.ToUpper("Get"),
		"/keydphi/MsgSrvc/0.1/messages/{id}",
		GetMessage,
	},

	Route{
		"GetMessages",
		strings.ToUpper("Get"),
		"/keydphi/MsgSrvc/0.1/messages",
		GetMessages,
	},

	Route{
		"SendMessage",
		strings.ToUpper("Post"),
		"/keydphi/MsgSrvc/0.1/messages",
		SendMessage,
	},

	Route{
		"UpdateMessage",
		strings.ToUpper("Put"),
		"/keydphi/MsgSrvc/0.1/messages/{id}",
		UpdateMessage,
	},
}
