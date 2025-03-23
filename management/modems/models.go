package modems

type ModemInfo struct {
	// Third Generation Partnership Project - 3GPP
	ThreeGpp struct {
		FiveGnr struct {
			RegistrationSettings struct {
				DrxCycle  string `json:"drx-cycle"`
				MicroMode string `json:"micro-mode"`
			} `json:"registration-settings"`
		} `json:"5gnr"`
		EnableLocks []string `json:"enable-locks"`
		Eps         struct {
			InitialBearer struct {
				DBusPath string `json:"dbus-path"`
				Settings struct {
					Apn      string `json:"apn"`
					IpType   string `json:"ip-type"`
					User     string `json:"user"`
					Password string `json:"password"`
				} `json:"settings"`
			} `json:"initial-bearer"`
			UeModeOperation string `json:"ue-mode-operation"`
		} `json:"eps"`
		Imei               string `json:"imei"`
		OperatorCode       string `json:"operator-code"`
		OperatorName       string `json:"operator-name"`
		PacketServiceState string `json:"packet-service-state"`
		Pco                string `json:"pco"`
		RegistrationState  string `json:"registration-state"`
	} `json:"3gpp"`
	Cdma struct {
		ActivationState         string `json:"activation-state"`
		Cdma1xRegistrationState string `json:"cdma1x-registration-state"`
		Esn                     string `json:"esn"`
		EvdoRegistrationState   string `json:"evdo-registration-state"`
		Meid                    string `json:"meid"`
		Nid                     string `json:"nid"`
		Sid                     string `json:"sid"`
	} `json:"cdma"`
	DBusPath string `json:"dbus-path"`
	Generic  struct {
		AccessTechnologies           []string `json:"access-technologies"`
		Bearers                      []string `json:"bearers"`
		CarrierConfiguration         string   `json:"carrier-configuration"`
		CarrierConfigurationRevision string   `json:"carrier-configuration-revision"`
		CurrentBands                 []string `json:"current-bands"`
		CurrentCapabilities          []string `json:"current-capabilities"`
		CurrentModes                 string   `json:"current-modes"`
		Device                       string   `json:"device"`
		DeviceIdentifier             string   `json:"device-identifier"`
		Drivers                      []string `json:"drivers"`
		EquipmentIdentifier          string   `json:"equipment-identifier"`
		HardwareRevision             string   `json:"hardware-revision"`
		Manufacturer                 string   `json:"manufacturer"`
		Model                        string   `json:"model"`
		OwnNumbers                   []string `json:"own-numbers"`
		Plugin                       string   `json:"plugin"`
		Ports                        []string `json:"ports"`
		PowerState                   string   `json:"power-state"`
		PrimaryPort                  string   `json:"primary-port"`
		PrimarySimSlot               string   `json:"primary-sim-slot"`
		Revision                     string   `json:"revision"`
		SignalQuality                struct {
			Recent string `json:"recent"`
			Value  string `json:"value"`
		} `json:"signal-quality"`
		Sim                   string   `json:"sim"`
		SimSlots              []string `json:"sim-slots"`
		State                 string   `json:"state"`
		StateFailedReason     string   `json:"state-failed-reason"`
		SupportedBands        []string `json:"supported-bands"`
		SupportedCapabilities []string `json:"supported-capabilities"`
		SupportedIpFamilies   []string `json:"supported-ip-families"`
		SupportedModes        []string `json:"supported-modes"`
		UnlockRequired        string   `json:"unlock-required"`
		UnlockRetries         []string `json:"unlock-retries"`
	} `json:"generic"`
}

type ModemSignal struct {
	// 5G
	FiveG struct {
		Rsrp      string `json:"rsrp"`
		Rsrq      string `json:"rsrq"`
		Snr       string `json:"snr"`
		ErrorRate string `json:"error-rate"`
	} `json:"5g"`
	Cdma1x struct {
		Ecio      string `json:"ecio"`
		Rssi      string `json:"rssi"`
		ErrorRate string `json:"error-rate"`
	} `json:"cdma1x"`
	Evdo struct {
		Ecio      string `json:"ecio"`
		Rssi      string `json:"rssi"`
		Sinr      string `json:"sinr"`
		Io        string `json:"io"`
		ErrorRate string `json:"error-rate"`
	} `json:"evdo"`
	Lte struct {
		Rssi      string `json:"rssi"`
		Rsrp      string `json:"rsrp"`
		Rsrq      string `json:"rsrq"`
		Snr       string `json:"snr"`
		ErrorRate string `json:"error-rate"`
	} `json:"lte"`
	Refresh struct {
		Rate string `json:"rate"`
	} `json:"refresh"`
	Threshold struct {
		ErrorRate string `json:"error-rate"`
		Rssi      string `json:"rssi"`
	} `json:"threshold"`
	Umts struct {
		ErrorRate string `json:"error-rate"`
		Ecio      string `json:"ecio"`
		Rscp      string `json:"rscp"`
		Rssi      string `json:"rssi"`
	} `json:"umts"`
}

type BearerInfo struct {
	DBusPath   string `json:"dbus-path"`
	Properties struct {
		Apn      string `json:"apn"`
		IpType   string `json:"ip-type"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"properties"`
	Status struct {
		Connected string `json:"connected"` // 'yes' for true
		Interface string `json:"interface"`
	} `json:"status"`
	Type string `json:"type"`
}
