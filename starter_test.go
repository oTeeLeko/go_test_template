package starter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	greeting := SayHello("Meee")
	assert.Equal(t, "Hello Meee. Welcome!", greeting)

	another_greeting := SayHello("Fern")
	assert.Equal(t, "Hello Fern. Welcome!", another_greeting)
}

func TestEvenOrOdd(t *testing.T) {
	t.Run("Test Non Negative Even Number", func(t *testing.T) {
		assert.Equal(t, "50 is an even number", evenOrOdd(50))
		assert.Equal(t, "42 is an even number", evenOrOdd(42))
		assert.Equal(t, "0 is an even number", evenOrOdd(0))
	})
	t.Run("Test Negative Even Number", func(t*testing.T) {
		assert.Equal(t, "-50 is an even number", evenOrOdd(-50))
		assert.Equal(t, "-42 is an even number", evenOrOdd(-42))
	})
	t.Run("Test Non Negative Odd Number", func(t*testing.T) {
		assert.Equal(t, "3 is an odd number", evenOrOdd(3))
		assert.Equal(t, "9 is an odd number", evenOrOdd(9))
		assert.Equal(t, "55 is an odd number", evenOrOdd(55))
	})
	t.Run("Test Negative Odd Number", func(t*testing.T) {
		assert.Equal(t, "-3 is an odd number", evenOrOdd(-3))
		assert.Equal(t, "-9 is an odd number", evenOrOdd(-9))
		assert.Equal(t, "-55 is an odd number", evenOrOdd(-55))
	})
}