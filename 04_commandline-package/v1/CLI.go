package poker

import (
	"bufio"
	"io"
	"strings"
)

/*

os.Stdin is what we'll use in main to capture the user's input. It is
a *File under the hood which means it implements io.Reader
is a handy way of capturing text

*/
type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

func (cli *CLI) PlayPoker() {
	reader := bufio.NewScanner(cli.in) // use a bufio.Scanner to read the input from the io.Reader.
	reader.Scan()

	cli.playerStore.RecordWin(extractWinner(reader.Text())) // Scanner.Text() to return the string the scanner read to
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
