package main

import "fmt"
import "github.com/shunichi/cpuid"

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
	fmt.Printf("cpu brand name    = %s\n", cpuid.BrandName())

	fmt.Printf("max func          = %08x\n", cpuid.MaxFunctionId())
	fmt.Printf("max extended func = %08x\n", cpuid.MaxExtendedFunctionId())

	fmt.Printf("cpu physical core = %d\n", cpuid.PhysicalCore())
	fmt.Printf("cpu logical core  = %d\n", cpuid.LogicalCore())
}
