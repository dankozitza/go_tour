package main

import (
	"io"
	"os"
	"strings"
)

type rotReader struct {
	r       io.Reader
	letters []byte
	rotby   int
}

func (rn *rotReader) Read(p []byte) (n int, e error) {

	if len(p) == 0 {
		return 0, nil
	}

	n, e = rn.r.Read(p)

	l_map := make(map[byte]int, len(p)+1)

	for i, l := range rn.letters {
		l_map[l] = i
	}

	for i, l := range p {
		offset, ok := l_map[l]
		if ok {

			offset += rn.rotby

			if offset >= len(rn.letters) {
				offset -= len(rn.letters)
			}

			p[i] = rn.letters[offset]
		}
	}

	return
}

func main() {
	s := strings.NewReader(
		strings.ToLower("Lbh penxrq gur pbqr!\n"))
	l := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	rn := rotReader{s, l, 13}

	io.Copy(os.Stdout, &rn)
}
