package v0

type Version struct {
	API      string `json:"api"`
	Server   string `json:"server"`
	Original string `json:"original"`
	Text     string `json:"text"`
	Hostname string `json:"hostname"`
	Firmware string `json:"firmware"`
	SDK      string `json:"sdk"`
}
