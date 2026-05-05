package movistarapi

import (
	"fmt"
	"os"
	"testing"

	hgu2 "github.com/nuriofernandez/movistarapi/hgu"
)

var routerPass = os.Getenv("MOVISTAR_ROUTER_PASS")

func TestEnsureSyntax(t *testing.T) {
	hgu, err := HGULogin(routerPass)
	if err != nil {
		t.Fatal(err)
	}

	localMap, err := hgu.LocalMap()
	if err != nil {
		return
	}

	fmt.Println(localMap)
}

func TestHGUSession_OpenPort(t *testing.T) {
	hgu, err := HGULogin(routerPass)
	if err != nil {
		t.Fatal(err)
	}

	port := hgu2.OpenPort{
		Name:              "test-23",
		Protocol:          hgu2.TCP,
		Address:           "192.168.1.165",
		ExternalPortStart: 12315,
		ExternalPortEnd:   0,
		InternalPortStart: 12312,
		Enabled:           true,
		Interface:         "ppp0.1",
	}

	// Ensure is not already opened
	ports, err := hgu.OpenPorts()
	if err != nil {
		t.Fatal(err)
	}
	for _, existingPort := range ports {
		if existingPort.Name == port.Name {
			t.Fatalf("Port %s already exists, must be deleted manually", port.Name)
		}
	}

	// Open it!
	err = hgu.OpenPort(port)
	if err != nil {
		t.Fatal(err)
	}

	// List open ports
	ports, err = hgu.OpenPorts()
	if err != nil {
		t.Fatal(err)
	}

	// Check if new port is present
	found := false
	for _, existingPort := range ports {
		if existingPort.Name == port.Name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("After opening the port, was not found on the list of open ports")
	}
}
