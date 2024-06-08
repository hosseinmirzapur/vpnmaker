package main

type BaseConfig struct {
	Outbounds []Outbound `json:"outbounds"`
}

type Outbound struct {
	Type          string      `json:"type"`
	Tag           string      `json:"tag"`
	Server        string      `json:"server"`
	ServerPort    int         `json:"server_port"`
	LocalAddress  []string    `json:"local_address"`
	PrivateKey    string      `json:"private_key"`
	PeerPublicKey string      `json:"peer_public_key"`
	Reserved      interface{} `json:"reserved"`
	Mtu           int         `json:"mtu"`
	FakePackets   string      `json:"fake_packets"`
}

type WarpAccount struct {
	IPV6       string   `json:"ipv6"`
	PrivateKey string   `json:"private_key"`
	Reserved   []string `json:"reserved"`
}

type Result struct {
	IpPort string `csv:"IP:PORT"`
	Loss   string `csv:"LOSS"`
	Delay  string `csv:"DELAY"`
}
