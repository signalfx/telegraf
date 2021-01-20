package procstat

import (
	"regexp"
	"unsafe"

	"golang.org/x/sys/windows"
)

// Pattern matches on the process name
func (pg *NativeFinder) Pattern(pattern string) ([]PID, error) {
	var pids []PID
	regxPattern, err := regexp.Compile(pattern)
	if err != nil {
		return pids, err
	}
	procsName, err := getProcessesName()
	if err != nil {
		return pids, err
	}
	for pid, name := range procsName {
		if name == "" {
			continue
		}
		if regxPattern.MatchString(name) {
			pids = append(pids, PID(pid))
		}
	}
	return pids, err
}

func getProcessesName() (map[PID]string, error) {
	results := make(map[PID]string)
	snapshot, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = windows.CloseHandle(snapshot)
	}()
	var pe32 windows.ProcessEntry32
	pe32.Size = uint32(unsafe.Sizeof(pe32))
	if err = windows.Process32First(snapshot, &pe32); err != nil {
		return nil, err
	}
	for {
		results[PID(pe32.ProcessID)] = windows.UTF16ToString(pe32.ExeFile[:])
		if err = windows.Process32Next(snapshot, &pe32); err != nil {
			// ERROR_NO_MORE_FILES we reached the end of the snapshot
			break
		}
	}
	return results, nil
}
