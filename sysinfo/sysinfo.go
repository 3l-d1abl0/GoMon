package sysinfo

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

func GetMemInfo() (*mem.VirtualMemoryStat, error) {

	vMemInfo, err := mem.VirtualMemory()

	//If we encounter error while fetching memory info
	if err != nil {
		return nil, fmt.Errorf("[mem.VirtualMemory() Error] %v", err.Error())
	}

	//check if its a valid Json
	_, errMarshal := json.Marshal(vMemInfo)
	if errMarshal != nil {
		return nil, fmt.Errorf("[json.Marshal Error] %v", errMarshal.Error())
	}

	return vMemInfo, nil
}
