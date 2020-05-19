package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// This breaks the sequence sleep, write, etc
// func Countdown(out io.Writer, sleeper Sleeper) {
// 	for i := countdownStart; i > 0; i-- {
// 		sleeper.Sleep()
// 	}

// 	for i := countdownStart; i > 0; i-- {
// 		fmt.Fprintln(out, i)
// 	}

// 	sleeper.Sleep()
// 	fmt.Fprint(out, finalWord)
// }

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	//sleeper := &DefaultSleeper{}
	sleeper := &ConfigurableSleeper{1000 * time.Millisecond, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
