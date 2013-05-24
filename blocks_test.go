package main

import (
	"bytes"
	"testing"
)

func TestSimpleBlocks(t *testing.T) {
	lines := [][]byte{
		[]byte("From joe01@test.com Thu May 23 22:00:04 2013\n"),
		[]byte("From joe02@test.com Thu May 23 22:00:04 2013\n"),
		[]byte("From joe03@test.com Thu May 23 22:00:05 2013\n"),
		[]byte("From joe04@test.com Thu May 23 22:00:06 2013\n"),
		[]byte("From joe05@test.com Thu May 23 22:00:07 2013\n"),
		[]byte("From joe06@test.com Thu May 23 22:00:07 2013\n"),
		[]byte("From joe07@test.com Thu May 23 22:00:08 2013\n"),
		[]byte("From joe08@test.com Thu May 23 22:00:08 2013\n"),
		[]byte("From joe09@test.com Thu May 23 22:00:09 2013\n"),
		[]byte("From joe10@test.com Thu May 23 22:00:10 2013\n"),
	}
	r := bytes.NewReader(bytes.Join(lines, []byte{}))

	blocks := make(chan []byte, 0)
	go mboxMessageBlocks(r, blocks)

	i := 0
	for block := range blocks {
		exp := string(lines[i])
		if exp != string(block) {
			t.Errorf("Incorrect block.\nExpect: '%s'\nGot: '%s'", exp, string(block))
		}
		i++
	}
}
