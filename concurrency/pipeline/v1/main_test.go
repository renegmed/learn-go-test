package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPipeline(t *testing.T) {
	t.Run("push an item through threee-stage pipelines.", func(t *testing.T) {

		// gen channel for a given set of items
		// receive item from third stage pipeline
		// assert

		want := []int{2, 8, 18, 32}
		//done := make(chan bool)
		numsChan := generate(1, 2, 3, 4)
		squareChan := square(numsChan)
		doubleChan := double(squareChan)

		var got []int
		// for {
		// 	select {
		// 	case n := <-doubleChan:
		// 		//fmt.Println(n)
		// 		got = append(got, n)

		// 	case <-done:
		// 		break //	t.Fatalf("not expecting default case.")
		// 	}
		// }
		got = append(got, <-doubleChan)
		got = append(got, <-doubleChan)
		got = append(got, <-doubleChan)
		got = append(got, <-doubleChan)

		fmt.Println(got)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d not equal", got, want)
		}

	})
}
