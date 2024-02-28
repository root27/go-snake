// Snake game in Go

package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

type Point struct {
	x int

	y int
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

	// Draw the snake

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for _, p := range snake {

		termbox.SetCell(p.x, p.y, '0', termbox.ColorGreen, termbox.ColorDefault)

	}

	termbox.SetCell(food.x, food.y, 'X', termbox.ColorRed, termbox.ColorDefault)

	termbox.Flush()

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

			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

			// Move the snake

			head := snake[0]

			newHead := Point{head.x + dx, head.y + dy}

			if newHead == food {

				// Grow the snake

				snake = append([]Point{newHead}, snake...)

				// Generate new food

				food = Point{rand.Intn(50), rand.Intn(20)}

			} else {

				// Move the snake

				snake = append([]Point{newHead}, snake...)

				snake = snake[:len(snake)-1]

			}

			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

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
