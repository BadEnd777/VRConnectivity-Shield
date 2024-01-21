package main

import (
	"fmt"
	"os/exec"
	"strings"
)

var hostsToAdd = []string{
	"api.amplitude.com",
	"api2.amplitude.com",
	"api.lab.amplitude.com",
	"api.eu.amplitude.com",
	"regionconfig.amplitude.com",
	"regionconfig.eu.amplitude.com",
	"o1125869.ingest.sentry.io",
	"api3.amplitude.com",
	"cdn.amplitude.com",
	"info.amplitude.com",
	"static.amplitude.com",
	"api.uca.cloud.unity3d.com",
	"config.uca.cloud.unity3d.com",
	"perf-events.cloud.unity3d.com",
	"public.cloud.unity3d.com",
	"cdp.cloud.unity3d.com",
	"data-optout-service.uca.cloud.unity3d.com",
	"ecommerce.iap.unity3d.com",
}

func main() {
	// Get logger
	log := GetLogger()

	const hostsFilePath = "C:\\Windows\\System32\\drivers\\etc\\hosts" // C:\Windows\System32\drivers\etc\hosts
	existingHosts, err := readHostsFile(hostsFilePath)
	if err != nil {
		log.Error("Error reading hosts file in readHostsFile:", err)
		return
	}

	backups, err := createBackups(hostsFilePath)
	if err != nil {
		log.Error(fmt.Sprintf("Error creating backups: %s", err))
	} else {
		log.Success(fmt.Sprintf("Successfully created backups: %s", strings.Join(backups, ", ")))
	}

	for _, host := range hostsToAdd {
		if contains(existingHosts, host) {
			log.Warn(fmt.Sprintf("Host %s already exists in hosts file", host))
		} else {
			err := addHostToHostsFile(hostsFilePath, host)
			if err != nil {
				log.Error(fmt.Sprintf("Error adding host %s to hosts file: %s", host, err))
			} else {
				log.Success(fmt.Sprintf("Successfully added host %s to hosts file", host))
			}
		}
	}

	// Flush DNS cache
	err = flushDNSCache()
	if err != nil {
		log.Error(fmt.Sprintf("Error flushing DNS cache: %s", err))
	} else {
		log.Success("Successfully flushed DNS cache")
	}

	comfirmation := log.Prompt("Do you want to test the connection? (y/n): ")
	for comfirmation != "y" && comfirmation != "n" {
		comfirmation = log.Prompt("Please enter y or n: ")
	}

	if comfirmation == "y" {
		// Test connection
		log.Info("Testing connection...")
		for _, host := range hostsToAdd {
			cmd := exec.Command("ping", host, "-n", "1", "-w", "1000") // Ping host once with 1 second timeout
			err := cmd.Run()
			if err != nil {
				log.Error(fmt.Sprintf("Error pinging host %s: %s", host, err))
			} else {
				log.Success(fmt.Sprintf("Successfully pinged host %s", host))
			}
		}

		log.Info("This program will block the a VRChat analytics and telemetry servers.")
		log.Info("If it works, you should see a lot of \"Error pinging host\" messages.")
		log.Info("Because the program blocks the servers, it will not be able to connect to them.")
	}
}
