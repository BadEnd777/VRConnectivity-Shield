# VRConnectivity Shield

VRConnectivity Shield is a DNS blocking utility implemented in Go, designed to block specific hosts by responding to DNS requests with 0.0.0.0. This can be useful for preventing unwanted connections to analytics services, telemetry servers, etc.

## Usage

### Prerequisites

- [Go](https://golang.org/doc/install) installed on your system.

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/Onyx-Innovators/VRConnectivity-Shield.git
    ```

2. Change into the project directory:

    ```bash
    cd VRConnectivity-Shield
    ```

3. Build the project:

    ```bash
    go build
    ```

4. Run the program as root:
   - Double-click the executable file, or
   - Run the executable from the command line:

    ```bash
    sudo ./VRConnectivity-Shield
    ```

### Configuration

The list of blocked hosts is sourced from [VRChat-Analytics-Blocker](https://github.com/DubyaDude/VRChat-Analytics-Blocker). The blocked hosts are defined in the utils/blockedHosts map in the utils package. You can modify this map to include or exclude specific hosts as needed.

```go
var blockedHosts = map[string]struct{}{
    "api.amplitude.com":             {},
    "api2.amplitude.com":            {},
    // Add or remove hosts as necessary
}
```

## How it works

The program sets up a DNS server listening on UDP port 53. When a DNS request is received, it checks if the requested host is in the list of blocked hosts. If it is, the request is blocked by responding with 0.0.0.0. Otherwise, the request is allowed to proceed.

## License

This project is licensed under the [MIT License](LICENSE).

---
