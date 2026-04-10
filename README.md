> _Currently under development, expect breaking changes in the future._

# Movistar unofficial router api

This library allows to manipulate the movistar router from GO code.

- Tested on Askey RTF3505VW

<img width="836" height="393" alt="image" src="https://github.com/user-attachments/assets/11ac9c94-ef63-41e7-b467-4d841def9bb1" />

# Implemented features

- movistarapi.HGULogin(routerPassword string) (*HGUSession, error)
- HGUSession#Restart() (error)
- HGUSession#LocalMap() (string, error)
- HGUSession#OpenPorts() (string, error)

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
	hgu, err := movistarapi.HGULogin("password123")
	if err != nil {
		fmt.Println("Unable to login, " + err.Error())
		return
	}
	fmt.Println("Successfully logged in")

	// Prepare to restart router
	fmt.Println("Restarting Movistar HGU router...")
	devices, err := hgu.Restart()
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
