package main

import (
	"bytes"
	"testing"
)

func TestSimpleBlocks(t *testing.T) {
	lines := [][]byte{
		[]byte("From joe@test.com Thu May 23 22:00:04 2013"),
		[]byte("From joe@test.com Thu May 23 22:00:04 2013"),
		[]byte("From joe@test.com Thu May 23 22:00:05 2013"),
		[]byte("From joe@test.com Thu May 23 22:00:06 2013"),
		[]byte("From joe@test.com Thu May 23 22:00:07 2013"),
		[]byte("From joe@test.com Thu May 23 22:00:07 2013"),
		[]byte("From joe@test.com Thu May 23 22:00:08 2013"),
		[]byte("From joe@test.com Thu May 23 22:00:08 2013"),
		[]byte("From joe@test.com Thu May 23 22:00:09 2013"),
		[]byte("From joe@test.com Thu May 23 22:00:10 2013"),
	}
	r := bytes.NewReader(bytes.Join(lines, []byte("\n\n")))
	msgs, err := mboxMessageBlocks(r)
	if err != nil {
		t.Fatal(err)
	}
	for msg := range msgs {
		t.Log(string(msg))
	}
}
