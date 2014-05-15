package cpuid

// +build 386 amd64

/*
#include <string.h>

void cpuid( unsigned op, unsigned* regs )
{
	unsigned eax, ebx, ecx, edx;
	asm volatile(
		"pushl %%ebx   \n\t"
	    "cpuid         \n\t"
	    "movl %%ebx, %1\n\t"
	    "popl %%ebx    \n\t"
	    : "=a"(eax), "=r"(ebx), "=c"(ecx), "=d"(edx)
	    : "a"(op)
	    : "cc" );
	regs[0] = eax;
	regs[1] = ebx;
	regs[2] = ecx;
	regs[3] = edx;
}

void cpuidex( unsigned op, unsigned op2, unsigned* regs )
{
	unsigned eax, ebx, ecx, edx;
	asm volatile(
		"pushl %%ebx   \n\t"
	    "cpuid         \n\t"
	    "movl %%ebx, %1\n\t"
	    "popl %%ebx    \n\t"
	    : "=a"(eax), "=r"(ebx), "=c"(ecx), "=d"(edx)
	    : "a"(op), "c"(op2)
	    : "cc" );
	regs[0] = eax;
	regs[1] = ebx;
	regs[2] = ecx;
	regs[3] = edx;
}

unsigned cpuid_max_extended_function(void)
{
	unsigned reg[4];
	cpuid(0x80000000, reg);
	return reg[0];
}

void cpu_brand_name( char* name )
{
	if( cpuid_max_extended_function() >= 0x80000004 )
	{
		unsigned regs[4];
		unsigned i;
		for( i = 0; i < 3u; ++i )
		{
			cpuid( 0x80000002 + i, regs );
			memcpy( name + i * 16, regs, 16 );
		}
	}
	else
	{
		strcpy( name, "unknown" );
	}
}
*/
import "C"

const (
	Intel = iota
	AMD
	Other
)

func CpuId(funcId uint32, regs [4]uint32) {
	var cregs [4]C.unsigned
	C.cpuid(C.uint(funcId), &cregs[0])
	regs[0] = uint32(cregs[0])
	regs[1] = uint32(cregs[1])
	regs[2] = uint32(cregs[2])
	regs[3] = uint32(cregs[3])
}

func CpuIdEx(funcId uint32, funcId2 uint32, regs [4]uint32) {
	var cregs [4]C.unsigned
	C.cpuidex(C.uint(funcId), C.uint(funcId2), &cregs[0])
	regs[0] = uint32(cregs[0])
	regs[1] = uint32(cregs[1])
	regs[2] = uint32(cregs[2])
	regs[3] = uint32(cregs[3])
}

func VendorId() int {
	var regs [4]C.unsigned
	C.cpuid(0, &regs[0])
	ebx, ecx, edx := uint32(regs[1]), uint32(regs[2]), uint32(regs[3])
	if ebx == 0x756e6547 && edx == 0x49656e69 && ecx == 0x6c65746e {
		return Intel
	} else if ebx == 0x68747541 && edx == 0x69746e65 && ecx == 0x444d4163 {
		return AMD
	} else {
		return Other
	}
}

func BrandName() string {
	var nameBuf [64]C.char
	C.cpu_brand_name(&nameBuf[0])
	return C.GoString(&nameBuf[0])
}

func ThreadPerCore() int {
	if MaxFunctionId() < 0xb {
		return 1
	}

	var regs [4]C.unsigned
	C.cpuidex(0xb, 0, &regs[0])
	return int(regs[1]) & 0xffff
}

func LogicalCore() int {
	var regs [4]C.unsigned
	switch VendorId() {
	case Intel:
		if MaxFunctionId() < 0xb {
			return 1
		}
		C.cpuidex(0xb, 1, &regs[0])
		return int(regs[1]) & 0xffff
	case AMD:
		C.cpuid(1, &regs[0])
		return (int(regs[1]) >> 16) & 0xff
	default:
		return 1
	}
}

func PhysicalCore() int {
	switch VendorId() {
	case Intel:
		return LogicalCore() / ThreadPerCore()
	case AMD:
		var regs [4]C.unsigned
		C.cpuid(0x80000008, &regs[0])
		return (int(regs[2]) & 0xff) + 1
	default:
		return 1
	}
}

func MaxFunctionId() uint32 {
	var regs [4]C.unsigned
	C.cpuid(0, &regs[0])
	return uint32(regs[0])
}

func MaxExtendedFunctionId() uint32 {
	var regs [4]C.unsigned
	C.cpuid(0x80000000, &regs[0])
	return uint32(regs[0])
}
