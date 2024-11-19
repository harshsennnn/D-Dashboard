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

 
