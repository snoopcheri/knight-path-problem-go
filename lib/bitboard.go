package lib

import "math/bits"

type Bitboard struct {
    value uint64
}

func NewBitboard(value uint64) Bitboard {
    return Bitboard{value: value}
}

func (bb Bitboard) countOnes() int {
    return bits.OnesCount64(bb.value)
}

func (bb Bitboard) Get(square Square) bool {
    return bb.get(int(square))
}

func (bb Bitboard) get(bit int) bool {
    return bb.value&(1<<bit) != 0
}

func (bb Bitboard) Set(square Square) Bitboard {
    return bb.set(int(square))
}

func (bb Bitboard) set(bit int) Bitboard {
    return Bitboard{value: bb.value | (1 << bit)}
}

func (bb Bitboard) Cleared(square Square) Bitboard {
    return bb.cleared(int(square))
}

func (bb Bitboard) cleared(bit int) Bitboard {
    return Bitboard{value: bb.value &^ (1 << bit)}
}

func (bb Bitboard) or(other Bitboard) Bitboard {
    return Bitboard{value: bb.value | other.value}
}

func (bb Bitboard) not() Bitboard {
    return Bitboard{value: ^bb.value}
}

func (bb Bitboard) nextOne() int {
    return bits.TrailingZeros64(bb.value)
}

func (bb Bitboard) isEmpty() bool {
    return bb.value == 0
}
