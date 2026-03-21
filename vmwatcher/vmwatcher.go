package vmwatcher

import "fmt"

//VM Vars, eg. ram allocated, cpu cores allocated, etc.
var ramAlloc int = 4 
var cpuCoresAlloc int = 2

//all of these in the future will need to be adopted so the functions need a vm id to check the status of a specific vm, but for now im just doing a simple mockup to get the basic structure down and work on other more important stuff for now

func VMRunningStatus() {
	// Placeholder for VM watching logic
	var vmRunning bool = false // placeholder until i implement actual monitoring logic
	fmt.Printf("Vm running? Status: %v", vmRunning)
}

func VMResourceAllocation() {
	// Placeholder for VM resource allocation logic
	fmt.Printf("Watching VM resource allocation - RAM allocated: %v GB, CPU cores allocated: %v", ramAlloc, cpuCoresAlloc)
}
