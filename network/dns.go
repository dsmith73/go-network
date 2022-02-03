// DNS resolution

package dns

import (
	"fmt"
	"net"
)

// Function to resolve a Website (string) to an IP address ([]string)
// Returns an array of IP if multiple are found.
func Lookup(Website string) (ip []string) {

	ip, err := net.LookupHost(Website)
	if err != nil {
		fmt.Println(err)
	}
	return ip
}
