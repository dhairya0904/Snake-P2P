package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gdamore/tcell"
	mapstructure "github.com/mitchellh/mapstructure"
)

type Game struct {
	Screen            tcell.Screen
	snakeBody         SnakeBody
	opponentSnakeBody SnakeBody
	FoodPos           Part
	OpponentFoodPos   Part
	Score             int
	GameOver          bool
	Node              *Node
}

func drawParts(s tcell.Screen, snakeParts []Part, foodPos Part, snakeStyle tcell.Style, foodStyle tcell.Style) {
	s.SetContent(foodPos.X, foodPos.Y, '\u25CF', nil, foodStyle)
	for _, part := range snakeParts {
		s.SetContent(part.X, part.Y, ' ', nil, snakeStyle)
	}
}

func drawText(s tcell.Screen, x1, y1, x2, y2 int, text string) {
	row := y1
	col := x1
	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	for _, r := range text {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func checkCollision(parts []Part, otherPart Part) bool {
	for _, part := range parts {
		if part.X == otherPart.X && part.Y == otherPart.Y {
			return true
		}
	}
	return false
}

func (g *Game) UpdateFoodPos(width int, height int) {
	g.FoodPos.X = rand.Intn(width)
	g.FoodPos.Y = rand.Intn(height)
	if g.FoodPos.Y == 1 && g.FoodPos.X < 10 {
		g.UpdateFoodPos(width, height)
	}
}

func (g *Game) UpdateOFoodPos(width int, height int) {
	g.OpponentFoodPos.X = rand.Intn(width)
	g.OpponentFoodPos.Y = rand.Intn(height)
	if g.FoodPos.Y == 1 && g.FoodPos.X < 10 {
		g.UpdateOFoodPos(width, height)
	}
}

func (g *Game) Run() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	g.Screen.SetStyle(defStyle)
	width, height := g.Screen.Size()
	g.snakeBody.ResetPos(width, height)
	g.opponentSnakeBody.ResetPos(width/2, height/2)
	g.UpdateFoodPos(width, height)
	g.UpdateOFoodPos(width, height)
	g.GameOver = false
	g.Score = 0
	snakeStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)
	for {

		longerSnake := false
		g.Screen.Clear()
		if checkCollision(g.snakeBody.Parts[len(g.snakeBody.Parts)-1:], g.FoodPos) {
			g.UpdateFoodPos(width, height)
			longerSnake = true
			g.Score++
		}
		if checkCollision(g.snakeBody.Parts[:len(g.snakeBody.Parts)-1], g.snakeBody.Parts[len(g.snakeBody.Parts)-1]) {
			break
		}
		g.snakeBody.Update(width, height, longerSnake)

		newState := GameStateUpdade{
			FoodPos: g.FoodPos,
			Parts:   g.snakeBody.Parts,
			Xspeed:  g.snakeBody.Xspeed,
			Yspeed:  g.snakeBody.Yspeed,
		}
		g.Node.writeChannel <- newState

		var newSnakeState GameStateUpdade
		tmo := <-g.Node.readChannel
		err := mapstructure.Decode(tmo, &newSnakeState)
		if err != nil {
			panic(err)
		}

		g.opponentSnakeBody.Parts = newSnakeState.Parts

		g.opponentSnakeBody.Update(width, height, longerSnake)
		drawParts(g.Screen, g.snakeBody.Parts, g.FoodPos, snakeStyle, defStyle)
		drawParts(g.Screen, g.opponentSnakeBody.Parts, g.OpponentFoodPos, snakeStyle, defStyle)
		drawText(g.Screen, 1, 1, 8+len(strconv.Itoa(g.Score)), 1, "Score: "+strconv.Itoa(g.Score))
		time.Sleep(40 * time.Millisecond)
		g.Screen.Show()
	}
	g.GameOver = true
	drawText(g.Screen, width/2-20, height/2, width/2+20, height/2, "Game Over, Score: "+strconv.Itoa(g.Score)+", Play Again? y/n")
	g.Screen.Show()
}
