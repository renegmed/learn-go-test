package poker_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	poker "github.com/renegmed/learn-go-test/05_time/v1"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

type GameSpy struct {
	StartedWith  int
	FinishedWith string
	StartCalled  bool
	FinishCalled bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
	g.StartCalled = true
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
	g.FinishCalled = true
}

func TestCLI(t *testing.T) {
	// t.Run("record chris win from user input", func(t *testing.T) {
	// 	in := strings.NewReader("1\nChris wins\n")

	// 	playerStore := &poker.StubPlayerStore{}
	// 	game := poker.NewTexasHoldem(dummyBlindAlerter, playerStore)

	// 	cli := poker.NewCLI(in, dummyStdOut, game)
	// 	cli.PlayPoker()

	// 	poker.AssertPlayerWin(t, playerStore, "Chris")
	// })

	// t.Run("record cleo win from user input", func(t *testing.T) {
	// 	in := strings.NewReader("1\nCleo wins\n")

	// 	playerStore := &poker.StubPlayerStore{}
	// 	game := poker.NewTexasHoldem(dummyBlindAlerter, playerStore)
	// 	cli := poker.NewCLI(in, dummyStdOut, game)
	// 	cli.PlayPoker()

	// 	poker.AssertPlayerWin(t, playerStore, "Cleo")
	// })

	// t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
	// 	stdout := &bytes.Buffer{}
	// 	in := strings.NewReader("7\n")
	// 	game := &GameSpy{}

	// 	cli := poker.NewCLI(in, stdout, game)
	// 	cli.PlayPoker()

	// 	gotPrompt := stdout.String()
	// 	wantPrompt := poker.PlayerPrompt + poker.BadPlayerInputErrMsg

	// 	if gotPrompt != wantPrompt {
	// 		t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
	// 	}

	// 	if game.StartedWith != 7 {
	// 		t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
	// 	}
	// })

	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &GameSpy{}
		stdout := &bytes.Buffer{}

		in := userSends("3", "Chris wins")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertMessageSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &GameSpy{}

		in := userSends("8", "Cleo wins")
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non number value is entered and does not start the game", func(t *testing.T) {
		game := &GameSpy{}

		stdout := &bytes.Buffer{}
		in := userSends("Pies")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessageSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("it prints an error when the winner is declared incorrectly", func(t *testing.T) {
		game := &GameSpy{}

		stdout := &bytes.Buffer{}
		in := userSends("8", "Lloyd is a killer")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotFinished(t, game)
		assertMessageSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputMsg)
	})
}

func userSends(args ...string) io.Reader {
	var result = ""
	for _, p := range args {
		result = result + fmt.Sprintf("%s\n", p)
	}
	//return strings.NewReader(fmt.Sprintf("%s\n%s\n", numOfPlayers, name))
	return strings.NewReader(result)
}

func assertGameNotStarted(t *testing.T, game *GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func assertGameNotFinished(t *testing.T, game *GameSpy) {
	t.Helper()
	if game.FinishCalled {
		t.Errorf("game should not have finished")
	}
}

func assertGameStartedWith(t *testing.T, game *GameSpy, numberOfPlayersWanted int) {
	t.Helper()
	if game.StartedWith != numberOfPlayersWanted {
		t.Errorf("wanted Start called with %d but got %d", numberOfPlayersWanted, game.StartedWith)
	}
}

func assertFinishCalledWith(t *testing.T, game *GameSpy, winner string) {
	t.Helper()
	if game.FinishedWith != winner {
		t.Errorf("expected finish called with %q but got %s", winner, game.FinishedWith)
	}
}

func assertMessageSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, want)
	}
}
