package snake

import (
    "gosnake/util"
)

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
