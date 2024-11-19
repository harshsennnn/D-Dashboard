package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/mem"
    "github.com/shirou/gopsutil/v3/disk"
    "github.com/shirou/gopsutil/v3/net"
    "github.com/rs/cors"
)

func getSystemMetrics(w http.ResponseWriter, r *http.Request) {
    // Get CPU usage
    cpuPercent, err := cpu.Percent(0, false)
    if err != nil {
        http.Error(w, fmt.Sprintf("Unable to fetch CPU data: %v", err), http.StatusInternalServerError)
        return
    }

    // Get Memory usage
    memStats, err := mem.VirtualMemory()
    if err != nil {
        http.Error(w, fmt.Sprintf("Unable to fetch memory data: %v", err), http.StatusInternalServerError)
        return
    }

    // Get Disk usage
    diskStats, err := disk.Usage("/")
    if err != nil {
        http.Error(w, fmt.Sprintf("Unable to fetch disk usage: %v", err), http.StatusInternalServerError)
        return
    }

    // Get Network stats
    netStats, err := net.IOCounters(false)
    if err != nil {
        http.Error(w, fmt.Sprintf("Unable to fetch network status: %v", err), http.StatusInternalServerError)
        return
    }

    // Send JSON response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `{
        "cpu_usage": %v,
        "memory_usage": %v,
        "disk_usage": %v,
        "network_usage": %v
    }`, cpuPercent, memStats.UsedPercent, diskStats.UsedPercent, netStats[0].BytesSent)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/api/stats", getSystemMetrics).Methods("GET")

    // Enable CORS
    handler := cors.Default().Handler(r)
    log.Fatal(http.ListenAndServe(":8080", handler))
}
