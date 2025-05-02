package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	typesPath = "test/data/types.go"
)

func Test_generate(t *testing.T) {
	err := ConfigureLogger("debug")
	require.NoError(t, err)

	err = generate(typesPath)
	assert.NoError(t, err)
}
