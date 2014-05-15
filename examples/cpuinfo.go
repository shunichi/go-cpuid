package main

import "fmt"
import "go/build"
import "github.com/shunichi/cpuid"

func main() {
	context := build.Default
	fmt.Printf("GOARCH = %s\n", context.GOARCH)
	fmt.Printf("GOOS   = %s\n", context.GOOS)
	fmt.Printf("GOPATH = %s\n", context.GOPATH)
	fmt.Printf("GOROOT = %s\n", context.GOROOT)

	fmt.Printf("Intel = %d\n", cpuid.Intel)
	fmt.Printf("AMD   = %d\n", cpuid.AMD)
	fmt.Printf("Other = %d\n", cpuid.Other)

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
