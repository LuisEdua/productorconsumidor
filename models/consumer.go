package models

import (
	"fmt"
)

func Consumer(ch <-chan int, ech chan<- bool) {
	for data := range ch {
		fmt.Printf("Consumidor recibió: %d\n", data)
	}
	fmt.Println("El canal está cerrado. Consumidor terminado.")

	ech <- true
}
