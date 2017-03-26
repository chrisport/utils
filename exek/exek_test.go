package exek_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"com/github/chrisport/utils/exek"
)

func TestSimpleLSCommand(t *testing.T) {
	out := exek.Call("touch kspo.yxc && ls && rm kspo.yxc")

	assert.Contains(t, out, "kspo.yxc")
}

func TestMultipleCommands(t *testing.T) {
	sampleContent := "test-goqt"
	out := exek.Call("echo \"%v\" > /tmp/testgoqt && cat /tmp/testgoqt && rm /tmp/testgoqt", sampleContent)

	assert.Equal(t, sampleContent, out)
}

func TestPipe(t *testing.T) {
	sampleContent := "day"
	out := exek.Call("echo \"%v\" | sed s/day/night/", sampleContent)

	assert.Equal(t, "night", out)
}

func TestCalltoLines(t *testing.T) {
	exek.Call("touch kspo.zxc0 && touch kspo.zxc1 && touch kspo.zxc2")
	out := exek.CallToLines("ls | grep kspo")
	exek.Call("rm kspo.zxc*")

	assert.Equal(t, "kspo.zxc0", out[0])
	assert.Equal(t, "kspo.zxc1", out[1])
	assert.Equal(t, "kspo.zxc2", out[2])
}
