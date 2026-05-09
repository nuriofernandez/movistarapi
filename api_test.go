package movistarapi

import (
	"os"
	"testing"

	"github.com/nuriofernandez/movistarapi/hgu"
)

var routerPass = os.Getenv("MOVISTAR_ROUTER_PASS")

func TestHGUSession_OpenPort(t *testing.T) {
	hguRouter, err := HGULogin(routerPass)
	if err != nil {
		t.Fatal(err)
	}

	tmpRuleName := "goapi-test-crea"
	port := hgu.OpenPort{
		Name:              tmpRuleName,
		Protocol:          hgu.TCP,
		Address:           "192.168.1.100",
		ExternalPortStart: 54142,
		ExternalPortEnd:   0,
		InternalPortStart: 54142,
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
	var newOpenPort hgu.OpenPort
	for _, existingPort := range ports {
		if existingPort.Name == port.Name {
			newOpenPort = existingPort
			break
		}
	}

	// Assert it was created
	if newOpenPort == (hgu.OpenPort{}) {
		t.Fatalf("After opening the port, was not found on the list of open ports")
	}

	// Delete the testing port
	err = hguRouter.DeletePort(newOpenPort.Id, newOpenPort.Interface)
	if err != nil {
		t.Fatal(err)
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

	// Delete the testing port
	err = hguRouter.DeletePort(openPort.Id, openPort.Interface)
	if err != nil {
		t.Fatal(err)
	}
}

func TestHGUSession_DeletePort(t *testing.T) {
	hguRouter, err := HGULogin(routerPass)
	if err != nil {
		t.Fatal(err)
	}

	// Add a testing port to manipulate it later.
	tmpRuleName := "goapi-test-del"
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

	err = hguRouter.DeletePort(openPort.Id, openPort.Interface)
	if err != nil {
		t.Fatal(err)
	}

	// Fetch rules and confirm the rule is gone
	ports, err = hguRouter.OpenPorts()
	if err != nil {
		t.Fatal(err)
	}
	var openPortAfterDelete hgu.OpenPort
	for _, existingPort := range ports {
		if existingPort.Name == tmpRuleName {
			openPortAfterDelete = existingPort
		}
	}

	// Assert it's gone

	if openPortAfterDelete != (hgu.OpenPort{}) {
		t.Fatalf("Open port was found! should be deleted")
	}
}
