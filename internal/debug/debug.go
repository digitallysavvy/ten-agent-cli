package debug

import (
    "fmt"
    "os/exec"
)

func ViewLogs() error {
    // TODO: Implement log viewing logic
    cmd := exec.Command("docker", "logs", "ten-agent")
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("failed to view logs: %w", err)
    }
    fmt.Println(string(output))
    return nil
}