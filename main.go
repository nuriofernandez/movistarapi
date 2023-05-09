package main

import (
	"fmt"
	"github.com/xXNurioXx/movistar-router-api/integration/mapalocal"
)

func main() {
	fmt.Println("HEY!")
	devices, err := mapalocal.GetLocalDevices("199970505")
	if err != nil {
		return
	}

	fmt.Println(devices)
}
