package cpuwatcher

import "fmt"

var pollingRate float64 = 1.0 // Default polling rate in seconds

func WatchCPU() { 
	
	var cpuUsage float64

	cpuUsage = 75.5 // Example CPU usage value
	fmt.Printf("Watching CPU usage - Current usage: %v - polling rate: %v seconds", cpuUsage, pollingRate)
}

func SetPollingRate(rate float64) {
	// In a real implementation, you would validate the input and update the polling rate accordingly.
	pollingRate = rate
	fmt.Printf("Setting CPU watcher polling rate to %v seconds\n", rate)
}
