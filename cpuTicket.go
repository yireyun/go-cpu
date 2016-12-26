// cpuTicket.go
package cpu

import (
	"runtime"
)

func CPUTicks() int64 {
	return runtime.CPUTicks()
}
