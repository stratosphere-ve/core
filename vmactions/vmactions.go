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

if strings.ContainsAny(vmname, "/\\.") {
	fmt.Println("Invalid VM name. Contains (one or more) invalid characters: /, \\, .")
	return
} else {
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

	fmt.Println("")
	fmt.Println("Successfully created VM!")}
}

func VMDelete() {
	vmname := ""
	fmt.Println("What VM would you like to delete?")
	fmt.Scanln(&vmname)
	err := os.Remove(fmt.Sprintf("vms/%v.json", vmname))
	if err != nil {
		fmt.Println("Error deleting VM:", err)
		return
	}

	fmt.Printf("Successfully deleted VM: %v", vmname)
}






func VMManage() {
option := 0
	fmt.Println(`What would you like to do with your VM?
1. Start
2. Stop
3. Restart
4. Rename

50+. Exit`)
fmt.Scanln(&option)

switch option {

case 1:

case 2:

case 3:

case 4:
VMRename()

case 50:
	fmt.Println("Exiting VM management...")
	return

default:
	println("Invalid option, please choose a valid option.")
}
}
// Everything function below this will be used in the "VMManage" function for things such as starting, stopping, restarting, renaming, etc.

func VMRename() {
	vmname := ""
	newname := ""
	fmt.Println("What VM would you like to rename?")
	fmt.Scanln(&vmname)
	fmt.Println("What would you like to rename it to?")
	fmt.Scanln(&newname)

if strings.ContainsAny (newname, "/\\.") {
	fmt.Println("Invalid VM name. VM names cannot contain these characters: /, \\, .")
	return
} else {
	err := os.Rename(fmt.Sprintf("vms/%v.json", vmname), fmt.Sprintf("vms/%v.json", newname))
	if err != nil {
		fmt.Println("Error renaming VM:", err)
		return
	}

	fmt.Printf("Successfully renamed VM from %v to %v", vmname, newname)
}
}
