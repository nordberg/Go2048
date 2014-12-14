package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Tuple struct {
	a, b int
}

const (
	width  = 4
	height = 4
)

func main() {
	score := 0

	fmt.Println("Welcome to 2048!")
	fmt.Println("R: Right, U: Up, D: Down, L: Left (lower case works)")

	// Two dimensional array (the board)
	board := make([][]int, height)
	for i := range board {
		board[i] = make([]int, width)
	}

	for {
		if !addTwo(board) {
			// Failed to add a new 2, game over
			break
		}

		fmt.Println(printBoard(board))

		userInput(board)
	}

	fmt.Println("Game over! Score: ", score)

}

func addTwo(board [][]int) bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// A list with possible spots for the new 2.
	openSpots := list.New()

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

	// No open spots to add 2, game over
	if openSpots.Len() == 0 {
		return false
	}

	// Add the new 2 at random place
	index := r.Intn(openSpots.Len())

	// Iterate through open spots until index is found
	// Is there a better way to do this?
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
	// Read user input for correct tile movement
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.ToLower(input)
	input = strings.TrimSpace(input)
	updateBoard(board, input)
}

func updateBoard(board [][]int, dir string) {
	switch dir {
	case "u":

	case "d":

	case "l":

	case "r":

	default:

	}

	// THREE nested for-loops, I like to live dangerously
	for j := 0; j < width; j++ {
		change := true
		for change {
			change = false
			for i := height - 1; i > 0; i-- {
				if board[i][j] != 0 && board[i][j] == board[i-1][j] {
					// Two equal to be merged into 2x
					board[i][j] *= -2
					board[i-1][j] = 0
					change = true
				} else if board[i][j] == 0 && board[i-1][j] != 0 {
					// A number has 0 below, move it down
					board[i][j] = board[i-1][j]
					board[i-1][j] = 0
					change = true
				}
			}
		}

		for i := height - 1; i > 0; i-- {
			if board[i][j] < 0 {
				board[i][j] *= -1
			}
		}
	}
}

// Print the board with format:
// [X][X][X][X]
// [X][X][X][X]
// [X][X][X][X]
// [X][X][X][X]
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
