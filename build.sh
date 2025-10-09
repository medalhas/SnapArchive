#!/bin/bash

# Local build script for testing cross-platform builds
set -e

echo "Building SnapArchive for multiple platforms..."

# Clean previous builds
rm -f snaparchive-*

# Build for Linux
echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o snaparchive-linux-amd64 main.go
GOOS=linux GOARCH=arm64 go build -o snaparchive-linux-arm64 main.go

# Build for macOS
echo "Building for macOS..."
GOOS=darwin GOARCH=amd64 go build -o snaparchive-darwin-amd64 main.go
GOOS=darwin GOARCH=arm64 go build -o snaparchive-darwin-arm64 main.go

# Build for Windows
echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o snaparchive-windows-amd64.exe main.go
GOOS=windows GOARCH=arm64 go build -o snaparchive-windows-arm64.exe main.go

echo "Build completed! Binaries:"
ls -la snaparchive-*

echo ""
echo "To create a release:"
echo "1. git tag v1.0.0"
echo "2. git push origin v1.0.0"
echo "3. GitHub Action will automatically create the release with archives"