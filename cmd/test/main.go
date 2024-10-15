package main

import (
	"fmt"
	"os"
)

func main() {
	// Try opening the file for reading
	file, err := os.Open("ip.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Naive approach to read the file line by line, save to hash map and print hash map length
	set := map[string]struct{}{}
	for {
		var ip string
		_, err := fmt.Fscanln(file, &ip)
		if err != nil {
			break
		}
		set[ip] = struct{}{}
	}

	fmt.Println("Number of unique IP addresses:", len(set))
}
