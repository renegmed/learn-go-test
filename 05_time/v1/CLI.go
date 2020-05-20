package poker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

/*

os.Stdin is what we'll use in main to capture the user's input. It is
a *File under the hood which means it implements io.Reader
is a handy way of capturing text

*/
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
	alerter     BlindAlerter
}

func NewCLI(store PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in), // use a bufio.Scanner to read the input from the io.Reader.
		alerter:     alerter,
	}
}
func (cli *CLI) PlayPoker() {
	cli.alerter.ScheduleAlertAt(5*time.Second, 100)
	userInput := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text() // Scanner.Text() to return the string the scanner read to
}
