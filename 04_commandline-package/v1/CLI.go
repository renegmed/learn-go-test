package poker

import "io"

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
	cli.playerStore.RecordWin("Chris")
}
