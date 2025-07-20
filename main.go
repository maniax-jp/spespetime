package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 800
	screenHeight = 600
	ballRadius   = 8
	paddleWidth  = 100
	paddleHeight = 15
	blockWidth   = 60
	blockHeight  = 20
	blockRows    = 5
	blockCols    = 10
)

type Ball struct {
	x, y   float64
	vx, vy float64
}

type Paddle struct {
	x, y float64
}

type Block struct {
	x, y      float64
	destroyed bool
}

type Game struct {
	ball     Ball
	paddle   Paddle
	blocks   [][]Block
	score    int
	gameOver bool
	won      bool
}

func NewGame() *Game {
	g := &Game{
		ball: Ball{
			x:  screenWidth / 2,
			y:  screenHeight - 100,
			vx: 3,
			vy: -3,
		},
		paddle: Paddle{
			x: screenWidth/2 - paddleWidth/2,
			y: screenHeight - 50,
		},
	}

	// Initialize blocks
	g.blocks = make([][]Block, blockRows)
	for row := 0; row < blockRows; row++ {
		g.blocks[row] = make([]Block, blockCols)
		for col := 0; col < blockCols; col++ {
			g.blocks[row][col] = Block{
				x: float64(col*(blockWidth+5) + 40),
				y: float64(row*(blockHeight+5) + 50),
			}
		}
	}

	return g
}

func (g *Game) Update() error {
	if g.gameOver {
		// Restart game on space key
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			*g = *NewGame()
		}
		return nil
	}

	// Paddle movement
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.paddle.x > 0 {
		g.paddle.x -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.paddle.x < screenWidth-paddleWidth {
		g.paddle.x += 5
	}

	// Ball movement
	g.ball.x += g.ball.vx
	g.ball.y += g.ball.vy

	// Ball collision with walls
	if g.ball.x <= ballRadius || g.ball.x >= screenWidth-ballRadius {
		g.ball.vx = -g.ball.vx
	}
	if g.ball.y <= ballRadius {
		g.ball.vy = -g.ball.vy
	}

	// Ball collision with paddle
	if g.ball.y+ballRadius >= g.paddle.y &&
		g.ball.x >= g.paddle.x && g.ball.x <= g.paddle.x+paddleWidth {
		g.ball.vy = -math.Abs(g.ball.vy)
		// Add some horizontal velocity based on where ball hits paddle
		hitPos := (g.ball.x - g.paddle.x) / paddleWidth
		g.ball.vx = (hitPos - 0.5) * 6
	}

	// Ball collision with blocks
	for row := 0; row < blockRows; row++ {
		for col := 0; col < blockCols; col++ {
			block := &g.blocks[row][col]
			if !block.destroyed &&
				g.ball.x+ballRadius >= block.x &&
				g.ball.x-ballRadius <= block.x+blockWidth &&
				g.ball.y+ballRadius >= block.y &&
				g.ball.y-ballRadius <= block.y+blockHeight {
				block.destroyed = true
				g.ball.vy = -g.ball.vy
				g.score += 10
			}
		}
	}

	// Check win condition
	allDestroyed := true
	for row := 0; row < blockRows; row++ {
		for col := 0; col < blockCols; col++ {
			if !g.blocks[row][col].destroyed {
				allDestroyed = false
				break
			}
		}
		if !allDestroyed {
			break
		}
	}
	if allDestroyed {
		g.won = true
		g.gameOver = true
	}

	// Game over if ball falls off screen
	if g.ball.y > screenHeight {
		g.gameOver = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 50, 255})

	if g.gameOver {
		var message string
		if g.won {
			message = fmt.Sprintf("You Won! Score: %d\nPress SPACE to restart", g.score)
		} else {
			message = fmt.Sprintf("Game Over! Score: %d\nPress SPACE to restart", g.score)
		}
		ebitenutil.DebugPrint(screen, message)
		return
	}

	// Draw ball
	vector.DrawFilledCircle(screen, float32(g.ball.x), float32(g.ball.y), ballRadius, color.White, false)

	// Draw paddle
	vector.DrawFilledRect(screen, float32(g.paddle.x), float32(g.paddle.y), paddleWidth, paddleHeight, color.RGBA{255, 255, 255, 255}, false)

	// Draw blocks
	for row := 0; row < blockRows; row++ {
		for col := 0; col < blockCols; col++ {
			block := g.blocks[row][col]
			if !block.destroyed {
				// Different colors for different rows
				blockColor := color.RGBA{
					R: uint8(255 - row*40),
					G: uint8(100 + row*30),
					B: uint8(150 + col*10),
					A: 255,
				}
				vector.DrawFilledRect(screen, float32(block.x), float32(block.y), blockWidth, blockHeight, blockColor, false)
			}
		}
	}

	// Draw score
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.score))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("ブロック崩し (Block Breaking Game)")

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}