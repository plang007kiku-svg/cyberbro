package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "strings"
    "time"
)

type IOCRequest struct {
    Text string `json:"text"`
}

type IOCResponse struct {
    IOC        string   `json:"ioc"`
    Verdict    string   `json:"verdict"`
    RiskScore  int      `json:"risk_score"`
    Sources    []string `json:"sources"`
    Timestamp  string   `json:"timestamp"`
}

func analyzeHandler(w http.ResponseWriter, r *http.Request) {
    body, _ := io.ReadAll(r.Body)
    var req IOCRequest
    json.Unmarshal(body, &req)
    
    response := IOCResponse{
        IOC:       req.Text,
        Verdict:   "unknown",
        RiskScore: 0,
        Sources:   []string{},
        Timestamp: time.Now().Format(time.RFC3339),
    }
    
    // Simulate analysis
    if strings.Contains(req.Text, ".") {
        response.Verdict = "suspicious"
        response.RiskScore = 50
        response.Sources = append(response.Sources, "GoAnalyzer")
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status":"ok"}`))
}

func main() {
    http.HandleFunc("/analyze", analyzeHandler)
    http.HandleFunc("/health", healthHandler)
    
    fmt.Println("Go Cyberbro Service running on :8080")
    http.ListenAndServe(":8080", nil)
}
