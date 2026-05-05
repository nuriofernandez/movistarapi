package hgu

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
 * OpenPort maps a port forward rule on the router
 * Example rule record = 40,test-1,TCP,192.168.1.165,12345,12346,12390,T,ppp0.1
 */

type Protocol string

const (
	TCP  Protocol = "TCP"
	UDP  Protocol = "UDP"
	BOTH Protocol = "BOTH"
)

type OpenPort struct {
	Id                int      `json:"id"`
	Name              string   `json:"name"`     // up to 16char
	Protocol          Protocol `json:"protocol"` // TCP/UDP/BOTH
	Address           string   `json:"address"`  // 192.168.1.165
	ExternalPortStart int      `json:"externalPortStart"`
	ExternalPortEnd   int      `json:"externalPortEnd"`
	InternalPortStart int      `json:"internalPortStart"` // Map ExternalPortStart, if there is a range, will be calculated automatically
	Enabled           bool     `json:"enabled"`           // T/F
	Interface         string   `json:"interface"`         // ppp0.1
}

func (o *OpenPort) Serialize() (string, error) {
	line := fmt.Sprintf("%d,", o.Id)

	if o.Name == "" || len(o.Name) > 16 {
		return "", fmt.Errorf("invalid name: must be between 1 and 16 characters long")
	}
	line += fmt.Sprintf("%s,", o.Name)

	if o.Protocol != TCP && o.Protocol != UDP && o.Protocol != BOTH {
		return "", fmt.Errorf("protocol '%s' is not supported, (supported protocols: TCP/UDP/BOTH)", o.Protocol)
	}
	line += fmt.Sprintf("%s,", o.Protocol)

	ipv4Pattern := `^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`
	re := regexp.MustCompile(ipv4Pattern)
	if !re.MatchString(o.Address) {
		return "", fmt.Errorf("address '%s' is not a valid IPv4 address", o.Address)
	}
	line += fmt.Sprintf("%s,", o.Address)

	if o.ExternalPortStart <= 0 || o.ExternalPortStart >= 65535 {
		return "", fmt.Errorf("external port start '%d' is not a valid integer between (1-65535)", o.ExternalPortStart)
	}
	line += fmt.Sprintf("%d,", o.ExternalPortStart)

	// This CAN be 0, in case is not defined.
	if o.ExternalPortEnd < 0 || o.ExternalPortEnd >= 65535 {
		return "", fmt.Errorf("external port end '%d' is not a valid integer between (0-65535)", o.ExternalPortEnd)
	}
	line += fmt.Sprintf("%d,", o.ExternalPortEnd)

	if o.InternalPortStart <= 0 || o.InternalPortStart >= 65535 {
		return "", fmt.Errorf("internal port start '%d' is not a valid integer between (1-65535)", o.InternalPortStart)
	}
	line += fmt.Sprintf("%d,", o.InternalPortStart)

	if o.Enabled {
		line += "T,"
	} else {
		line += "F,"
	}

	// TODO validate
	line += o.Interface

	return line, nil
}

func Parse(line string) (OpenPort, error) {
	parts := strings.Split(line, ",")

	id, err := strconv.Atoi(parts[0])
	if err != nil {
		return OpenPort{}, fmt.Errorf("record id '%s' is not a valid number", parts[0])
	}

	name := parts[1]
	if len(name) > 16 {
		return OpenPort{}, fmt.Errorf("name '%s' length exceeds 16 characters", name)
	}

	prot := Protocol(strings.ToUpper(parts[2]))
	if prot != TCP && prot != UDP && prot != BOTH {
		return OpenPort{}, fmt.Errorf("protocol '%s' is not supported", parts[2])
	}

	address := parts[3]
	ipv4Pattern := `^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`
	re := regexp.MustCompile(ipv4Pattern)
	if !re.MatchString(address) {
		return OpenPort{}, fmt.Errorf("address '%s' is not a valid IPv4 address", address)
	}

	externalPortStart, err := strconv.Atoi(parts[4])
	if err != nil || externalPortStart <= 0 || externalPortStart >= 65535 {
		return OpenPort{}, fmt.Errorf("external port start '%s' is not a valid integer between (1-65535)", parts[4])
	}

	externalPortEnd, err := strconv.Atoi(parts[5])
	// This CAN be 0, in case is not defined.
	if err != nil || externalPortEnd < 0 || externalPortEnd >= 65535 {
		return OpenPort{}, fmt.Errorf("external port end '%s' is not a valid integer between (0-65535)", parts[5])
	}

	internalPortStart, err := strconv.Atoi(parts[6])
	if err != nil || internalPortStart <= 0 || internalPortStart >= 65535 {
		return OpenPort{}, fmt.Errorf("internal port start '%s' is not a valid integer between (1-65535)", parts[5])
	}

	enabled, err := strconv.ParseBool(parts[7])
	if err != nil {
		return OpenPort{}, fmt.Errorf("enabled '%s' is not a valid boolean", parts[7])
	}

	wanInterface := parts[8] // TODO not validated

	return OpenPort{
		Id:                id,
		Name:              name,
		Protocol:          prot,
		Address:           address,
		ExternalPortStart: externalPortStart,
		ExternalPortEnd:   externalPortEnd,
		InternalPortStart: internalPortStart,
		Enabled:           enabled,
		Interface:         wanInterface,
	}, nil
}
