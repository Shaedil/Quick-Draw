#include "connection.h"

#define PUSHBUTTON D1
const byte RGB[] = { D5, D6, D8, };

volatile unsigned long reaction_time = 0;

ICACHE_RAM_ATTR void playerReaction();

ICACHE_RAM_ATTR void playerReaction()
{
	reaction_time = millis() - reaction_time;
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
}

void loop()
{
}
