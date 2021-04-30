package lib

type Square int

func (square Square) String() string {
    var x, y = fromSquare(square)
    return string(rune('A'+x)) + string(rune('1'+y))
}

func toSquare(x int, y int) Square {
    return Square(x + (y * 8))
}

func fromSquare(square Square) (int, int) {
    var x = int(square % 8)
    var y = int(square / 8)

    return x, y
}

var A1 = toSquare(0, 0)
var A2 = toSquare(0, 1)
var A3 = toSquare(0, 2)
var A4 = toSquare(0, 3)
var A5 = toSquare(0, 4)
var A6 = toSquare(0, 5)
var A7 = toSquare(0, 6)
var A8 = toSquare(0, 7)

var B1 = toSquare(1, 0)
var B2 = toSquare(1, 1)
var B3 = toSquare(1, 2)
var B4 = toSquare(1, 3)
var B5 = toSquare(1, 4)
var B6 = toSquare(1, 5)
var B7 = toSquare(1, 6)
var B8 = toSquare(1, 7)

var C1 = toSquare(2, 0)
var C2 = toSquare(2, 1)
var C3 = toSquare(2, 2)
var C4 = toSquare(2, 3)
var C5 = toSquare(2, 4)
var C6 = toSquare(2, 5)
var C7 = toSquare(2, 6)
var C8 = toSquare(2, 7)

var D1 = toSquare(3, 0)
var D2 = toSquare(3, 1)
var D3 = toSquare(3, 2)
var D4 = toSquare(3, 3)
var D5 = toSquare(3, 4)
var D6 = toSquare(3, 5)
var D7 = toSquare(3, 6)
var D8 = toSquare(3, 7)

var E1 = toSquare(4, 0)
var E2 = toSquare(4, 1)
var E3 = toSquare(4, 2)
var E4 = toSquare(4, 3)
var E5 = toSquare(4, 4)
var E6 = toSquare(4, 5)
var E7 = toSquare(4, 6)
var E8 = toSquare(4, 7)

var F1 = toSquare(5, 0)
var F2 = toSquare(5, 1)
var F3 = toSquare(5, 2)
var F4 = toSquare(5, 3)
var F5 = toSquare(5, 4)
var F6 = toSquare(5, 5)
var F7 = toSquare(5, 6)
var F8 = toSquare(5, 7)

var G1 = toSquare(6, 0)
var G2 = toSquare(6, 1)
var G3 = toSquare(6, 2)
var G4 = toSquare(6, 3)
var G5 = toSquare(6, 4)
var G6 = toSquare(6, 5)
var G7 = toSquare(6, 6)
var G8 = toSquare(6, 7)

var H1 = toSquare(7, 0)
var H2 = toSquare(7, 1)
var H3 = toSquare(7, 2)
var H4 = toSquare(7, 3)
var H5 = toSquare(7, 4)
var H6 = toSquare(7, 5)
var H7 = toSquare(7, 6)
var H8 = toSquare(7, 7)
