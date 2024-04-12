# pritunl-api-go

Pritunl API Client for Go

> [!NOTE]
> The project is currently in development, so some features may be limited or unavailable at this time.

## API Usage

Load your Pritunl API Credentials in our Environment Variables.

```bash
export PRITUNL_BASE_URL="https://vpn.domain.tld/"
export PRITUNL_API_TOKEN="<PRITUNL API TOKEN>"
export PRITUNL_API_SECRET="<PRITUNL API SECRET>"
```

Get the Pritunl API Client for Go Package/Library

```bash
go get github.com/nathanielvarona/pritunl-api-go
```

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nathanielvarona/pritunl-api-go"
)

func main() {
	/* INITIALIZATION */
	// Provide authentication credentials as needed for client creation
	// Automaticlly sets from environment variables if present
	client, err := pritunl.NewClient()

	// You can also initialize an instance by manually providing the arguments.
	// client := pritunl.NewClient(&pritunl.Client{
	// 	BaseUrl:   "<PRITUNL API URL>",
	// 	ApiToken:  "<PRITUNL API TOKEN>",
	// 	ApiSecret: "<PRITUNL API SECRET>",
	// })

	if err != nil {
		log.Fatal(err)
	}

	// Create a context for the request
	ctx := context.Background()

	// Retrieve the server status
	status, err := client.StatusGet(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Struct Output
	for _, stat := range status {
		fmt.Println("Server Version", stat.ServerVersion)
		fmt.Println("Local Networks:", stat.LocalNetworks)
		fmt.Println("Host Online:", stat.HostsOnline)
	}

	// JSON Output
	statusBytes, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		log.Println("Error marshalling status:", err)
	} else {
		fmt.Println(string(statusBytes))
	}
}
```

## Features

### Core Pritunl API Client

| Feature Function   | Description                             | Status                 |
|--------------------|-----------------------------------------|------------------------|
| StatusGet          | Status of Pritunl Server                | :white_check_mark: Yes |
| KeyGet             | Generate or Retrieve a Key for the User | :white_check_mark: Yes |
| UserGet            | Get the Information of Existing User    | :white_check_mark: Yes |
| UserCreate         | Create a New User                       | :white_check_mark: Yes |
| UserUpdate         | Update an Existing User                 | :white_check_mark: Yes |
| UserDelete         | Delete an User                          | :white_check_mark: Yes |
| OrganizationGet    | Get the Information of Existing Org     | :white_check_mark: Yes |
| OrganizationCreate | Create a New Org                        | :white_check_mark: Yes |
| OrganizationUpdate | Update an Existing Org                  | :white_check_mark: Yes |
| OrganizationDelete | Delete an Org                           | :white_check_mark: Yes |
| ServerGet          | Get the Information of Existing Server  | Not Yet                |
| ServerCreate       | Create a New Server                     | Not Yet                |
| ServerUpdate       | Update an Existing Server               | Not Yet                |
| ServerDelete       | Delete an Server                        | Not Yet                |

### Future Enhancements (CLI)

1. **CLI Framework:** Consider using a popular framework like Cobra (https://github.com/spf13/cobra) or Viper (https://github.com/google/go-github) to simplify the command structure, argument parsing, and flag handling.
2. **Build Distribution Workflow:** Implement a CI/CD workflow (e.g., using GitHub Actions) to automate building and distributing the CLI tool across various platforms (Windows, macOS, Linux) and architectures (32-bit, 64-bit). This will streamline setup for users on different systems.
