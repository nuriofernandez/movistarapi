> _Currently under development, expect breaking changes in the future._

# Movistar unofficial router api

This library allows to manipulate the movistar router from GO code.

- Tested on Askey RTF3505VW

<img width="836" height="393" alt="image" src="https://github.com/user-attachments/assets/11ac9c94-ef63-41e7-b467-4d841def9bb1" />

# Implemented features

- Login(str)
- Restart()
- GetLocalDevices()

## Example CLI:

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
```

## Example output

```bash
$ movistarcli routerpass123

Welcome to the Unofficial Movistar router control CLI!
Logging in...
Successfully logged in
Retrieving 'Local Map'...
[['1','CoolDeviceName','0','192.168.1.XX','WIFI','no','XX:XX:XX:XX:XX:XX']]
```
