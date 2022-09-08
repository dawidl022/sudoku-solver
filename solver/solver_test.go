package solver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialiseCells_UsesArrayValuesAsCellValues(t *testing.T) {
	matrix := [9][9]int{
		{7, 8, 7, 3, 6, 3, 7, 6, 1},
		{6, 4, 6, 1, 9, 1, 8, 9, 2},
		{3, 3, 4, 5, 5, 2, 9, 7, 6},
		{2, 6, 6, 7, 9, 3, 9, 6, 1},
		{3, 8, 6, 1, 8, 1, 4, 2, 9},
		{5, 2, 4, 9, 9, 8, 5, 1, 5},
		{9, 4, 2, 6, 9, 3, 4, 1, 3},
		{3, 6, 2, 6, 3, 8, 4, 4, 2},
		{8, 4, 2, 5, 9, 4, 4, 3, 6},
	}

	cells := initialiseCells(matrix)

	for i := 0; i < SUDOKU_SIZE; i++ {
		for j := 0; j < SUDOKU_SIZE; j++ {
			assert.Equal(t, cell{value: &matrix[i][j]}, cells[i][j])
		}
	}
}

func TestInitialiseCells_TurnsZerosIntoNilValues(t *testing.T) {
	matrix := [9][9]int{
		{7, 8, 7, 0, 6, 3, 0, 6, 0},
	}

	cells := initialiseCells(matrix)

	assert.Equal(t, [9][9]cell{
		{cell{value: &matrix[0][0]}, cell{value: &matrix[0][1]}, cell{value: &matrix[0][2]},
			cell{value: nil}, cell{value: &matrix[0][4]}, cell{value: &matrix[0][5]},
			cell{value: nil}, cell{value: &matrix[0][7]}, cell{value: nil}},
	}, cells)
}

func TestInitialiseCells_StartsWithNoPossibleValues(t *testing.T) {
	matrix := [9][9]int{
		{7, 8, 7, 3, 6, 3, 7, 6, 1},
		{6, 4, 6, 1, 9, 1, 8, 9, 2},
		{3, 3, 4, 5, 5, 2, 9, 7, 6},
		{2, 6, 6, 7, 9, 3, 9, 6, 1},
		{3, 8, 6, 1, 8, 1, 4, 2, 9},
		{5, 2, 4, 9, 9, 8, 5, 1, 5},
		{9, 4, 2, 6, 9, 3, 4, 1, 3},
		{3, 6, 2, 6, 3, 8, 4, 4, 2},
		{8, 4, 2, 5, 9, 4, 4, 3, 6},
	}

	cells := initialiseCells(matrix)
	var empty []int

	for i := 0; i < SUDOKU_SIZE; i++ {
		for j := 0; j < SUDOKU_SIZE; j++ {
			assert.Equal(t, empty, cells[i][j].possible_values)
		}
	}
}
