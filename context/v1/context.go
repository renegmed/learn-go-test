package context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select { // select effectively race to the two asynchronous processes
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done(): // returns a channel which gets sent a signal when context is 'done' or 'cancelled'
			store.Cancel()
		}

	}
}
