package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// The address to receive data from (e.g., 127.0.0.1:9999)
	sourceAddr := ":9999"

	// The addresses to forward the received data to
	destAddr1 := ":8888"
	destAddr2 := ":7777"

	// Resolve the source and destination addresses
	srcAddr, err := net.ResolveUDPAddr("udp", sourceAddr)
	if err != nil {
		fmt.Println("Error resolving source address:", err)
		os.Exit(1)
	}

	dest1Addr, err := net.ResolveUDPAddr("udp", destAddr1)
	if err != nil {
		fmt.Println("Error resolving first destination address:", err)
		os.Exit(1)
	}

	dest2Addr, err := net.ResolveUDPAddr("udp", destAddr2)
	if err != nil {
		fmt.Println("Error resolving second destination address:", err)
		os.Exit(1)
	}

	// UDP connection for receiving data
	conn, err := net.ListenUDP("udp", srcAddr)
	if err != nil {
		fmt.Println("Error starting UDP server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Buffer for receiving data
	buf := make([]byte, 1024)

	fmt.Printf("Waiting for data on port %s...\n", sourceAddr)

	for {
		// Receive data from the source port
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			continue
		}

		// Forward the received data to both destination ports
		data := buf[:n]

		// Send to the first destination port (destAddr1)
		err = sendData(dest1Addr, data)
		if err != nil {
			fmt.Println("Error sending to", destAddr1, ":", err)
		} else {
			fmt.Printf("Sent to %v: %s\n", destAddr1, string(data))
		}

		// Send to the second destination port (destAddr2)
		err = sendData(dest2Addr, data)
		if err != nil {
			fmt.Println("Error sending to", destAddr2, ":", err)
		} else {
			fmt.Printf("Sent to %v: %s\n", destAddr2, string(data))
		}

		// Optional: Display the received data
		fmt.Printf("Received from %v: %s\n", remoteAddr, string(data))
	}
}

// Helper function to send data to the specified destination port
func sendData(addr *net.UDPAddr, data []byte) error {
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(data)
	return err
}
