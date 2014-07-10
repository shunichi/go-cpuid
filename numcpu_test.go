package cpuid

import (
	"runtime"
	"testing"
)

func TestNumCpu(t *testing.T) {
	actual := LogicalCores()
	expected := runtime.NumCPU()
	if actual != expected {
		t.Errorf("LogicalCores() expected to return %d but %d", expected, actual)
	}
}
