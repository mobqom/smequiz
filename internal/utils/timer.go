package utils

import "time"

func Timer(second int) chan int {
	counter := make(chan int)
	if second < 0 {
		close(counter)
		return counter
	}
	go func() {
		for i := second; i >= 0; i-- {
			counter <- i
			time.Sleep(1 * time.Second)
		}
		close(counter)
	}()
	return counter
}
