package main

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/magefile/mage/sh"
)

// Ensures dependencies are up to date.
func Vendor() {
	log.Info("Dependency management", "state", "starting")
	started := time.Now()
	sh.Run("go", "mod", "tidy")
	log.Info("Dependency management", "state", "complete", "elapsed", time.Since(started).String())
}
