package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	// Read sample data from a file
	filePath := "data.txt"

	// Send data to localhost
	serverAddr := "127.0.0.1:9199"

	// Parse the address
	addr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		log.Fatalf("Error parsing the address: %v", err)
	}

	// Create a UDP connection
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("Error establishing the UDP connection: %v", err)
	}
	defer conn.Close()

	for {
		// Open the file
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("Error opening the file: %v", err)
		}

		// Read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			// Send data over the UDP stream
			_, err := conn.Write([]byte(line))
			if err != nil {
				log.Printf("Error sending data: %v", err)
				continue
			}

			fmt.Printf("Sent: %s\n", line)

			// Sleep for 500 milliseconds between lines
			time.Sleep(500 * time.Millisecond)
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading the file: %v", err)
		}

		// Close the file before reopening it for the next loop
		file.Close()

		// Inform the user that the file will be read again
		fmt.Println("File reached its end. Reading again ...")
	}
}
