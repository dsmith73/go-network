package main

import (
	"fmt"

	dns "../../network"
)

func main() {
	t1 := "cnn.com"
	t2 := "abc.123.xyz"

	fmt.Println(dns.WebLookup(t1))    // Pass - Array returned
	fmt.Println(dns.WebLookup(t1)[0]) // Pass - 1st IP printed
	fmt.Println(dns.WebLookup(t2))    // Fail - no such host
}
