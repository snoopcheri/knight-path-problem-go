package lib

import (
    "fmt"
    "strconv"
)

type Board struct {
    whiteKnights Bitboard
    blackKnights Bitboard
}

func NewBoard(whiteKnights Bitboard, blackKnights Bitboard) Board {
    return Board{
        whiteKnights: whiteKnights,
        blackKnights: blackKnights,
    }
}

func (b Board) knights() Bitboard {
    return b.whiteKnights.or(b.blackKnights)
}

func (b Board) WithSwitchedKnights() Board {
    return Board{
        whiteKnights: b.blackKnights,
        blackKnights: b.whiteKnights,
    }
}

func (b Board)withKnight(square Square, colour Colour) Board {
    switch colour {
    case WHITE:
        return Board{whiteKnights: b.whiteKnights.set(int(square)), blackKnights: b.blackKnights}
    case BLACK:
        return Board{whiteKnights: b.whiteKnights, blackKnights: b.blackKnights.set(int(square))}
    default:
        panic("Unsupported colour")
    }
}

func (b Board) afterMove(move Move) Board {
    if b.whiteKnights.Get(move.from) {
        return Board{
            whiteKnights: b.whiteKnights.Cleared(move.from).Set(move.to),
            blackKnights: b.blackKnights,
        }
    } else if b.blackKnights.Get(move.from) {
        return Board{
            whiteKnights: b.whiteKnights,
            blackKnights: b.blackKnights.Cleared(move.from).Set(move.to),
        }
    } else {
        panic(fmt.Sprint("Cannot execute move ", move, " as square ", move.from, " is not occupied"))
    }
}

func (b Board) squareIsEmpty(square Square) bool {
    return !b.whiteKnights.Get(square) && !b.blackKnights.Get(square)
}

func (b Board) emptySquares() Bitboard {
    var occupied = b.knights()
    return occupied.not()
}

func (b Board) String() string {
    var str = ""

    for y := 7; y >= 0; y-- {
        str += strconv.Itoa(y + 1)

        for x :=0; x <=7; x++ {
            str += " "
            var square = toSquare(x, y)
            if b.whiteKnights.get(int(square)) {
                str += "N"
            } else if (b.blackKnights.get(int(square))) {
                str += "n"
            } else {
                str += "."
            }
        }

        str += "\n"
    }

    str += " "
    for x :=0; x <=7; x++ {
        str = str + " " + string(rune('a'+x))
    }

    str += "\n"

    return str
}
