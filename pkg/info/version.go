package info

import (
	"fmt"
	"strings"
)

// record system version
type version struct {
	ver string
	tag string
	upt string
	env string
}

func (v *version) ToString() string {
	p1 := strings.Join([]string{v.ver, v.upt}, ".")
	ap := strings.Join([]string{p1, v.tag}, "@")
	return ap
}

func (v *version) Print() {
	fmt.Printf("Version: %s \n", v.ver)
	fmt.Printf("Git Tag: %s \n", v.tag)
	fmt.Printf("Update Time: %s \n", v.upt)
	fmt.Printf("Compiler Environment: %s \n", v.env)
}

func getVersion() *version {
	return &version{
		ver: verStr,
		tag: tagStr,
		upt: uptStr,
		env: envStr,
	}
}

var Version = getVersion()
