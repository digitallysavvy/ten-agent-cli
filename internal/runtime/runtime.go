package runtime

import (
    "fmt"
    "os/exec"
)

func RunAgent() error {
    // TODO: Implement agent running logic
    cmd := exec.Command("docker", "run", "ten-agent")
    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("failed to run agent: %w", err)
    }
    return nil
}