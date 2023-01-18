package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	t.Parallel()
	assert.Equal(t, Sub(1, 2), -1)
}
