package hgu

import (
	"github.com/nuriofernandez/movistarapi/api"
)

func (h *HGUSession) Login(pass string) (string, error) {
	// Authenticate with api.Login
	sessionId, err := api.Login(pass)
	if err != nil {
		return "", err
	}

	// Store sessionId and return a successful response
	h.sessionId = sessionId
	h.IsValid = true
	return sessionId, nil
}
