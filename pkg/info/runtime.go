package info

import (
	"fmt"
	rt "runtime"
	"strings"
)

// record system runtime
type runtime struct {
	os    string
	arch  string
	goVer string
}

func (r *runtime) ToString() string {
	ap := strings.Join([]string{r.os, r.arch, r.goVer}, "  ")
	return ap
}

func (r *runtime) Print() {
	fmt.Printf("Os: %s \n", r.os)
	fmt.Printf("Arch: %s \n", r.arch)
	fmt.Printf("Go Version: %s \n", r.goVer)
}

func getRuntime() *runtime {
	return &runtime{
		os:    rt.GOOS,
		arch:  rt.GOARCH,
		goVer: rt.Version(),
	}
}

var Runtime = getRuntime()
