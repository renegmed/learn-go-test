package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

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

		mergedCh := merge(doubleChan1, doubleChan2)

		got = append(got, <-mergedCh)
		got = append(got, <-mergedCh)
		got = append(got, <-mergedCh)
		got = append(got, <-mergedCh) // comment this last number won't be received

		sort.Ints(got)

		fmt.Println(want)
		fmt.Println(got)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d not equal", got, want)
		}

	})
}
