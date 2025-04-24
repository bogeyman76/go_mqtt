
Rought draft:

This script assumes you have a .env file, located in the root folder of the project, with the following type defined in a JSON string:

type MQTT_creds struct {
	KeyPath  string
	CertPath string
	CaPath   string
	ClientID string
	Host     string
	Port     string
}

MQTT_CREDS={"KeyPath":"path","CertPath":"path to the cert", ...} etc.

Make sure to add tls:// to your AWS IOT endpoint!

The cert files will be generated when you create your AWS "thing".  

Debugging:

You will not be able to connect if you have another device connected using the same credentials.  Please use a unique "thing" to connect.

You can publish from the AWS IOT test client to "topic/test" to ensure communication to AWS is happening.
