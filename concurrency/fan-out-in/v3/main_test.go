package main

import (
	"fmt"
	"testing"
)

/*
The main can unblock all the senders simply by closing the done channel.

This close is effectively a broadcast signal to the senders. We extend each
of our pipeline functions to accept done as a parameter and arrange for the
close to happen via a defer statement, so that all return paths from main
will signal the pipeline stages to exit.
*/

func TestFanout(t *testing.T) {

	t.Run("1. Read all channel inputs", func(t *testing.T) {

		/*

			Set up a done channel that's shared by the whole pipeline,
			and close that channel whan this pipeline exits, as a signal
			for all the goroutines we started to exit.

		*/
		done := make(chan struct{})
		defer close(done)

		numsChan := generate(done, 1, 2, 3, 4)

		// Distribute the square and double works across multiple goroutines the beoth read from in.
		squareChan1 := square(done, numsChan)
		squareChan2 := square(done, numsChan)

		doubleChan1 := double(done, squareChan1)
		doubleChan2 := double(done, squareChan2)

		// Consume the merged output from doubleChan1 and doubleChan2

		var got []int

		mergedCh := merge(done, doubleChan1, doubleChan2)
		for n := range mergedCh {
			got = append(got, n) // all receiving channels are accounted for
		}
		//fmt.Println(got)

		if len(got) != 4 {
			t.Errorf("result slice length should be 4 not %d", len(got))
		}

	})

	t.Run("2. Read channel inputs but not all", func(t *testing.T) {

		/*

			Set up a done channel that's shared by the whole pipeline,
			and close that channel whan this pipeline exits, as a signal
			for all the goroutines we started to exit.

		*/
		done := make(chan struct{})
		defer close(done)

		numsChan := generate(done, 1, 2, 3, 4)

		// Distribute the square and double works across multiple goroutines the beoth read from in.
		squareChan1 := square(done, numsChan)
		squareChan2 := square(done, numsChan)

		doubleChan1 := double(done, squareChan1)
		doubleChan2 := double(done, squareChan2)

		// Consume the merged output from doubleChan1 and doubleChan2

		var got []int

		mergedCh := merge(done, doubleChan1, doubleChan2)

		got = append(got, <-mergedCh) // no locked up even though only 2 channels are received
		got = append(got, <-mergedCh)
		// got = append(got, <-mergedCh)
		// got = append(got, <-mergedCh)
		// got = append(got, <-mergedCh)
		// got = append(got, <-mergedCh)
		// got = append(got, <-mergedCh)

		fmt.Println(got)

		if len(got) != 2 {
			t.Errorf("result slice length should be 2 not %d", len(got))
		}

	})
}
