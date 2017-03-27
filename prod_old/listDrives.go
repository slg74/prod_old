package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func getDriveLocations() []string {
	cmd := "storcli /c0/eall/sall show all|awk '/Device attributes/ {print $2}'"
	res, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Sprintf("Failed to exec command: %s", cmd)
	}
	return strings.Fields(fmt.Sprintf("%s", res))
}

func getDriveSerialNums() []string {
	cmd := "storcli /c0/eall/sall show all|awk '/SN =/ {print $3}'"
	res, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Sprintf("Failed to exec command: %s", cmd)
	}
	return strings.Fields(fmt.Sprintf("%s", res))
}

func main() {
	locations := getDriveLocations()
	serials := getDriveSerialNums()

	for i := range locations {
		fmt.Printf("%s\t%s\n", locations[i], serials[i])
	}
}
