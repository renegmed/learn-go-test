package poker

/*

* We are creating our two components we are trying to integrate
with: InMemoryPlayerStore and PlayerServer.

* We then fire off 3 requests to record 3 wins for player. We're
not too concerned about the status codes in this test as it's
not relevant to whether they are integrating well.

* The next response we do care about (so we store a variable response)
because we are going to try and get the player's score.

*/
import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "")
	defer cleanDatabase()

	store, err := NewFileSystemPlayerStore(database)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server := NewPlayerServer(store)

	player := "Pepper"

	// put three player records to be stored
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("1. get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("2. get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})

}
