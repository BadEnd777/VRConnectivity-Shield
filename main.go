package main

import (
	"os"

	"github.com/Onyx-Innovators/VRConnectivity-Shield/logger"
	"github.com/Onyx-Innovators/VRConnectivity-Shield/utils"
	"github.com/miekg/dns"
)

// main is the entry point of the program
func main() {
	log := logger.GetLogger() // Get the logger

	port := "53"

	// Set up the DNS server
	dns.HandleFunc(".", utils.HandleDNSRequest)
	// Start the DNS server
	go func() {
		err := dns.ListenAndServe(":"+port, "udp", nil) // Listen on UDP port 53
		if err != nil {
			log.Error("Failed to set udp listener: ", err)
			os.Exit(1)
		}
	}()

	log.Info("DNS blocker is running on port ", port)

	select {} // Block forever
}
