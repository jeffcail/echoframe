package procs

import "runtime"

// GroRuntimeMaxCpu
func GroRuntimeMaxCpu() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
