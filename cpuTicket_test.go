// cpuTicket_test.go
package cpu

import (
	"runtime"
	"testing"
	"time"
)

func Test(t *testing.T) {
	t.Logf("CPUTicks: %v", runtime.CPUTicks())
	t.Logf("CPUTicks: %v", runtime.CPUTicks())
}

func TestCpuTicks(t *testing.T) {
	N := 10000 * 1000
	start := time.Now()
	for i := 0; i < N; i++ {
		_ = runtime.CPUTicks()
	}
	end := time.Now()
	use := end.Sub(start)
	op := use / time.Duration(N)
	t.Logf("Times: %10v, Use: %14v %10v/op\n", N, use, op)
}
