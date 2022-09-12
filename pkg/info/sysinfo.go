package info

// app system status : dynamic record system status
type sysInfo struct {
	sysInfo map[string]interface{}
}

func (s *sysInfo) Set(k string, v interface{}) {
	s.sysInfo[k] = v
}
func (s *sysInfo) Get(k string) (v interface{}) {
	return s.sysInfo[k]
}

func getSysInfo() *sysInfo {
	return &sysInfo{
		sysInfo: make(map[string]interface{}),
	}
}

var SysInfo = getSysInfo()
