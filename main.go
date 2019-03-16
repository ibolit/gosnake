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

	// l := iter2.NewLinkedList()
	// for iter := l.Iter(); iter != nil; iter = iter.Next() {
	// 	if iter == nil {
	// 		fmt.Println("Iter is nil")
	// 	}
	// 	fmt.Println(iter)
	// }

	// for ; node != nil; node = node.Next() {
	// 	fmt.Println(node)
	// }

	// return

	// rootSegment := &snake.Segment{}
	// // fmt.Println(rootSegment, &rootSegment)
	// currentSegment := rootSegment
	// for i := 0; i < 5; i++ {
	// 	newSegment := snake.NewSegment(i, i, currentSegment)
	// 	currentSegment.SetNext(newSegment)
	// 	currentSegment = newSegment
	// }

	aField := field.NewField(10, 10)
	aField.Print()

	aSnake := snake.NewSnake()
	// printSnake(a_snake)
	aSnake.MoveTo(util.Point{0, 1})
	// a_field.Print()
	printSnake(aSnake)

	// currentSegment = rootSegment
	// for currentSegment.HasNext() {
	// 	util.PrintAt(currentSegment.Point(), currentSegment.Value())
	// 	// fmt.Println(currentSegment.point.X)
	// 	currentSegment = currentSegment.Next()
	// }

	channel := make(chan keyListener.Direction)

	util.PrintAt(util.Point{11, 0}, ">>")
	go keyListener.Listen(channel)

	for i := range channel {
		fmt.Println(i)
	}

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
