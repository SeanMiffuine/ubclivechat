package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "io"
)

func main() {
    go startServer()

    serviceHost := os.Getenv("SERVICE_HOST")
    if serviceHost == "" {
        // Default to localhost if not set
        serviceHost = "localhost"
    }
    // Connect to the peer
    fmt.Println("connect to peer: " + serviceHost)
    connectToPeer(serviceHost + ":8080")

    // Send messages
    for {
        sendMessages()
    }
}

func startServer() {
    // Listen for incoming connections.
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        // Handle the error properly; for now, we'll just log it and exit
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    defer ln.Close()
    for {
        conn, err := ln.Accept()
        if err != nil {
            // Handle the error properly; for now, we'll log it
            fmt.Println("Error accepting: ", err.Error())
            // Optionally continue to accept other connections even if one fails
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    reader := bufio.NewReader(conn)
    for {
        message, err := reader.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                // End of file (or stream), which means the connection was closed
                fmt.Println("Connection closed by peer")
            } else {
                // An actual error occurred
                fmt.Println("Error reading:", err.Error())
            }
            break // Exit the loop and end the goroutine
        }
        if message != "" {
            fmt.Print("Message received: ", message)
        }
    }
}

func connectToPeer(address string) {
    conn, err := net.Dial("tcp", address)
    if err != nil {
        fmt.Printf("Failed to connect to peer at %s: %v\n", address, err)
        os.Exit(1) // or handle the error in a more sophisticated way
    }
    defer conn.Close()

    // Proceed with using `conn`
}

func sendMessages() {
    // Send messages to the peer
    // Similar to connectToPeer but might want to maintain the 
}
