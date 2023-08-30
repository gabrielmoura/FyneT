package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Mqtt struct {
	Client  mqtt.Client
	Message mqtt.Message
}
type MqttConfig struct {
	Broker   string
	ClientId string
	Username string
	Password string
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func (m *Mqtt) Connect(config MqttConfig) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(config.Broker)

	opts.SetClientID(config.ClientId)
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	m.Client = mqtt.NewClient(opts)

	if token := m.Client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
func (m *Mqtt) Sub(topic string, callback func(string)) {
	Message := ""
	m.Client.Subscribe(topic, 0, func(client mqtt.Client, message mqtt.Message) {
		Message = string(message.Payload())
		fmt.Printf("%v: Received message: %s from topic: %s\n", message.MessageID(), message.Payload(), message.Topic())
		callback(Message)
		//message.Ack()
	})
}
func (m *Mqtt) Pub(topic string, message string) {
	m.Client.Publish(topic, 0, false, message)
}
