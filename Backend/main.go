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
