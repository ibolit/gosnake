package snake

import (
    "fmt"
    "gosnake/util"
)

// SNAKE_S The char representing a snake segment
const SNAKE_S = "▄"

type Segment struct {
    point    util.Point
    next     *Segment
    previous *Segment
}

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
    return SNAKE_S
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
    // fmt.Println(snake.head.previous)
    if false {
        fmt.Println(point.IsAdjacentTo(snake.head.point))
        fmt.Println(snake.head.previous == nil)
    }
    // if point.IsAdjacentTo(snake.head.point) &&
    //  snake.head.previous == nil ||
    //  point != snake.head.previous.point {
    snake.addHead(point)
    util.PrintAt(util.Point{13, 0}, ">> Moving, head is "+snake.head.point.String())
    // snake.removeTail()
    // }
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
