package main

import (
	"fmt"
	"net/http"
	"strings"
)

/*

This moved the score calculation out of the main body of our handler
into a function GetPlayerScore

*/
type PlayerStore interface {
	GetPlayerScore(name string) int
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
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}
