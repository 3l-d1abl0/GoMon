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

	//var networkInfo commondata.NetworkInfo
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
