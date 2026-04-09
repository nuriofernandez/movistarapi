package main

import (
	"fmt"
	"os"

	"github.com/xXNurioXx/movistarapi/integration/login"
	"github.com/xXNurioXx/movistarapi/integration/mapalocal"
)

func main() {
	fmt.Println("Welcome to the Unofficial Movistar router control CLI!")

	// Ensure the user actually typed something
	if len(os.Args) < 2 {
		fmt.Println("Invalid command")
		fmt.Println("Usage: binary <router-password>")
		return
	}

	// Store router password
	password := os.Args[1]

	// Prepare login
	fmt.Println("Logging in...")
	sessionId, err := login.Login(password)
	if err != nil {
		fmt.Println("Unable to login, " + err.Error())
		return
	}
	fmt.Println("Successfully logged in")

	// Prepare local map retrieve
	fmt.Println("Retrieving 'Local Map'...")
	devices, err := mapalocal.GetLocalDevices(sessionId)
	if err != nil {
		return
	}

	fmt.Println(devices)
}
