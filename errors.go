package main

import "errors"

var (
	ErrFailedToGetNetworkInterfaces     = errors.New("failed to get network interfaces")
	ErrFailedToGetAddressesForInterface = errors.New("failed to get addresses for interface")
	ErrFailedToGetIPAddress             = errors.New("failed to get IP address")
	ErrPingFailed                       = errors.New("ping failed")
)
