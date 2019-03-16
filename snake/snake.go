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

// SetNext Sets the next segment in the linked structure
func (segment *Segment) SetNext(next *Segment) {
    segment.next = next
}

// SetPrevious Sets the previous thing
func (segment *Segment) SetPrevious(previous *Segment) {
    segment.previous = previous
}

// Next Get the next thing
func (segment *Segment) Next() *Segment {
    return segment.next
}

// HasNext Returns true of there are more things in the thing
func (segment *Segment) HasNext() bool {
    return segment.next != nil
}

// Point Returns the point of the setment
func (segment *Segment) Point() util.Point {
    return segment.point
}

// Value Returns the value of the segment
func (segment *Segment) Value() string {
    return SnakeSymbol
}

// Snake The structure holding our snake
type Snake struct {
    head      *Segment
    tail      *Segment
    direction util.Direction
}

// NewSnake Create a new snake instance
func NewSnake() *Snake {
    snake := &Snake{}
    snake.head = &Segment{}
    snake.tail = snake.head
    snake.direction = util.Right
    return snake
}

// SetCandidateMove Sets the next point the snake should move to
func (snake *Snake) SetCandidateMove(direction util.Direction) {
    snake.direction = direction
}

func (snake *Snake) nextPoint() util.Point {
    startingPoint := snake.head.point
    switch snake.direction {
    case util.Up:
        return util.Point{startingPoint.X - 1, startingPoint.Y}
    case util.Down:
        return util.Point{startingPoint.X + 1, startingPoint.Y}
    case util.Left:
        return util.Point{startingPoint.X, startingPoint.Y - 1}
    case util.Right:
        return util.Point{startingPoint.X, startingPoint.Y + 1}
    }
    return startingPoint
}

// MoveTo Move to the next (candidate) point
func (snake *Snake) Move() {
    snake.addHead(snake.nextPoint())
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

// Iterator Get the iterator to iterate over the snake
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
