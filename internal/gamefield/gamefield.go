package gamefield

import "errors"

type Gamefield struct {
	//
	size uint // Size * Size
}

func NewGamefield(size uint) (Gamefield, error) {
	if size == 0 {
		return Gamefield{}, errors.New("Нельзя создавать поле размером 0х0")
	}
	return Gamefield{size}, nil
}

func (g *Gamefield) SetCell(value string) error {

	return nil
}

func (g Gamefield) HaveEmptyCells() bool {

	return false
}

func (g *Gamefield) CheckForWin() (string, bool) {

	return "", false
}
