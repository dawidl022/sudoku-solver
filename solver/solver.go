package solver

const SUDOKU_SIZE = 9

type cell struct {
	value           *int
	possible_values []int
}

func Solve(board [9][9]int) ([9][9]int, error) {
	cells := initialiseCells(board)
	populatePossibleValues(cells)
	return [9][9]int{}, nil
}

func initialiseCells(board [9][9]int) [9][9]cell {
	cells := [9][9]cell{}

	for i := 0; i < SUDOKU_SIZE; i++ {
		for j := 0; j < SUDOKU_SIZE; j++ {
			var value *int
			if board[i][j] != 0 {
				value = &board[i][j]
			}
			cells[i][j] = cell{value: value}
		}
	}
	return cells
}

func possibleValuesForCell(cells [9][9]cell, i int, j int) []int {

}
