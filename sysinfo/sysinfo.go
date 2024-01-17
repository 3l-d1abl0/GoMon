package sysinfo

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
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

func GetCPUInfo() ([]cpu.InfoStat, error) {

	cpuInfo, err := cpu.Info()
	//1. If error encountered while fetching cpu indo
	if err != nil {
		return nil, fmt.Errorf("[cpu.Info() Error] %v", err.Error())
	}
	//2. If valid json format
	_, errMarshal := json.Marshal(cpuInfo)
	if errMarshal != nil {
		return nil, fmt.Errorf("[json.Marshal Error] %v", errMarshal.Error())
	}

	return cpuInfo, nil
}

func GetHostInfo() (*host.InfoStat, error) {

	hostInfo, err := host.Info()
	//1. If errpr occured while fetching host Info
	if err != nil {
		return nil, fmt.Errorf("[host.Info() Error] %v", err.Error())
	}

	//2. If valid json format
	_, errMarshal := json.Marshal(hostInfo)
	if err != nil {
		return nil, fmt.Errorf("[json.Marshal Error] %v", errMarshal.Error())
	}

	return hostInfo, nil
}

func GetLoadInfo() (*load.AvgStat, error) {

	loadInfo, err := load.Avg()
	//1. If errpr occured while fetching load Info
	if err != nil {
		return nil, fmt.Errorf("[load.Avg() Error] %v", err.Error())
	}

	//2. If valid json format
	_, errMarshal := json.Marshal(loadInfo)
	if err != nil {
		return nil, fmt.Errorf("[json.Marshal Error] %v", errMarshal.Error())
	}

	return loadInfo, nil
}
