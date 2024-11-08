package main

import (
	"path/filepath"
	"runtime"
	"time"

	"github.com/charmbracelet/log"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Builds the binary for the current platform.
func Build() error {
	log.Info("Building", "binary", ApplicationName, "os", runtime.GOOS)
	started := time.Now()
	mg.Deps(Vendor, ensureDirs)
	outputPath := filepath.Join("bin", binaryName(runtime.GOOS))
	if err := sh.Run("go", "build", "-o", outputPath); err != nil {
		return err
	}

	log.Info("Build complete",
		"state", "complete",
		"binary", outputPath,
		"elapsed", time.Since(started).String(),
	)

	return nil
}
