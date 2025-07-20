# ブロック崩しゲーム (Block Breaking Game)

This repository contains a block breaking game implemented in Go.

## Requirements

- Go 1.18 or later
- X11 development libraries (for graphics)

### Installing X11 libraries (Linux)

```bash
# Ubuntu/Debian
sudo apt-get install libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev libgl1-mesa-dev libglu1-mesa-dev libasound2-dev

# CentOS/RHEL/Fedora
sudo yum install libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel libXi-devel mesa-libGL-devel mesa-libGLU-devel alsa-lib-devel
```

## Building and Running

```bash
# Clone the repository
git clone https://github.com/maniax-jp/spespetime.git
cd spespetime

# Install dependencies
go mod tidy

# Build the game
go build -o blockbreaker main.go

# Run the game
./blockbreaker
```

## Game Controls

- **Arrow Left/Right**: Move paddle
- **Space**: Restart game (when game over)

## Game Features

- Ball physics with realistic bouncing
- Paddle movement with ball angle control
- 5 rows of colorful blocks to destroy
- Score tracking
- Win/lose conditions
- Restart functionality

## Implementation Details

The game is implemented using the Ebitengine (formerly Ebiten) graphics library for Go, which provides:
- 2D graphics rendering
- Input handling
- Game loop management
- Cross-platform compatibility

The game includes:
- Ball movement and collision detection
- Paddle controls
- Block destruction mechanics
- Score management
- Game state handling (playing, game over, won)