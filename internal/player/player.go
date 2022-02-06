package player

import "errors"

type Player struct {
	name string
	side string
	account float64 // количество денег на счету пользователя (в рублях)
}

func NewPlayer(name, side string, account float64) (Player, error) {
	if err := validateName(name); err != nil {
		return Player{}, err
	}
	return Player{name,side,account}, nil
}

func (p *Player) Rename(name string) error {
	if err := validateName(name); err != nil {
		return err
	}
	p.name = name
	return nil
}

func (p *Player) Income(x float64) {
	p.account += x
}

func (p Player) GetAccountAsUSD() float64 {
	return p.account / 75.0
}

func validateName(name string) error {
	if len(name) < 3 {
		return errors.New("Нельзя ставить такое короткое имя")
	}
	for _, x := range name {
		if x == ' ' {
			return errors.New("В имени не может быть пробела")
		}
	}
	return nil
}