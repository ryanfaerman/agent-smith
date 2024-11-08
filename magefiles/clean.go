package main

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

// Cleans up build artifacts.
func Clean() error {
	log.Info("Cleaning", "state", "starting")
	started := time.Now()
	if err := os.RemoveAll("./bin"); err != nil {
		return err
	}
	if err := os.RemoveAll("./release"); err != nil {
		return err
	}
	log.Info("Cleaning", "state", "complete", "elapsed", time.Since(started).String())
	return nil
}
