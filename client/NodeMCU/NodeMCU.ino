#include "connection.h"
#include <ESP8266WiFi.h>

#define PUSHBUTTON D1
const byte RGB[] = { D5, D6, D8, };

WiFiClient client;
volatile unsigned long reaction_time = 0;

struct communication {
	const byte magic   = 0xBB;
	const byte hello   = 0x00;
	const byte readyUp = 0x01;
	const byte fired   = 0x02;
	const byte result  = 0x03;
} packet;

ICACHE_RAM_ATTR void playerReaction();
void connect2AP();
int handshake();

ICACHE_RAM_ATTR void playerReaction()
{
	reaction_time = millis() - reaction_time;
}

void connect2AP() {
	unsigned long connection_start = millis();
	WiFi.begin(SSID, PASSWORD);

	Serial.println("\033[2J");  // clear previous serial output

	Serial.print("Connecting to: ");
	Serial.println(SSID);

	unsigned long wait = millis() - connection_start;

	// Live update time waiting for connection
	Serial.print("Waiting(ms): ");
	Serial.print(wait);
	do {
		byte digits = 0;

		while (wait) {
			wait /= 10;
			digits++;
		}

		for (byte i = 0; i < digits; i++) {
			Serial.print("\b");
		}

		wait = millis() - connection_start;
		Serial.print(wait);
		delay(20);

	} while (WiFi.status() != WL_CONNECTED);

	Serial.println();

	Serial.print("IP address: ");
	Serial.println(WiFi.localIP());
}

int handshake() {
	Serial.print("Connecting to: ");
	Serial.print(HOST);
	Serial.print(":");
	Serial.println(PORT);

	Serial.print("Connection state: ");
	if (client.connect(HOST, PORT)) {
		Serial.println("connected");
	} else {
		Serial.println("failed");
		goto error;
	}

	// Send handshake to server
	if (client.connected()) {
		Serial.println("Sending: handshake");
		client.write(packet.magic);
		client.write(packet.hello);
	}

	// Check if handshake was successful
	Serial.print("Handshake: ");
	while (client.read() != packet.magic);
	if (client.read() == packet.hello) {
		Serial.println("successful");
	} else {
		Serial.println("failed");
		goto error;
	}

	return 1;

error:
	return 0;
}

void setup()
{
	pinMode(PUSHBUTTON, INPUT);

	for (byte i = 0; i < sizeof(RGB) / sizeof(byte); i++) {
		pinMode(RGB[i], OUTPUT);
	}

	Serial.flush();
	Serial.begin(9600);

	attachInterrupt(digitalPinToInterrupt(PUSHBUTTON), playerReaction, RISING);

	connect2AP();          // Connect to specified access point
	while (!handshake());  // Try to handshake with game server
}

void loop()
{
}
