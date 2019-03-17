package keyListener

import (
	term "github.com/nsf/termbox-go"
	"gosnake/util"
)

// Listen Listens for key events and sends the detected direction to the
// given channel
func Listen(c chan util.Direction, quit chan int) {
	err := term.Init()
	if err != nil {
		panic(err)
	}

	defer term.Close()

keyPressListenerLoop:
	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				break keyPressListenerLoop
				quit <- 1

			case term.KeyArrowUp:
				c <- util.Up
			case term.KeyArrowDown:
				c <- util.Down
			case term.KeyArrowLeft:
				c <- util.Left
			case term.KeyArrowRight:
				c <- util.Right
			}
		case term.EventError:
			panic(ev.Err)
		}
	}
}
