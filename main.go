package main

import (
	"ProductorConsumidor/models"
)

func main() {
	datachannel := make(chan int, 10)
	endChannel := make(chan bool)

	go models.Producer(datachannel)

	go models.Consumer(datachannel, endChannel)

	<-endChannel
}
