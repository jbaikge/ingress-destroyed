package main

import (
	"bufio"
	"bytes"
	"io"
)

// Reads an mbox-formatted file and returns the message blocks. The format
// dictates that each message begins with a "From " line
func mboxMessageBlocks(r io.Reader, blocks chan []byte) {
	defer close(blocks)

	br := bufio.NewReader(r)
	buf := bytes.NewBuffer(make([]byte, 0, 4096))

	for {
		line, err := br.ReadBytes('\n')
		if err != nil && err != io.EOF {
			break
		}

		// Send buffered message when encountering a new message
		if len(line) > 5 && string(line[0:5]) == "From " && buf.Len() > 0 {
			block := make([]byte, buf.Len())
			copy(block, buf.Bytes())
			blocks <- block
			buf.Reset()
		}
		buf.Write(line)

		// Send last buffered message
		if err == io.EOF {
			blocks <- buf.Bytes()
			break
		}
	}
}
