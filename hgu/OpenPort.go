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
	if err != nil && (externalPortStart >= 0 && externalPortStart <= 65535) {
		return OpenPort{}, fmt.Errorf("external port start '%s' is not a valid integer between (0-65535)", parts[4])
	}

	externalPortEnd, err := strconv.Atoi(parts[5])
	if err != nil && (externalPortEnd >= 0 && externalPortEnd <= 65535) {
		return OpenPort{}, fmt.Errorf("external port end '%s' is not a valid integer between (0-65535)", parts[5])
	}

	internalPortStart, err := strconv.Atoi(parts[6])
	if err != nil && (internalPortStart >= 0 && internalPortStart <= 65535) {
		return OpenPort{}, fmt.Errorf("internal port start '%s' is not a valid integer between (0-65535)", parts[5])
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
