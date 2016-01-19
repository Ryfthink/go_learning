package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
包 http 通过任何实现了 http.Handler 的值来响应 HTTP 请求：

package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}

*/
func main() {
	// 1. test hello
	// testHello()

	// 2. test SubWeb
	testSubWeb()
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (s String) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Fprint(w, s)
}

func (s *Struct) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
}

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

func testSubWeb() {
	http.Handle("/string", String("I'm a frayed knot."))
 	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
 	http.ListenAndServe("localhost:4000", nil)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Fprint(w, "Hello Web Service !")
}

func testHello() {
	var h Hello = Hello{}
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}
}



