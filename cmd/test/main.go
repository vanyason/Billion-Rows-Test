package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math/bits"
	"net"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
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
	word, bit := pos/64, uint(pos%64)
	mask := uint64(1) << bit

	// Atomically set the bit using OR operation
	for {
		old := atomic.LoadUint64(&b.bits[word])
		new := old | mask
		if atomic.CompareAndSwapUint64(&b.bits[word], old, new) {
			break
		}
	}
}

func (b *BitSet) GetUniqueIPCount() int {
	count := 0
	for _, block := range b.bits {
		count += bits.OnesCount64(block)
	}
	return count
}

type part struct {
	offset, size int64
}

func splitFile(inputPath string, numParts int) ([]part, error) {
	const maxLineLength = 100

	f, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer f.Close() // Ensure file is closed after the function ends

	st, err := f.Stat()
	if err != nil {
		return nil, err
	}
	size := st.Size()
	splitSize := size / int64(numParts)

	parts := make([]part, 0, numParts)
	offset := int64(0)

	for i := 0; i < numParts && offset < size; i++ {
		seekOffset := offset + splitSize

		// Move the seek offset back to the start of a new line
		if seekOffset >= size {
			seekOffset = size
		} else {
			if _, err := f.Seek(seekOffset, io.SeekStart); err != nil {
				return nil, err
			}

			buf := make([]byte, maxLineLength)
			n, err := io.ReadFull(f, buf)
			if err != nil && err != io.ErrUnexpectedEOF {
				return nil, err
			}

			chunk := buf[:n]
			newline := bytes.IndexByte(chunk, '\n')
			if newline < 0 {
				return nil, fmt.Errorf("newline not found near offset %d", seekOffset)
			}

			// Adjust the seekOffset to the first newline after splitSize
			seekOffset += int64(newline + 1)
		}

		// Append the part with start offset and length
		parts = append(parts, part{offset: offset, size: seekOffset - offset})
		offset = seekOffset
	}

	return parts, nil
}

func processRange(fileOffset int64, fileSize int64, bitset *BitSet, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open("ip.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Seek(fileOffset, io.SeekStart)
	if err != nil {
		panic(err)
	}
	f := io.LimitedReader{R: file, N: fileSize}

	scanner := bufio.NewScanner(&f)

	// Process lines in the range
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
}

func main() {
	const MAX_IP = 4294967296 // 2^32

	numThreads := runtime.NumCPU() - 1
	var wg sync.WaitGroup
	bitset := NewBitSet(MAX_IP)

	parts, err := splitFile("ip.txt", numThreads)
	if err != nil {
		fmt.Println("Error splitting file:", err)
		return
	}

	wg.Add(numThreads)
	for i := 0; i < numThreads; i++ {
		go processRange(parts[i].offset, parts[i].size, bitset, &wg)
	}

	wg.Wait()

	fmt.Println("Number of unique IP addresses:", bitset.GetUniqueIPCount())
}
