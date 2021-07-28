package rps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PlayRound(t *testing.T) {
	expected := Round{1, "", ""}
	actual := PlayRound(1)

	assert.Equal(t, expected, actual, "The result should be the same")
}
