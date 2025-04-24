
<h2>Rought draft</h2>

This script assumes you have a .env file, located in the root folder of the project, with the following type defined in a JSON string:

<code>
type MQTT_creds struct { <br>
&nbsp; &nbsp;KeyPath&nbsp; &nbsp;&nbsp;&nbsp;string  <br>
&nbsp; &nbsp;CertPath&nbsp;&nbsp;&nbsp;&nbsp;string  <br>
&nbsp; &nbsp;CaPath&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;string  <br>
&nbsp; &nbsp;ClientID&nbsp; &nbsp;&nbsp;&nbsp;string  <br>
&nbsp; &nbsp;Host&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;string  <br>
&nbsp; &nbsp;Port&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;string  <br>
}
</code>

<br>

<code>MQTT_CREDS={"KeyPath":"path","CertPath":"path to the cert", ...} </code> etc.

Make sure to add tls:// to your AWS IOT endpoint!

The cert files will be generated when you create your AWS "thing".  

<h3>Implementation</h3>

import the package and then execute the following in your root function: 	

<code>mqtt.MQTT_connect()</code>

<h3>Debugging</h3>

You will not be able to connect if you have another device connected using the same credentials.  Please use a unique "thing" to connect.

You can publish from the AWS IOT test client to "topic/test" to ensure communication to AWS is happening.
