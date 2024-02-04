package commondata

type NetworkAddr struct {
	Addr string `json:"addr"`
}

type NetworkInterface struct {
	Index        int           `json:"index"`
	MTU          int           `json:"mtu"`
	Name         string        `json:"name"`
	HardwareAddr string        `json:"hardwareaddr"`
	Flags        []string      `json:"flags"`
	Addresses    []NetworkAddr `json:"addrs"`
}

type NetworkInfo struct {
	Interfaces []NetworkInterface `json:"networkInfo"`
}

type CPUInfo struct {
	CPU        int      `json:"cpu"`
	VendorID   string   `json:"vendorId"`
	Family     string   `json:"family"`
	Model      string   `json:"model"`
	Stepping   int      `json:"stepping"`
	PhysicalID string   `json:"physicalId"`
	CoreID     string   `json:"coreId"`
	Cores      int      `json:"cores"`
	ModelName  string   `json:"modelName"`
	MHz        int      `json:"mhz"`
	CacheSize  int      `json:"cacheSize"`
	Flags      []string `json:"flags"`
	Microcode  string   `json:"microcode"`
}

type CPUInfoList struct {
	CPUInfoList []CPUInfo `json:"cpuInfo"`
}
