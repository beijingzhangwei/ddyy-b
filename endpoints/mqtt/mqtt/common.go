package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var broker = "144.168.63.73"
var port = 1883
var userName = "emqx"
var passwd = "2139511"
var topic = "topic/test"

func sub(client mqtt.Client, producer bool) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	if producer {
		fmt.Printf("Producer subscribed to topic %s", topic)
	} else {
		fmt.Printf("Consumer subscribed to topic %s", topic)
	}
}
