package models

func Producer(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
}
