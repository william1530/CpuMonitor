package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/gorilla/websocket"
    "github.com/shirou/gopsutil/cpu"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer conn.Close()

        for {
            cpuPercent, _ := cpu.Percent(time.Second, false)
            if err := conn.WriteJSON(cpuPercent[0]); err != nil {
                fmt.Println(err)
                return
            }
            time.Sleep(time.Second)
        }
    })

    http.ListenAndServe(":8080", nil)
}