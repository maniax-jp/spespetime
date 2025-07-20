package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	width  = 40
	height = 20
)

type ConsoleGame struct {
	ballX, ballY     int
	ballVX, ballVY   int
	paddleX          int
	blocks           [][]bool
	score            int
	gameOver         bool
	won              bool
}

func NewConsoleGame() *ConsoleGame {
	g := &ConsoleGame{
		ballX:   width / 2,
		ballY:   height - 3,
		ballVX:  1,
		ballVY:  -1,
		paddleX: width/2 - 2,
	}

	// Initialize blocks
	g.blocks = make([][]bool, 3)
	for i := range g.blocks {
		g.blocks[i] = make([]bool, 8)
		for j := range g.blocks[i] {
			g.blocks[i][j] = true
		}
	}

	return g
}

func (g *ConsoleGame) Update() {
	if g.gameOver {
		return
	}

	// Move ball
	g.ballX += g.ballVX
	g.ballY += g.ballVY

	// Ball collision with walls
	if g.ballX <= 0 || g.ballX >= width-1 {
		g.ballVX = -g.ballVX
	}
	if g.ballY <= 0 {
		g.ballVY = -g.ballVY
	}

	// Ball collision with paddle
	if g.ballY == height-2 && g.ballX >= g.paddleX && g.ballX <= g.paddleX+4 {
		g.ballVY = -g.ballVY
	}

	// Ball collision with blocks
	blockY := g.ballY / 2
	blockX := g.ballX / 5
	if blockY >= 2 && blockY < 5 && blockX >= 0 && blockX < 8 {
		if g.blocks[blockY-2][blockX] {
			g.blocks[blockY-2][blockX] = false
			g.ballVY = -g.ballVY
			g.score += 10
		}
	}

	// Check win condition
	allDestroyed := true
	for i := range g.blocks {
		for j := range g.blocks[i] {
			if g.blocks[i][j] {
				allDestroyed = false
				break
			}
		}
	}
	if allDestroyed {
		g.won = true
		g.gameOver = true
	}

	// Game over if ball falls off screen
	if g.ballY >= height {
		g.gameOver = true
	}
}

func (g *ConsoleGame) Draw() {
	// Clear screen
	fmt.Print("\033[2J\033[H")

	if g.gameOver {
		if g.won {
			fmt.Printf("You Won! Score: %d\n", g.score)
		} else {
			fmt.Printf("Game Over! Score: %d\n", g.score)
		}
		fmt.Println("This is a simple console demo.")
		fmt.Println("Run the full graphics version with: go run main.go")
		return
	}

	// Draw game field
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Draw ball
			if x == g.ballX && y == g.ballY {
				fmt.Print("O")
				continue
			}

			// Draw paddle
			if y == height-2 && x >= g.paddleX && x <= g.paddleX+4 {
				fmt.Print("=")
				continue
			}

			// Draw blocks
			if y >= 2 && y < 8 && x%5 == 0 && x/5 < 8 {
				blockRow := (y - 2) / 2
				blockCol := x / 5
				if blockRow < 3 && g.blocks[blockRow][blockCol] {
					fmt.Print("█")
					continue
				}
			}

			// Draw walls
			if x == 0 || x == width-1 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	fmt.Printf("Score: %d | Use 'a'/'d' to move paddle, 'q' to quit\n", g.score)
}

func (g *ConsoleGame) MovePaddle(direction int) {
	g.paddleX += direction
	if g.paddleX < 1 {
		g.paddleX = 1
	}
	if g.paddleX > width-6 {
		g.paddleX = width - 6
	}
}

func runConsoleDemo() {
	game := NewConsoleGame()
	
	// Set up non-blocking input
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			
			switch input {
			case "a":
				game.MovePaddle(-1)
			case "d":
				game.MovePaddle(1)
			case "q":
				os.Exit(0)
			}
		}
	}()

	for {
		game.Draw()
		game.Update()
		
		if game.gameOver {
			time.Sleep(3 * time.Second)
			break
		}
		
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	fmt.Println("ブロック崩し (Block Breaking Game) - Console Demo")
	fmt.Println("================================================")
	fmt.Println()
	fmt.Println("This is a simple console version for demonstration.")
	fmt.Println("For the full graphics version, install X11 libraries and run: go run main.go")
	fmt.Println()
	fmt.Println("Press Enter to start demo...")
	fmt.Scanln()
	
	runConsoleDemo()
}