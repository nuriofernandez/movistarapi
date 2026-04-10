package hgu

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (h *HGUSession) OpenPorts() (string, error) {
	// Validate session
	if !h.IsValid {
		return "", fmt.Errorf("invalid session, must call Login() first")
	}

	req, err := http.NewRequest("GET", "http://192.168.1.1/te_ppp_pm.cmd?action=retrieve&type=identifier&sessionKey="+h.sessionId, nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(&http.Cookie{Name: "sessionID", Value: h.sessionId})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	//Convert the body to type string
	sb := string(body)

	// Clean results
	removeScript := strings.Split(sb, "</script>\ninst=")[1]
	removeSession := strings.Split(removeScript, "&")[0]

	// Fetch all ids one by one and add them to list string
	var list string
	for _, id := range strings.Split(removeSession, ",") {
		s, err := fetchPort(h.sessionId, id)
		if err != nil {
			continue
		}
		if len(list) != 0 {
			list += ","
		}
		list += "[" + s + "]"
	}

	// Return the list of open ports
	return "[" + list + "]", nil
}

func fetchPort(sessionId, id string) (string, error) {
	req, err := http.NewRequest("GET", "http://192.168.1.1/te_ppp_pm.cmd?action=retrieve&type=instance&id="+id, nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(&http.Cookie{Name: "sessionID", Value: sessionId})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	//Convert the body to type string
	sb := string(body)

	// Clean response
	removeScript := strings.Split(sb, "</script>\ninst=")[1]
	removeSession := strings.Split(removeScript, "&")[0]

	// return id list
	return removeSession, nil
}
