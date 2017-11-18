package handler

import (
	"net/http"
	"log"
	"strings"
)

func CheckRequest(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		host := r.Host
		queryParas := r.URL.Query()
		headers := r.Header
		if len(headers["Content-Type"]) != 0 && !strings.Contains(headers["Content-Type"][0], "application/json") {
			http.Error(w, "Bad request - Go away!", 400)
		} else {
			next.ServeHTTP(w, r)
		}
		log.Println(host, path, queryParas, headers)
	}
	return http.HandlerFunc(fn)
}
