package executor

import (
	"os/exec"
	"fmt"
	"time"
)

func ExecuteCommand(command string) (string, time.Duration, error) {
	startTime := time.Now()
	cmd := exec.Command("cmd", "/c", command)
	err := cmd.Run()
	if err != nil {
		return "", time.Since(startTime), fmt.Errorf("failed to execute command: %w", err)
	}
	return "Command executed successfully", time.Since(startTime), nil
}