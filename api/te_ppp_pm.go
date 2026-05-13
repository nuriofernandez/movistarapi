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

func OpenPort(sessionId, portValue string) error {
	req, err := http.NewRequest("GET", "http://192.168.1.1/te_ppp_pm.cmd?action=create&inst="+portValue+"&sessionKey="+sessionId, nil)
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

// ListPortsIds returns a list of entries ids '1,2,3,4,5,6,7,8,9,10'
func ListPortsIds(sessionId string) (string, error) {
	req, err := http.NewRequest("GET", "http://192.168.1.1/te_ppp_pm.cmd?action=retrieve&type=identifier&sessionKey="+sessionId, nil)
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

	// Clean results
	removeScript := strings.Split(sb, "</script>\ninst=")[1]
	return strings.Split(removeScript, "&")[0], nil
}

// FetchPort provides a string with the port properties: "40,test-1,TCP,192.168.1.165,12345,12346,12390,T,ppp0.1"
func FetchPort(sessionId, id string) (string, error) {
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
	rawPortData := strings.Split(removeScript, "&")[0]

	// return id list
	return rawPortData, nil
}
