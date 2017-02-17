package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)
	A, Z, a, z := byte('A'), byte('Z'), byte('a'), byte('z')
	for i := range b {
		if b[i] >= A && b[i] <= Z {
			b[i] = byte((b[i]-A+13)%26 + A)
		} else if b[i] >= a && b[i] <= z {
			b[i] = byte((b[i]-a+13)%26 + a)
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
