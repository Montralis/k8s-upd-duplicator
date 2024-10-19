package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "os"
)

func main() {
    // Die Datei, aus der die Daten gelesen werden
    filePath := "data.txt"

    // Zieladresse für den UDP-Stream (IP und Port)
    serverAddr := "127.0.0.1:9999"

    // Adresse parsen
    addr, err := net.ResolveUDPAddr("udp", serverAddr)
    if err != nil {
        log.Fatalf("Fehler beim Parsen der Adresse: %v", err)
    }

    // UDP-Verbindung erstellen
    conn, err := net.DialUDP("udp", nil, addr)
    if err != nil {
        log.Fatalf("Fehler beim Herstellen der UDP-Verbindung: %v", err)
    }
    defer conn.Close()

    for {
        // Datei öffnen
        file, err := os.Open(filePath)
        if err != nil {
            log.Fatalf("Fehler beim Öffnen der Datei: %v", err)
        }

        // Datei zeilenweise lesen
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            line := scanner.Text()

            // Daten über den UDP-Stream senden
            _, err := conn.Write([]byte(line))
            if err != nil {
                log.Printf("Fehler beim Senden der Daten: %v", err)
                continue
            }

            fmt.Printf("Gesendet: %s\n", line)
        }

        if err := scanner.Err(); err != nil {
            log.Fatalf("Fehler beim Lesen der Datei: %v", err)
        }

        // Datei schließen, bevor wir sie erneut öffnen
        file.Close()

        // Hinweis für den Benutzer, dass die Datei wiederholt wird
        fmt.Println("Dateiende erreicht, wiederhole das Lesen von vorne.")
    }
}
