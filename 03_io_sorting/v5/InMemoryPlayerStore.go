package main

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// implements PlayerStore interface
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

// implements PlayerStore interface
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

// implements PlayerStore interface
func (i *InMemoryPlayerStore) GetLeague() League {
	var league League

	// league := make([]Player, 0, len(i.store))
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}
