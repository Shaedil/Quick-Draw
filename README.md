# Quick-Draw
[![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)]()

*Have yourself a high noon standoff right in your own home*

**To be submitted as part of "YeeHaw Hacks" hackathon by MLH**

## What is this?
We hacked a nerf gun to "shoot" information to a web server, containing the gun's orientation in 3D space, and the timestamp at which the trigger was pulled. 

If we have two of these guns... well, we have ourselves a showdown!

### TODO:
- [ ] MPU9250 sensor (for one gun) works and outputs orientation of sensor
- [ ] Check if 4.8V AA battery powers Arduino Nano AND other components (for one gun)
- [ ] Check if pushbutton works and lights up an LED
- [ ] Check if WiFi module works by uploading onboard-LED blink sketch
- [ ] Check if WiFi module works by sending request to server
- [ ] Set up PHP server via youtube video
- [ ] WiFi module sends a request to a server and persistently logs it via PHP

## How?
### Hardware Wise
A standard nerf gun has space inside for more than just nerf bullets, it has space for the smallest electronics: an Arduino Nano, 4 AA batteries, and some sensors. Utilizing an ESP8266 WiFi module for the Arduino (the Arduino does not have WiFi built in) and a MPU9250 9-axial gyroscope/accelerometer/magnetometer, we've built a nerf gun that can send requests over the internet to our specific server.

### Server Wise
Our server grabs these 'requests' from two nerf guns and compares the orientation of the gun first and then the timestamp to see who shot first and survived. To prevent any spam trigger pulling, we fitted the orientation sensor so that if the gun is pointed downwards, the shot 'misses' and you may 'die' if the other gun shoots you before you shoot again. When you get shot, the server sends a byte to the Arduino telling it to light up an LED indicating to the shooter you've died.

### Troubleshooting/Limitations
A big limitation was the entire size of the electronics. The nerf gun is small. It's about the length of an adult hand. We could not fit an Arduino Uno into the gun, but we could fit an Arduino Nano inside. The only problem with that is the Arduino Nano doesn't have as many ports as the former which worried us a bit. We overcame this by joining wires together via intersection and cutting wires down to length so that it doesn't clutter the space inside.

Another limitation was our power source. We've had a bit of trouble deciding on our portable source of power as this was fairly new to us. We eventually decided on a 4.8V AA battery pack, regularly used for an RC car. We hoped this was enough for all the sensors to be powered for at least an hour at 5V.

## Why?
I suppose there really is no _concrete_ reason why we did this, we just thought it was really cool to do, especially so because the hackathon was "Wild West" themed and we were inspired by Clint Eastwood films like the Dollars Trilogies, "The Good, The Bad & The Ugly" and "Once Upon A Time in the West". Yippee Ki-Yay!
