package main

import (
	"fmt"
	"gosnake/field"
	// "gosnake/iter2"
	"gosnake/keyListener"
	"gosnake/snake"
	"gosnake/util"
	"time"
)

// Rabbit Possible symbols for a rabbit to be eaten by a snake
const Rabbit = "🐇🐁🐀🐔🐓🐊🐉🐤🐥🐦🐭🐯🐹🙀😼"

// Rabbit2 Alternative symbols
const Rabbit2 = "🐰🐅🐙☠☣☢"

var i int

func printSnake(aSnake *snake.Snake) {
	for it := aSnake.Iterator(); it.Next(); i++ {
		util.PointPrint(it.Value())
	}
}

func main() {
	fmt.Println("Hello	")
	aField := field.NewField(10, 10)
	aField.Print()

	aSnake := snake.NewSnake()

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for range ticker.C {
			aSnake.Move()
			print("\033[H\033[2J")
			printSnake(aSnake)
		}
	}()

	channel := make(chan util.Direction)
	quit := make(chan int)

	go keyListener.Listen(channel, quit)

	util.PrintAt(util.Point{11, 0}, ">>")

mainLoop:
	for {

		select {
		case direction := <-channel:
			aSnake.SetCandidateMove(direction)
		case <-quit:
			ticker.Stop()
			close(channel)
			close(quit)
			break mainLoop
		}
	}
}
