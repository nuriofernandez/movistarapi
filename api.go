package movistarapi

import "github.com/xXNurioXx/movistarapi/hgu"

// HGULogin generates a session to operate the HGU router
func HGULogin(password string) (hgu.HGUSession, error) {
	router := hgu.HGUSession{}

	_, err := router.Login(password)
	if err != nil {
		return router, err
	}

	return router, nil
}
