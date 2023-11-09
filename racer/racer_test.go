package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	// slowURL := "http://www.facebook.com"
	// fastURL := "http://www.quii.dev"

	// Mocking - START
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0)
	defer slowServer.Close()
	defer fastServer.Close()
	slowURL := slowServer.URL
	fastURL := fastServer.URL
	// Mocking - END
	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
