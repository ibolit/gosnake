package main

import (
	"fmt"
	"gosnake/field"
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
	aField := field.NewField(10, 10)
	aField.Print()
	aSnake := snake.NewSnake()
	printSnake(aSnake)

	i := 0

	ticker := time.NewTicker(time.Millisecond * 500)
	wait := true

	go func() {
		for t := range ticker.C {
			i++
			aSnake.MoveTo(util.Point{X: 0, Y: i})
			printSnake(aSnake)
			if i > 12 {
				ticker.Stop()
				wait = false
			}
			if false {
				fmt.Print(t)
			}
		}
	}()

	for wait {
	}

}
