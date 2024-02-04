package sysinfo

import (
	commondata "GoMon/commonData"
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetMemInfo(t *testing.T) {

	memInfo, err := GetMemInfo()

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	type MemoryInfo struct {
		Total       uint64  `json:"total"`
		Available   uint64  `json:"available"`
		Used        uint64  `json:"used"`
		UsedPercent float64 `json:"usedPercent"`
		Free        uint64  `json:"free"`
		Active      uint64  `json:"active"`
		Inactive    uint64  `json:"inactive"`
		SwapTotal   uint64  `json:"swapTotal"`
		SwapFree    uint64  `json:"swapFree"`
	}

	type MemoryInfoResponse struct {
		MemoryInfo MemoryInfo `json:"memoryInfo"`
	}

	var memoryInfoContainer MemoryInfoResponse

	marshallErr := json.Unmarshal([]byte(fmt.Sprintln(memInfo)), &memoryInfoContainer)
	if marshallErr != nil {
		t.Errorf("Error decoding JSON: %s", marshallErr)
	}
}

func TestGetNetinfo(t *testing.T) {

	netInfo, err := GetNetInfo()

	if err != nil {
		t.Errorf("Error fetching Network Info: %s", err)
	}

	var networkInterfaces []commondata.NetworkInterface

	jsonData, marshallErr := json.MarshalIndent(netInfo, "", "  ")
	if marshallErr != nil {
		t.Errorf("Error marshalling NetworkInfo data: %s", marshallErr)
	}
	marshallErr = json.Unmarshal([]byte(string(jsonData)), &networkInterfaces)
	if marshallErr != nil {
		t.Errorf("Error decoding networkInfo JSON: %s", marshallErr)
	}

}

func TestGetCpuInfo(t *testing.T) {

	cpuInfo, err := GetCPUInfo()

	if err != nil {
		t.Errorf("Error fetching CPU Info: %s", err)
	}

	//Convert cpu.InfoStat to Json
	jsonData, marshallErr := json.MarshalIndent(cpuInfo, "", "  ")
	if marshallErr != nil {
		t.Errorf("Error marshalling CpuInfo data: %s", marshallErr)
	}

	var cpuInfoList []commondata.CPUInfo
	//Check fo valid JSON Structure
	marshallErr = json.Unmarshal([]byte(string(jsonData)), &cpuInfoList)
	if marshallErr != nil {
		t.Errorf("Error decoding cpu.InfoStat JSON: %s", marshallErr)
	}
}
func TestGetHostInfo(t *testing.T) {

	hostInfoData, err := GetHostInfo()

	if err != nil {
		t.Errorf("Error fetching Host Info: %s", err)
	}

	var hostInfo commondata.CPUInfo
	//Check fo valid JSON Structure
	marshallErr := json.Unmarshal([]byte(fmt.Sprintln(hostInfoData)), &hostInfo)
	if marshallErr != nil {
		t.Errorf("Error decoding host.InfoStat JSON: %s", marshallErr)
	}
}

func TestGetLoadInfo(t *testing.T) {

	loadInfoData, err := GetLoadInfo()
	if err != nil {
		t.Errorf("Error fetching Load Info: %s", err)
	}

	var loadInfo commondata.LoadInfo
	//Check fo valid JSON Structure
	marshallErr := json.Unmarshal([]byte(fmt.Sprintln(loadInfoData)), &loadInfo)
	if marshallErr != nil {
		t.Errorf("Error decoding load.AvgStat JSON: %s", marshallErr)
	}
}
