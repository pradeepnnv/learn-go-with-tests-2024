package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of two servers, returns url of the fastest one", func(t *testing.T) {
		// slowURL := "https://facebook.com"
		// fastURL := "https://quii.dev"
		slowServer := makeDelayedServer(20 * time.Millisecond)
		defer slowServer.Close()
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer fastServer.Close()
		slowURL := slowServer.URL
		fastURL := fastServer.URL
		want := fastURL
		got, err := Racer(slowURL, fastURL)
		if err != nil {
			t.Errorf("got an error when expected none")
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if server does not respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)
		if err == nil {
			t.Errorf("expecting error but got none")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
