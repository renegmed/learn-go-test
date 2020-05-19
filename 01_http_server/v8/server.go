package main

import (
	"fmt"
	"net/http"
	"strings"
)

/*

   To be able to store new scores.

*/
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

/*
implement the Handler interface by adding a method to our new struct
and putting in our existing handler code
*/
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}

}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	//fmt.Println("++++ player:", player)
	score := p.store.GetPlayerScore(player)
	//fmt.Println("++++++ score:", score)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
