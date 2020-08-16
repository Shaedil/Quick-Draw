#include "connection.h"
#include <ESP8266WiFi.h>

#define PUSHBUTTON D1
const byte RGB[] = { D5, D6, D8, };

WiFiClient client;

struct communication {
	const byte magic   = 0xBB;
	const byte hello   = 0x00;
	const byte readyUp = 0x01;
	const byte fired   = 0x02;
	const byte result  = 0x03;
	const byte ffs     = 0xCA;
} packet;

volatile bool canShoot = false;

ICACHE_RAM_ATTR void playerShoot();
void connect2AP();
int handshake();
void readyUp();
void duel();

ICACHE_RAM_ATTR void playerShoot()
{
	if (canShoot) {
		Serial.println("You have shot!");
		client.write(packet.magic);
		client.write(packet.fired);
		canShoot = false;
	}
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

void readyUp() {
	while (client.read() != packet.magic);
	if (client.read() == packet.readyUp) {
		// Hamilton Easter egg :3
		Serial.println("Summon all the courage you require");
		Serial.println("Then count");
		Serial.println("one");
		delay(1000);
		Serial.println("two");
		delay(1000);
		Serial.println("three");
		delay(1000);
		Serial.println("four");
		delay(1000);
		Serial.println("five");
		delay(1000);
		Serial.println("six");
		delay(1000);
		Serial.println("seven");
		delay(1000);
		Serial.println("eight");
		delay(1000);
		Serial.println("nine");
		delay(1000);
		Serial.println("Number");
		Serial.println("Ten paces!");
		delay(1000);
		Serial.println("Fire!");

		canShoot = true;
	}
}

void duel()
{
	while (client.read() != packet.magic);
	if (client.read() == packet.result) {
		Serial.print("The results are in: ");

		while (client.read() != packet.ffs);
		Serial.println(client.read() ? "you won!" : "you lost!");
	}
}

void setup()
{
	pinMode(PUSHBUTTON, INPUT);

	for (byte i = 0; i < sizeof(RGB) / sizeof(byte); i++) {
		pinMode(RGB[i], OUTPUT);
	}

	Serial.flush();
	Serial.begin(9600);

	attachInterrupt(digitalPinToInterrupt(PUSHBUTTON), playerShoot, RISING);

	connect2AP();          // Connect to specified access point
	while (!handshake());  // Try to handshake with game server
}

void loop()
{
	readyUp(); // Wait for opponent to connect to server
	duel();
}
