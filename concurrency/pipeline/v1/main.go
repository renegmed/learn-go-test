package main

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
