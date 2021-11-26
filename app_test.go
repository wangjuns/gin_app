package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestRandomPassword(t *testing.T) {
	pwd := randomPassword(10)
	assert.Equal(t, 10, len(pwd))
}
