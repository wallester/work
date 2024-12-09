package runtime

import (
	"fmt"
	"runtime"
)

type MemoryUsageMessage struct {
	Alloc      string
	TotalAlloc string
	Sys        string
	NumGC      string
}

func NewMemoryUsageMessage() *MemoryUsageMessage {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return &MemoryUsageMessage{
		Alloc:      fmt.Sprintf("Alloc = %v MiB", m.Alloc/1024/1024),
		TotalAlloc: fmt.Sprintf("TotalAlloc = %v MiB", m.TotalAlloc/1024/1024),
		Sys:        fmt.Sprintf("Sys = %v MiB", m.Sys/1024/1024),
		NumGC:      fmt.Sprintf("NumGC = %v MiB", m.NumGC),
	}
}
