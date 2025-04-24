package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var client mqtt.Client

type MQTT_creds struct {
	KeyPath  string
	CertPath string
	CaPath   string
	ClientID string
	Host     string
	Port     string
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

func NewTlsConfig(mqtt_creds MQTT_creds) *tls.Config {
	certpool := x509.NewCertPool()
	ca, err := os.ReadFile(mqtt_creds.CaPath)
	if err != nil {
		log.Fatalln(err.Error())
	}

	certpool.AppendCertsFromPEM(ca)
	// Import client certificate/key pair
	clientKeyPair, err := tls.LoadX509KeyPair(mqtt_creds.CertPath, mqtt_creds.KeyPath)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		RootCAs:            certpool,
		ClientAuth:         tls.NoClientCert,
		ClientCAs:          nil,
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{clientKeyPair},
	}
}

func MQTTSub(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}

func MQTTPub(client mqtt.Client, topic string, msg any) {
	tkn := client.Publish(topic, 1, false, msg)
	tkn.Wait()
}

func GetClient() *mqtt.Client {  // returns the client address so you can apply it to MQTTPub in other packages
	return &client
}

func MQTT_connect() {

	mqtt_creds := MQTT_creds{}
	err := json.Unmarshal([]byte(os.Getenv("MQTT_CREDS")), &mqtt_creds)
	if err != nil {
		log.Println(err)
		return
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s:%s", mqtt_creds.Host, mqtt_creds.Port))
	tlsConfig := NewTlsConfig(mqtt_creds)
	opts.SetTLSConfig(tlsConfig)
	opts.SetClientID(mqtt_creds.ClientID)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	fmt.Printf("Connecting to MQTT...")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	MQTTSub(client, `testing`)  // remove this if you want to subscribe in other packages. For example only.
}
