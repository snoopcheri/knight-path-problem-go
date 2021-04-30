package lib

import "fmt"

func GenerateDistinctBoards(whiteKnightsCount int, blackKnightsCount int) []Board {
    var boards []Board
    boards = append(boards, NewBoard(NewBitboard(0), NewBitboard(0)))

    for wk := 0; wk <whiteKnightsCount; wk++ {
        boards = boardsWithKnight(boards, WHITE)
    }

    for bk := 0; bk <blackKnightsCount; bk++ {
        boards = boardsWithKnight(boards, BLACK)
    }

    var distinctBoards = unique(boards)

    fmt.Println("Generated", len(boards), "boards (#distinct=", len(distinctBoards), ")")

    return boards
}

func boardsWithKnight(boards []Board, colour Colour) []Board {
    var newBoards []Board

    for _, board := range boards {
        newBoards = append(newBoards, boardsForBoardWithKnight(board, colour)...)
    }

    return newBoards
}

func boardsForBoardWithKnight(board Board, colour Colour) []Board {
    var boards []Board
    var emptySquares = NewBitboardIterator(board.emptySquares())

    for emptySquares.HasNext() {
        var emptySquare = emptySquares.Next()
        boards = append(boards, board.withKnight(Square(emptySquare), colour))
    }

    return boards
}

func unique(boardsSlice []Board) []Board {
    keys := make(map[Board]bool)
    list := []Board{}
    for _, entry := range boardsSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}
