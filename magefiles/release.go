package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/log"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Builds the binary and installers for all supported platforms.
func Release() error {
	log.Info("Release Build", "state", "starting")

	mg.Deps(Clean)
	mg.Deps(Vendor, ensureDirs)

	started := time.Now()

	{
		log.Info("Building for MacOS", "state", "starting", "goos", "darwin")
		started := time.Now()

		outputPath := filepath.Join(
			"./release", "darwin", "usr", "local", "bin",
			binaryName("darwin"),
		)
		err := sh.RunWith(
			map[string]string{"GOOS": "darwin"},
			"go", "build", "-o", outputPath,
		)
		if err != nil {
			return err
		}

		// Create the launch agent directory
		if err := os.MkdirAll("release/darwin/Library/LaunchAgents", 0755); err != nil {
			return err
		}

		if err := sh.Copy("release/darwin/Library/LaunchAgents/com.github.ryanfaerman.agent-smith.plist", "agent-smith.plist"); err != nil {
			return err
		}

		log.Info("Creating installer", "state", "starting", "goos", "darwin")
		err = sh.Run("pkgbuild",
			"--root", "release/darwin",
			"--identifier", "com.github.ryanfaerman.agent-smith",
			"--version", "0.1",
			"--install-location", "/",
			"release/agent-smith.pkg",
		)
		if err != nil {
			return err
		}
		log.Info("Creating installer", "state", "complete", "goos", "darwin")

		log.Info("Building for MacOS", "state", "complete", "goos", "darwin", "elapsed", time.Since(started).String())
	}

	{
		log.Info("Building for Windows", "state", "starting", "goos", "windows")
		started := time.Now()

		outputPath := filepath.Join(
			"./release", "windows",
			binaryName("windows"),
		)
		err := sh.RunWith(
			map[string]string{"GOOS": "windows"},
			"go", "build", "-o", outputPath,
		)
		if err != nil {
			return err
		}

		if err := sh.Copy("release/windows/installer.nsi", "installer.nsi"); err != nil {
			return err
		}

		log.Info("Creating installer", "state", "starting", "goos", "darwin")
		err = sh.Run("makensis",
			"release/windows/installer.nsi")
		if err != nil {
			return err
		}
		log.Info("Creating installer", "state", "complete", "goos", "darwin")

		log.Info("Building for Windows", "state", "complete", "goos", "windows", "elapsed", time.Since(started).String())
	}

	log.Info("Release Build",
		"state", "complete",
		"elapsed", time.Since(started).String(),
	)
	return nil
}
