package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1.
	var greeting = "Hello go !"
	fmt.Printf("normal: %s\n",greeting)
	fmt.Printf("hex: ")
	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%x ",greeting[i])
	}
	fmt.Printf("\n")
	const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" 
	/*q flag escapes unprintable characters, with + flag it escapses non-ascii characters as well 
	 to make output unambigous  */
	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", sampleText)
	fmt.Printf("\n")

	// 2.Concatenating strings
	greetings := []string{"Hello","go!","Hahaha"}
	fmt.Println(strings.Join(greetings," "))


}