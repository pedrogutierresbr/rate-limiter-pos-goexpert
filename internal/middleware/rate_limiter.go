package middleware

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/pedrogutierresbr/rate-limiter-pos-goexpert/internal/limiter"
)

func RateLimiter(rl *limiter.RateLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = r.RemoteAddr
		}
		token := r.Header.Get("API_KEY")
		ctx := context.Background()
		if token != "" {
			if !rl.Allow(ctx, "token:"+token, rl.LimitToken) {
				log.Println("Token limit exceeded:", token)
				http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
				return
			}
		} else {
			if !rl.Allow(ctx, "ip:"+ip, rl.LimitIP) {
				log.Println("IP limit exceeded:", ip)
				http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
