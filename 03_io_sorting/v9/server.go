package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Player stores a name with a number of wins
type Player struct {
	Name string
	Wins int
}

// PlayerStore stores score information about players
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

// PlayerServer is a HTTP interface for player information
type PlayerServer struct {
	store        PlayerStore
	http.Handler // an interface with method ServeHTTP(ResponseWriter, *Request)
}

const jsonContentType = "application/json"

// NewPlayerServer creates a PlayerServer with routing configured
func NewPlayerServer(store PlayerStore) *PlayerServer {

	pServer := new(PlayerServer)

	pServer.store = store

	router := http.NewServeMux() // ServeMux has implementation of ServeHTTP(w ResponserWriter, r *Request)
	router.Handle("/league", http.HandlerFunc(pServer.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(pServer.playersHandler))

	pServer.Handler = router // assign emplementor of SeverHTTP(w ResponseWriter, r *Request)

	return pServer
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(p.store.GetLeague()) // ask playerServer's storage for league data
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player) // ask playerServer's storage for certain player's score
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player) // save player win record to playerSever's store
	w.WriteHeader(http.StatusAccepted)
}
