package context

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				//log.Logger("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
				//fmt.Println("--result:", result)
			}
		}

		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

// to simulate http.ResponseWriter
type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestHandler(t *testing.T) {

	data := "hello, world"

	t.Run("returns data from store", func(t *testing.T) {

		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		//store.assertWasNotCancelled()

	})

	t.Run("tells store to cancel work if request is cancelled.", func(t *testing.T) {

		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// The context package provides functions to derive new Context
		// values from existing ones. These values form a tree: when
		// a Context is canceled, all Contexts derived from it are also canceled.

		// It's important that you derive your contexts so that cancellations
		// are propagated throughout the call stack for a given request.

		// What we do is derive a new cancellingCtx from our request which
		// returns us a cancel function. We then schedule that function to
		// be called in 5 milliseconds by using time.AfterFunc.

		// Finally we use this new context in our request by calling request.WithContext.

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(35*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written, processed was cancelled.")
		}
		//store.assertWasCancelled()

	})
}

// func (s *SpyStore) assertWasCancelled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Errorf("store was not told to cancel")
// 	}
// }

// func (s *SpyStore) assertWasNotCancelled() {
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Errorf("store was told to cancel")
// 	}
// }
