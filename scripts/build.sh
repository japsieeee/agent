#!/bin/bash
# scripts/build.sh

set -e

APP_NAME="o8s-agent"
BUILD_DIR="../build"
CONFIG_DIR="../configs"
VERSION=${1:-"v0.1.0"}   # pass version as first argument, default v0.1.0

echo "Cleaning previous builds..."
rm -rf $BUILD_DIR
mkdir -p $BUILD_DIR

echo "Building $APP_NAME binary with version $VERSION..."
# Use -ldflags to embed version into binary
GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$VERSION" -o $BUILD_DIR/$APP_NAME ./cmd/agent

echo "Copying configs..."
cp -r $CONFIG_DIR $BUILD_DIR/

echo "Build complete. Binary and configs in $BUILD_DIR"