package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getErr() error {
	return errors.New("some error")
}

func TestSimple(t *testing.T) {
	if !assert.Equal(t, 1, 1) {
		return
	}
	if !assert.Error(t, getErr()) {
		return
	}
}