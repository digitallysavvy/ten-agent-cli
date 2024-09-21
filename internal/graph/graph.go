package graph

import (
    "fmt"
    "os/exec"
)

func LaunchEditor() error {
    // TODO: Implement Graph Editor launching logic
    cmd := exec.Command("open", "http://localhost:3000/graph-editor")
    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("failed to launch Graph Editor: %w", err)
    }
    return nil
}