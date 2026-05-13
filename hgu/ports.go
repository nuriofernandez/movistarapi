package hgu

import (
	"fmt"
	"strings"

	"github.com/nuriofernandez/movistarapi/api"
)

func (h *HGUSession) OpenPorts() ([]OpenPort, error) {
	// Validate session
	if !h.IsValid {
		return []OpenPort{}, fmt.Errorf("invalid session, must call Login() first")
	}

	// Fetch ports Ids from api
	portsIds, err := api.ListPortsIds(h.sessionId)
	if err != nil {
		return []OpenPort{}, err
	}

	// Fetch all ids one by one and add them to list string
	var list []OpenPort
	for _, id := range strings.Split(portsIds, ",") {
		s, err := api.FetchPort(h.sessionId, id)
		if err != nil {
			continue
		}
		openPort, err := Parse(s)
		if err != nil {
			// Don't break the flow since it's data coming from the router.
			fmt.Printf("WARNING: (ignored) unable to parse port '%s' due to '%f'\n", s, err)
			continue
		}
		list = append(list, openPort)
	}

	// Return the list of open ports
	return list, nil
}
