package main

import (
	"fmt"
	"github.com/buger/goterm"
	// "gosnake/iter"
	"gosnake/keyListener"
	"gosnake/snake"
	"gosnake/util"
	"math/rand"
	"time"
)

// Rabbit Possible symbols for a rabbit to be eaten by a snake
const Rabbit = "v" //"🐇" //🐁🐀🐔🐓🐊🐉🐤🐥🐦🐭🐯🐹🙀😼"

// Rabbit2 Alternative symbols
const Rabbit2 = "🐰🐅🐙☠☣☢"

var i int

func printSnake(aSnake *snake.Snake) {
	printSnakeHelper(aSnake, snake.SnakeSymbol)
}

func deleteSnake(aSnake *snake.Snake) {
	printSnakeHelper(aSnake, " ")
}

func printSnakeHelper(aSnake *snake.Snake, val string) {
	for it := aSnake.Iterator(); it.Next(); i++ {
		util.PrintAt(it.Value().Point(), val)
	}
}

func makeRabbit(oldRabbit util.Point, width, height int) util.Point {
	util.PrintAt(oldRabbit, " ")
	newRabbit := util.Point{X: rand.Intn(width-5) + 3, Y: rand.Intn(height-2) + 1}
	util.PrintAt(newRabbit, "v")
	return newRabbit
}

func getTermSize() (width, height int) {
	return goterm.Width(), goterm.Height()
}

func main() {
	width, height := getTermSize()
	// fmt.Println("width, height: ", width, height)
	// return

	aSnake := snake.NewSnake()
	shouldMakeRabbit := true
	rabbitLifespan := 0

	ticker := time.NewTicker(time.Millisecond * 500)
	var rabbit util.Point

	channel := make(chan util.Direction)
	quit := make(chan int)
	shouldGoOn := true

	go func() {
		for range ticker.C {
			if shouldMakeRabbit {
				rabbit = makeRabbit(rabbit, width, height)
				rabbitLifespan = rand.Intn(100)
				shouldMakeRabbit = false
			}
			rabbitLifespan--
			deleteSnake(aSnake)
			eaten, err := aSnake.Move(rabbit)
			if err > 0 {
				ticker.Stop()
				close(channel)
				close(quit)
				shouldGoOn = false
			}
			shouldMakeRabbit = eaten || rabbitLifespan == 0
			printSnake(aSnake)
		}
	}()

	go keyListener.Listen(channel, quit)

	util.PrintAt(util.Point{11, 0}, ">>")

	for shouldGoOn {

		select {
		case direction := <-channel:
			aSnake.SetCandidateMove(direction)
		case <-quit:
			fmt.Println("Quitting")
			ticker.Stop()
			close(channel)
			close(quit)
			shouldGoOn = false
		}
	}
}
