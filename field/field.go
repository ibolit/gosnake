package field
import (
    "gosnake/util"
    "fmt"
)

type field struct {
    width int
    height int
    field []int
}

func NewField(width int, height int) *field {
    return &field{width: width, height: height, field: make([]int, width * height)}
}

func (f *field) Print() {
    for i := 0; i < f.width; i++ {
        for j:= 0; j< f.height; j++ {
            util.PrintAt(util.Point{i, j}, fmt.Sprintf("%d", f.field[i * f.width + j]))
        }
    }
}
