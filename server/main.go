package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"sync"
)

const (
	packetMagic       = 0xBB
	packetTypeHello   = 0x0
	packetTypeReadyUp = 0x1
	packetTypeFired   = 0x2
)

var (
	state struct {
		sync.Mutex
		connectionCount int
		firedTimes      *int
	}
)

type packetHeader struct {
	Magic      byte
	PacketType byte
}

func main() {
	ln, err := net.Listen("tcp", listenAddress)
	if err != nil {
		panic(err)
	}

	fmt.Printf("up at %s\n", listenAddress)

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}
}

func read(conn net.Conn, len uint64) []byte {
	x := make([]byte, len)

	_, err := conn.Read(x)
	if err != nil {
		panic(err)
	}

	return x
}

func writeErrorToConn(conn net.Conn, err string) {
	conn.Write([]byte(fmt.Sprintf("error: %s\n", err)))
	conn.Close()
}

func readPacketHeader(conn net.Conn) packetHeader {
	var header packetHeader
	binary.Read(conn, binary.LittleEndian, &header)
	return header
}

func writePacketHeader(conn net.Conn, header packetHeader) {
	binary.Write(conn, binary.LittleEndian, header)
}

func handleConnection(conn net.Conn) {
	state.Lock()
	if state.connectionCount == 2 {
		writeErrorToConn(conn, "too many players are on this server! please try again later")
		return
	}
	state.Unlock()

	for {
		header := readPacketHeader(conn)
		if header.Magic != packetMagic {
			writeErrorToConn(conn, "magic is not 0xBB")
			return
		}

		switch header.PacketType {
		case packetTypeHello:
			state.Lock()
			state.connectionCount++
			state.Unlock()

			count := 0

			for count != 2 {
				state.Lock()
				count = state.connectionCount
				state.Unlock()
			}

			writePacketHeader(conn, packetHeader{
				Magic:      packetMagic,
				PacketType: packetTypeReadyUp,
			})

		case packetTypeFired:
			state.Lock()
			//
			state.Unlock()

		default:
			writeErrorToConn(conn, fmt.Sprintf("unknown packet type: 0x%X", header.PacketType))
			return
		}
	}
}
