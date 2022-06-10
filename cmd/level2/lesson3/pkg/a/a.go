package a

import "test/pkg/b"

var AVar string

func Print() string {
	return AVar + b.BVar
}