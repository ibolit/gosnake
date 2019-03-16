package snake

import (
    "fmt"
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
func NewSegment(x, y int, previous *Segment) *Segment {
    return &Segment{point: util.Point{x, y}, previous: previous}
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

type Snake struct {
    head *Segment
    tail *Segment
}

func NewSnake() *Snake {
    snake := &Snake{}
    snake.head = &Segment{}
    snake.tail = snake.head
    return snake
}

func (snake *Snake) MoveTo(point util.Point) {
    // fmt.Println(snake.head.previous)
    if false {
        fmt.Println(point.IsAdjacentTo(snake.head.point))
        fmt.Println(snake.head.previous == nil)
    }
    if point.IsAdjacentTo(snake.head.point) &&
        snake.head.previous == nil ||
        point != snake.head.previous.point {
        // fmt.Println("Moving")
        snake.addHead(point)
        // snake.removeTail()
    }
}

func (snake *Snake) addHead(point util.Point) {
    newHead := &Segment{point: point, previous: snake.head}
    snake.head.next = newHead
    snake.head = newHead
    // fmt.Println(snake.head)
}

func (snake *Snake) removeTail() {
    snake.tail = snake.tail.next
    snake.tail.previous = nil
}

func (snake *Snake) Iterator() *snakeIterator {
    // fmt.Println(">>", snake.head)
    return &snakeIterator{&Segment{previous: snake.head}}
}

type snakeIterator struct {
    segment *Segment
}

func (x *snakeIterator) Next() bool {
    x.segment = x.segment.previous
    return x.segment != nil
}

func (x *snakeIterator) Value() util.Printable {
    return x.segment
}
