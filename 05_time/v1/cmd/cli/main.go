package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/renegmed/learn-go-test/05_time/v1"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	/*
		Note:

		StdOutAlerter has the method signature of function type BlindAlerterFunc
		BlindAlerterFunc implements the interface BlindAlerter, thus
		StdOutAlerter function can be passed as parameter of NewCLI
	*/
	poker.NewCLI(store, os.Stdin, poker.BlindAlerterFunc(poker.StdOutAlerter)).PlayPoker()
}
