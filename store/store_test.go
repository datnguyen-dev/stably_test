package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	val := checkIsPrime(7)
	assert.True(t, val)
	val = checkIsPrime(17)
	assert.True(t, val)
	val = checkIsPrime(70)
	assert.False(t, val)
}

func TestFindPrime(t *testing.T) {
	val := FindPrime(55)
	assert.True(t, val == 53)

	val = FindPrime(101)
	assert.True(t, val == 97)

	val = FindPrime(2)
	assert.False(t, val == 2)
}
