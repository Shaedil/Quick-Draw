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
