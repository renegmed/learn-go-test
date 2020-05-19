package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")

		playerStore := &StubPlayerStore{}
		cli := &CLI{playerStore, in}
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Chris")
	})
}
