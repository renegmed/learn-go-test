package main

import "sync"

func generate(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}

		close(out)
	}()

	return out
}

func square(numsCh chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range numsCh {

			out <- n * n
		}
		close(out)
	}()
	return out
}

func double(nums chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range nums {

			out <- n * 2
		}
		close(out)
	}()
	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {

	var wg sync.WaitGroup
	out := make(chan int)

	// start an output goroutine for each input channel in
	// cs.  output copies values from c to out until c is closed,
	// then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			select {
			case out <- n:
			case <-done:
			}
		}

		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines
	// are done. This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out

}
