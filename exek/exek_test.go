package exek_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"com/github/chrisport/exek"
)

func TestSimpleLSCommand(t *testing.T) {
	result := exek.Call("touch asd.yxc && ls && rm asd.yxc")

	assert.Contains(t, result, "asd.yxc")
}

func TestMultipleCommands(t *testing.T) {
	sampleContent := "test-goqt"
	result := exek.Call("echo \"%v\" > /tmp/testgoqt && cat /tmp/testgoqt && rm /tmp/testgoqt", sampleContent)

	assert.Equal(t, sampleContent, result)
}

func TestPipe(t *testing.T) {
	sampleContent := "day"
	result := exek.Call("echo \"%v\" | sed s/day/night/", sampleContent)

	assert.Equal(t, "night", result)
}
