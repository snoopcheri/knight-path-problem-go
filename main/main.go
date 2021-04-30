package main

import (
    "GoTest/lib"
    "time"
)
import "fmt"

func main() {
    fmt.Println("STARTED AT", time.Now())

    var startBoard = lib.NewBoard(
        lib.NewBitboard(0).Set(lib.B1).Set(lib.G1),
        lib.NewBitboard(0).Set(lib.B8).Set(lib.G8),
    )

    var endBoard = startBoard.WithSwitchedKnights()

    var allBoards = lib.GenerateDistinctBoards(2, 2)
    var solver = lib.NewSolver(allBoards)

    var solutions = solver.SolveFor(endBoard)

    fmt.Println("From board")
    fmt.Println(startBoard)
    fmt.Println("To board")
    fmt.Println(endBoard)
    fmt.Println("#steps required = ", solutions[startBoard])

    var printer = lib.NewPathFinder(solutions)
    printer.PrintPath(startBoard, endBoard)

    fmt.Println("FINISHED AT", time.Now())
}
