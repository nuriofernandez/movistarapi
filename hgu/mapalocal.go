package hgu

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nuriofernandez/movistarapi/api"
)

type ConnectedDevice struct {
	Unknown1       string // 1
	Name           string
	Unknown2       string // 0
	IPAddress      string
	ConnectionType string // Cable Ethernet, WIFI
	OpenPorts      bool
	MacAddress     string
}

func ParseConnectedDevice(line string) ConnectedDevice {
	// Split and clean string parts
	// '1','emby','0','192.168.1.37','Cable Ethernet','yes','90:2b:34:33:ff:9e'
	parts := strings.Split(line, ",")
	for i, v := range parts {
		parts[i] = strings.Trim(v, "'")
	}

	return ConnectedDevice{
		Unknown1:       parts[0],
		Name:           parts[1],
		Unknown2:       parts[2],
		IPAddress:      parts[3],
		ConnectionType: parts[4],
		OpenPorts:      parts[5] == "yes",
		MacAddress:     parts[6],
	}
}

func (h *HGUSession) LocalMap() ([]ConnectedDevice, error) {
	// Validate session
	if !h.IsValid {
		return []ConnectedDevice{}, fmt.Errorf("invalid session, must call Login() first")
	}

	body, err := api.LocalMap(h.sessionId)
	if err != nil {
		return []ConnectedDevice{}, err
	}

	// Clean response into a flat "'1','emby','0','192.168.1.37','Cable Ethernet','yes','90:2b:34:33:ff:9e'"
	trimmed := strings.Trim(body, "[]")
	reArray := regexp.MustCompile(`\[([^\]]+)\]`)
	matches := reArray.FindAllStringSubmatch(trimmed, -1)

	// Parse lines into ConnectedDevice struct
	devices := make([]ConnectedDevice, 0)
	for _, match := range matches {
		// match[1]: '1','emby',...
		rowRaw := match[1]

		device := ParseConnectedDevice(rowRaw)
		devices = append(devices, device)
	}

	return devices, nil
}
