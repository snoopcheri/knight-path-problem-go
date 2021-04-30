package lib

import "fmt"

type Move struct {
    from Square
    to   Square
}

func NewMove(from Square, to Square) Move {
    return Move{from: from, to: to}
}

func (mv Move) String() string {
    return fmt.Sprint(mv.from, "-", mv.to)
}
