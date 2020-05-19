package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	// ReadSeeker has ability to point the reader rom the beginning unlike io.Reader
	database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	// var league []Player

	// err := json.NewDecoder(f.database).Decode(&league)
	// if err != nil {
	// 	fmt.Printf("Error on decoding file system player store: %q", err)
	// }
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int
	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}

	return wins

	// f.database.Seek(0, 0)
	// league, _ := NewLeague(f.database)
	// for _, player := range league {
	// 	if player.Name == name {
	// 		return player.Wins
	// 	}
	// }
	// return 0
}
