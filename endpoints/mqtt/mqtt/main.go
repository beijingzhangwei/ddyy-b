package main

import "time"

func main() {
	go consumerPoint()
	go producerPoint()
	time.Sleep(300 * time.Second)
}
