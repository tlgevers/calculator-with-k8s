package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	args := "1+1"
	sum := Add(args)
	if sum != "2.000000" {
		t.Error(Add(args))
	}
}
