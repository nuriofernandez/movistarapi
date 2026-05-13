package hgu

import (
	"fmt"

	"github.com/nuriofernandez/movistarapi/api"
)

func (h *HGUSession) Reboot() error {
	// Validate session
	if !h.IsValid {
		return fmt.Errorf("invalid session, must call Login() first")
	}

	return api.Reboot(h.sessionId)
}
