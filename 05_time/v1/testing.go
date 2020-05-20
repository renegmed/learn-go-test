package poker

import "testing"

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

// implements PlayerStore interface
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

// implements PlayerStore interface
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

// implements PlayerStore interface
func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

//checks to see if wins are recorded
func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], winner)
	}
}
