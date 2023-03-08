package machineid

import (
	"fmt"
	"runtime"
	"strings"
)

// ID returns the platform specific machine id of the current host OS.
// Regard the returned id as "confidential" and consider using ProtectedID() instead.
func ID() (id string, err error) {
	switch runtime.GOOS {
	case "linux":
		id, err = machineIdLinux()
	case "darwin":
		id, err = machineIdDarwin()
	}

	if err != nil {
		return "", fmt.Errorf("machineid: %v", err)
	}
	return strings.ToLower(id), nil
}
