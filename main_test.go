package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {
	// Generate a random ordering for 5 users
	numUsers := 5
	o := makeOrder(numUsers)
	// The order should consist of a forward and backward pass
	assert.Equal(t, numUsers*2, len(o))
	// The values should be the same
	for i := 0; i < numUsers; i++ {
		assert.Equal(t, o[i], o[len(o)-(i+1)])
	}
	// The values should be distinct and in range
	assert.ElementsMatch(t, []int{0, 1, 2, 3, 4}, o[:numUsers])
}

func TestValidPrefs(t *testing.T) {
	assert.Equal(t, true, validPrefs([]int{0}))
	assert.Equal(t, true, validPrefs([]int{0, 1, 2, 3, 4}))
	assert.Equal(t, true, validPrefs([]int{4, 3, 2, 1, 0}))
	assert.Equal(t, false, validPrefs([]int{0, 0}))
	assert.Equal(t, false, validPrefs([]int{0, 1, 2, 3, -1}))
	assert.Equal(t, false, validPrefs([]int{0, 1, 2, 3, 5}))
}

func TestPickItem(t *testing.T) {
	prefs := []int{4, 3, 0, 1, 2}
	taken := make(map[int]bool)
	p, i := pickItem(prefs, taken)
	assert.Equal(t, 4, p)
	assert.Equal(t, 0, i)
	p, i = pickItem(prefs, taken)
	assert.Equal(t, 3, p)
	assert.Equal(t, 1, i)
	taken[0] = true
	p, i = pickItem(prefs, taken)
	assert.Equal(t, 1, p)
	assert.Equal(t, 3, i)
	taken[2] = true
	p, i = pickItem(prefs, taken)
	assert.Equal(t, -1, p)
	assert.Equal(t, -1, i)
}
