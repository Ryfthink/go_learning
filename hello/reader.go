package main

import (
	"fmt"
	"strings"
	"io"
)

func main() {
	// 1. 
	r := strings.NewReader("Hello, Reader!")
	fmt.Printf("r type is %T \n",r)

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("read : b[:n] = %q \n", b[:n])
		if err == io.EOF {
			break
		}
	}

	// 2. 实现一个 Reader 类型，它不断生成 ASCII 字符 'A' 的流。
	b = make([]byte, 8)
	r2 := MyReader{}
	for i := 0; i < 16; i++ {
		n, err := r2.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("read : b[:n] = %q \n", b[:n])
	}
}


type MyReader struct{}


func (m MyReader) Read(b []byte) (int,error) {
    for x := range b {
        b[x] = 'A'
    }
    return len(b), nil
}
