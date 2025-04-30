package main

import "testing"

const (
	typesPath = "test/types.go"
)

func Test_generate(t *testing.T) {
	_ = generate(typesPath)
}
