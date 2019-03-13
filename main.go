package main

import (
	"fmt"
	"gosnake/field"
	"gosnake/snake"
	"gosnake/util"
)

// Rabbit Possible symbols for a rabbit to be eaten by a snake
const Rabbit = "🐇🐁🐀🐔🐓🐊🐉🐤🐥🐦🐭🐯🐹🙀😼"

// Rabbit2 Alternative symbols
const Rabbit2 = "🐰🐅🐙☠☣☢"

var i int

func printSnake(aSnake *snake.Snake) {
	for it := aSnake.Iterator(); it.Next(); i++ {
		util.PrintAt(it.Value().Point(), it.Value().Value())
		util.PrintAt(util.Point{X: 14 + i, Y: 0}, fmt.Sprintf("%p", it))
	}
}

func main() {
	fmt.Println("Hello	")

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
	printSnake(aSnake)
	aSnake.MoveTo(util.Point{X: 0, Y: 1})
	// a_field.Print()
	printSnake(aSnake)

	// currentSegment = rootSegment
	// for currentSegment.HasNext() {
	// 	util.PrintAt(currentSegment.Point(), currentSegment.Value())
	// 	// fmt.Println(currentSegment.point.X)
	// 	currentSegment = currentSegment.Next()
	// }

	util.PrintAt(util.Point{X: 11, Y: 0}, ">>")

}
