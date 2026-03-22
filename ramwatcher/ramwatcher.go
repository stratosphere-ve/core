package cpuwatcher

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/mem"
)

var ramPollingRate time.Duration = 1 * time.Second


func WatchRAM() {

	for {
		vmStat, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("Error fetching RAM info:", err)
			return
		}
		var totalRAMGB float64 = float64(vmStat.Total) / (1024 * 1024 * 1024)
		var usedRAMGB float64 = float64(vmStat.Used) / (1024 * 1024 * 1024)
		var usedRAMPercent float64 = vmStat.UsedPercent

		fmt.Printf("Watching RAM usage - Total RAM: %.2f GB, Used RAM: %.2f GB, Used RAM Percent: %.2f%%\n",
			totalRAMGB, usedRAMGB, usedRAMPercent)

		time.Sleep(ramPollingRate)
	}
}

func SetPollingRate(rate float64) {
	// In a real implementation, you would validate the input and update the polling rate accordingly.
	ramPollingRate = time.Duration(rate * float64(time.Second))
	fmt.Printf("Setting RAM watcher polling rate to %v seconds\n", rate)
}
