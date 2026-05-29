package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
)

type IOCResult struct {
    IOC        string    `json:"ioc"`
    RiskScore  int       `json:"risk_score"`
    Timestamp  time.Time `json:"timestamp"`
}

func main() {
    http.HandleFunc("/api/realtime", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/event-stream")
        w.Header().Set("Cache-Control", "no-cache")
        w.Header().Set("Connection", "keep-alive")

        for {
            resp, _ := http.Get("http://localhost:5000/api/recent")
            if resp != nil {
                body, _ := io.ReadAll(resp.Body)
                fmt.Fprintf(w, "data: %s\n\n", body)
                w.(http.Flusher).Flush()
                resp.Body.Close()
            }
            time.Sleep(5 * time.Second)
        }
    })

    fmt.Println("Go service running on :8080")
    http.ListenAndServe(":8080", nil)
}
