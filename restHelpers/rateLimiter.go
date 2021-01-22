package restHelpers

import (
	"../../global"
	"../middlewares"
	"golang.org/x/time/rate"
	"net/http"
)

// Setup the rate limiter for network calls.
var limiter = middlewares.NewIPRateLimiter(rate.Limit(global.RateLimiterRefillRate), global.RateLimiterDefaultStock)

/*
	limitMiddleware
	Attaches a rate limiter to network requests (5 requests per second cap).
*/
func LimitMiddleware(arg func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {

		// Get the ip address from cloudflare or raw ip.
		ipAddress := GetIP(r)

		limiter := limiter.GetLimiter(ipAddress) // r.RemoteAddr)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		arg(w, r)
	}
}
