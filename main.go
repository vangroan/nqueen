package main

import (
	"fmt"
	"strings"
)

// Board dimensions
const dim int = 16

type board struct {
	data [][]uint8
	ld   []uint8
	rd   []uint8
	cl   []uint8
}

func newBoard(n int) board {
	b := make([][]uint8, 0)
	for i := 0; i < n; i++ {
		row := make([]uint8, n)
		b = append(b, row)
	}
	return board{
		data: b,
		ld:   make([]uint8, n*2),
		rd:   make([]uint8, n*2),
		cl:   make([]uint8, n),
	}
}

func (b board) String() string {
	sb := make([]string, 0)

	for _, col := range b.data {
		for _, val := range col {
			if val == 0 {
				sb = append(sb, "- ")
			} else {
				sb = append(sb, "0 ")
			}
		}
		sb = append(sb, "\n")
	}

	return strings.Join(sb, "")
}

type coord struct {
	x int
	y int
}

func isSafe(b *board, col, row, n int) bool {
	if b.cl[row] == 1 {
		return false
	}

	// diagonal \
	if b.ld[row-col+n-1] == 1 {
		return false
	}

	// diagonal /
	if b.rd[col+row] == 1 {
		return false
	}

	return true
}

func solve(b *board, col, n int) bool {
	// base case for termination
	if col >= n {
		return true
	}

	// consider the current column and attempt to
	// place the queen in each row
	for row := 0; row < n; row++ {
		if isSafe(b, col, row, n) {
			// light diagonal index
			ld := row - col + n - 1
			// fmt.Printf("row %d - col %d + N %d - 1 = %d\n", col, row, n, rd)

			// place queen
			b.data[col][row] = 1
			b.rd[col+row] = 1
			b.ld[ld] = 1
			b.cl[row] = 1

			// recur to place rest of the queens
			if solve(b, col+1, n) {
				// keep solution
				return true
			}

			// reset queen for backtrace
			b.data[col][row] = 0
			b.rd[col+row] = 0
			b.ld[ld] = 0
			b.cl[row] = 0
		}
	}

	return false
}

func main() {
	fmt.Println("Starting")
	b := newBoard(dim)

	if solve(&b, 0, dim) {
		fmt.Println("Solution found")
		fmt.Println(b.String())
	} else {
		fmt.Println("Failed")
	}
}
