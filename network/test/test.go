package main

import (
	"fmt"

	dns "../../network"
)

func main() {
	t1 := "cnn.com"
	t2 := "abc.123.xyz"

	fmt.Println(dns.Lookup(t1))    // Pass - Array returned
	fmt.Println(dns.Lookup(t1)[0]) // Pass - 1st postion printed
	fmt.Println(dns.Lookup(t2))    // Fail - no such host
}
