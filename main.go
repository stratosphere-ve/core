package main

import (
	"fmt"
	"time"
	"example.com/m/v2/cpuwatcher"
)

func main() {
	fmt.Println(`Welcome to Stratosphere!
This is a simple CLI mockup of how it will work in the future.`)
	cpuwatcher.WatchCPU()
	fmt.Println(`
What would you like to do next? 
1. Change CpuWatcher Polling rate
2. Exit`)

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
	default:
		fmt.Println("Invalid choice. Please enter 1 or 2.")
	}

}
