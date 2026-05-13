package api

import (
	"io"
	"net/http"
	"regexp"
)

// LocalMap Produces an array of devices properties [['1','hostname','0','192.168.1.10','Cable Ethernet','yes','ff:ff:ff:ff:ff:ff']]
func LocalMap(sessionId string) (string, error) {
	req, err := http.NewRequest("GET", "http://192.168.1.1/te_mapa_red_local.html", nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(&http.Cookie{Name: "sessionID", Value: sessionId})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// We Read the response body
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	//Convert the body to type string
	sb := string(body)
	var re = regexp.MustCompile(`(?m)deviceData=(.*)(;\n\nvar gatewa)`)

	// Extract group 1 (deviceData=value)
	return re.FindStringSubmatch(sb)[1], nil
}
