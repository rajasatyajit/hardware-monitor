# Hardware Monitor

A lightweight, cross-platform system monitoring tool for the command line, written in Go.

## Features

- **Real-time monitoring** of key hardware parameters
- **Cross-platform** support (Linux, macOS, Windows)
- **Color-coded output** for easy reading
- **Minimal dependencies** for lightweight operation
- **Refreshes in-place** without scrolling

### Monitored Parameters

- **CPU Usage**: Total and per-core usage percentages
- **Memory Usage**: RAM usage with used/total and percentage
- **Disk Usage**: Usage statistics for all mounted drives
- **Network Usage**: Upload/download rates for all network interfaces
- **Battery Status**: Charge percentage and charging state (where available)
- **CPU Temperature**: Current CPU temperature (where available)

## Installation

### Prerequisites

- Go 1.21 or higher

### Build from Source

```bash
# Clone the repository
git clone https://github.com/rajasatyajit/hardware-monitor.git
cd hardware-monitor

# Download dependencies
go mod download

# Build the application
go build -o hardware-monitor

# Run the application
./hardware-monitor