.PHONY: all console graphics clean run run-console help

# Default target
all: console

# Build console demo (works without graphics libraries)
console:
	@echo "Building console demo..."
	go build -o console_demo console_demo.go

# Build graphics version (requires X11 libraries)
graphics:
	@echo "Building graphics version..."
	go build -o blockbreaker main.go

# Clean built binaries
clean:
	@echo "Cleaning..."
	rm -f console_demo blockbreaker

# Run console demo
run-console: console
	@echo "Running console demo..."
	./console_demo

# Run graphics version
run: graphics
	@echo "Running graphics version..."
	./blockbreaker

# Show help
help:
	@echo "Available targets:"
	@echo "  all        - Build console demo (default)"
	@echo "  console    - Build console demo"
	@echo "  graphics   - Build graphics version"
	@echo "  clean      - Remove built binaries"
	@echo "  run-console- Build and run console demo"
	@echo "  run        - Build and run graphics version"
	@echo "  help       - Show this help"