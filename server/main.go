package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"sync"
)

const (
	packetTypeHello = 0x0
)

var (
	state struct {
		sync.Mutex
		connectionCount uint8
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
	writePacketHeader(conn, packetHeader{
		Magic:      0xBB,
		PacketType: packetTypeHello,
	})

	for {
		header := readPacketHeader(conn)

		switch header.PacketType {
		case packetTypeHello:
			state.Lock()
			state.connectionCount++
			state.Unlock()

		default:
			writeErrorToConn(conn, fmt.Sprintf("unknown packet type: 0x%X", header.PacketType))
			return
		}
	}
}
