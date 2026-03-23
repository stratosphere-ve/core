package main

import (
	"fmt"
	"time"

	"example.com/m/v2/cpuwatcher"
	ramwatcher "example.com/m/v2/ramwatcher"
	"example.com/m/v2/vmwatcher"
)

func main() {
	fmt.Println(`Welcome to the Stratosphere CLI!
This is just a simple mockup/placeholder of how everything will work in the future

What would you like to do next?`) 

fmt.Println(`-WATCHERS-
1. Change CpuWatcher Polling rate
2. View CPU usage

3. Change RAMWatcher Polling rate
4. View RAM usage

5. View VM Resource Allocation
6. View VM Status
7. View “ View VM Resource Allocation” + “View VM Status”


-OTHER-
50. Exit`)

	var intChoice int
	fmt.Scanln(&intChoice)

	switch intChoice {
	case 1:
		fmt.Println("Changing CpuWatcher Polling rate... (2 secs)")
		cpuwatcher.CpuSetPollingRate(2)
		time.Sleep(3 * time.Second) // Simulate waiting for the new polling rate to take effect
		cpuwatcher.WatchCPU()

	case 2:
		cpuwatcher.WatchCPU()

	case 3:
		fmt.Println("Changing RAM Watcher Polling rate... (2 secs)")
		ramwatcher.RamSetPollingRate(2)
		time.Sleep(3 * time.Second)

	case 4:
		fmt.Println("Checking system RAM usage...")
		ramwatcher.WatchRAM()

	case 5:
		vmwatcher.VMResourceAllocation()

	case 6:
		vmwatcher.VMRunningStatus()

	case 7:
		vmwatcher.VMResourceAllocation()
		fmt.Println() // Add a newline for better readability
		vmwatcher.VMRunningStatus()

	case 50:
		fmt.Println("Exiting...")
		return

	default:
		fmt.Println("Invalid choice. Please enter 1 or 2.")
	}
}
