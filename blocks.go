package main

import (
	"bytes"
	"io"
)

// Reads an mbox-formatted file and returns the message blocks. The format
// dictates that each message begins with a "From " line
func mboxMessageBlocks(r io.Reader) (msgs <-chan []byte, err error) {
	msgs = make(chan []byte, 1)
	return
}
