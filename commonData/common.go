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
