package main

import (
	"errors"
	"fmt"
	"net"
	"os"
)

type commander struct{}

func NewCommander() Commander {
	return &commander{}
}

func (c *commander) GetSystemInfo() (SystemInfo, error) {
	var (
		out SystemInfo
		err error
	)

	out.Hostname, err = os.Hostname()
	if err != nil {
		return out, err
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		return out, errors.Join(
			ErrFailedToGetNetworkInterfaces,
			err,
		)
	}

	for _, iface := range interfaces {
		// Ignore down interfaces and loopback interfaces
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addresses, err := iface.Addrs()
		if err != nil {
			return out, errors.Join(
				fmt.Errorf("interface '%s' failure", iface.Name),
				ErrFailedToGetAddressesForInterface,
				err,
			)
		}

		for _, addr := range addresses {
			// Check if it’s an IP address (IPv4 or IPv6)
			ip, ok := addr.(*net.IPNet)
			if !ok || ip.IP.IsLoopback() {
				continue
			}

			// Return the first non-loopback IPv4 address
			if ip.IP.To4() != nil {
				out.IPAddress = ip.IP.String()
				break
			}
		}
	}

	if out.IPAddress == "" {
		err = ErrFailedToGetIPAddress
	}

	return out, err
}
