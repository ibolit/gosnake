package snake

import (
    "gosnake/util"
)

// SnakeSymbol A symbol that makes up the snake
const SnakeSymbol = "▄"

// Segment a segment of the snake
type Segment struct {
    point    util.Point
    next     *Segment
    previous *Segment
}

// NewSegment Make a new snake segment
func NewSegment(x, y int, next *Segment) *Segment {
    return &Segment{point: util.Point{x, y}, next: next}
}

func (segment *Segment) SetNext(next *Segment) {
    segment.next = next
}

func (segment *Segment) SetPrevious(previous *Segment) {
    segment.previous = previous
}

func (segment *Segment) Next() *Segment {
    return segment.next
}

func (segment *Segment) HasNext() bool {
    return segment.next != nil
}

func (segment *Segment) Point() util.Point {
    return segment.point
}

func (segment *Segment) Value() string {
    return SnakeSymbol
}

// Snake The structure holding our snake
type Snake struct {
    head *Segment
    tail *Segment
}

// NewSnake Create a new snake instance
func NewSnake() *Snake {
    snake := &Snake{}
    snake.head = &Segment{}
    snake.tail = snake.head
    return snake
}

func (snake *Snake) MoveTo(point util.Point) {
    snake.addHead(point)
    util.PrintAt(util.Point{13, 0}, ">> Moving, head is "+snake.head.point.String())
    // snake.removeTail()
    // }
}

func (snake *Snake) addHead(point util.Point) {
    newHead := &Segment{point: point, next: snake.head}
    snake.head.previous = newHead
    snake.head = newHead
}

func (snake *Snake) removeTail() {
    snake.tail = snake.tail.next
    snake.tail.previous = nil
}

func (snake *Snake) Iterator() *snakeIterator {
    return &snakeIterator{&Segment{next: snake.head}}
}

type snakeIterator struct {
    segment *Segment
}

func (x *snakeIterator) Next() bool {
    x.segment = x.segment.next
    return x.segment != nil
}

func (x *snakeIterator) Value() util.Printable {
    return x.segment
}
