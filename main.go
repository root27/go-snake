// Snake game in Go

package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	WIDTH  = 100
	HEIGHT = 30
)

type Point struct {
	x int

	y int
}

func printText(x int, y int, text string) {

	for _, c := range text {

		termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorDefault)

		x++

	}

}

func drawWalls(xStart int, yStart int, width int, height int) {

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	printText(105, 10, "Snake Game")

	for xStart < width {

		termbox.SetCell(xStart, 0, '-', termbox.ColorWhite, termbox.ColorDefault)

		termbox.SetCell(xStart, height, '-', termbox.ColorWhite, termbox.ColorDefault)

		xStart++

	}

	for yStart < height {

		termbox.SetCell(0, yStart, '#', termbox.ColorWhite, termbox.ColorDefault)

		termbox.SetCell(width, yStart, '#', termbox.ColorWhite, termbox.ColorDefault)

		yStart++

	}

	termbox.Flush()

}

func main() {

	err := termbox.Init()

	if err != nil {

		log.Fatalf("Error: %s", err)

	}

	defer termbox.Close()

	// initital direction

	dx, dy := 1, 0

	// Set the snake

	snake := []Point{{4, 4}, {4, 3}, {4, 2}}

	food := Point{5, 5}

	score := 0

	// Draw the snake

	for _, p := range snake {

		termbox.SetCell(p.x, p.y, '0', termbox.ColorGreen, termbox.ColorDefault)

	}

	termbox.SetCell(food.x, food.y, 'X', termbox.ColorRed, termbox.ColorDefault)

	// Set the game loop

	eventChannel := make(chan termbox.Event)

	go func() {

		for {

			eventChannel <- termbox.PollEvent()

		}

	}()

	for {

		select {
		case ev := <-eventChannel:

			if ev.Type == termbox.EventKey {

				switch ev.Key {

				case termbox.KeyArrowUp:

					dx, dy = 0, -1

				case termbox.KeyArrowDown:

					dx, dy = 0, 1

				case termbox.KeyArrowLeft:

					dx, dy = -1, 0

				case termbox.KeyArrowRight:

					dx, dy = 1, 0

				case termbox.KeyEsc:

					return

				}

			}

		default:

			drawWalls(0, 0, WIDTH, HEIGHT)

			printText(105, 15, fmt.Sprintf("Score: %d", score))

			// Move the snake

			head := snake[0]

			newHead := Point{head.x + dx, head.y + dy}

			// Check if the snake has collided with the wall

			if newHead.x <= 0 || newHead.x >= WIDTH || newHead.y <= 0 || newHead.y >= HEIGHT {

				printText(105, 20, "Game Over")

				termbox.Flush()

				time.Sleep(2 * time.Second)

				return

			}

			if newHead == food {

				// Grow the snake

				snake = append([]Point{newHead}, snake...)

				score += 10

				// Generate new food

				food = Point{1 + rand.Intn(30), 1 + rand.Intn(20)}

			} else {

				// Move the snake

				snake = append([]Point{newHead}, snake...)

				snake = snake[:len(snake)-1]

			}

			// drawWalls(0, 0, WIDTH, HEIGHT, score)

			// Draw the snake

			for _, p := range snake {

				termbox.SetCell(p.x, p.y, 'O', termbox.ColorGreen, termbox.ColorDefault)

			}

			// Draw the food

			termbox.SetCell(food.x, food.y, 'X', termbox.ColorRed, termbox.ColorDefault)

			termbox.Flush()

			time.Sleep(100 * time.Millisecond)

		}

	}

}
