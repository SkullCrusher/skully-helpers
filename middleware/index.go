package middleware

import (
	"net/http"
)


/*
	forceHeaders
	Handle forcing the headers for the rest calls.

	@param {http.ResponseWriter} w - The request we want to force the headers on.

	@return null
*/
func ForceHeaders(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
}
