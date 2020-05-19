package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/*
This approach has a problem: each downstream receiver needs to know the
number of potentially blocked upstream senders and arrange to signal
those senders on early return. Keeping track of these counts is tedious
and error-prone.
*/
func TestFanout(t *testing.T) {

	t.Run("read same channel from multiple functions (fan-out).", func(t *testing.T) {

		want := []int{2, 8, 18, 32}

		numsChan := generate(1, 2, 3, 4)

		// Distribute the square and double works across multiple goroutines the beoth read from in.
		squareChan1 := square(numsChan)
		squareChan2 := square(numsChan)

		doubleChan1 := double(squareChan1)
		doubleChan2 := double(squareChan2)

		// Consume the merged output from doubleChan1 and doubleChan2

		var got []int

		done := make(chan struct{}, 2)
		mergedCh := merge(done, doubleChan1, doubleChan2)
		got = append(got, <-mergedCh)
		got = append(got, <-mergedCh)
		got = append(got, <-mergedCh)
		got = append(got, <-mergedCh)

		// Tell the remaining senders we are leaving
		done <- struct{}{}
		done <- struct{}{}

		sort.Ints(got)

		fmt.Println(want)
		fmt.Println(got)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d not equal", got, want)
		}

	})

	t.Run("read channel inputs but not all", func(t *testing.T) {

		//want := []int{2, 8, 18, 32}

		numsChan := generate(1, 2, 3, 4)

		// Distribute the square and double works across multiple goroutines the beoth read from in.
		squareChan1 := square(numsChan)
		squareChan2 := square(numsChan)

		doubleChan1 := double(squareChan1)
		doubleChan2 := double(squareChan2)

		// Consume the merged output from doubleChan1 and doubleChan2

		var got []int

		done := make(chan struct{}, 2)
		mergedCh := merge(done, doubleChan1, doubleChan2)
		got = append(got, <-mergedCh) // for 2
		got = append(got, <-mergedCh) // for 8
		// got = append(got, <-mergedCh)   // for 18
		// got = append(got, <-mergedCh)  // for 32

		// Tell the remaining senders we are leaving
		// potentially four blocked senders
		done <- struct{}{} // for first merged Ch
		done <- struct{}{} // for second merged Ch
		done <- struct{}{} // for first in the buffer
		done <- struct{}{} // for second in the buffer
		//done <- struct{}{} uncomment this will cause to lock up, nothing is being produced. 2 are produced and 2  are in the buffer

		// sort.Ints(got)
		fmt.Println(got)

		if len(got) != 2 {
			t.Errorf("result slice length should be 2 not %d", len(got))
		}

	})
}
