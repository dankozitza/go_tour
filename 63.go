package main

import (
	//	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r       io.Reader
	letters []byte
}

type rotReader struct {
	r       io.Reader
	letters []byte
	rotby   int
}

func (r13 *rot13Reader) Read(p []byte) (n int, e error) {

	if len(p) == 0 {
		return 0, nil
	}

	n, e = r13.r.Read(p)

	l_map := make(map[byte]int, len(p)+1)

	for i, l := range r13.letters {
		//fmt.Println("letter: [", string(l), "], index: [", i, "]")
		l_map[l] = i
	}

	for i, l := range p {
		offset, ok := l_map[l]
		if ok {

			offset += 13

			if offset >= len(r13.letters) {
				offset -= len(r13.letters)
			}

			//fmt.Println("position: [", l_map[l], "], offset: [", offset, "], i: [", i, "], l: [", string(l), "]")

			p[i] = r13.letters[offset]
		}
	}

	return
}

func main() {
	s := strings.NewReader(
		"lbh penxrq gur pbqr!\n")
	r := rot13Reader{s, []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}}

	io.Copy(os.Stdout, &r)
}
