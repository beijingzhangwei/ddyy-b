package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

var ProducerMqtt mqtt.Client
var broker = "144.168.63.73"
var port = 1883
var userName = "emqx"
var passwd = "2139511"
var topic = "topic/test"

func InitProducerClient() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_producer")
	opts.SetUsername(userName)
	opts.SetPassword(passwd)
	opts.SetKeepAlive(8 * time.Second)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	ProducerMqtt = mqtt.NewClient(opts)
	if token := ProducerMqtt.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	sub(ProducerMqtt, true)
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Producer Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func sub(client mqtt.Client, producer bool) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	if producer {
		fmt.Printf("Producer subscribed to topic %s", topic)
	} else {
		fmt.Printf("Consumer subscribed to topic %s", topic)
	}
}
