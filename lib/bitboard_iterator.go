package lib

type BitboardIterator struct {
    bitboard Bitboard
}

func NewBitboardIterator(bb Bitboard) BitboardIterator {
    return BitboardIterator{bitboard: bb}
}

func (bi BitboardIterator) HasNext() bool {
    return !bi.bitboard.isEmpty()
}

func (bi *BitboardIterator) Next() int {
    var next = bi.bitboard.nextOne()
    bi.bitboard = bi.bitboard.cleared(next)

    return next
}
