package main

import (
    "fmt"
    "math/rand"
    "time"
    "container/list"
)

type Tuple struct {
    a, b interface{}
}

const (
    width = 4
    height = 4
)

func main() {
    fmt.Println("Welcome to 2048!")

    //round := 1
    board := make([][]int, height)

    for i := range board {
        board[i] = make([]int, width)
    }

    addTwo(board)
}

func addTwo(board [][]int) {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    openSpots := list.New()

    board[0][1] = 2

    for i := range board {
        for j := range board[i] {
            if board[i][j] == 0 {
                tupl := Tuple{i, j}
                openSpots.PushFront(tupl)
            }
        }
    }

    index := r.Intn(openSpots.Len())

    e := openSpots.Front()
    for i := 0; i < openSpots.Len(); i++ {
        if i == index {
            fmt.Println(e.Value)
            return
        }
        e = e.Next()
    }
}

