package middlewares

import (
	"Regret/domain"
	"Regret/src/clock"
	"log"
	"log/slog"
	"net/http"
	"strconv"
)

func UpdateServerTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		tick := request.Header.Get(domain.TickHeaderString)
		if tick == "" {
			slog.Warn("clock tick Not received in the request:  ", *request)
			next.ServeHTTP(writer, request)
		}
		clockInstance := clock.GetClockInstance()
		val, err := strconv.ParseUint(tick, 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse string: %v", err)
		}
		clockInstance.UpdateTick(val)
		next.ServeHTTP(writer, request)
	})
}
