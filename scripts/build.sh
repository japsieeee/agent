#!/bin/bash
# scripts/build.sh
# Build the Go agent binary with versioning

set -e  # exit immediately if a command fails
set -o pipefail

APP_NAME="agent"
BUILD_DIR="../build"
CONFIG_DIR="../configs"
VERSION=${1:-"v0.1.0"}   # pass version as first argument, default v0.1.0

echo "Starting build for $APP_NAME version $VERSION..."

# Ensure script is run from the 'scripts' directory
SCRIPT_DIR=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)
PROJECT_ROOT="$SCRIPT_DIR/.."

# Clean previous builds
echo "Cleaning previous builds..."
rm -rf "$PROJECT_ROOT/build"
mkdir -p "$PROJECT_ROOT/build"

# Build binary
echo "Building $APP_NAME binary..."
GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$VERSION" -o "$PROJECT_ROOT/build/$APP_NAME" "$PROJECT_ROOT/cmd/agent"

# Copy configs
echo "Copying configs..."
cp -r "$PROJECT_ROOT/configs" "$PROJECT_ROOT/build/"

echo "Build complete. Binary and configs are in $PROJECT_ROOT/build"