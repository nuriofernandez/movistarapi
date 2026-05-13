package hgu

import (
	"fmt"

	"github.com/nuriofernandez/movistarapi/api"
)

func (h *HGUSession) OpenPort(port OpenPort) error {
	// Validate session
	if !h.IsValid {
		return fmt.Errorf("invalid session, must call Login() first")
	}

	portValue, err := port.Serialize()
	if err != nil {
		return err
	}

	// Forward it to api.OpenPort
	return api.OpenPort(h.sessionId, portValue)
}
