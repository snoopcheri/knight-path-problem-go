package lib

import "fmt"

type PathFinder struct {
    depth     map[Board]int
    solutions [][]Move
}

func NewPathFinder(solutions map[Board]int) PathFinder {
    return PathFinder{depth: solutions}
}

func (printer PathFinder) PrintPath(start Board, end Board) {

    var empty []Move
    printer.printPath(start, end, empty)
}

func (printer PathFinder) printPath(start Board, end Board, path []Move) {
    if end == start {
        printer.storeSolution(path)
        return
    }

    var successors = printer.successorsOf(start, printer.depth[start])

    var lastElem *Move = nil
    if len(path) > 0 {
        var lastIdx = len(path) - 1
        lastElem = &path[lastIdx]
    }
    var successorMove, successorBoard = printer.successorWithFollowupMove(successors, lastElem)
    printer.printPath(successorBoard, end, append(path, successorMove))
}

func (printer PathFinder) storeSolution(path []Move) {
    printer.solutions = append(printer.solutions, path)

    var cnt = len(printer.solutions)
    fmt.Println("#", cnt, path)
}

func (printer PathFinder) successorsOf(board Board, depth int) []MoveAndBoard {
    var moves = board.GenerateMoves()
    var successors []MoveAndBoard

    for _, move := range moves {
        var moveAndBoard = MoveAndBoard{move: move, board: board.afterMove(move)}
        if printer.depth[moveAndBoard.board] == depth-1 {
            successors = append(successors, moveAndBoard)
        }
    }

    return successors
}

func (printer PathFinder) successorWithFollowupMove(successors []MoveAndBoard, lastMove *Move) (Move, Board) {
    if lastMove != nil {
        for _, successor := range successors {
            if successor.move.from == lastMove.to {
                return successor.move, successor.board
            }
        }

    }

    var first = successors[0]
    return first.move, first.board
}

type MoveAndBoard struct {
    move  Move
    board Board
}
