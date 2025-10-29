package main

import (
    "bufio"
    "fmt"
    "log"
    "net/rpc"
    "os"
    "strings"
    "time"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        client, err := rpc.Dial("tcp", "localhost:1234")
        if err != nil {
            fmt.Println(" Could not connect to server. Retrying in 2s...")
            time.Sleep(2 * time.Second)
            continue
        }

        fmt.Println(" Connected! Type your message (or 'exit' to quit):")

        for {
            fmt.Print("You: ")
            msg, _ := reader.ReadString('\n')
            msg = strings.TrimSpace(msg)

            if msg == "exit" {
                fmt.Println(" Exiting chat...")
                client.Close()
                return
            }

            var reply []string
            err = client.Call("ChatServer.SendMessage", msg, &reply)
            if err != nil {
                log.Println("Error sending message:", err)
                break
            }

            fmt.Println("\n Chat history:")
            for _, m := range reply {
                fmt.Println("•", m)
            }
            fmt.Println()
        }

        client.Close()
    }
}
