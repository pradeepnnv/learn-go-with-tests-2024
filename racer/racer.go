package racer

import (
	"net/http"
	"time"
)

func Racer(s, f string) (w string) {

	if measureResponseTime(s) > measureResponseTime(f) {
		return f
	}
	return s
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
