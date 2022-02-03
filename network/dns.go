// DNS resolution

package dns

import (
	"fmt"
	"net"
	"strings"
)

// Function to resolve a Website (string) to an IP address (string)
func Lookup(Website string) string {

	ip, err := net.LookupHost(Website)

	if err != nil {
		fmt.Println(err)
	}

	address := strings.Join(ip, "")

	fmt.Println(address)

	return string(address)

}
