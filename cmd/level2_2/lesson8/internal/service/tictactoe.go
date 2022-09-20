package service

import "errors"

type Sign int

const (
	None Sign = 0
	X    Sign = 1
	Zero Sign = 2
)

var (
	ErrOutOfBounds        = errors.New("out of bounds")
	ErrIncorrectFieldSize = errors.New("incorrect field size")
	ErrCellIsAlreadyBusy  = errors.New("cell is already busy")
)

type TicTacToe struct {
	size    int
	field   [][]Sign
	current Sign
}

func NewTicTacToe(size int) (*TicTacToe, error) {
	if size < 2 {
		return nil, ErrIncorrectFieldSize
	}
	field := make([][]Sign, size)
	for i := 0; i < size; i++ {
		field[i] = make([]Sign, size)
	}
	return &TicTacToe{size: size, field: field, current: X}, nil
}

func NewTicTacToeFromField(field [][]Sign, current Sign) (*TicTacToe, error) {
	return &TicTacToe{size: len(field), field: field, current: current}, nil
}

func (t *TicTacToe) GetField() [][]Sign {
	return t.field
}

func (t *TicTacToe) Current() Sign {
	return t.current
}

func (t *TicTacToe) Move(row, col int) (Sign, error) {
	if row < 0 || col < 0 || row >= len(t.field) || col >= len(t.field[0]) {
		return None, ErrOutOfBounds
	}

	if t.field[row][col] != None {
		return None, ErrCellIsAlreadyBusy
	}

	t.field[row][col] = t.current

	if win := t.checkWin(row, col); win != None {
		return win, nil
	}

	if t.current == X {
		t.current = Zero
	} else {
		t.current = X
	}

	return None, nil
}

func (t *TicTacToe) checkWin(row, col int) Sign {
	c := 0
	for i := 0; i < row; i++ {
		if t.field[i][col] == t.current {
			c++
		} else {
			break
		}
	}
	for i := row; i < t.size; i++ {
		if t.field[i][col] == t.current {
			c++
		} else {
			break
		}
	}

	if c == t.size {
		return t.current
	}
	c = 0
	for i := 0; i < col; i++ {
		if t.field[row][i] == t.current {
			c++
		} else {
			break
		}
	}
	for i := col; i < t.size; i++ {
		if t.field[row][i] == t.current {
			c++
		} else {
			break
		}
	}
	if c == t.size {
		return t.current
	}
	return None
}
