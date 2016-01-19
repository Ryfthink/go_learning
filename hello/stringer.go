package main

import "fmt"

/*
一个普遍存在的接口是 fmt 包中定义的 Stringer。

type Stringer interface {
    String() string
}
Stringer 是一个可以用字符串描述自己的类型。`fmt`包 （还有许多其他包）使用这个来进行输出。
*/

func main() {
	// 1.
	testPerson()

	// 2.
	testIp()
}

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func testPerson() {
	a := Person{"John",21}
	b := Person{"Meimei",19}
	fmt.Println(a)
	fmt.Println(b)
}

//////////////////////////////////////////////////////////////////////////////////////////////////

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d",ip[0],ip[1],ip[2],ip[3])
}

// IPAddr{1, 2, 3, 4} 应当输出 "1.2.3.4"。
func testIp() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}