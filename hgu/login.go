package hgu

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

// Generate a 9 digits random number to be used as a sessionId
func randomSessionId() string {
	minLimit := int(math.Pow10(9))
	maxLimit := int(math.Pow10(9 - 1))
	randInt := int(rand.Float64() * float64(minLimit))
	if randInt < maxLimit {
		randInt += maxLimit
	}
	return fmt.Sprint(randInt)
}

func (h *HGUSession) Login(pass string) (string, error) {
	// Generate a random sessionId
	sessionId := randomSessionId()

	// Prepare login form
	data := url.Values{}
	data.Set("loginPassword", pass)

	// Prepare request
	req, err := http.NewRequest("POST", "http://192.168.1.1/te_acceso_router.cgi", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.AddCookie(&http.Cookie{Name: "sessionID", Value: sessionId})

	// Execute the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// Validate status code
	if res.StatusCode != 200 {
		return "", errors.New(res.Status)
	}

	// Store sessionId and return a successful response
	h.sessionId = sessionId
	h.IsValid = true
	return sessionId, nil
}
