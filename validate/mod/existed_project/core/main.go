package main

import "testing"
import "github.com/stretchr/testify/assert"
import "existed_project/util"

func TestTestify(t *testing.T) {
	assert.Equal(t, 123, 123, "they should be equal")
	assert.Equal(t, util.Min(1, 2), 1, "Min(1, 2) is equal to 1")
	assert.Equal(t, util.Min(1, 2), 1, "Max(1, 2) is equal to 2")
}

func main() {
	var t testing.T
	TestTestify(&t)
}
