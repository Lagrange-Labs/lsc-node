package crypto

import "github.com/Lagrange-Labs/lagrange-node/utils"

var (
	LeafPrefix  = []byte{1}
	InnerPrefix = []byte{2}
)

// MerkleRoot returns the Merkle root of the given data array.
func MerkleRoot(data [][]byte) []byte {

	if len(data) == 0 {
		return utils.Hash([]byte{})
	}

	emptyLeaf := leafHash([]byte{})
	// expand the leaf nodes to a power of 2
	count := nextPowerOfTwo(len(data))
	nodes := make([][]byte, 0, count)

	for _, d := range data {
		nodes = append(nodes, leafHash(d))
	}
	for i := len(data); i < count; i++ {
		nodes = append(nodes, emptyLeaf)
	}

	// calculate the root
	for count > 1 {
		for i := 0; i < count; i += 2 {
			left := nodes[i]
			right := nodes[i+1]
			nodes[i/2] = innerHash(left, right)
		}
		count /= 2
	}

	return nodes[0]
}

func nextPowerOfTwo(n int) int {
	if n <= 0 {
		return 1
	}
	// Subtract 1 from n and bitwise OR it with the result to set all bits to the right of the highest set bit to 1
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	// Increment the result to get the next power of two
	return n + 1
}

func leafHash(data []byte) []byte {
	return utils.Hash(LeafPrefix, data)
}

func innerHash(left, right []byte) []byte {
	return utils.Hash(InnerPrefix, left, right)
}
