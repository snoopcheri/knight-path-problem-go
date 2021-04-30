package lib

import (
    "fmt"
)

type Solver struct {
    boards []Board
    solutions map[Board]map[Board]int
}

func NewSolver(boards []Board) Solver {
    return Solver{
        boards: boards,
    }
}

func (solver Solver) SolveFor(endBoard Board) map[Board]int {
    var unsolved = make(map[Board]bool)
    for _,board := range solver.boards {
        unsolved[board] = true
    }
    var solved = make(map[Board]int)
    var depth = 0

    solver.markAsSolved(endBoard, 0, unsolved, solved)

    for true {
        fmt.Print("depth=", depth, ", #unsolved=", len(unsolved), ", #solved=", len(solved), "")

        var solvedAtDepth []Board

        for board := range unsolved {
            var successors = successorsOf(board)
            var minDepth = minDepth(successors, solved)

            if minDepth != -1 {
                solvedAtDepth = append(solvedAtDepth, board)
            }
        }

        fmt.Println(" -> found", len(solvedAtDepth), "new solutions")

        if len(solvedAtDepth) == 0 {
            break;
        }

        for _, solvedBoard := range solvedAtDepth {
            solver.markAsSolved(solvedBoard, depth+1, unsolved, solved)
        }

        depth++
    }

    return solved
}

func minDepth(successors []Board, solved map[Board]int) int {
    var min = -1

    for _, successor := range successors {
        if depth, ok := solved[successor]; ok {
            if min == -1 {
                min = depth
            } else if depth < min {
                min = depth
            }
        }
    }

    return min
}

func successorsOf(board Board) []Board {
    var successors []Board
    var moves = board.GenerateMoves()

    for _, move := range moves {
        successors = append(successors, board.afterMove(move))
    }

    return successors
}

func (solver Solver) markAsSolved(board Board, depth int, unsolved map[Board]bool, solved map[Board]int) {
    delete(unsolved, board)
    solved[board] = depth
}
