package restHelpers

import (
	"github.com/gorilla/mux"
	"net/http"
)

/*
	bindRoute
	handles binding with middleware that applies to everything

	@param {*mux.Router} r       - Router pointer.
	@param {string}      route   - The route we want to bind.
	@param {...}         arg     - The callback function to handle the request.
	@param {string       methods - The http method that can be used to access the route.

	@returns null
*/
func bindRoute(r *mux.Router, route string, handler func(w http.ResponseWriter, r *http.Request), methods string){
	r.HandleFunc(route, LimitMiddleware(handler)).Methods(methods)
	r.HandleFunc(route, LimitMiddleware(handler)).Methods("OPTIONS")
}

/*
	BindNoAuthRoutes
	Binds a new public route that doesn't require auth.

	@param {*mux.Router} r       - Router pointer.
	@param {string}      route   - The route we want to bind.
	@param {...}         arg     - The callback function to handle the request.
	@param {string       methods - The http method that can be used to access the route.

	@returns null
*/
func BindNoAuthPublicRoutes(r *mux.Router, route string, arg func(w http.ResponseWriter, r *http.Request, userId string, namespace string), methods string){
	bindRoute(r, route, handleNoAuth(arg), methods)
}

/*
	BindAuthRoute
	Binds a new public route that requires auth.

	@param {*mux.Router} r       - Router pointer.
	@param {string}      route   - The route we want to bind.
	@param {...}         arg     - The callback function to handle the request.
	@param {string       methods - The http method that can be used to access the route.

	@returns null
*/
func BindAuthPublicRoute(r *mux.Router, route string, arg func(w http.ResponseWriter, r *http.Request, userId string, namespace string), methods string){
	bindRoute(r, route, handleDiscordAuth(arg), methods)
}