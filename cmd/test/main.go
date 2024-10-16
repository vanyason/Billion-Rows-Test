package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"runtime"
	"sync"
)

type BitSet struct {
	bits []uint64
}

func NewBitSet(size int) *BitSet {
	return &BitSet{
		bits: make([]uint64, (size+63)/64), // +63 ensures we round up to the nearest 64-bit block
	}
}

func (b *BitSet) Set(pos uint32) {
	index := pos / 64            // Determine which 64-bit block the bit falls into
	offset := pos % 64           // Determine the bit's position within the block
	b.bits[index] |= 1 << offset // Use bitwise OR to set the bit
}

func (b *BitSet) GetUniqueIPCount() int {
	count := 0
	for _, block := range b.bits {
		for block != 0 {
			count += int(block & 1)
			block >>= 1
		}
	}
	return count
}

const (
	MAX_IP      = 4294967296 // 2^32
	TOTAL_LINES = 1000000000
)

func processLineRange(startLine, endLine int, bitset *BitSet, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open("ip.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Skip to the start line
	for i := 0; i < startLine && scanner.Scan(); i++ {
	}

	// Process lines in the range
	for i := startLine; i < endLine && scanner.Scan(); i++ {
		ipStr := scanner.Text()
		ip := net.ParseIP(ipStr).To4()
		if ip == nil {
			continue
		}
		bitset.Set(binary.BigEndian.Uint32(ip))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func main() {
	numThreads := runtime.NumCPU()
	linesPerGoroutine := TOTAL_LINES / numThreads
	var wg sync.WaitGroup
	bitset := NewBitSet(MAX_IP)

	for i := 0; i < numThreads; i++ {
		startLine := i * linesPerGoroutine
		endLine := startLine + linesPerGoroutine
		if i == numThreads-1 {
			endLine = TOTAL_LINES + 1 // Ensure the last goroutine reads any remaining lines
		}
		wg.Add(1)
		go processLineRange(startLine, endLine, bitset, &wg)
	}

	wg.Wait()
	fmt.Println("Number of unique IP addresses:", bitset.GetUniqueIPCount())
}
