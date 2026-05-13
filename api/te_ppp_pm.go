package api

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func DeletePort(sessionId string, portId int, wanInterface string) error {
	req, err := http.NewRequest("GET", "http://192.168.1.1/te_ppp_pm.cmd?action=delete&id="+strconv.Itoa(portId)+"&intf="+wanInterface+"&sessionKey="+sessionId, nil)
	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{Name: "sessionID", Value: sessionId})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// We read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	//Convert the body to type string
	stringBody := string(body)
	removeScript := strings.Split(stringBody, "</script>\n")[1]

	responseOk := strings.Contains(removeScript, "inst=true")
	if !responseOk {
		return fmt.Errorf("failed to perform operation, movistar denied it")
	}

	// All right!
	return nil
}
