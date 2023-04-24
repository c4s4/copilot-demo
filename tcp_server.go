package main

import (
    "fmt"
    "net"
)

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Listening on port 8080...")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }

        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn) {
    buffer := make([]byte, 1024)
    _, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(string(buffer))

    conn.Close()
}
