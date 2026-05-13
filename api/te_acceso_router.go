package api

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

func Login(password string) (string, error) {
	// Generate a random sessionId
	minLimit := int(math.Pow10(9))
	maxLimit := int(math.Pow10(9 - 1))
	randInt := int(rand.Float64() * float64(minLimit))
	if randInt < maxLimit {
		randInt += maxLimit
	}
	sessionId := fmt.Sprint(randInt)

	// Prepare login form
	data := url.Values{}
	data.Set("loginPassword", password)

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
		return "", fmt.Errorf("unsuccessful response from Movistar: %s\n", res.Status)
	}

	return sessionId, nil
}
