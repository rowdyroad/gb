package b

import "test/pkg/a"

var BVar string

func Print() string {
	return a.AVar + BVar
}