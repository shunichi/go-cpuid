package main

import "fmt"
import "github.com/shunichi/go-cpuid"

func main() {
	fmt.Print("cpu vendor        = ")
	switch cpuid.VendorId() {
	case cpuid.Intel:
		fmt.Println("Intel")
	case cpuid.AMD:
		fmt.Println("AMD")
	default:
		fmt.Println("Unknown")
	}
	fmt.Printf("cpu brand name       = %s\n", cpuid.BrandName())

	fmt.Printf("max func id          = %08x\n", cpuid.MaxFunctionId())
	fmt.Printf("max extended func id = %08x\n", cpuid.MaxExtendedFunctionId())

	fmt.Printf("cpu physical cores   = %d\n", cpuid.PhysicalCores())
	fmt.Printf("cpu logical cores    = %d\n", cpuid.LogicalCores())
	fmt.Printf("threads / core       = %d\n", cpuid.ThreadsPerCore())
}
