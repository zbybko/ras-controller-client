package wifi

type Manager interface {
	Status() (*WiFiInfo, error)
	Enable() error
	Disable() error
	SetSSID(name string) error
	SetSSIDHidden(hidden bool) error
	SetPassword(password string) error
	SetSecurityType(wpa3 bool) error
	SetChannel(channel int) error
	ServiceExists() bool
}
