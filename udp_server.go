package main

import (
    "fmt"
    "net"
)

func main() {
    addr, err := net.ResolveUDPAddr("udp", ":8080")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer conn.Close()

    fmt.Println("Listening on port 8080...")

    buffer := make([]byte, 1024)

    for {
        n, addr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }

        fmt.Printf("Received message from %s: %s\n", addr, string(buffer[:n]))
    }
}
