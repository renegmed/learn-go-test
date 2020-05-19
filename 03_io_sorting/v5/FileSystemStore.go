package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	// ReadWriteSeeker has ability to point the reader from the beginning unlike io.Reader
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() League {

	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	// var wins int
	// for _, player := range f.GetLeague() {
	// 	if player.Name == name {
	// 		wins = player.Wins
	// 		break
	// 	}
	// }
	// return wins
	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	// league := f.GetLeague()
	// for i, player := range league {
	// 	if player.Name == name {
	// 		// why league[i] not player.Wins+++
	// 		// range over a slice you are returned the current index
	// 		// of the loop (in our case i) and a copy of the element
	// 		// at that index
	// 		league[i].Wins++
	// 	}
	// }

	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)

	// THIS ALSO WORKS
	// players := []Player{}
	// for _, player := range f.GetLeague() {
	// 	if player.Name == name {
	// 		player.Wins++
	// 	}
	// 	players = append(players, player)
	// }
	// f.database.Seek(0, 0)
	// json.NewEncoder(f.database).Encode(players)
}
