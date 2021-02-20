package main
import "fmt"

type board [8][8]string

func (b board) print() {
  for _, row := range b {
    fmt.Println(row)
  }
}

func main() {
  var board board// 8 x 8 2D array

  for y, row := range board {
    for x := range row {
      if (y + x) % 2 == 0 {
        board[y][x] = "X"
      } else {
        board[y][x] = "_"
      }
    }
  }
  board.print()
}
