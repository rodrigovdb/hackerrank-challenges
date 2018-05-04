package main

import "fmt"

type hourglass struct {
  i     int
  j     int
  board   [6][6]int
}

func (h hourglass) sum() int {
  response := 0

  for l := 0; l < 3; l++ {
    response += h.board[h.i][h.j + l]
    response += h.board[h.i + 2][h.j + l]
  }

  response += h.board[h.i + 1][h.j + 1]

  return response
}

func (h hourglass) show(){
  fmt.Println(h.board[h.i][h.j], h.board[h.i][h.j + 1], h.board[h.i][h.j + 2])
  fmt.Println(" ", h.board[h.i + 1][h.j + 1], " ")
  fmt.Println(h.board[h.i + 2][h.j], h.board[h.i + 2][h.j + 1], h.board[h.i + 2][h.j + 2])
  fmt.Println("\n")
}

func read_input() [6][6]int{
  var input int
  var hourglasses [6][6]int

  for i := 0; i < 6; i++ {
    for j := 0; j < 6; j++ {
      fmt.Scan(&input)
      hourglasses[i][j] = input
    }
  }

  return hourglasses
}

func make_hourglasses() []hourglass{
  input     := read_input()
  response  := make([]hourglass, 0)

  for i := 0; i < 4; i++ {
    for j := 0; j < 4; j++ {
      item     := hourglass{ i: i, j: j, board: input }
      response  = append(response, item)
      //item.show()
    }
  }

  return response
}

func main() {
  hourglasses := make_hourglasses()
  higher    := hourglasses[0].sum()

  for _, item := range hourglasses {
    if item.sum() > higher {
      higher = item.sum()
    }
  }

  fmt.Println(higher)
}
