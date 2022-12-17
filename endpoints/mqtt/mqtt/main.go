package main

import "time"

func main() {
	go consumerPoint()
	go producerPoint()
	time.Sleep(30 * time.Second)
}
