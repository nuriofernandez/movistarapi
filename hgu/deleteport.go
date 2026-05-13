package hgu

import (
	"fmt"

	"github.com/nuriofernandez/movistarapi/api"
)

func (h *HGUSession) DeletePort(portId int, wanInterface string) error {
	// Validate session
	if !h.IsValid {
		return fmt.Errorf("invalid session, must call Login() first")
	}

	// Forward to api.DeletePort
	return api.DeletePort(h.sessionId, portId, wanInterface)
}
