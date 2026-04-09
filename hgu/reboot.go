package hgu

import (
	"fmt"
	"net/http"
)

func (h *HGUSession) Reboot() error {
	// Validate session
	if !h.IsValid {
		return fmt.Errorf("invalid session, must call Login() first")
	}

	req, err := http.NewRequest("GET", "http://192.168.1.1/rebootinfo.cgi?sessionKey="+h.sessionId, nil)
	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{Name: "sessionID", Value: h.sessionId})

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
