#!/bin/bash

echo "Building ブロック崩し (Block Breaking Game)..."

# Build console demo (works without graphics libraries)
echo "Building console demo..."
go build -o console_demo console_demo.go
if [ $? -eq 0 ]; then
    echo "✓ Console demo built successfully: ./console_demo"
else
    echo "✗ Console demo build failed"
    exit 1
fi

# Try to build graphics version
echo "Building graphics version..."
go build -o blockbreaker main.go
if [ $? -eq 0 ]; then
    echo "✓ Graphics version built successfully: ./blockbreaker"
    echo ""
    echo "Run './blockbreaker' for the full graphics experience"
    echo "Run './console_demo' for a simple console version"
else
    echo "✗ Graphics version build failed (probably missing X11 libraries)"
    echo "  Install X11 development libraries for graphics version"
    echo "  For now, you can run the console demo: ./console_demo"
fi