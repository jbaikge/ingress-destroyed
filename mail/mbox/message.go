package mbox

import (
	"bytes"
	"errors"
	"github.com/jbaikge/ingress-destroyed/message"
	"io/ioutil"
	"mime/multipart"
	"net/mail"
	"strings"
)

func toMessage(b []byte, m *message.Message) (err error) {
	msg, err := mail.ReadMessage(bytes.NewReader(b))
	if err != nil {
		return
	}

	m.Id = msg.Header.Get("Message-ID")
	m.From = msg.Header.Get("From")
	if m.Date, err = msg.Header.Date(); err != nil {
		return
	}

	bound, err := getBoundary(msg.Header)
	// err will be "not found", assume message body is plain text
	if err != nil {
		m.Text, err = ioutil.ReadAll(msg.Body)
		return
	}

	mr := multipart.NewReader(msg.Body, bound)
	for {
		part, err := mr.NextPart()
		if err != nil {
			break
		}

		ct, err := getContentType(part.Header)
		if err != nil {
			return err
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
