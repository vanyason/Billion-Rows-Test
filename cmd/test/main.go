package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("ip.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buffer := make([]byte, 100*1024*1024) // 100 MB buffer
	scanner.Buffer(buffer, len(buffer))

	set := map[string]struct{}{}
	for scanner.Scan() {
		ip := scanner.Text()
		set[ip] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Number of unique IP addresses:", len(set))
}
