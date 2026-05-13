package api

import "net/http"

func Reboot(sessionId string) error {
	req, err := http.NewRequest("GET", "http://192.168.1.1/rebootinfo.cgi?sessionKey="+sessionId, nil)
	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{Name: "sessionID", Value: sessionId})

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
