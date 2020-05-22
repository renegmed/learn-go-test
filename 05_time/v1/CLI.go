package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const PlayerPrompt = "Please enter the number of players: "
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with number"
const BadWinnerInputMsg = "invalid winner input, expect format of 'PlayerName wins'"

/*

os.Stdin is what we'll use in main to capture the user's input. It is
a *File under the hood which means it implements io.Reader
is a handy way of capturing text

*/

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

// CLI being concerned with input/output and delegating game specific actions to a Game
func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayersInput := cli.readLine()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))
	if err != nil {
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.game.Start(numberOfPlayers)

	winnerInput := cli.readLine()
	winner, err := extractWinner(winnerInput)
	if err != nil {
		fmt.Fprint(cli.out, BadWinnerInputMsg)
		return
	}

	cli.game.Finish(winner)
}

func extractWinner(userInput string) (string, error) {
	if !strings.HasSuffix(userInput, " wins") {
		return "", fmt.Errorf(BadWinnerInputMsg)
	}
	return strings.Replace(userInput, " wins", "", 1), nil
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text() // Scanner.Text() to return the string the scanner read to
}
