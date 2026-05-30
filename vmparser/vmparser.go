package vmparser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type VMInfo struct {
	Name      string `json:"name"`
	UUID      string `json:"uuid"`
	CreatedAt string `json:"created_at"`
}

type CPU struct {
	Cores int    `json:"cores"`
	Type  string `json:"type"`
}

type Memory struct {
	SizeMB int `json:"size_mb"`
}

type Network struct {
	Interface string `json:"interface"`
	Type      string `json:"type"`
	MAC       string `json:"mac"`
}

type Disk struct {
	Name   string  `json:"name"`
	Path   string  `json:"path"`
	SizeGB float64 `json:"size_gb"`
	Format string  `json:"format"`
}

type VM struct {
	VM      VMInfo  `json:"vm"`
	CPU     CPU     `json:"cpu"`
	Memory  Memory  `json:"memory"`
	Network Network `json:"network"`
	Disk    Disk    `json:"disk"`
}

// func VMParser(vmname string) (*VM, error) {
//	path := filepath.Join("vms", fmt.Sprintf("%s.json", vmname))
//	data, err := os.ReadFile(path)
//	if err != nil {
//		return nil, fmt.Errorf("read vm file %s: %w", path, err)
//	}
//	var vm VM
//	if err := json.Unmarshal(data, &vm); err != nil {
//		return nil, fmt.Errorf("unmarshal vm json: %w", err)
//	}
//	return &vm, nil
// }

func VMParserVMInfo(vmname string) (*VMInfo, error) {
	path := filepath.Join("vms", fmt.Sprintf("%s.json", vmname))
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read vm file %s: %w", path, err)
	}
	var vm VM
	if err := json.Unmarshal(data, &vm); err != nil {
		return nil, fmt.Errorf("unmarshal vm json: %w", err)
	}
	return &vm.VM, nil
}

func VMParserCPU(vmname string) (*CPU, error) {
	path := filepath.Join("vms", fmt.Sprintf("%s.json", vmname))
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read vm file %s: %w", path, err)
	}
	var vm VM
	if err := json.Unmarshal(data, &vm); err != nil {
		return nil, fmt.Errorf("unmarshal cpu json: %w", err)
	}
	return &vm.CPU, nil
}

func VMParserMemory(vmname string) (*Memory, error) {
	path := filepath.Join("vms", fmt.Sprintf("%s.json", vmname))
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read vm file %s: %w", path, err)
	}
	var vm VM
	if err := json.Unmarshal(data, &vm); err != nil {
		return nil, fmt.Errorf("unmarshal memory json: %w", err)
	}
	return &vm.Memory, nil
}

func VMParserNetwork(vmname string) (*Network, error) {
	path := filepath.Join("vms", fmt.Sprintf("%s.json", vmname))
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read vm file %s: %w", path, err)
	}
	var vm VM
	if err := json.Unmarshal(data, &vm); err != nil {
		return nil, fmt.Errorf("unmarshal network json: %w", err)
	}
	return &vm.Network, nil
}

func VMParserDisk(vmname string) (*Disk, error) {
	path := filepath.Join("vms", fmt.Sprintf("%s.json", vmname))
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read vm file %s: %w", path, err)
	}
	var vm VM
	if err := json.Unmarshal(data, &vm); err != nil {
		return nil, fmt.Errorf("unmarshal disk json: %w", err)
	}
	return &vm.Disk, nil
}

//func VMParserWriter(vm VM) error {
//	content, err := json.MarshalIndent(vm, "", "  ")
//	if err != nil {
//		return fmt.Errorf("marshal vm json: %w", err)
//	}
//	path := filepath.Join("vms", fmt.Sprintf("%s.json", vm.VM.Name))
//	if err := os.WriteFile(path, content, 0644); err != nil {
//		return fmt.Errorf("write vm file %s: %w", path, err)
//	}
//	return nil
// }

func VMParserWriterInfo(info VMInfo) error {
	path := filepath.Join("vms", fmt.Sprintf("%s.json", info.Name))
	var vm VM
	// try to read existing file to preserve other fields
	if data, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(data, &vm)
	}
	vm.VM = info
	content, err := json.MarshalIndent(vm, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal vm json: %w", err)
	}
	if err := os.WriteFile(path, content, 0644); err != nil {
		return fmt.Errorf("write vm file %s: %w", path, err)
	}
	return nil
}

func VMParserWriterCPU(vmname string, cpu CPU) error {
	path := filepath.Join("vms", fmt.Sprintf("%s.json", vmname))
	var vm VM
	if data, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(data, &vm)
	}
	vm.CPU = cpu
	content, err := json.MarshalIndent(vm, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal vm json: %w", err)
	}
	if err := os.WriteFile(path, content, 0644); err != nil {
		return fmt.Errorf("write vm file %s: %w", path, err)
	}
	return nil
}

func VMParserWriterMemory(vmname string, memory Memory) error {
	path := filepath.Join("vms", fmt.Sprintf("%s.json", vmname))
	var vm VM
	if data, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(data, &vm)
	}
	vm.Memory = memory
	content, err := json.MarshalIndent(vm, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal vm json: %w", err)
	}
	if err := os.WriteFile(path, content, 0644); err != nil {
		return fmt.Errorf("write vm file %s: %w", path, err)
	}
	return nil
}

func VMParserWriterDisk(vmname string, disk Disk) error {
	path := filepath.Join("vms", fmt.Sprintf("%s.json", vmname))
	var vm VM
	if data, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(data, &vm)
	}
	vm.Disk = disk
	content, err := json.MarshalIndent(vm, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal vm json: %w", err)
	}
	if err := os.WriteFile(path, content, 0644); err != nil {
		return fmt.Errorf("write vm file %s: %w", path, err)
	}
	return nil
}

func VMParserWriterNetwork(vmname string, network Network) error {
	path := filepath.Join("vms", fmt.Sprintf("%s.json", vmname))
	var vm VM
	if data, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(data, &vm)
	}
	vm.Network = network
	content, err := json.MarshalIndent(vm, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal vm json: %w", err)
	}
	if err := os.WriteFile(path, content, 0644); err != nil {
		return fmt.Errorf("write vm file %s: %w", path, err)
	}
	return nil
}
