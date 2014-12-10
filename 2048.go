package main

import (
    "fmt"
    "math/rand"
    "time"
    "container/list"
    "strconv"
    "bufio"
    "os"
    "strings"
)

type Tuple struct {
    a, b int
}

type key struct {
    x  int
    y  int
    ch rune
}

const (
    width = 4
    height = 4
)

var K_ARROW_UP = []key{{54, 10, '('}, {55, 10, 0x2191}, {56, 10, ')'}}
var K_ARROW_LEFT = []key{{50, 12, '('}, {51, 12, 0x2190}, {52, 12, ')'}}
var K_ARROW_DOWN = []key{{54, 12, '('}, {55, 12, 0x2193}, {56, 12, ')'}}
var K_ARROW_RIGHT = []key{{58, 12, '('}, {59, 12, 0x2192}, {60, 12, ')'}}

func main() {
    gameOver := false
    score := 0

    fmt.Println("Welcome to 2048!")
    fmt.Println("R: Right, U: Up, D: Down, L: Left (lower case works)")

    //round := 1
    board := make([][]int, height)

    for i := range board {
        board[i] = make([]int, width)
    }

    for !gameOver {
        if !addTwo(board) {
            break
        }

        fmt.Println(printBoard(board))

        userInput(board)
    }

    fmt.Println("Game over! Score: ", score)

    
}

func addTwo(board [][]int) bool {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    // A list with possible spots for the new number. If length = 0 -> Game Over
    openSpots := list.New();

    for i := range board {
        for j := range board[i] {
            // If empty, add to openSpots list
            if board[i][j] == 0 {
                //tupl := []int{i, j}
                tupl := Tuple{i, j}
                openSpots.PushFront(tupl)
            }
        }
    }

    if (openSpots.Len() == 0) {
        return false
    }

    index := r.Intn(openSpots.Len())

    e := openSpots.Front()
    for i := 0; i < openSpots.Len(); i++ {
        if i == index {
            var a int = e.Value.(Tuple).a
            var b int = e.Value.(Tuple).b
            board[a][b] = 2
            return true
        }

        e = e.Next()
    }

    return true
}

func userInput(board [][]int) {
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')

    input = strings.ToLower(input)

    switch input {
    case "u":
        
    case "d":

    case "l":

    case "r":

    default:

    }

}

func printBoard(board [][]int) string {

    var output string

    for i := 0; i < width; i++ {
        for j := 0; j < height; j++ {
            output += "[" + strconv.Itoa(board[i][j]) + "]"
            //output += "[ " + strconv.Itoa(board[i][j]) + " ]"
        }
        output += "\n"
    }

    return output
}