package storagewatcher

import (
	"fmt"
	"time"
	"strings"
	"github.com/shirou/gopsutil/v3/disk"
)

var storagePollingRate time.Duration = 1 * time.Second


func WatchStorageUsage() {
	prev, err := disk.IOCounters()
	if err != nil {
		fmt.Println("Error fetching disk IO counters:", err)
		return
	}

	for {
		time.Sleep(storagePollingRate)

		curr, err := disk.IOCounters()
		if err != nil {
			fmt.Println("Error fetching disk IO counters:", err)
			return
		}

		for name, now := range curr {

			// skip loop devices, cdroms, partitions, etc.
			if strings.HasPrefix(name, "loop") ||
				strings.HasPrefix(name, "sr") ||
				strings.HasSuffix(name, "1") ||
				strings.HasSuffix(name, "14") ||
				strings.HasSuffix(name, "15") {
				continue
			}

			old := prev[name]

			readBytes := now.ReadBytes - old.ReadBytes
			writeBytes := now.WriteBytes - old.WriteBytes

			readMB := float64(readBytes) / (1024 * 1024)
			writeMB := float64(writeBytes) / (1024 * 1024)

			fmt.Printf(
				"Disk %s - Read: %.2f MB/s | Write: %.2f MB/s\n",
				name,
				readMB,
				writeMB,
			)
		}

		prev = curr
	}
}

func WatchStorageSpace() {

	for {
		vmStat, err := disk.Usage("/")
		if err != nil {
			fmt.Println("Error fetching storage info:", err)
			return
		}
		var totalStorageGB float64 = float64(vmStat.Total) / (1024 * 1024 * 1024)
		var usedStorageGB float64 = float64(vmStat.Used) / (1024 * 1024 * 1024)
		var usedStoragePercent float64 = vmStat.UsedPercent

		fmt.Printf("Watching storage usage - Total Storage: %.2f GB, Used Storage: %.2f GB, Used Storage Percent: %.2f%%\n",
			totalStorageGB, usedStorageGB, usedStoragePercent)

		time.Sleep(storagePollingRate)
	}
}

func StorageSetPollingRate(rate float64) {
	// In a real implementation, you would validate the input and update the polling rate accordingly.
	storagePollingRate = time.Duration(rate * float64(time.Second))
	fmt.Printf("Setting storage watcher polling rate to %v seconds\n", rate)
}
