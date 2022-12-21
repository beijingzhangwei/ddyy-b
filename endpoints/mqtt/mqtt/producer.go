package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

var producerMqtt mqtt.Client

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Producer Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func producerPoint() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_producer")
	opts.SetUsername(userName)
	opts.SetPassword(passwd)
	opts.SetKeepAlive(8 * time.Second)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client, true)
	publish(client)

	time.Sleep(30 * time.Second)
	client.Disconnect(250)
}

func publish(client mqtt.Client) {
	num := 1000
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("%d", i)
		token := client.Publish(topic, 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}
