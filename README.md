# esphome_api

A Go library to manage [ESPHome](https://esphome.io/) devices.

This Go library provides a client implementation for interacting with ESPHome
devices using the native ESPHome API. It enables developers to control and
monitor ESPHome devices programmatically from Go applications. The library
offers functionalities for establishing connections, authenticating, sending
commands, subscribing to state updates, and receiving device information.

## Installation

Library:

```bash
go get github.com/mycontroller-org/esphome_api@latest
```

CLI (`esphomectl`):

```bash
go build -trimpath -ldflags "-s -w" -o esphomectl ./cli/main.go
```

Release binaries (multi-platform, version ldflags):

```bash
./scripts/generate_executables.sh
```

## Usage

### 1. Import the Library

```go
import (
        "fmt"
        "log"
        "time"

        "github.com/mycontroller-org/esphome_api/pkg/api"
        "github.com/mycontroller-org/esphome_api/pkg/client"
        "google.golang.org/protobuf/proto"
)
```

### 2. Create a Client

```go
// Replace with your ESPHome device's address and encryption key (if applicable)
address := "esphome.local:6053"
encryptionKey := "YOUR_ENCRYPTION_KEY"

// Create a new ESPHome API client
client, err := client.GetClient("my-client-id", address, encryptionKey, 10*time.Second, handleFunc)
if err != nil {
        log.Fatalln(err)
}
defer client.Close()
```

### 3. Hello

```go
// Required after GetClient on modern devices.
if _, err := client.Hello(); err != nil {
        log.Fatalln(err)
}

// Deprecated (pre-2026.1 password auth only):
// if err := client.Login("YOUR_PASSWORD"); err != nil { ... }
```

See [docs/API_PROTO_UPDATE_2024.5.0_to_2026.7.0.md](docs/API_PROTO_UPDATE_2024.5.0_to_2026.7.0.md) for protocol changes and breaking changes.

### 4. Send Commands and Receive State Updates

```go
// Example: Subscribe to state updates
if err := client.SubscribeStates(); err != nil {
        log.Fatalln(err)
}

// Example: Send a command to a light entity
lightKey := uint32(12) // Replace with the key of your light entity
if err := client.Send(&api.LightCommandRequest{
        Key:   lightKey,
        State: true, // Turn the light on
}); err != nil {
        log.Fatalln(err)
}
```

### 5. Handle Incoming Messages

```go
// Define a handler function to process incoming messages
func handleFunc(msg proto.Message) {
        switch msg := msg.(type) {
        case *api.LightStateResponse:
                fmt.Printf("Light state update: Key=%d, State=%t\n", msg.Key, msg.State)

        case *api.BinarySensorStateResponse:
                fmt.Printf("Binary sensor state update: Key=%d, State=%t\n", msg.Key, msg.State)

        // Handle other message types as needed
        default:
                fmt.Printf("Received message of type: %T\n", msg)
        }
}
```

## Configuration Example

The following is a complete configuration example for the `esphomectl` CLI
tool, which utilizes the `esphome_api` library:

**File: `~/.esphomectl.yaml`**

```yaml
active: esphome.local:6053
devices:
    - address: esphome.local:6053
      encryptionKey: YOUR_ENCRYPTION_KEY
      timeout: 10s
      info:
          name: My ESPHome Device
          model: NodeMCU
          macAddress: AC:BC:32:89:0E:A9
          esphomeVersion: "2026.7.0"
          compilationTime: "2026-01-15T10:00:00"
          usesPassword: false
          hasDeepSleep: false
          apiEncryptionSupported: true
          statusOn: 2026-07-18T12:00:00+05:30
```

**Note:** Prefer `encryptionKey`. API password auth was removed in ESPHome 2026.1.0.
If you still need a password for older firmware, store it Base64-encoded:

```bash
echo -n "YOUR_PASSWORD" | base64
```

## Examples

The [examples](/examples/) directory contains various examples demonstrating the usage of the `esphome_api` library for different purposes, such as:

-   **camera:** Capture images from an ESPHome camera.
-   **device_info:** Retrieve device information from an ESPHome device.
-   **tail_logs:** Subscribe to and display logs from an ESPHome device.
