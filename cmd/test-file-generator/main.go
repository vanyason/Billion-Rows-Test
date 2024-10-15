package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// GenerateRandomIP generates a random IP address in the format X.X.X.X
func GenerateRandomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		rand.Intn(256),
		rand.Intn(256),
		rand.Intn(256),
		rand.Intn(256))
}

func main() {
	// Check command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/test-file-generator/main.go <number_of_ips>")
		return
	}

	// Parse the number of IPs to generate
	numIPs, err := strconv.Atoi(os.Args[1])
	if err != nil || numIPs <= 0 {
		fmt.Println("Please enter a valid positive integer for the number of IP addresses.")
		return
	}

	// Open the file for writing
	file, err := os.Create("ip.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Generate and write random IP addresses to the file
	for i := 0; i < numIPs; i++ {
		ip := GenerateRandomIP()
		_, err := file.WriteString(ip + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}
