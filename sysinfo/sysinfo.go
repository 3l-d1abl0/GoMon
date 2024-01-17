package sysinfo

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/net"
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

func GetNetInfo() ([]net.InterfaceStat, error) {

	netInfo, err := net.Interfaces()
	//1. If error encountered while fetching mem Info
	if err != nil {
		return nil, fmt.Errorf("[net.Interfaces() Error] %v", err.Error())
	}
	//2. Check if valid Json
	_, errMarshal := json.Marshal(netInfo)
	if errMarshal != nil {
		return nil, fmt.Errorf("[json.Marshal Error] %v", errMarshal.Error())
	}

	return netInfo, nil
}
