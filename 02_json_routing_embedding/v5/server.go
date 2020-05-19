package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

/*

Go does not provide the typical, type-driven notion of subclassing, but it
does have the ability to “borrow” pieces of an implementation by embedding
types within a struct or interface.

What this means is that our PlayerServer now has all the methods that
http.Handler has, which is just ServeHTTP

WARNING: You must be careful with embedding types because you will
         expose all public methods and fields of the type you embed
*/
type PlayerServer struct {
	store        PlayerStore
	http.Handler // an interface with method ServeHTTP(ResponseWriter, *Request)
}

/*

It's quite odd (and inefficient) to be setting up a router as a request comes
in and then calling it. What we ideally want to do is have some kind of
NewPlayerServer function which will take our dependencies and do the
one-time setup of creating the router. Each request can then just use
that one instance of the router.


To "fill in" the http.Handler we assign it to the router we create in
NewPlayerServer. We can do this because http.ServeMux has the method ServeHTTP

*/
func NewPlayerServer(store PlayerStore) *PlayerServer {

	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux() // ServeMux has implementation of ServeHTTP(w ResponserWriter, r *Request)
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router // assign emplementor of SeverHTTP(w ResponseWriter, r *Request)

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(p.store.GetLeague())
	w.WriteHeader(http.StatusOK)
}

// func (p *PlayerServer) getLeagueTable() []Player {
// 	return []Player{
// 		{"Chirs", 20},
// 	}
// }

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
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
