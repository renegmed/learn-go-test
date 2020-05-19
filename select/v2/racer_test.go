package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		// run the servers
		slowServer := makeDelayedServer(20 * time.Second)
		fastServer := makeDelayedServer(0 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		// slowURL := slowServer.URL
		// fastURL := fastServer.URL

		want := fastServer.URL //fastURL

		// call the urls
		got, err := Racer(slowServer.URL, fastServer.URL)
		if err != nil {
			t.Errorf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {

		server := makeDelayedServer(25 * time.Second)
		defer server.Close()

		// call the urls
		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one.")
		}
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay) // sleep when client(url) calls the server
			w.WriteHeader(http.StatusOK)
		},
	))
}
