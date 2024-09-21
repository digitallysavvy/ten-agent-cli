package container

import (
    "os"
)

func IsInsideContainer() bool {
    _, err := os.Stat("/.dockerenv")
    return err == nil
}