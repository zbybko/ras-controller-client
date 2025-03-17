package wifi

import (
	"errors"
	"fmt"
	"sync"
)

type MockWiFiManager struct {
	mu         sync.Mutex
	active     bool
	ssid       string
	hiddenSSID bool
	password   string
	security   string
	channel    int
}

func NewMockWiFiManager() *MockWiFiManager {
	return &MockWiFiManager{
		active:     false,
		ssid:       "MockSSID",
		hiddenSSID: false,
		password:   "mockpassword",
		security:   "wpa-psk",
		channel:    6,
	}
}

func (m *MockWiFiManager) Status() (*WiFiInfo, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	status := fmt.Sprintf(
		"Active: %v\nSSID: %s\nHidden SSID: %v\nPassword: %s\nSecurity: %s\nChannel: %d",
		m.active, m.ssid, m.hiddenSSID, m.password, m.security, m.channel,
	)

	// Выводим информацию о статусе
	fmt.Println(status)

	// Возвращаем все поля в структуре WiFiInfo
	return &WiFiInfo{
		Active:     m.active,
		SSID:       m.ssid,
		HiddenSSID: m.hiddenSSID,
		Password:   m.password,
		Security:   m.security,
		Channel:    m.channel,
	}, nil
}

func (m *MockWiFiManager) Enable() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Симулируем успешное включение Wi-Fi
	if m.active {
		return errors.New("Wi-Fi is already enabled") // Если Wi-Fi уже включен, возвращаем ошибку
	}
	m.active = true
	fmt.Println("Mock: Wi-Fi enabled successfully")
	return nil
}

func (m *MockWiFiManager) Disable() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Симулируем успешное выключение Wi-Fi
	if !m.active {
		return errors.New("Wi-Fi is already disabled") // Если Wi-Fi уже выключен, возвращаем ошибку
	}
	m.active = false
	fmt.Println("Mock: Wi-Fi disabled successfully")
	return nil
}

func (m *MockWiFiManager) SetSSID(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Симулируем успешную смену SSID
	if name == "" {
		return errors.New("SSID cannot be empty")
	}
	m.ssid = name
	fmt.Printf("Mock: SSID set to %s\n", name)
	return nil
}

func (m *MockWiFiManager) SetSSIDHidden(hidden bool) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Симулируем успешную смену состояния скрытости SSID
	m.hiddenSSID = hidden
	fmt.Printf("Mock: SSID hidden state set to %v\n", hidden)
	return nil
}

func (m *MockWiFiManager) SetPassword(password string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Симулируем успешную смену пароля
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	m.password = password
	fmt.Println("Mock: Password set successfully")
	return nil
}

func (m *MockWiFiManager) SetSecurityType(wpa3 bool) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Симулируем успешную смену типа безопасности
	if wpa3 {
		m.security = "sae"
	} else {
		m.security = "wpa-psk"
	}
	fmt.Printf("Mock: Security type set to %s\n", m.security)
	return nil
}

func (m *MockWiFiManager) SetChannel(channel int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Симулируем успешную смену канала
	if channel < 1 || channel > 11 {
		return errors.New("invalid channel number")
	}
	m.channel = channel
	fmt.Printf("Mock: Channel set to %d\n", channel)
	return nil
}

func (m *MockWiFiManager) ServiceExists() bool {
	// Симулируем существование сервиса Wi-Fi
	return true
}

func (m *MockWiFiManager) GetState() string {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Возвращаем подробное состояние Wi-Fi
	return fmt.Sprintf(
		"Wi-Fi Active: %v\nSSID: %s\nHidden: %v\nPassword: %s\nSecurity: %s\nChannel: %d",
		m.active, m.ssid, m.hiddenSSID, m.password, m.security, m.channel,
	)
}
