package vmactions

import (
	"fmt"
	mrand "math/rand"
	"os"
	"strings"
)

var vmname string

func VMCreate() {

	var uuid string
	var cpucores int
	var cputype string
	var memorysize int
	var networkinterface string
	var networktype string
	var networkmac string
	var diskname string
	var diskpath string
	var disksize float64
	var format string

	os.MkdirAll("vms", os.ModePerm)

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
	}

	for _, f := range files {
		name := strings.TrimSuffix(f.Name(), ".json")
		if name == vmname {
			fmt.Printf("A VM with that name (%v) already exists. Please choose a different name.\n", vmname)
			return
		}
	}

	randomNumber1 := mrand.Intn(10)
	letters1 := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randomNumber2 := mrand.Intn(10)
	letters2 := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randomNumber3 := mrand.Intn(10)
	letters3 := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randomNumber4 := mrand.Intn(10)
	letters4 := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randomNumber5 := mrand.Intn(10)
	letters5 := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randomNumber6 := mrand.Intn(10)
	letters6 := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	uuid = fmt.Sprintf("%d%c-%d%c-%d%c-%d%c-%d%c-%d%c",
		randomNumber1, letters1[mrand.Intn(len(letters1))],
		randomNumber2, letters2[mrand.Intn(len(letters2))],
		randomNumber3, letters3[mrand.Intn(len(letters3))],
		randomNumber4, letters4[mrand.Intn(len(letters4))],
		randomNumber5, letters5[mrand.Intn(len(letters5))],
		randomNumber6, letters6[mrand.Intn(len(letters6))])

	fmt.Println("Generated UUID:", uuid)

	fmt.Println("How many CPU cores would you like to allocate? ")
	fmt.Scanln(&cpucores)
	fmt.Println("Enter CPU type: ")
	fmt.Scanln(&cputype)
	fmt.Println("How much memory (in MB) would you like to allocate? ")
	fmt.Scanln(&memorysize)
	fmt.Println("Enter network interface: ")
	fmt.Scanln(&networkinterface)
	fmt.Println("Enter network type: ")
	fmt.Scanln(&networktype)
	fmt.Println("Enter network MAC address: ")
	fmt.Scanln(&networkmac)
	fmt.Println("Enter disk name: ")
	fmt.Scanln(&diskname)
	fmt.Println("Enter disk path: ")
	fmt.Scanln(&diskpath)
	fmt.Println("Enter disk size (in GB): ")
	fmt.Scanln(&disksize)
	fmt.Println("Enter disk format: ")
	fmt.Scanln(&format)

	fmt.Printf("Creating VM with name: %v\n", vmname)

	file, err := os.Create(fmt.Sprintf("vms/%v.json", vmname))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	if err := file.Close(); err != nil {
		fmt.Println("Error closing file:", err)
		return
	}

	fmt.Println("Successfully created VM!")
	NewVMConfigure()
}

func VMDelete() {
	vmname := ""
	fmt.Println("What VM would you like to delete?")
	fmt.Scanln(&vmname)

	if strings.ContainsAny(vmname, "/\\.") {
		fmt.Println("Invalid VM name. Contains (one or more) invalid characters: /, \\, .")
		return
	}

	err := os.Remove(fmt.Sprintf("vms/%v.json", vmname))
	if err != nil {
		fmt.Println("Error deleting VM:", err)
		return
	}

	fmt.Printf("Successfully deleted VM: %v\n", vmname)
}

func VMManage() {
	option := 0
	fmt.Println(`What would you like to do with your VM?
1. Start
2. Stop
3. Restart
4. Rename

50. Exit`)
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

	if strings.ContainsAny(newname, "/\\.") {
		fmt.Println("Invalid VM name. VM names cannot contain these characters: /, \\, .")
		return
	}

	err := os.Rename(fmt.Sprintf("vms/%v.json", vmname), fmt.Sprintf("vms/%v.json", newname))
	if err != nil {
		fmt.Println("Error renaming VM:", err)
		return
	}

	fmt.Printf("Successfully renamed VM from %v to %v\n", vmname, newname)
}

func NewVMConfigure() {

	var uuid string
	var cpucores int
	var cputype string
	var memorysize int
	var networkinterface string
	var networktype string
	var networkmac string
	var diskname string
	var diskpath string
	var disksize float64
	var format string

	randomNumber1 := mrand.Intn(10)
	letters1 := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randomNumber2 := mrand.Intn(10)
	letters2 := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randomNumber3 := mrand.Intn(10)
	letters3 := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randomNumber4 := mrand.Intn(10)
	letters4 := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randomNumber5 := mrand.Intn(10)
	letters5 := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randomNumber6 := mrand.Intn(10)
	letters6 := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	uuid = fmt.Sprintf("%d%c-%d%c-%d%c-%d%c-%d%c-%d%c",
		randomNumber1, letters1[mrand.Intn(len(letters1))],
		randomNumber2, letters2[mrand.Intn(len(letters2))],
		randomNumber3, letters3[mrand.Intn(len(letters3))],
		randomNumber4, letters4[mrand.Intn(len(letters4))],
		randomNumber5, letters5[mrand.Intn(len(letters5))],
		randomNumber6, letters6[mrand.Intn(len(letters6))])

	fmt.Println("Generated UUID:", uuid)

	fmt.Println("How many CPU cores would you like to allocate? ")
	fmt.Scanln(&cpucores)
	fmt.Println("Enter CPU type: ")
	fmt.Scanln(&cputype)
	fmt.Println("How much memory (in MB) would you like to allocate? ")
	fmt.Scanln(&memorysize)
	fmt.Println("Enter network interface: ")
	fmt.Scanln(&networkinterface)
	fmt.Println("Enter network type: ")
	fmt.Scanln(&networktype)
	fmt.Println("Enter network MAC address: ")
	fmt.Scanln(&networkmac)
	fmt.Println("Enter disk name: ")
	fmt.Scanln(&diskname)
	fmt.Println("Enter disk path: ")
	fmt.Scanln(&diskpath)
	fmt.Println("Enter disk size (in GB): ")
	fmt.Scanln(&disksize)
	fmt.Println("Enter disk format: ")
	fmt.Scanln(&format)

	content := fmt.Sprintf(`{
  "vm": {
    "name": "%s",
    "uuid": "%s",
    "created_at": "1970-01-01T12:00:00Z"
  },
  "cpu": {
    "cores": %d,
    "type": "%s"
  },
  "memory": {
    "size_mb": %d
  },
  "network": {
    "interface": "%s",
    "type": "%s",
    "mac": "%s"
  },
  "disk": {
    "name": "%s",
    "path": "%s",
    "size_gb": %f,
    "format": "%s"
  }
}`, vmname, uuid, cpucores, cputype, memorysize, networkinterface, networktype, networkmac, diskname, diskpath, disksize, format)

	err := os.WriteFile(fmt.Sprintf("vms/%s.json", vmname), []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}
