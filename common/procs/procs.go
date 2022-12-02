package _procs

import "runtime"

// GroRuntimeMaxCpu
func GroRuntimeMaxCpu() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
