package main

import (
	"gosnake/field"
	"gosnake/util"
	"gosnake/snake"
	"fmt"
)



const RABBIT = "🐇🐁🐀🐔🐓🐊🐉🐤🐥🐦🐭🐯🐹🙀😼"
const R_2 = "🐰🐅🐙☠☣☢"


func printSnake(a_snake *snake.Snake) {
	for it := a_snake.Iterator(); it.Next(); {
		util.PrintAt(it.Value().Point(), ">")
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


	a_field := field.NewField(10, 10)
	a_field.Print()

	a_snake := snake.NewSnake()
	// printSnake(a_snake)
	a_snake.MoveTo(util.Point{0, 1})
	// a_field.Print()
	printSnake(a_snake)
	




	// currentSegment = rootSegment
	// for currentSegment.HasNext() {
	// 	util.PrintAt(currentSegment.Point(), currentSegment.Value())
	// 	// fmt.Println(currentSegment.point.X)
	// 	currentSegment = currentSegment.Next()
	// }

	util.PrintAt(util.Point{11, 0}, ">>")


}
