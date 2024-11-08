package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/log"
)

const (
	ApplicationName = "agent-smith"
)

func binaryName(goos string) string {
	switch goos {
	case "windows":
		return fmt.Sprintf("%s.exe", ApplicationName)
	default:
		return ApplicationName
	}
}

// exists returns whether the given file or directory exists or not
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func ensureDirs() error {
	var b strings.Builder

	dirs := []string{"./bin", "./release", "./tmp", "./release/darwin", "./release/windows"}

	for _, dir := range dirs {
		if !exists("./" + dir) {
			fmt.Fprintf(&b, "(+) %s\n", dir)
			if err := os.MkdirAll("./"+dir, 0755); err != nil {
				return err
			}
		} else {
			fmt.Fprintf(&b, "(✓) %s\n", dir)
		}
	}
	log.Info("Ensured output directories", "directories", b.String())
	return nil
}
