package keyListener

import (
	term "github.com/nsf/termbox-go"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

// func reset() {
// 	term.Sync() // cosmestic purpose
// }

func Listen(c chan Direction) {
	err := term.Init()
	if err != nil {
		panic(err)
	}

	defer term.Close()

	// fmt.Println("Enter any key to see their ASCII code or press ESC button to quit")

keyPressListenerLoop:
	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				break keyPressListenerLoop
				close(c)

			case term.KeyArrowUp:
				// reset()
				// fmt.Println("Arrow Up pressed")
				c <- Up
			case term.KeyArrowDown:
				// reset()
				// fmt.Println("Arrow Down pressed")
				c <- Down
			case term.KeyArrowLeft:
				// reset()
				// fmt.Println("Arrow Left pressed")
				c <- Left
			case term.KeyArrowRight:
				// reset()
				// fmt.Println("Arrow Right pressed")
				c <- Right
			}
		case term.EventError:
			panic(ev.Err)
		}
	}
}
