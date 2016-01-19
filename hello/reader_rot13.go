package main

import (
	"os"
	"io"
	"strings"
)

type Rot13Reader struct {
	r io.Reader
}

func (rot *Rot13Reader) Read (p []byte) (n int, err error) {
	n,err = rot.r.Read(p)
	for i := 0; i < len(p); i++ {
		if (p[i] >= 'A' && p[i] < 'N') || (p[i] >='a' && p[i] < 'n') {
			p[i] += 13
		} else if (p[i] > 'M' && p[i] <= 'Z') || (p[i] > 'm' && p[i] <= 'z'){
			p[i] -= 13
		}
	}
	return
}

// 自定义过滤流，实现 rot13 对称加密
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := Rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
