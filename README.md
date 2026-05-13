> ⚠️ _Currently under development, expect breaking changes in the future._

# Movistar unofficial router api

This library allows to manipulate the movistar router from GO code.

- Tested on Askey RTF3505VW

<img width="836" height="393" alt="image" src="https://github.com/user-attachments/assets/11ac9c94-ef63-41e7-b467-4d841def9bb1" />

# Implemented features

- movistarapi.HGULogin(routerPassword string) (*HGUSession, error)
- HGUSession#Restart() (error)
- HGUSession#LocalMap() ([]ConnectedDevice, error)
- HGUSession#OpenPorts() ([]OpenPort, error)
- HGUSession#OpenPort(OpenPort) (error)
- HGUSession#UpdatePort(OpenPort) (error)
- HGUSession#DeletePort(int, string) (error)
## Example usage to open a port (creating a NAT rule)

```go
hguRouter, err := movistarapi.HGULogin(routerPass)
if err != nil {
	fmt.Println("invalid pass")
    return
}
err = hguRouter.OpenPort(hgu.OpenPort{
		Name:              "rule-name",
		Protocol:          hgu.TCP, // TCP/UDP/BOTH
		Address:           "192.168.1.100",
		ExternalPortStart: 80,
		ExternalPortEnd:   0, // optional
		InternalPortStart: 80,
		Enabled:           true,
		Interface:         "ppp0.1",
})
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println("Port was open!")
```

## Example usage to update a NAT rule

```go
hguRouter, err := movistarapi.HGULogin(routerPass)
if err != nil {
	fmt.Println("invalid pass")
    return
}

// Fetch rules and search a rule we want to update by name
ports, err := hguRouter.OpenPorts()
if err != nil {
    fmt.Println(err)
    return
}
var openPort hgu.OpenPort
for _, existingPort := range ports {
    if existingPort.Name == "RULE-NAME" {
        openPort = existingPort
    }
}
if openPort == (hgu.OpenPort{}) {
    fmt.Println(err)
    return
}

// Modify the rule
openPort.Enabled = false
openPort.Name = "NEW-RULE-NAME"

// Update the rule with HGUSession#UpdatePort(OpenPort)
err = hguRouter.UpdatePort(openPort)
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println("NAT rule was disabled and renamed!")
```

## Example usage to delete a NAT rule

```go
hguRouter, err := movistarapi.HGULogin(routerPass)
if err != nil {
	fmt.Println("invalid pass")
    return
}

// Fetch rules and search a rule we want to delete by name
ports, err := hguRouter.OpenPorts()
if err != nil {
    fmt.Println(err)
    return
}
var openPort hgu.OpenPort
for _, existingPort := range ports {
    if existingPort.Name == "RULE-NAME" {
        openPort = existingPort
    }
}
if openPort == (hgu.OpenPort{}) {
    fmt.Println(err)
    return
}

// Delete the rule with HGUSession#DeletePort(int, string)
err = hguRouter.DeletePort(openPort.Id, openPort.Interface)
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println("NAT rule was deleted!")
```

## Example usage as a CLI:

```go
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
	hguRouter, err := movistarapi.HGULogin(password)
	if err != nil {
		fmt.Println("Unable to login, " + err.Error())
		return
	}
	fmt.Println("Successfully logged in")

	// Prepare to restart router
	fmt.Println("Restarting Movistar HGU router...")
	err = hguRouter.Reboot()
	if err != nil {
		return
	}

	fmt.Println("Gone!")
}
```

## Example output

```bash
$ movistarcli routerpass123

Welcome to the Unofficial Movistar router control CLI!
Logging in...
Successfully logged in
Restarting Movistar HGU router...
Gone!
```
