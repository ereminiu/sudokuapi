package service

import (
	"fmt"

	"github.com/ereminiu/apitest/pkg/models"
	"github.com/ereminiu/apitest/pkg/service/tools"
)

var n int = 9

var row []map[int]bool
var col []map[int]bool
var area [][]map[int]bool

func find_empty(board [][]rune) []int {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				return []int{i, j}
			}
		}
	}
	return nil
}

func isValid(x int, y int, val int) bool {
	return !(row[x][val] || col[y][val] || area[x/3][y/3][val])
}

func solve(board [][]rune) bool {
	empty_cell := find_empty(board)
	if empty_cell == nil {
		return true
	}

	x, y := empty_cell[0], empty_cell[1]
	for val := 1; val < 10; val++ {
		if !isValid(x, y, val) {
			continue
		}

		board[x][y] = rune('0' + val)
		row[x][val] = true
		col[y][val] = true
		area[x/3][y/3][val] = true

		if solve(board) {
			return true
		}

		board[x][y] = '.'
		row[x][val] = false
		col[y][val] = false
		area[x/3][y/3][val] = false
	}

	return false
}

func printBoard(board [][]rune) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%q ", board[i][j])
		}
		fmt.Println()
	}
}

func Solve(board models.Grid) [][]rune {
	row = make([]map[int]bool, n)
	col = make([]map[int]bool, n)
	area = make([][]map[int]bool, n/3)

	for i := 0; i < n; i++ {
		row[i] = make(map[int]bool)
		col[i] = make(map[int]bool)
	}

	for i := 0; i < n/3; i++ {
		area[i] = make([]map[int]bool, n/3)
		for j := 0; j < n/3; j++ {
			area[i][j] = make(map[int]bool)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board.Array[i][j] == "." {
				continue
			}

			val := tools.ToInt(board.Array[i][j])

			row[i][val] = true
			col[j][val] = true
			area[i/3][j/3][val] = true
		}
	}

	ans := make([][]rune, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			if board.Array[i][j] == "." {
				ans[i][j] = '.'
				continue
			}
			ans[i][j] = rune('0' + tools.ToInt(board.Array[i][j]))
		}
	}

	solve(ans)

	printBoard(ans)

	return ans
}
