# VRConnectivity Shield

<div align="center">
    <img src="https://goreportcard.com/badge/github.com/Onyx-Innovators/VRConnectivity-Shield" alt="Go Report Card">
    <img src="https://img.shields.io/github/license/Onyx-Innovators/VRConnectivity-Shield" alt="License">
    <img src="https://img.shields.io/github/v/release/Onyx-Innovators/VRConnectivity-Shield" alt="Release">
</div>

VRConnectivity Shield is a simple DNS blocking utility for Windows. It is designed to block hosts related to VRChat analytics. This utility is intended to be used in conjunction with a VPN to prevent VRChat from collecting analytics data while playing the game.

## Usage

### Prerequisites

- [Go](https://golang.org/doc/install) installed on your system.

### Installation

To use VRConnectivity Shield, follow these simple steps:

1. **Download:** Obtain the executable file for your platform from the [Releases](https://github.com/Onyx-Innovators/VRConnectivity-Shield/releases) page.

2. **Run:** Execute the VRConnectivity-Shield.exe executable to start the DNS blocking utility.

## Development

To build the project from source, follow these steps:

1. **Clone:** Clone the repository to your local machine.
    ```bash
    git clone https://github.com/Onyx-Innovators/VRConnectivity-Shield.git
    ```
2. **Navigate:** Navigate to the project directory.
    ```bash
    cd VRConnectivity-Shield/src
    ```

3. **Build:** Build the project using the Go compiler.
    ```bash
    go build -o VRConnectivity-Shield.exe
    ```
4. **Run:** Execute the VRConnectivity-Shield.exe executable to start the DNS blocking utility.

## Configuration

The list of blocked hosts is sourced from [VRChat-Analytics-Blocker](https://github.com/DubyaDude/VRChat-Analytics-Blocker). The blocked hosts are defined in the `main.go` map in the utils package. You can modify this map to include or exclude specific hosts as needed.

```go
var hostsToAdd = []string{
 "api.amplitude.com",
 "api2.amplitude.com",
 // Add specific hosts here
}
```

## How it works

The program will add the specified hosts to the hosts file located at `C:\Windows\System32\drivers\etc\hosts`. The hosts file is a text file that maps hostnames to IP addresses. When a DNS request is made for a hostname that is defined in the hosts file, the DNS resolver will use the IP address specified in the hosts file instead of querying a DNS server.

## License

This project is licensed under the [MIT License](LICENSE).

---
