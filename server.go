package main

import (
    "fmt"
    "log"
    "net"
    "net/rpc"
    "sync"
)

type ChatServer struct {
    messages []string
    mu       sync.Mutex
}

func (cs *ChatServer) SendMessage(message string, reply *[]string) error {
    cs.mu.Lock()
    defer cs.mu.Unlock()
    cs.messages = append(cs.messages, message)
    *reply = append([]string(nil), cs.messages...)
    return nil
}

func (cs *ChatServer) GetHistory(dummy string, reply *[]string) error {
    cs.mu.Lock()
    defer cs.mu.Unlock()
    *reply = append([]string(nil), cs.messages...)
    return nil
}

func main() {
    chat := new(ChatServer)
    rpc.Register(chat)

    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
    defer listener.Close()

    fmt.Println("💬 Chat server running on port 1234...")

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println("Connection error:", err)
            continue
        }
        go rpc.ServeConn(conn)
    }
}
