package main

// func Encode(msg, blockSize int) <-chan []int {
// 	ch := make()

// 	for block in msg {
// 		encodeBlock(block)
// 	}

// 	return ch
// }

func EncodeBlock() []int {
	msg := 472
	blockSize := 4
	result := make([]int, blockSize*blockSize)
	for i := 3; i < blockSize*blockSize; i++ {
		if i&(i-1) == 0 {
			continue
		}
		mod := msg % 2
		result[i] = mod
		msg >>= 1

		// parity update
		if mod == 1 {
			result[0] ^= 1
			for j := 1; i>>j > 0; j++ {
				result[1<<j] ^= (i >> j) % 2
				result[0] ^= (i >> j) % 2
			}
		}
	}
	return result
}
