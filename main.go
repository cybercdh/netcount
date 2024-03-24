/*
read in a list of CIDR ranges and outputs how many IPs are in each range
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		cidr := strings.TrimSpace(scanner.Text())
		if cidr == "" {
			continue
		}

		_, ipNet, err := net.ParseCIDR(cidr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing CIDR %s: %v\n", cidr, err)
			continue
		}

		// Calculate the number of IP addresses
		ones, bits := ipNet.Mask.Size()
		numIPs := 1 << (bits - ones)

		fmt.Printf("%s,%d\n", cidr, numIPs)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}
