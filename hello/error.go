package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0,errors.New("Zero parameter!")
	}
	return math.Sqrt(value),nil
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

/*
Go 程序使用 error 值来表示错误状态。

与 fmt.Stringer 类似， error 类型是一个内建接口：

type error interface {
    Error() string
}
*/
func main() {
	result,err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result,err = Sqrt(9)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}

}