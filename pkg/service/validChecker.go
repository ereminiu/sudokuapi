package service

import (
	"github.com/ereminiu/sudokuapi/pkg/models"
	"github.com/ereminiu/sudokuapi/pkg/service/tools"
)

// 1 - it's correct
// 2 - same digits at the same row
// 3 - same digits at the same column
func IsValid(grid models.Grid) (code int) {
	n := 9
	if len(grid.Array) != n {
		return models.UnvalidRowAmount
	}
	if len(grid.Array[0]) != n {
		return models.UnvalidColAmount
	}

	row := make([]map[int]bool, n)
	col := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		row[i] = make(map[int]bool, n)
		col[i] = make(map[int]bool, n)
	}

	for i, r := range grid.Array {
		for j, val := range r {
			if val == "." {
				continue
			}

			me := tools.ToInt(val)
			if row[i][me] {
				return 2
			}
			if col[j][me] {
				return 3
			}

			row[i][me] = true
			col[j][me] = true
		}
	}

	return 1
}
