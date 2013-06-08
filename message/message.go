package message

import (
	"bytes"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/mail"
	"strings"
)

type Message struct {
	Message *mail.Message
	Text    []byte
	HTML    []byte
}

func Parse(b []byte) (m *Message, err error) {
	return toMessage(b)
}

func toMessage(b []byte) (m *Message, err error) {
	m = &Message{}
	m.Message, err = mail.ReadMessage(bytes.NewReader(b))
	if err != nil {
		return
	}

	bound, err := getBoundary(m.Message.Header)
	// err will be "not found", assume message body is plain text
	if err != nil {
		m.Text, err = ioutil.ReadAll(m.Message.Body)
		return
	}

	mr := multipart.NewReader(m.Message.Body, bound)
	for {
		part, err := mr.NextPart()
		if err != nil {
			break
		}

		ct, err := getContentType(part.Header)
		if err != nil {
			return m, err
		}

		switch ct {
		case "text/plain":
			m.Text, err = ioutil.ReadAll(part)
		case "text/html":
			m.HTML, err = ioutil.ReadAll(part)
		}
	}
	return
}

func getBoundary(h map[string][]string) (b string, err error) {
	header, ok := h["Content-Type"]
	if !ok {
		err = errors.New("No Content-Type header found")
		return
	}
	s := header[0]
	seek := "boundary="
	idx := strings.Index(s, seek)
	if idx == -1 {
		err = errors.New("No boundary defined in Content-Type")
		return
	}
	b = s[idx+len(seek):]
	return
}

func getContentType(h map[string][]string) (t string, err error) {
	header, ok := h["Content-Type"]
	if !ok {
		err = errors.New("No Content-Type header found")
		return
	}
	s := header[0]
	if idx := strings.Index(s, ";"); idx > -1 {
		t = s[:idx]
	} else {
		t = s
	}
	return
}
