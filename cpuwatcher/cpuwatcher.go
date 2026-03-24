package cpuwatcher

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

var cpuPollingRate time.Duration = 1 * time.Second

func WatchCPU() {

	for {
		cpuPercent, err := cpu.Percent(cpuPollingRate, false)
		if err != nil {
			fmt.Println("Error fetching CPU info:", err)
			return
		}
		fmt.Printf("Watching CPU usage - Current usage: %.2f%% - polling rate: %v seconds\n", cpuPercent[0], cpuPollingRate)
	}
}

func CPUSetPollingRate(rate float64) {
	// In a real implementation, you would validate the input and update the polling rate accordingly.
	cpuPollingRate = time.Duration(rate * float64(time.Second))
	fmt.Printf("Setting CPU watcher polling rate to %v seconds\n", rate)
}
