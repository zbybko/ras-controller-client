package sim

type SimInfo struct {
	DBusPath   string `json:"dbus-path"`
	Properties struct {
		Active            string   `json:"active"`
		Eid               string   `json:"eid"`
		EmergencyNumbers  []string `json:"emergency-numbers"`
		EsimStatus        string   `json:"esim-status"`
		Gid1              string   `json:"gid1"`
		Gid2              string   `json:"gid2"`
		Iccid             string   `json:"iccid"`
		Imsi              string   `json:"imsi"`
		OperatorCode      string   `json:"operator-code"`
		OperatorName      string   `json:"operator-name"`
		PreferredNetworks []string `json:"preferred-networks"`
		Removability      string   `json:"removability"`
		SimType           string   `json:"sim-type"`
	} `json:"properties"`
}
