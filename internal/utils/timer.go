package utils

import "time"

func Timer(second int) chan int {
	counter := make(chan int)
	go func() {
		for i := second; i >= 0; i-- {
			counter <- i
			time.Sleep(1 * time.Second)
		}
		close(counter)
	}()
	return counter
}
