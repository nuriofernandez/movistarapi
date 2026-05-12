package hgu

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type ConnectedDevice struct {
	Unknown1       string
	Name           string
	Unknown2       string
	IPAddress      string
	ConnectionType string
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

	req, err := http.NewRequest("GET", "http://192.168.1.1/te_mapa_red_local.html", nil)
	if err != nil {
		return []ConnectedDevice{}, err
	}

	req.AddCookie(&http.Cookie{Name: "sessionID", Value: h.sessionId})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []ConnectedDevice{}, err
	}

	// We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []ConnectedDevice{}, err
	}

	//Convert the body to type string
	sb := string(body)
	var re = regexp.MustCompile(`(?m)deviceData=(.*)(;\n\nvar gatewa)`)

	// Extract group 1 (deviceData=value)
	localMapStr := re.FindStringSubmatch(sb)[1]
	localMapStr = strings.Trim(localMapStr, "[]")

	reArray := regexp.MustCompile(`\[([^\]]+)\]`)
	matches := reArray.FindAllStringSubmatch(localMapStr, -1)

	devices := make([]ConnectedDevice, 0)
	for _, match := range matches {
		// match[1]: '1','emby',...
		rowRaw := match[1]

		device := ParseConnectedDevice(rowRaw)
		devices = append(devices, device)
	}

	return devices, nil
}
