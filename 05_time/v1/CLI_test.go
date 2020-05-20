package poker_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	poker "github.com/renegmed/learn-go-test/05_time/v1"
)

type scheduleAlert struct {
	at     time.Duration
	amount int
}

func (s scheduleAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []scheduleAlert
}

// implements interface BlindAlerter
func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduleAlert{duration, amount})
}

var dummySpyAlerter = &SpyBlindAlerter{}

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")

		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")

		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		cases := []scheduleAlert{
			{0 * time.Second, 100},
			{10 * time.Second, 200},
			{20 * time.Second, 300},
			{30 * time.Second, 400},
			{40 * time.Second, 500},
			{50 * time.Second, 600},
			{60 * time.Second, 800},
			{70 * time.Second, 1000},
			{80 * time.Second, 2000},
			{90 * time.Second, 4000},
			{100 * time.Second, 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {

				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduleAlert(t, got, want)
			})
		}
	})
}

func assertScheduleAlert(t *testing.T, got, want scheduleAlert) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got scheduled alert of %v, want %v", got, want)
	}
	// amountGot := alert.amount
	// 			if amountGot != c.expectedAmount {
	// 				t.Errorf("got amount %d, want %d", amountGot, c.expectedAmount)
	// 			}

	// 			gotScheduledTime := alert.scheduleAt
	// 			if gotScheduledTime != c.expectedScheduleTime {
	// 				t.Errorf("got scheduled time of %v, want %v", gotScheduledTime, c.expectedScheduleTime)
	// 			}

}
