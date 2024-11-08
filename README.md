# Agent Smith

Agent-Smith is a demo agent that performs some simple commands.

## Building

This application is built using [mage](https://magefile.org/). To build the application, run the following command:

```bash
go run mage.go build
```

if you already have mage installed, you can run the following command:

```bash
mage build
````

Throughout this document, we will use the `mage` command to run the magefile, but you can 
also use `go run mage.go` if you prefer.

Running `mage` will list all the available targets:
```bash
$ mage

Targets:
  build      Builds the binary for the current platform.
  clean      Cleans up build artifacts.
  release    Builds the binary and installers for all supported platforms.
  vendor     Ensures dependencies are up to date.
```
### External Dependencies
This project assumes that you're running on a Mac and have the following tools installed:

- [golang](https://golang.org/)
- makensis (`brew install makensis`)
- [mage](https://magefile.org/) (optional)

## Release

This application supports building installers for MacOS and Windows. 

Note: I have not tested the Windows installer as I only have a Mac, but it should work.

These can be created with the `mage release` command. Once complete, you will find the installers 
in the `release` directory along with their artifacts.

```bash 
$ tree release
release
├── agent-smith.pkg
├── darwin
│   ├── Library
│   │   └── LaunchAgents
│   │       └── com.github.ryanfaerman.agent-smith.plist
│   └── usr
│       └── local
│           └── bin
│               └── agent-smith
└── windows
    ├── agent-smith.exe
    ├── agent_smith_installer.exe
    └── installer.nsi
```

## Installation

This can be installed by running `mage release` and running one of the installers. If you'd prefer 
not to install it locally (for testing) run `mage build` and then it can be executed directly from the `bin` directory.
  
## Testing

Tests are executed in the standard golang manner, with `go test ./...`. There are currently tests for all major 
functions except the build code.

## Usage
After installation or running directly with `mage build && ./bin/agent-smith`, the server should be running on port `8080`. 

From there, you can interact with the server using the `execute` endpoint. 
```bash
$ curl -XPOST http://localhost:8080/execute -d '{"type": "ping", "payload": "google.com"}'
{"success":true,"data":{"successful":false,"time":56136050}}

$ curl -XPOST http://localhost:8080/execute -d '{"type": "sysinfo"}'
{"success":true,"data":{"hostname":"rfnewbookpro.local","ip_address":"100.86.23.26"}}
```

The following commands are supported:

- `ping` - Pings a host and returns the time it took and if it was successful.
- `sysinfo` - Returns the hostname and IP address of the machine.

The payload of any request is the following:

```json
{
    "type": "ping|sysinfo",
    "payload": "google.com"
}
```

The `payload` field is ignored for the `sysinfo` command.

You can see this in action below:

[![asciicast](https://asciinema.org/a/tK3VweSO7y43shWPCqai1DR7N.svg)](https://asciinema.org/a/tK3VweSO7y43shWPCqai1DR7N)
