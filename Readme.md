# Ras Controller

**Ras Controller** is a modular web-based system for managing core router functions on Linux-based embedded devices.

It provides a REST API and lightweight web interface for configuring Wi-Fi, Ethernet, DHCP, SSH, and other system-level services. Built with extensibility in mind, it is ideal for routers, gateways, and embedded Linux platforms.

---

## 🧑‍💻 My Contribution

This project is developed by a team.  
I personally implemented the following key modules:

- **Wi-Fi management**: SSID visibility, password, encryption (WPA2/WPA3), channel selection (`hostapd`)
- **DHCP server control**: enable/disable, IP range configuration, static lease support (`dhcpd`)
- **SSH service**: enable/disable over `systemd`, firewall integration via `firewalld`
- **Ethernet status**: display active ports, IPs, MAC addresses, and interface state
- **System service control**: interaction with `systemctl`, including restarts and service detection
- **API backend**: designed and implemented REST endpoints using the `gin-gonic` framework in Go

Other modules such as SIM support, VPN, or cloud sync are developed by other contributors.

---

## ✨ Features Overview

- Wi-Fi configuration using `hostapd`
- DHCP server management with support for static and dynamic leases
- SSH toggle and access management
- Ethernet interfaces status monitoring
- System control integrations (reboot, firewall, restart services)
- JSON-based REST API for frontend or remote access

---

## 💡 Technologies Used

| Layer       | Stack                                                                 |
|-------------|-----------------------------------------------------------------------|
| Language    | Go (Golang)                                                           |
| Framework   | [Gin Gonic](https://github.com/gin-gonic/gin) – Fast HTTP web framework |
| Services    | `hostapd`, `dhcpd`, `systemd`, `firewalld`                            |
| OS Support  | Fedora, Debian, Ubuntu, OpenWRT, Yocto, and most Linux-based systems |
| Interface   | RESTful API (JSON), planned: Vue.js-based web interface              |

---

## 🚀 How to Run Locally

> **Note:** root permissions are required for controlling system services.

1. Clone the repository:

```bash
git clone git@gitlab.com:ras995910/ras-controller-client.git
cd ras-controller-client
```

2. Build the Go binary:

```bash
go build -o ras-controller .
```

3. Run:

```bash
go run ./cmd/server/main.go
```

4. Open in browser or via curl:

```bash
http://localhost:8080/api/
```

5. To check Wi-Fi status:

```bash
curl http://localhost:8080/api/wifi/status
```

---

## 🔐 Example API Endpoints

| Endpoint                        | Method | Description                         |
|--------------------------------|--------|-------------------------------------|
| `/api/wifi/enable`             | POST   | Enable Wi-Fi                        |
| `/api/wifi/status`             | GET    | Get Wi-Fi configuration             |
| `/api/dhcp/leases`             | GET    | View active DHCP leases             |
| `/api/dhcp/static/add`         | POST   | Add static DHCP lease               |
| `/api/ssh/enable`              | POST   | Enable SSH                          |
| `/api/ethernet/status`         | GET    | Get Ethernet port statuses          |

---

> _Focused on system-level development, API design, and router logic in Golang._