package main

import (
	"fmt"
	"time"

	"example.com/m/v2/cpuwatcher"
	ramwatcher "example.com/m/v2/ramwatcher"
	"example.com/m/v2/vmwatcher"
)

func main() {
	fmt.Println(`Welcome to Stratosphere!
This is a simple CLI mockup of how it will work in the future.`)
	fmt.Println(`
What would you like to do next? 

1. Change CpuWatcher Polling rate
2. Exit
3. View VM Resource Allocation
4. View VM Running Status
5. View both VM Resource Allocation and Running Status
6. View RAM usage
7. Change RAM Watcher Polling rate
8. View CPU usage`)

	var intChoice int
	fmt.Scanln(&intChoice)

	switch intChoice {
	case 1:
		fmt.Println("Changing CpuWatcher Polling rate... (2 secs)")
		cpuwatcher.SetPollingRate(2)
		time.Sleep(3 * time.Second) // Simulate waiting for the new polling rate to take effect
		cpuwatcher.WatchCPU()

	case 2:
		fmt.Println("Exiting...")
		return

	case 3:
		vmwatcher.VMResourceAllocation()
	case 4:
		vmwatcher.VMRunningStatus()
	case 5:
		vmwatcher.VMResourceAllocation()
		fmt.Println() // Add a newline for better readability
		vmwatcher.VMRunningStatus()
	case 6:
		fmt.Println("Checking system RAM usage...")
		ramwatcher.WatchRAM()
	case 7:
		fmt.Println("Changing RAM Watcher Polling rate... (2 secs)")
		ramwatcher.SetPollingRate(2)
		time.Sleep(3 * time.Second)
	case 8:
		cpuwatcher.WatchCPU()
	default:
		fmt.Println("Invalid choice. Please enter 1 or 2.")
	}

}
