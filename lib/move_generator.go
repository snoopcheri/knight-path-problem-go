package lib

func (b Board) GenerateMoves() []Move {
    var moves []Move
    var knightSquares = NewBitboardIterator(b.knights())

    for knightSquares.HasNext() {
        var square = knightSquares.Next()
        moves = append(moves, b.generateMoveFromSquare(Square(square))...)
    }

    return moves
}

func (b Board) generateMoveFromSquare(from Square) []Move {
    var moves []Move

    for _, distance := range knightDistances {
        var to = destinationSquare(from, distance)

        if to != nil && b.squareIsEmpty(*to) {
            moves = append(moves, NewMove(from, *to))
        }
    }

    return moves
}

func destinationSquare(from Square, distance Distance) *Square {
    var x, y = fromSquare(from)
    x += distance.x
    y += distance.y

    if x >= 0 && x <= 7 && y >= 0 && y <= 7 {
        var to = toSquare(x, y)
        return &to
    } else {
        return nil
    }

}

type Distance struct {
    x int
    y int
}

var knightDistances = [...]Distance{
    Distance{x: 1, y: 2},
    Distance{x: 2, y: 1},
    Distance{x: 2, y: -1},
    Distance{x: 1, y: -2},
    Distance{x: -1, y: -2},
    Distance{x: -2, y: -1},
    Distance{x: -2, y: 1},
    Distance{x: -1, y: 2},
}
