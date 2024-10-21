package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Get the source port from the environment variable
	sourcePort := os.Getenv("SOURCE_PORT")
	if sourcePort == "" {
		fmt.Println("SOURCE_PORT not set in environment")
		os.Exit(1)
	}
	sourceAddr := ":" + sourcePort

	// Get the destination ports from the environment variable
	destPortsStr := os.Getenv("DESTINATION_PORTS")
	if destPortsStr == "" {
		fmt.Println("DESTINATION_PORTS not set in environment")
		os.Exit(1)
	}

	// Split the comma-separated destination ports (in format host:port)
	destPorts := strings.Split(destPortsStr, ",")

	fmt.Println(destPorts)
	// Resolve the source address for receiving data
	srcAddr, err := net.ResolveUDPAddr("udp", sourceAddr)
	if err != nil {
		fmt.Println("Error resolving source address:", err)
		os.Exit(1)
	}

	// Resolve the destination addresses
	var destAddrs []*net.UDPAddr
	for _, port := range destPorts {
		addr, err := net.ResolveUDPAddr("udp", port)
		if err != nil {
			fmt.Println("Error resolving destination address:", err)
			os.Exit(1)
		}
		destAddrs = append(destAddrs, addr)
	}

	// Create a UDP connection to listen on the source port
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

		// Data to forward
		data := buf[:n]

		// Forward the received data to all destination addresses
		for _, destAddr := range destAddrs {
			err := sendData(destAddr, data)
			if err != nil {
				fmt.Println("Error sending to", destAddr, ":", err)
			} else {
				fmt.Printf("Sent to %v: %s\n", destAddr, string(data))
			}
		}

		// Optional: Display the received data
		fmt.Printf("Received from %v: %s\n", remoteAddr, string(data))
	}
}

// Helper function to send data to the specified destination address
func sendData(addr *net.UDPAddr, data []byte) error {
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(data)
	return err
}
