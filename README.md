# VRConnectivity Shield

<div align="center">
    <img src="https://goreportcard.com/badge/github.com/Onyx-Innovators/VRConnectivity-Shield" alt="Go Report Card">
    <img src="https://img.shields.io/github/license/Onyx-Innovators/VRConnectivity-Shield" alt="License">
    <img src="https://img.shields.io/github/v/release/Onyx-Innovators/VRConnectivity-Shield" alt="Release">
</div>

VRConnectivity Shield is a DNS blocking utility implemented in Go, designed to block specific hosts by responding to DNS requests with 0.0.0.0. This can be useful for preventing unwanted connections to analytics services, telemetry servers, etc.

> [!NOTE]
> This project is focused on blocking hosts related to VRChat. If you are looking for a more general purpose DNS blocking utility, check out [HostsMan](https://www.abelhadigital.com/hostsman/). It is a great tool for managing your hosts file.

## Usage

### Prerequisites

- [Go](https://golang.org/doc/install) installed on your system.

### Installation

To use VRConnectivity Shield, follow these simple steps:

1. **Download:** Obtain the executable file for your platform from the [Releases](https://github.com/Onyx-Innovators/VRConnectivity-Shield/releases) page.
   
2. **Run:** Execute the VRConnectivity-Shield.exe executable to start the DNS blocking utility.

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
