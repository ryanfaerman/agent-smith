//go:build windows

// This file contains the Ping method for the commander struct. This
// method is only available on Windows.

package main

import (
	"errors"
	"os/exec"
	"time"
)

// Ping the given host and return the result. Even when there is an error
// the PingResult will be returned, with the Successful flag set to false.
//
// This is the Windows implementation.
func (c *commander) Ping(host string) (PingResult, error) {
	var out PingResult
	cmd := exec.Command("ping", "-n", "1", host)

	start := time.Now()
	err := cmd.Run()
	out.Time = time.Since(start)

	if err != nil {
		out.Successful = false
		return out, errors.Join(ErrPingFailed, err)
	}

	return out, nil
}
