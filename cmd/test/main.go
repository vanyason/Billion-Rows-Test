package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
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

const MAX_IP = 4294967296 // 2^32

func main() {
	file, err := os.Open("ip.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buffer := make([]byte, 64*1024*1024) // 64 MB buffer
	scanner.Buffer(buffer, len(buffer))

	bitset := NewBitSet(MAX_IP)

	for scanner.Scan() {
		ipStr := scanner.Text()
		ip := net.ParseIP(ipStr).To4()

		if ip == nil {
			fmt.Println("Invalid IP address:", ipStr)
			continue
		}

		bitset.Set(binary.BigEndian.Uint32(ip))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Number of unique IP addresses:", bitset.GetUniqueIPCount())
}
