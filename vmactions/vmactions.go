package vmactions

import (
	"fmt"
	"os"
	"strings"
)

func VMCreate() {
	os.MkdirAll("vms", os.ModePerm)

	vmname := ""

	files, err := os.ReadDir("vms")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Enter VM name: ")
	fmt.Scanln(&vmname)

	for _, f := range files {
		name := strings.TrimSuffix(f.Name(), ".json")
		if name == vmname {
			fmt.Printf("A VM with that name (%v) already exists. Please choose a different name.", vmname)
			return
		}
	}

	fmt.Printf("Creating VM with name: %v", vmname)

	file, err := os.Create(fmt.Sprintf("vms/%v.json", vmname))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("Successfully created VM!")
}
