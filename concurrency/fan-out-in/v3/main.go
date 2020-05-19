package main

import "sync"

/*

Each of our pipeline stages is now free to return as soon as done is closed.
The output routine in merge can return without draining its inbound channel,
since it knows the upstream sender, square and double, will stop attempting to send when done
is closed. output ensures wg.Done is called on all return paths via a defer
statement:

*/
func generate(done <-chan struct{}, nums ...int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()

	return out
}

func square(done <-chan struct{}, numsCh chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range numsCh {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func double(done <-chan struct{}, nums chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range nums {
			select {
			case out <- n * 2:
			case <-done:
				return
			}
		}

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
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
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
