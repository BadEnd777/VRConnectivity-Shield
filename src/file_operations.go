package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func createBackups(filePath string) ([]string, error) {
	var backups []string
	for i := 0; ; i++ {
		backupPath := fmt.Sprintf("%s.backup-%d", filePath, i)
		if _, err := os.Stat(backupPath); os.IsNotExist(err) {
			err := exec.Command("cmd", "/C", "copy", filePath, backupPath).Run()
			if err != nil {
				return nil, err
			}
			backups = append(backups, backupPath)
			break
		}
	}
	return backups, nil
}

func readHostsFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var hosts []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 && !strings.HasPrefix(line, "#") {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				hosts = append(hosts, fields[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return hosts, nil
}

func addHostToHostsFile(filePath, host string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend) // Open file in append mode
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("0.0.0.0 %s\n", host)); err != nil {
		return err
	}

	return nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func flushDNSCache() error {
	cmd := exec.Command("ipconfig", "/flushdns")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
