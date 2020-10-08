package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (m rot13Reader) Read(l []byte) (int, error) {
	length, err := m.r.Read(l)
	if err != nil {
		return length, err
	}

	for i := 0; i < length; i++ {
		switch {
		case 'a' <= l[i] && l[i] <= 'm':
			l[i] += 13
		case 'A' <= l[i] && l[i] <= 'M':
			l[i] += 13
		case 'n' <= l[i] && l[i] <= 'z':
			l[i] -= 13
		case 'N' <= l[i] && l[i] <= 'Z':
			l[i] -= 13
		}
	}

	return length, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
