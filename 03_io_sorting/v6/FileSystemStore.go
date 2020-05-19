package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	// ReadWriteSeeker has ability to point the reader from the beginning unlike io.Reader
	database io.ReadWriteSeeker
	league   League
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{
		database: database,
		league:   league,
	}
}

func (f *FileSystemPlayerStore) GetLeague() League {

	// f.database.Seek(0, 0)
	// league, _ := NewLeague(f.database)
	// return league
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(f.league)

	// league := f.GetLeague()
	// player := league.Find(name)

	// if player != nil {
	// 	player.Wins++
	// } else {
	// 	league = append(league, Player{name, 1})
	// }

	// f.database.Seek(0, 0)
	// json.NewEncoder(f.database).Encode(league)
}
