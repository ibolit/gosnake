package util

import (
    "fmt"
    "math"
)

type Point struct {
    X, Y int
}

func (point Point) String() string {
    return fmt.Sprintf("x: %d, y: %d", point.X, point.Y)
}

func PrintAt(point Point, value interface{}) {
    fmt.Printf("\033[%d;%df%s", point.X, point.Y+1, value)
}

type Printable interface {
    Point() Point
    Value() string
}

func (point Point) IsAdjacentTo(another Point) bool {
    xDiff := math.Abs(float64(point.X - another.X))
    yDiff := math.Abs(float64(point.Y - another.Y))
    return xDiff+yDiff == 1
}

type Iterator interface {
    Next() bool
    Value() Printable
}
