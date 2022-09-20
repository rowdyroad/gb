package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHappyPath(t *testing.T) {
	ttt, err := NewTicTacToe(3)
	if !assert.NoError(t, err) {
		return
	}
	if !assert.Equal(t, 3, len(ttt.field)) {
		return
	}
	if !assert.Equal(t, 3, len(ttt.field[0])) {
		return
	}

	s, err := ttt.Move(1, 1)
	if !assert.NoError(t, err) {
		return
	}
	if !assert.Equal(t, None, s) {
		return
	}

	s, err = ttt.Move(0, 0)
	if !assert.NoError(t, err) {
		return
	}
	if !assert.Equal(t, None, s) {
		return
	}
	s, err = ttt.Move(0, 1) //x
	if !assert.NoError(t, err) {
		return
	}
	if !assert.Equal(t, None, s) {
		return
	}
	s, err = ttt.Move(1, 0)
	if !assert.NoError(t, err) {
		return
	}
	if !assert.Equal(t, None, s) {
		return
	}
	s, err = ttt.Move(2, 1) //x
	if !assert.NoError(t, err) {
		return
	}
	if !assert.Equal(t, X, s) {
		return
	}
}

func TestCreateErrors(t *testing.T) {
	ttt, err := NewTicTacToe(1)
	if !assert.Error(t, err) {
		return
	}
	assert.Nil(t, ttt)

	ttt, err = NewTicTacToe(0)
	if !assert.Error(t, err) {
		return
	}
	assert.Nil(t, ttt)

	ttt, err = NewTicTacToe(-1)
	if !assert.Error(t, err) {
		return
	}
	assert.Nil(t, ttt)
}

func TestCellError(t *testing.T) {
	ttt, err := NewTicTacToe(3)
	if !assert.NoError(t, err) {
		return
	}
	_, err = ttt.Move(-1, 0)
	if !assert.Error(t, err) {
		return
	}

}

func TestBusyCellError(t *testing.T) {
	ttt, err := NewTicTacToe(3)
	if !assert.NoError(t, err) {
		return
	}
	_, err = ttt.Move(10, 0)
	if !assert.Error(t, err) {
		return
	}

	_, err = ttt.Move(0, 0)
	if !assert.NoError(t, err) {
		return
	}

	_, err = ttt.Move(0, 0)
	if !assert.Error(t, err) {
		return
	}
}
