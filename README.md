<div align="center">
	<h1>Quick-Draw</h1>
	<blockquote>
		<p><i>Have yourself a high noon standoff right in your own home</i></p>
		<p><b>To be submitted as part of "YeeHaw Hacks" hackathon by MLH</b></p>
	<img src="https://img.shields.io/badge/Built%20w%2F-Arduino%20C%20%26%20Go-blue?style=for-the-badge">
	<img src="https://img.shields.io/badge/License-MIT-orange?style=for-the-badge">
	</blockquote>
</div>

## Who Lives, Who Dies, Who Tells Your Story

Now, who doesn't love them a good ol' fashioned duel!
Sure is a shame they're outlawed... unless?

### Execution

```
+---------------------------+
|Server waits for connection|    Establish TCP connection with server
+-+-------------------------+
  |
  |
  v
+-+---------------------+
|Client-Server handshake|        Complete handshake to check
|   (for each client)   |        if communication is working
+-+---------------------+
  |
  |
  v
+-+--------------+
|Countdown begins|               The duel will commence after we count to ten
+-+--------------+
  |
  |
  v
+-+-----------+
|Clients shoot|                  Server recieves packets and determines winner
+-+-----------+
  |
  |
  v
+-+----------------------+
|Server broadcasts winner|       Send client unique packet with win/loss
+------------------------+
```

### Do It Yourself

#### Hardware

* NodeMCU
* Push Button
* 10K ohm resistor
* RGB LED
* Breadboard
* Wires
* Toy gun (optional)

#### Software

* Go
* Arduino CLI
* Make
* BASH (for automating installing NodeMCU cores)

---

```bash
# Building & running the server
###############################

cd server
go build

# Make sure to port forward port 1337
# before running the server

./qd-server


# Compiling and flashing the NodeMCU
####################################

cd client
chmod +x prepare.sh
./prepare.sh

# Edit NodeMCU/connection.h and set the
# variables to meet your connection needs

# Set PORT to the port of your NodeMCU
make PORT=/dev/ttyUSBX flash
```

<div align="center">
	<img src="https://i.imgur.com/oKokoPz.png">
</div>
