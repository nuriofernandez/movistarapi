package movistarapi

import (
	"fmt"
	"os"
	"testing"

	"github.com/nuriofernandez/movistarapi/hgu"
)

var routerPass = os.Getenv("MOVISTAR_ROUTER_PASS")

func TestEnsureSyntax(t *testing.T) {
	hguRouter, err := HGULogin(routerPass)
	if err != nil {
		t.Fatal(err)
	}

	localMap, err := hguRouter.LocalMap()
	if err != nil {
		return
	}

	fmt.Println(localMap)
}

func TestHGUSession_OpenPort(t *testing.T) {
	hguRouter, err := HGULogin(routerPass)
	if err != nil {
		t.Fatal(err)
	}

	port := hgu.OpenPort{
		Name:              "test-23",
		Protocol:          hgu.TCP,
		Address:           "192.168.1.165",
		ExternalPortStart: 12315,
		ExternalPortEnd:   0,
		InternalPortStart: 12312,
		Enabled:           true,
		Interface:         "ppp0.1",
	}

	// Ensure is not already opened
	ports, err := hguRouter.OpenPorts()
	if err != nil {
		t.Fatal(err)
	}
	for _, existingPort := range ports {
		if existingPort.Name == port.Name {
			t.Fatalf("Port %s already exists, must be deleted manually", port.Name)
		}
	}

	// Open it!
	err = hguRouter.OpenPort(port)
	if err != nil {
		t.Fatal(err)
	}

	// List open ports
	ports, err = hguRouter.OpenPorts()
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

func TestHGUSession_UpdatePort(t *testing.T) {
	hguRouter, err := HGULogin(routerPass)
	if err != nil {
		t.Fatal(err)
	}

	// Add a testing port to manipulate it later.
	tmpRuleName := "goapi-test-up"
	err = hguRouter.OpenPort(hgu.OpenPort{
		Name:              tmpRuleName,
		Protocol:          hgu.TCP,
		Address:           "192.168.1.100",
		ExternalPortStart: 54142,
		ExternalPortEnd:   0,
		InternalPortStart: 54142,
		Enabled:           true,
		Interface:         "ppp0.1",
	})
	if err != nil {
		t.Fatal(err)
	}

	// Fetch rules and get the just added rule
	ports, err := hguRouter.OpenPorts()
	if err != nil {
		t.Fatal(err)
	}
	var openPort hgu.OpenPort
	for _, existingPort := range ports {
		if existingPort.Name == tmpRuleName {
			openPort = existingPort
		}
	}
	if openPort == (hgu.OpenPort{}) {
		t.Fatalf("Open port not found")
	}

	// Modify the rule and send it
	openPort.Enabled = false
	openPort.Name = tmpRuleName + "x"
	err = hguRouter.UpdatePort(openPort)
	if err != nil {
		t.Fatal(err)
		return
	}

	// Fetch rules and get the just UPDATED rule
	ports, err = hguRouter.OpenPorts()
	if err != nil {
		t.Fatal(err)
	}
	var openPortSecond hgu.OpenPort
	for _, existingPort := range ports {
		if existingPort.Id == openPort.Id {
			openPortSecond = existingPort
		}
	}
	if openPortSecond == (hgu.OpenPort{}) {
		t.Fatalf("Open port not found")
	}

	// Assertions

	if openPortSecond.Enabled {
		t.Fatalf("Open port should be disabled, was not")
	}

	if openPortSecond.Name != tmpRuleName+"x" {
		t.Fatalf("Open port name should be updated")
	}

	// TODO Delete the port
}
