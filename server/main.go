package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"sync"
)

const (
	packetMagic       = 0xBB
	packetTypeHello   = 0x0
	packetTypeReadyUp = 0x1
	packetTypeFired   = 0x2
	packetTypeResults = 0x3
)

var (
	state struct {
		sync.Mutex
		connectionCount int
		fired           int
	}
)

type packetHeader struct {
	Magic      byte
	PacketType byte
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "client" {
		conn, err := net.Dial("tcp", "127.0.0.1"+listenAddress)
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		writePacketHeader(conn, packetHeader{
			Magic:      packetMagic,
			PacketType: packetTypeHello,
		})

		header := readPacketHeader(conn)

		if header.Magic != packetMagic || header.PacketType != packetTypeHello {
			return
		}

		fmt.Printf("handshake successful\n")

		header = readPacketHeader(conn)

		if header.Magic != packetMagic || header.PacketType != packetTypeReadyUp {
			return
		}

		fmt.Printf("ready up uwu\npress enter to fire\n")

		bufio.NewScanner(os.Stdin).Scan()

		writePacketHeader(conn, packetHeader{
			Magic:      packetMagic,
			PacketType: packetTypeFired,
		})

		header = readPacketHeader(conn)

		if header.Magic != packetMagic || header.PacketType != packetTypeResults {
			return
		}

		result := read(conn, 1)

		if result[0] == 1 {
			fmt.Printf("you won!\n")
		} else {
			fmt.Printf("you lost!\n")
		}
	} else {
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
	fmt.Printf("got connection with address %s\n", conn.RemoteAddr().String())

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

			writePacketHeader(conn, packetHeader{
				Magic:      packetMagic,
				PacketType: packetTypeHello,
			})

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

			writePacketHeader(conn, packetHeader{
				Magic:      packetMagic,
				PacketType: packetTypeResults,
			})

			if state.fired == 0 {
				conn.Write([]byte{1})
				state.fired = 1
			} else {
				conn.Write([]byte{0})
			}

			state.Unlock()

		default:
			writeErrorToConn(conn, fmt.Sprintf("unknown packet type: 0x%X", header.PacketType))
			return
		}
	}
}
